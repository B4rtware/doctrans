// Code generated by go-swagger; DO NOT EDIT.

package d_t_a_server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/theovassiliou/doctrans/models"
)

// TransformPipeReader is a Reader for the TransformPipe structure.
type TransformPipeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TransformPipeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTransformPipeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTransformPipeOK creates a TransformPipeOK with default headers values
func NewTransformPipeOK() *TransformPipeOK {
	return &TransformPipeOK{}
}

/*TransformPipeOK handles this case with default header values.

A successful response.
*/
type TransformPipeOK struct {
	Payload *models.DtaserviceTransformDocumentReply
}

func (o *TransformPipeOK) Error() string {
	return fmt.Sprintf("[POST /v1/document/transform-pipe][%d] transformPipeOK  %+v", 200, o.Payload)
}

func (o *TransformPipeOK) GetPayload() *models.DtaserviceTransformDocumentReply {
	return o.Payload
}

func (o *TransformPipeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtaserviceTransformDocumentReply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
