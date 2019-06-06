// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteTeamParams creates a new DeleteTeamParams object
// with the default values initialized.
func NewDeleteTeamParams() *DeleteTeamParams {
	var ()
	return &DeleteTeamParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTeamParamsWithTimeout creates a new DeleteTeamParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteTeamParamsWithTimeout(timeout time.Duration) *DeleteTeamParams {
	var ()
	return &DeleteTeamParams{

		timeout: timeout,
	}
}

// NewDeleteTeamParamsWithContext creates a new DeleteTeamParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteTeamParamsWithContext(ctx context.Context) *DeleteTeamParams {
	var ()
	return &DeleteTeamParams{

		Context: ctx,
	}
}

// NewDeleteTeamParamsWithHTTPClient creates a new DeleteTeamParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteTeamParamsWithHTTPClient(client *http.Client) *DeleteTeamParams {
	var ()
	return &DeleteTeamParams{
		HTTPClient: client,
	}
}

/*DeleteTeamParams contains all the parameters to send to the API endpoint
for the delete team operation typically these are written to a http.Request
*/
type DeleteTeamParams struct {

	/*TeamID
	  A team UUID or slug

	*/
	TeamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete team params
func (o *DeleteTeamParams) WithTimeout(timeout time.Duration) *DeleteTeamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete team params
func (o *DeleteTeamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete team params
func (o *DeleteTeamParams) WithContext(ctx context.Context) *DeleteTeamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete team params
func (o *DeleteTeamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete team params
func (o *DeleteTeamParams) WithHTTPClient(client *http.Client) *DeleteTeamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete team params
func (o *DeleteTeamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTeamID adds the teamID to the delete team params
func (o *DeleteTeamParams) WithTeamID(teamID string) *DeleteTeamParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the delete team params
func (o *DeleteTeamParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTeamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param team_id
	if err := r.SetPathParam("team_id", o.TeamID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
