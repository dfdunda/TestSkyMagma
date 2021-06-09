// Code generated by go-swagger; DO NOT EDIT.

package e2e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// PostTestsE2eEnodebdReader is a Reader for the PostTestsE2eEnodebd structure.
type PostTestsE2eEnodebdReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTestsE2eEnodebdReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostTestsE2eEnodebdCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPostTestsE2eEnodebdDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostTestsE2eEnodebdCreated creates a PostTestsE2eEnodebdCreated with default headers values
func NewPostTestsE2eEnodebdCreated() *PostTestsE2eEnodebdCreated {
	return &PostTestsE2eEnodebdCreated{}
}

/*PostTestsE2eEnodebdCreated handles this case with default header values.

Created
*/
type PostTestsE2eEnodebdCreated struct {
}

func (o *PostTestsE2eEnodebdCreated) Error() string {
	return fmt.Sprintf("[POST /tests/e2e/enodebd][%d] postTestsE2eEnodebdCreated ", 201)
}

func (o *PostTestsE2eEnodebdCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostTestsE2eEnodebdDefault creates a PostTestsE2eEnodebdDefault with default headers values
func NewPostTestsE2eEnodebdDefault(code int) *PostTestsE2eEnodebdDefault {
	return &PostTestsE2eEnodebdDefault{
		_statusCode: code,
	}
}

/*PostTestsE2eEnodebdDefault handles this case with default header values.

Unexpected Error
*/
type PostTestsE2eEnodebdDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post tests e2e enodebd default response
func (o *PostTestsE2eEnodebdDefault) Code() int {
	return o._statusCode
}

func (o *PostTestsE2eEnodebdDefault) Error() string {
	return fmt.Sprintf("[POST /tests/e2e/enodebd][%d] PostTestsE2eEnodebd default  %+v", o._statusCode, o.Payload)
}

func (o *PostTestsE2eEnodebdDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostTestsE2eEnodebdDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}