// Code generated by go-swagger; DO NOT EDIT.

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeletePaymentdefaultsIDParams creates a new DeletePaymentdefaultsIDParams object
// with the default values initialized.
func NewDeletePaymentdefaultsIDParams() *DeletePaymentdefaultsIDParams {
	var ()
	return &DeletePaymentdefaultsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePaymentdefaultsIDParamsWithTimeout creates a new DeletePaymentdefaultsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePaymentdefaultsIDParamsWithTimeout(timeout time.Duration) *DeletePaymentdefaultsIDParams {
	var ()
	return &DeletePaymentdefaultsIDParams{

		timeout: timeout,
	}
}

// NewDeletePaymentdefaultsIDParamsWithContext creates a new DeletePaymentdefaultsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePaymentdefaultsIDParamsWithContext(ctx context.Context) *DeletePaymentdefaultsIDParams {
	var ()
	return &DeletePaymentdefaultsIDParams{

		Context: ctx,
	}
}

// NewDeletePaymentdefaultsIDParamsWithHTTPClient creates a new DeletePaymentdefaultsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePaymentdefaultsIDParamsWithHTTPClient(client *http.Client) *DeletePaymentdefaultsIDParams {
	var ()
	return &DeletePaymentdefaultsIDParams{
		HTTPClient: client,
	}
}

/*DeletePaymentdefaultsIDParams contains all the parameters to send to the API endpoint
for the delete paymentdefaults ID operation typically these are written to a http.Request
*/
type DeletePaymentdefaultsIDParams struct {

	/*ID
	  Limit Id

	*/
	ID strfmt.UUID
	/*Version
	  Version

	*/
	Version int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete paymentdefaults ID params
func (o *DeletePaymentdefaultsIDParams) WithTimeout(timeout time.Duration) *DeletePaymentdefaultsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete paymentdefaults ID params
func (o *DeletePaymentdefaultsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete paymentdefaults ID params
func (o *DeletePaymentdefaultsIDParams) WithContext(ctx context.Context) *DeletePaymentdefaultsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete paymentdefaults ID params
func (o *DeletePaymentdefaultsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete paymentdefaults ID params
func (o *DeletePaymentdefaultsIDParams) WithHTTPClient(client *http.Client) *DeletePaymentdefaultsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete paymentdefaults ID params
func (o *DeletePaymentdefaultsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete paymentdefaults ID params
func (o *DeletePaymentdefaultsIDParams) WithID(id strfmt.UUID) *DeletePaymentdefaultsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete paymentdefaults ID params
func (o *DeletePaymentdefaultsIDParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithVersion adds the version to the delete paymentdefaults ID params
func (o *DeletePaymentdefaultsIDParams) WithVersion(version int64) *DeletePaymentdefaultsIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete paymentdefaults ID params
func (o *DeletePaymentdefaultsIDParams) SetVersion(version int64) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePaymentdefaultsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := swag.FormatInt64(qrVersion)
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}