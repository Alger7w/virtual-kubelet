package interaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewContainerSetStdinParams creates a new ContainerSetStdinParams object
// with the default values initialized.
func NewContainerSetStdinParams() *ContainerSetStdinParams {
	var ()
	return &ContainerSetStdinParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewContainerSetStdinParamsWithTimeout creates a new ContainerSetStdinParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewContainerSetStdinParamsWithTimeout(timeout time.Duration) *ContainerSetStdinParams {
	var ()
	return &ContainerSetStdinParams{

		timeout: timeout,
	}
}

// NewContainerSetStdinParamsWithContext creates a new ContainerSetStdinParams object
// with the default values initialized, and the ability to set a context for a request
func NewContainerSetStdinParamsWithContext(ctx context.Context) *ContainerSetStdinParams {
	var ()
	return &ContainerSetStdinParams{

		Context: ctx,
	}
}

// NewContainerSetStdinParamsWithHTTPClient creates a new ContainerSetStdinParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewContainerSetStdinParamsWithHTTPClient(client *http.Client) *ContainerSetStdinParams {
	var ()
	return &ContainerSetStdinParams{
		HTTPClient: client,
	}
}

/*ContainerSetStdinParams contains all the parameters to send to the API endpoint
for the container set stdin operation typically these are written to a http.Request
*/
type ContainerSetStdinParams struct {

	/*Deadline*/
	Deadline *strfmt.DateTime
	/*ID*/
	ID string
	/*RawStream*/
	RawStream io.ReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the container set stdin params
func (o *ContainerSetStdinParams) WithTimeout(timeout time.Duration) *ContainerSetStdinParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container set stdin params
func (o *ContainerSetStdinParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container set stdin params
func (o *ContainerSetStdinParams) WithContext(ctx context.Context) *ContainerSetStdinParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container set stdin params
func (o *ContainerSetStdinParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container set stdin params
func (o *ContainerSetStdinParams) WithHTTPClient(client *http.Client) *ContainerSetStdinParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container set stdin params
func (o *ContainerSetStdinParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeadline adds the deadline to the container set stdin params
func (o *ContainerSetStdinParams) WithDeadline(deadline *strfmt.DateTime) *ContainerSetStdinParams {
	o.SetDeadline(deadline)
	return o
}

// SetDeadline adds the deadline to the container set stdin params
func (o *ContainerSetStdinParams) SetDeadline(deadline *strfmt.DateTime) {
	o.Deadline = deadline
}

// WithID adds the id to the container set stdin params
func (o *ContainerSetStdinParams) WithID(id string) *ContainerSetStdinParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the container set stdin params
func (o *ContainerSetStdinParams) SetID(id string) {
	o.ID = id
}

// WithRawStream adds the rawStream to the container set stdin params
func (o *ContainerSetStdinParams) WithRawStream(rawStream io.ReadCloser) *ContainerSetStdinParams {
	o.SetRawStream(rawStream)
	return o
}

// SetRawStream adds the rawStream to the container set stdin params
func (o *ContainerSetStdinParams) SetRawStream(rawStream io.ReadCloser) {
	o.RawStream = rawStream
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerSetStdinParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Deadline != nil {

		// query param deadline
		var qrDeadline strfmt.DateTime
		if o.Deadline != nil {
			qrDeadline = *o.Deadline
		}
		qDeadline := qrDeadline.String()
		if qDeadline != "" {
			if err := r.SetQueryParam("deadline", qDeadline); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.RawStream); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
