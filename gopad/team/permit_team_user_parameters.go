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
	"github.com/go-openapi/strfmt"

	"github.com/gopad/gopad-go/v1/models"
)

// NewPermitTeamUserParams creates a new PermitTeamUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPermitTeamUserParams() *PermitTeamUserParams {
	return &PermitTeamUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPermitTeamUserParamsWithTimeout creates a new PermitTeamUserParams object
// with the ability to set a timeout on a request.
func NewPermitTeamUserParamsWithTimeout(timeout time.Duration) *PermitTeamUserParams {
	return &PermitTeamUserParams{
		timeout: timeout,
	}
}

// NewPermitTeamUserParamsWithContext creates a new PermitTeamUserParams object
// with the ability to set a context for a request.
func NewPermitTeamUserParamsWithContext(ctx context.Context) *PermitTeamUserParams {
	return &PermitTeamUserParams{
		Context: ctx,
	}
}

// NewPermitTeamUserParamsWithHTTPClient creates a new PermitTeamUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewPermitTeamUserParamsWithHTTPClient(client *http.Client) *PermitTeamUserParams {
	return &PermitTeamUserParams{
		HTTPClient: client,
	}
}

/* PermitTeamUserParams contains all the parameters to send to the API endpoint
   for the permit team user operation.

   Typically these are written to a http.Request.
*/
type PermitTeamUserParams struct {

	/* TeamID.

	   A team UUID or slug
	*/
	TeamID string

	/* TeamUser.

	   The team user data to update
	*/
	TeamUser *models.TeamUserParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the permit team user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PermitTeamUserParams) WithDefaults() *PermitTeamUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the permit team user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PermitTeamUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the permit team user params
func (o *PermitTeamUserParams) WithTimeout(timeout time.Duration) *PermitTeamUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the permit team user params
func (o *PermitTeamUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the permit team user params
func (o *PermitTeamUserParams) WithContext(ctx context.Context) *PermitTeamUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the permit team user params
func (o *PermitTeamUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the permit team user params
func (o *PermitTeamUserParams) WithHTTPClient(client *http.Client) *PermitTeamUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the permit team user params
func (o *PermitTeamUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTeamID adds the teamID to the permit team user params
func (o *PermitTeamUserParams) WithTeamID(teamID string) *PermitTeamUserParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the permit team user params
func (o *PermitTeamUserParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WithTeamUser adds the teamUser to the permit team user params
func (o *PermitTeamUserParams) WithTeamUser(teamUser *models.TeamUserParams) *PermitTeamUserParams {
	o.SetTeamUser(teamUser)
	return o
}

// SetTeamUser adds the teamUser to the permit team user params
func (o *PermitTeamUserParams) SetTeamUser(teamUser *models.TeamUserParams) {
	o.TeamUser = teamUser
}

// WriteToRequest writes these params to a swagger request
func (o *PermitTeamUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param team_id
	if err := r.SetPathParam("team_id", o.TeamID); err != nil {
		return err
	}
	if o.TeamUser != nil {
		if err := r.SetBodyParam(o.TeamUser); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
