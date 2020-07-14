/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package websocket_test

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	apiv1 "github.com/kubermatic/kubermatic/pkg/api/v1"
	v1 "github.com/kubermatic/kubermatic/pkg/crd/kubermatic/v1"
	"github.com/kubermatic/kubermatic/pkg/handler/test"
	"github.com/kubermatic/kubermatic/pkg/handler/test/hack"
)

func TestUserWatchEndpoint(t *testing.T) {
	t.Parallel()
	testcases := []struct {
		name                string
		userToUpdate        string
		userSettingsUpdate  *v1.UserSettings
		userUpdate          *apiv1.User
		existingAPIUser     *apiv1.User
		existingUsers       []*apiv1.User
		updateShouldTimeout bool
	}{
		{
			name:         "should be able to watch and notice user setting change on its own user",
			userToUpdate: test.GenDefaultAPIUser().Name,
			userSettingsUpdate: &v1.UserSettings{
				CollapseSidenav: true,
			},
			existingAPIUser:     test.GenDefaultAPIUser(),
			existingUsers:       []*apiv1.User{test.GenDefaultAPIUser()},
			updateShouldTimeout: false,
		},
		{
			name:         "should be able to watch and but not notice the user setting change on a different user",
			userToUpdate: test.GenAPIUser("john", "john@acme.com").Name,
			userSettingsUpdate: &v1.UserSettings{
				CollapseSidenav: true,
			},
			existingAPIUser:     test.GenDefaultAPIUser(),
			existingUsers:       []*apiv1.User{test.GenDefaultAPIUser(), test.GenAPIUser("john", "john@acme.com")},
			updateShouldTimeout: true,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			var runtimeObjectUsers []runtime.Object
			for _, user := range tc.existingUsers {
				runtimeObjectUsers = append(runtimeObjectUsers, test.APIUserToKubermaticUser(*user))
			}

			ep, cli, err := test.CreateTestEndpointAndGetClients(*tc.existingAPIUser, nil, []runtime.Object{}, nil,
				runtimeObjectUsers, nil, nil, hack.NewTestRouting)
			if err != nil {
				t.Fatalf("failed to create test endpoint due to %v", err)
			}
			server := httptest.NewServer(ep)
			defer server.Close()

			// setup ws client
			wsURL := "ws" + strings.TrimPrefix(server.URL, "http") + "/api/v1/ws/me"
			ch, err := createWSClient(t, wsURL)
			if err != nil {
				t.Fatalf("failed to initialize websocket client: %v", err)
			}

			var wsMsg wsMessage
			select {
			case <-time.After(time.Second):
				t.Fatalf("timeout waiting for ws message")
			case wsMsg = <-ch:
			}
			if wsMsg.err != nil {
				t.Fatalf("error reading ws message: %v", err)
			}

			var user *apiv1.User
			err = json.Unmarshal(wsMsg.p, &user)
			if err != nil {
				t.Fatalf("failed unmarshalling user: %v", err)
			}
			if user.Name != tc.existingAPIUser.Name {
				t.Fatalf("got wrong initial user from watch, expected: %s, got %s", tc.existingAPIUser.Name, user.Name)
			}

			// Update user to get watch notification
			userToUpdate, err := cli.FakeKubermaticClient.KubermaticV1().Users().Get(tc.userToUpdate, metav1.GetOptions{})
			if err != nil {
				t.Fatalf("error getting user to update: %v", err)
			}
			userToUpdate.Spec.Settings = tc.userSettingsUpdate

			time.Sleep(time.Second)
			_, err = cli.FakeKubermaticClient.KubermaticV1().Users().Update(userToUpdate)
			if err != nil {
				t.Fatalf("error updating user: %v", err)
			}

			// get the update notification
			select {
			case <-time.After(time.Second):
				if !tc.updateShouldTimeout {
					t.Fatal("Watch update notification didnt arrive in time")
				}
			case wsMsg = <-ch:
			}
			if wsMsg.err != nil {
				t.Fatalf("error reading ws message: %v", err)
			}

			if !tc.updateShouldTimeout {
				var userUpdate *apiv1.User
				err = json.Unmarshal(wsMsg.p, &userUpdate)
				if err != nil {
					t.Fatalf("failed unmarshalling user: %v", err)
				}

				if !reflect.DeepEqual(userUpdate.Settings, tc.userSettingsUpdate) {
					t.Fatalf("expected settings %v, got %v", tc.userSettingsUpdate, userUpdate.Settings)
				}
			}
		})
	}
}

type wsMessage struct {
	messageType int
	p           []byte
	err         error
}

func createWSClient(t *testing.T, url string) (chan wsMessage, error) {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize websocket dialer: %v", err)
	}

	ch := make(chan wsMessage, 5)

	go func() {
		for {
			ty, p, err := ws.ReadMessage()
			ch <- wsMessage{
				messageType: ty,
				p:           p,
				err:         err,
			}
			t.Logf("Read a message from WS, type: %d, contents: %s", ty, string(p))
			if err != nil {
				t.Logf("error reading ws message, closing ws channel: %v", err)
				close(ch)
				ws.Close()
				break
			}
		}
	}()

	return ch, nil
}
