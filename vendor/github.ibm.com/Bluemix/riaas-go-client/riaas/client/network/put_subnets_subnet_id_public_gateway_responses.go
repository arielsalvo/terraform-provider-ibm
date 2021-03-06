// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// PutSubnetsSubnetIDPublicGatewayReader is a Reader for the PutSubnetsSubnetIDPublicGateway structure.
type PutSubnetsSubnetIDPublicGatewayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutSubnetsSubnetIDPublicGatewayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPutSubnetsSubnetIDPublicGatewayCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutSubnetsSubnetIDPublicGatewayBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPutSubnetsSubnetIDPublicGatewayInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutSubnetsSubnetIDPublicGatewayCreated creates a PutSubnetsSubnetIDPublicGatewayCreated with default headers values
func NewPutSubnetsSubnetIDPublicGatewayCreated() *PutSubnetsSubnetIDPublicGatewayCreated {
	return &PutSubnetsSubnetIDPublicGatewayCreated{}
}

/*PutSubnetsSubnetIDPublicGatewayCreated handles this case with default header values.

dummy
*/
type PutSubnetsSubnetIDPublicGatewayCreated struct {
	Payload *models.PublicGateway
}

func (o *PutSubnetsSubnetIDPublicGatewayCreated) Error() string {
	return fmt.Sprintf("[PUT /subnets/{subnet_id}/public_gateway][%d] putSubnetsSubnetIdPublicGatewayCreated  %+v", 201, o.Payload)
}

func (o *PutSubnetsSubnetIDPublicGatewayCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicGateway)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutSubnetsSubnetIDPublicGatewayBadRequest creates a PutSubnetsSubnetIDPublicGatewayBadRequest with default headers values
func NewPutSubnetsSubnetIDPublicGatewayBadRequest() *PutSubnetsSubnetIDPublicGatewayBadRequest {
	return &PutSubnetsSubnetIDPublicGatewayBadRequest{}
}

/*PutSubnetsSubnetIDPublicGatewayBadRequest handles this case with default header values.

error
*/
type PutSubnetsSubnetIDPublicGatewayBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PutSubnetsSubnetIDPublicGatewayBadRequest) Error() string {
	return fmt.Sprintf("[PUT /subnets/{subnet_id}/public_gateway][%d] putSubnetsSubnetIdPublicGatewayBadRequest  %+v", 400, o.Payload)
}

func (o *PutSubnetsSubnetIDPublicGatewayBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutSubnetsSubnetIDPublicGatewayInternalServerError creates a PutSubnetsSubnetIDPublicGatewayInternalServerError with default headers values
func NewPutSubnetsSubnetIDPublicGatewayInternalServerError() *PutSubnetsSubnetIDPublicGatewayInternalServerError {
	return &PutSubnetsSubnetIDPublicGatewayInternalServerError{}
}

/*PutSubnetsSubnetIDPublicGatewayInternalServerError handles this case with default header values.

error
*/
type PutSubnetsSubnetIDPublicGatewayInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *PutSubnetsSubnetIDPublicGatewayInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /subnets/{subnet_id}/public_gateway][%d] putSubnetsSubnetIdPublicGatewayInternalServerError  %+v", 500, o.Payload)
}

func (o *PutSubnetsSubnetIDPublicGatewayInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
