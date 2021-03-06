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

package semver

import (
	"encoding/json"
	"flag"
	"fmt"
	"strconv"

	semverlib "github.com/Masterminds/semver/v3"
)

// ensure that Semver implements flag.Value interface.
var _ flag.Value = &Semver{}

// Semver is struct that encapsulates semver.Semver struct so we can use it in API
// +k8s:deepcopy-gen=true
type Semver struct {
	*semverlib.Version
}

// NewSemver creates new Semver version struct and returns pointer to it
func NewSemver(ver string) (*Semver, error) {
	v := &Semver{}
	err := v.Set(ver)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// NewSemverOrDie behaves similar to NewVersion, i.e. it creates new Semver version struct, but panics if an error happens
func NewSemverOrDie(ver string) *Semver {
	sv, err := NewSemver(ver)
	if err != nil {
		panic(err)
	}
	return sv
}

// Set initializes semver struct and sets version
func (s *Semver) Set(ver string) error {
	sv, err := semverlib.NewVersion(ver)
	if err != nil {
		return err
	}
	s.Version = sv
	return nil
}

// Semver returns library semver struct
func (s *Semver) Semver() *semverlib.Version {
	return s.Version
}

// Equal compares two version structs by comparing Semver values
func (s *Semver) Equal(b *Semver) bool {
	return s.Version.Equal(b.Version)
}

// String returns string representation of Semver version
func (s *Semver) String() string {
	if s.Version == nil {
		return ""
	}
	return s.Version.String()
}

// MajorMinor returns a string like "Major.Minor"
func (s *Semver) MajorMinor() string {
	return fmt.Sprintf("%d.%d", s.Major(), s.Minor())
}

// UnmarshalJSON converts JSON to Semver struct
func (s *Semver) UnmarshalJSON(data []byte) error {
	ver, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}
	if ver == "" {
		return nil
	}
	return s.Set(ver)
}

// MarshalJSON converts Semver struct to JSON
func (s Semver) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(s.String())), nil
}

var _ json.Marshaler = &Semver{}
var _ json.Unmarshaler = &Semver{}

// DeepCopy copies value of Semver struct and returns a new struct.
// If passed Semver struct is nil, it is assumed zero value is being copied
func (s Semver) DeepCopy() Semver {
	if s.Version == nil {
		return Semver{}
	}
	return *NewSemverOrDie(s.String())
}
