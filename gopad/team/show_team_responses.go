// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/gopad/gopad-go/models"
)

// ShowTeamReader is a Reader for the ShowTeam structure.
type ShowTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ShowTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewShowTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewShowTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewShowTeamDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewShowTeamOK creates a ShowTeamOK with default headers values
func NewShowTeamOK() *ShowTeamOK {
	return &ShowTeamOK{}
}

/*ShowTeamOK handles this case with default header values.

The fetched team details
*/
type ShowTeamOK struct {
	Payload *models.Team
}

func (o *ShowTeamOK) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}][%d] showTeamOK  %+v", 200, o.Payload)
}

func (o *ShowTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Team)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowTeamForbidden creates a ShowTeamForbidden with default headers values
func NewShowTeamForbidden() *ShowTeamForbidden {
	return &ShowTeamForbidden{}
}

/*ShowTeamForbidden handles this case with default header values.

User is not authorized
*/
type ShowTeamForbidden struct {
	Payload *models.GeneralError
}

func (o *ShowTeamForbidden) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}][%d] showTeamForbidden  %+v", 403, o.Payload)
}

func (o *ShowTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewShowTeamDefault creates a ShowTeamDefault with default headers values
func NewShowTeamDefault(code int) *ShowTeamDefault {
	return &ShowTeamDefault{
		_statusCode: code,
	}
}

/*ShowTeamDefault handles this case with default header values.

Some error unrelated to the handler
*/
type ShowTeamDefault struct {
	_statusCode int

	Payload *models.GeneralError
}

// Code gets the status code for the show team default response
func (o *ShowTeamDefault) Code() int {
	return o._statusCode
}

func (o *ShowTeamDefault) Error() string {
	return fmt.Sprintf("[GET /teams/{team_id}][%d] ShowTeam default  %+v", o._statusCode, o.Payload)
}

func (o *ShowTeamDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GeneralError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
