// Code generated by go-swagger; DO NOT EDIT.

package carrier_wifi_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// GetCwfNetworkIDGatewaysGatewayIDTierReader is a Reader for the GetCwfNetworkIDGatewaysGatewayIDTier structure.
type GetCwfNetworkIDGatewaysGatewayIDTierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCwfNetworkIDGatewaysGatewayIDTierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCwfNetworkIDGatewaysGatewayIDTierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCwfNetworkIDGatewaysGatewayIDTierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCwfNetworkIDGatewaysGatewayIDTierOK creates a GetCwfNetworkIDGatewaysGatewayIDTierOK with default headers values
func NewGetCwfNetworkIDGatewaysGatewayIDTierOK() *GetCwfNetworkIDGatewaysGatewayIDTierOK {
	return &GetCwfNetworkIDGatewaysGatewayIDTierOK{}
}

/*GetCwfNetworkIDGatewaysGatewayIDTierOK handles this case with default header values.

The ID of the upgrade tier
*/
type GetCwfNetworkIDGatewaysGatewayIDTierOK struct {
	Payload models.TierID
}

func (o *GetCwfNetworkIDGatewaysGatewayIDTierOK) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/gateways/{gateway_id}/tier][%d] getCwfNetworkIdGatewaysGatewayIdTierOK  %+v", 200, o.Payload)
}

func (o *GetCwfNetworkIDGatewaysGatewayIDTierOK) GetPayload() models.TierID {
	return o.Payload
}

func (o *GetCwfNetworkIDGatewaysGatewayIDTierOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCwfNetworkIDGatewaysGatewayIDTierDefault creates a GetCwfNetworkIDGatewaysGatewayIDTierDefault with default headers values
func NewGetCwfNetworkIDGatewaysGatewayIDTierDefault(code int) *GetCwfNetworkIDGatewaysGatewayIDTierDefault {
	return &GetCwfNetworkIDGatewaysGatewayIDTierDefault{
		_statusCode: code,
	}
}

/*GetCwfNetworkIDGatewaysGatewayIDTierDefault handles this case with default header values.

Unexpected Error
*/
type GetCwfNetworkIDGatewaysGatewayIDTierDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get cwf network ID gateways gateway ID tier default response
func (o *GetCwfNetworkIDGatewaysGatewayIDTierDefault) Code() int {
	return o._statusCode
}

func (o *GetCwfNetworkIDGatewaysGatewayIDTierDefault) Error() string {
	return fmt.Sprintf("[GET /cwf/{network_id}/gateways/{gateway_id}/tier][%d] GetCwfNetworkIDGatewaysGatewayIDTier default  %+v", o._statusCode, o.Payload)
}

func (o *GetCwfNetworkIDGatewaysGatewayIDTierDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCwfNetworkIDGatewaysGatewayIDTierDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}