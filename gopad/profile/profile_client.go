// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new profile API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for profile API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ShowProfile(params *ShowProfileParams, opts ...ClientOption) (*ShowProfileOK, error)

	TokenProfile(params *TokenProfileParams, opts ...ClientOption) (*TokenProfileOK, error)

	UpdateProfile(params *UpdateProfileParams, opts ...ClientOption) (*UpdateProfileOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ShowProfile retrieves an unlimited auth token
*/
func (a *Client) ShowProfile(params *ShowProfileParams, opts ...ClientOption) (*ShowProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewShowProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ShowProfile",
		Method:             "GET",
		PathPattern:        "/profile/self",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ShowProfileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ShowProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ShowProfileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  TokenProfile retrieves an unlimited auth token
*/
func (a *Client) TokenProfile(params *TokenProfileParams, opts ...ClientOption) (*TokenProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTokenProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "TokenProfile",
		Method:             "GET",
		PathPattern:        "/profile/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TokenProfileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TokenProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*TokenProfileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateProfile retrieves an unlimited auth token
*/
func (a *Client) UpdateProfile(params *UpdateProfileParams, opts ...ClientOption) (*UpdateProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateProfile",
		Method:             "PUT",
		PathPattern:        "/profile/self",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateProfileReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateProfileDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
