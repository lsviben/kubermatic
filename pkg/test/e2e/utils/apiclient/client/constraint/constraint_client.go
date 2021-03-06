// Code generated by go-swagger; DO NOT EDIT.

package constraint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new constraint API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for constraint API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateDefaultConstraint(params *CreateDefaultConstraintParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDefaultConstraintOK, error)

	GetDefaultConstraint(params *GetDefaultConstraintParams, authInfo runtime.ClientAuthInfoWriter) (*GetDefaultConstraintOK, error)

	ListDefaultConstraint(params *ListDefaultConstraintParams, authInfo runtime.ClientAuthInfoWriter) (*ListDefaultConstraintOK, error)

	PatchDefaultConstraint(params *PatchDefaultConstraintParams, authInfo runtime.ClientAuthInfoWriter) (*PatchDefaultConstraintOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateDefaultConstraint Creates default constraint
*/
func (a *Client) CreateDefaultConstraint(params *CreateDefaultConstraintParams, authInfo runtime.ClientAuthInfoWriter) (*CreateDefaultConstraintOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDefaultConstraintParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createDefaultConstraint",
		Method:             "POST",
		PathPattern:        "/api/v2/constraints",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDefaultConstraintReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateDefaultConstraintOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateDefaultConstraintDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetDefaultConstraint Gets an specified default constraint
*/
func (a *Client) GetDefaultConstraint(params *GetDefaultConstraintParams, authInfo runtime.ClientAuthInfoWriter) (*GetDefaultConstraintOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDefaultConstraintParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDefaultConstraint",
		Method:             "GET",
		PathPattern:        "/api/v2/constraints/{constraint_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDefaultConstraintReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDefaultConstraintOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetDefaultConstraintDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListDefaultConstraint lists default constraint
*/
func (a *Client) ListDefaultConstraint(params *ListDefaultConstraintParams, authInfo runtime.ClientAuthInfoWriter) (*ListDefaultConstraintOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListDefaultConstraintParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listDefaultConstraint",
		Method:             "GET",
		PathPattern:        "/api/v2/constraints",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListDefaultConstraintReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListDefaultConstraintOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListDefaultConstraintDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  PatchDefaultConstraint Patch a specified default constraint
*/
func (a *Client) PatchDefaultConstraint(params *PatchDefaultConstraintParams, authInfo runtime.ClientAuthInfoWriter) (*PatchDefaultConstraintOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchDefaultConstraintParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchDefaultConstraint",
		Method:             "PATCH",
		PathPattern:        "/api/v2/constraints/{constraint_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchDefaultConstraintReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchDefaultConstraintOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PatchDefaultConstraintDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
