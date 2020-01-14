// Code generated by go-swagger; DO NOT EDIT.

package d_t_a_server

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

	models "github.com/theovassiliou/doctrans/models"
)

// NewTransformPipeParams creates a new TransformPipeParams object
// with the default values initialized.
func NewTransformPipeParams() *TransformPipeParams {
	var ()
	return &TransformPipeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTransformPipeParamsWithTimeout creates a new TransformPipeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTransformPipeParamsWithTimeout(timeout time.Duration) *TransformPipeParams {
	var ()
	return &TransformPipeParams{

		timeout: timeout,
	}
}

// NewTransformPipeParamsWithContext creates a new TransformPipeParams object
// with the default values initialized, and the ability to set a context for a request
func NewTransformPipeParamsWithContext(ctx context.Context) *TransformPipeParams {
	var ()
	return &TransformPipeParams{

		Context: ctx,
	}
}

// NewTransformPipeParamsWithHTTPClient creates a new TransformPipeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTransformPipeParamsWithHTTPClient(client *http.Client) *TransformPipeParams {
	var ()
	return &TransformPipeParams{
		HTTPClient: client,
	}
}

/*TransformPipeParams contains all the parameters to send to the API endpoint
for the transform pipe operation typically these are written to a http.Request
*/
type TransformPipeParams struct {

	/*Body*/
	Body *models.DtaserviceTransformPipeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the transform pipe params
func (o *TransformPipeParams) WithTimeout(timeout time.Duration) *TransformPipeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the transform pipe params
func (o *TransformPipeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the transform pipe params
func (o *TransformPipeParams) WithContext(ctx context.Context) *TransformPipeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the transform pipe params
func (o *TransformPipeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the transform pipe params
func (o *TransformPipeParams) WithHTTPClient(client *http.Client) *TransformPipeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the transform pipe params
func (o *TransformPipeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the transform pipe params
func (o *TransformPipeParams) WithBody(body *models.DtaserviceTransformPipeRequest) *TransformPipeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the transform pipe params
func (o *TransformPipeParams) SetBody(body *models.DtaserviceTransformPipeRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TransformPipeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
