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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/form3tech/go-form3-web/httpclient/support/swagger-client/paymentapi/models"
)

// NewPostPaymentsIDReturnsParams creates a new PostPaymentsIDReturnsParams object
// with the default values initialized.
func NewPostPaymentsIDReturnsParams() *PostPaymentsIDReturnsParams {
	var ()
	return &PostPaymentsIDReturnsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPaymentsIDReturnsParamsWithTimeout creates a new PostPaymentsIDReturnsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPaymentsIDReturnsParamsWithTimeout(timeout time.Duration) *PostPaymentsIDReturnsParams {
	var ()
	return &PostPaymentsIDReturnsParams{

		timeout: timeout,
	}
}

// NewPostPaymentsIDReturnsParamsWithContext creates a new PostPaymentsIDReturnsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPaymentsIDReturnsParamsWithContext(ctx context.Context) *PostPaymentsIDReturnsParams {
	var ()
	return &PostPaymentsIDReturnsParams{

		Context: ctx,
	}
}

// NewPostPaymentsIDReturnsParamsWithHTTPClient creates a new PostPaymentsIDReturnsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPaymentsIDReturnsParamsWithHTTPClient(client *http.Client) *PostPaymentsIDReturnsParams {
	var ()
	return &PostPaymentsIDReturnsParams{
		HTTPClient: client,
	}
}

/*PostPaymentsIDReturnsParams contains all the parameters to send to the API endpoint
for the post payments ID returns operation typically these are written to a http.Request
*/
type PostPaymentsIDReturnsParams struct {

	/*ReturnCreationRequest*/
	ReturnCreationRequest *models.ReturnCreation
	/*ID
	  Payment Id

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post payments ID returns params
func (o *PostPaymentsIDReturnsParams) WithTimeout(timeout time.Duration) *PostPaymentsIDReturnsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post payments ID returns params
func (o *PostPaymentsIDReturnsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post payments ID returns params
func (o *PostPaymentsIDReturnsParams) WithContext(ctx context.Context) *PostPaymentsIDReturnsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post payments ID returns params
func (o *PostPaymentsIDReturnsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post payments ID returns params
func (o *PostPaymentsIDReturnsParams) WithHTTPClient(client *http.Client) *PostPaymentsIDReturnsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post payments ID returns params
func (o *PostPaymentsIDReturnsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReturnCreationRequest adds the returnCreationRequest to the post payments ID returns params
func (o *PostPaymentsIDReturnsParams) WithReturnCreationRequest(returnCreationRequest *models.ReturnCreation) *PostPaymentsIDReturnsParams {
	o.SetReturnCreationRequest(returnCreationRequest)
	return o
}

// SetReturnCreationRequest adds the returnCreationRequest to the post payments ID returns params
func (o *PostPaymentsIDReturnsParams) SetReturnCreationRequest(returnCreationRequest *models.ReturnCreation) {
	o.ReturnCreationRequest = returnCreationRequest
}

// WithID adds the id to the post payments ID returns params
func (o *PostPaymentsIDReturnsParams) WithID(id strfmt.UUID) *PostPaymentsIDReturnsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post payments ID returns params
func (o *PostPaymentsIDReturnsParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostPaymentsIDReturnsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ReturnCreationRequest != nil {
		if err := r.SetBodyParam(o.ReturnCreationRequest); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}