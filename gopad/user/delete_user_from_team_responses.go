// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/gopad/gopad-go/models"
)

// DeleteUserFromTeamReader is a Reader for the DeleteUserFromTeam structure.
type DeleteUserFromTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserFromTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteUserFromTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewDeleteUserFromTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewDeleteUserFromTeamPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewDeleteUserFromTeamUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteUserFromTeamDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteUserFromTeamOK creates a DeleteUserFromTeamOK with default headers values
func NewDeleteUserFromTeamOK() *DeleteUserFromTeamOK {
	return &DeleteUserFromTeamOK{}
}

/*DeleteUserFromTeamOK handles this case with default header values.

Plain success message
*/
type DeleteUserFromTeamOK struct {
	Payload *models.GeneralError
}

func (o *DeleteUserFromTeamOK) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/teams][%d] deleteUserFromTeamOK  %+v", 200, o.Payload)
}

func (o *DeleteUserFromTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserFromTeamForbidden creates a DeleteUserFromTeamForbidden with default headers values
func NewDeleteUserFromTeamForbidden() *DeleteUserFromTeamForbidden {
	return &DeleteUserFromTeamForbidden{}
}

/*DeleteUserFromTeamForbidden handles this case with default header values.

User is not authorized
*/
type DeleteUserFromTeamForbidden struct {
	Payload *models.GeneralError
}

func (o *DeleteUserFromTeamForbidden) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/teams][%d] deleteUserFromTeamForbidden  %+v", 403, o.Payload)
}

func (o *DeleteUserFromTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserFromTeamPreconditionFailed creates a DeleteUserFromTeamPreconditionFailed with default headers values
func NewDeleteUserFromTeamPreconditionFailed() *DeleteUserFromTeamPreconditionFailed {
	return &DeleteUserFromTeamPreconditionFailed{}
}

/*DeleteUserFromTeamPreconditionFailed handles this case with default header values.

Failed to parse request body
*/
type DeleteUserFromTeamPreconditionFailed struct {
	Payload *models.GeneralError
}

func (o *DeleteUserFromTeamPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/teams][%d] deleteUserFromTeamPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteUserFromTeamPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserFromTeamUnprocessableEntity creates a DeleteUserFromTeamUnprocessableEntity with default headers values
func NewDeleteUserFromTeamUnprocessableEntity() *DeleteUserFromTeamUnprocessableEntity {
	return &DeleteUserFromTeamUnprocessableEntity{}
}

/*DeleteUserFromTeamUnprocessableEntity handles this case with default header values.

Team is not assigned
*/
type DeleteUserFromTeamUnprocessableEntity struct {
	Payload *models.GeneralError
}

func (o *DeleteUserFromTeamUnprocessableEntity) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/teams][%d] deleteUserFromTeamUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *DeleteUserFromTeamUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserFromTeamDefault creates a DeleteUserFromTeamDefault with default headers values
func NewDeleteUserFromTeamDefault(code int) *DeleteUserFromTeamDefault {
	return &DeleteUserFromTeamDefault{
		_statusCode: code,
	}
}

/*DeleteUserFromTeamDefault handles this case with default header values.

Some error unrelated to the handler
*/
type DeleteUserFromTeamDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the delete user from team default response
func (o *DeleteUserFromTeamDefault) Code() int {
	return o._statusCode
}

func (o *DeleteUserFromTeamDefault) Error() string {
	return fmt.Sprintf("[DELETE /users/{user_id}/teams][%d] DeleteUserFromTeam default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteUserFromTeamDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
