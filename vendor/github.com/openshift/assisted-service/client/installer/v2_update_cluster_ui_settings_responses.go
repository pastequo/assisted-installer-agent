// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// V2UpdateClusterUISettingsReader is a Reader for the V2UpdateClusterUISettings structure.
type V2UpdateClusterUISettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *V2UpdateClusterUISettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewV2UpdateClusterUISettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewV2UpdateClusterUISettingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewV2UpdateClusterUISettingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewV2UpdateClusterUISettingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewV2UpdateClusterUISettingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewV2UpdateClusterUISettingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewV2UpdateClusterUISettingsOK creates a V2UpdateClusterUISettingsOK with default headers values
func NewV2UpdateClusterUISettingsOK() *V2UpdateClusterUISettingsOK {
	return &V2UpdateClusterUISettingsOK{}
}

/*
V2UpdateClusterUISettingsOK describes a response with status code 200, with default header values.

Success.
*/
type V2UpdateClusterUISettingsOK struct {
	Payload string
}

// IsSuccess returns true when this v2 update cluster Ui settings o k response has a 2xx status code
func (o *V2UpdateClusterUISettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this v2 update cluster Ui settings o k response has a 3xx status code
func (o *V2UpdateClusterUISettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 update cluster Ui settings o k response has a 4xx status code
func (o *V2UpdateClusterUISettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 update cluster Ui settings o k response has a 5xx status code
func (o *V2UpdateClusterUISettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 update cluster Ui settings o k response a status code equal to that given
func (o *V2UpdateClusterUISettingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *V2UpdateClusterUISettingsOK) Error() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/ui-settings][%d] v2UpdateClusterUiSettingsOK  %+v", 200, o.Payload)
}

func (o *V2UpdateClusterUISettingsOK) String() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/ui-settings][%d] v2UpdateClusterUiSettingsOK  %+v", 200, o.Payload)
}

func (o *V2UpdateClusterUISettingsOK) GetPayload() string {
	return o.Payload
}

func (o *V2UpdateClusterUISettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateClusterUISettingsBadRequest creates a V2UpdateClusterUISettingsBadRequest with default headers values
func NewV2UpdateClusterUISettingsBadRequest() *V2UpdateClusterUISettingsBadRequest {
	return &V2UpdateClusterUISettingsBadRequest{}
}

/*
V2UpdateClusterUISettingsBadRequest describes a response with status code 400, with default header values.

Error.
*/
type V2UpdateClusterUISettingsBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 update cluster Ui settings bad request response has a 2xx status code
func (o *V2UpdateClusterUISettingsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 update cluster Ui settings bad request response has a 3xx status code
func (o *V2UpdateClusterUISettingsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 update cluster Ui settings bad request response has a 4xx status code
func (o *V2UpdateClusterUISettingsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 update cluster Ui settings bad request response has a 5xx status code
func (o *V2UpdateClusterUISettingsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 update cluster Ui settings bad request response a status code equal to that given
func (o *V2UpdateClusterUISettingsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *V2UpdateClusterUISettingsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/ui-settings][%d] v2UpdateClusterUiSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *V2UpdateClusterUISettingsBadRequest) String() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/ui-settings][%d] v2UpdateClusterUiSettingsBadRequest  %+v", 400, o.Payload)
}

func (o *V2UpdateClusterUISettingsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UpdateClusterUISettingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateClusterUISettingsUnauthorized creates a V2UpdateClusterUISettingsUnauthorized with default headers values
func NewV2UpdateClusterUISettingsUnauthorized() *V2UpdateClusterUISettingsUnauthorized {
	return &V2UpdateClusterUISettingsUnauthorized{}
}

/*
V2UpdateClusterUISettingsUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type V2UpdateClusterUISettingsUnauthorized struct {
	Payload *models.InfraError
}

// IsSuccess returns true when this v2 update cluster Ui settings unauthorized response has a 2xx status code
func (o *V2UpdateClusterUISettingsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 update cluster Ui settings unauthorized response has a 3xx status code
func (o *V2UpdateClusterUISettingsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 update cluster Ui settings unauthorized response has a 4xx status code
func (o *V2UpdateClusterUISettingsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 update cluster Ui settings unauthorized response has a 5xx status code
func (o *V2UpdateClusterUISettingsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 update cluster Ui settings unauthorized response a status code equal to that given
func (o *V2UpdateClusterUISettingsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *V2UpdateClusterUISettingsUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/ui-settings][%d] v2UpdateClusterUiSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *V2UpdateClusterUISettingsUnauthorized) String() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/ui-settings][%d] v2UpdateClusterUiSettingsUnauthorized  %+v", 401, o.Payload)
}

func (o *V2UpdateClusterUISettingsUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2UpdateClusterUISettingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateClusterUISettingsForbidden creates a V2UpdateClusterUISettingsForbidden with default headers values
func NewV2UpdateClusterUISettingsForbidden() *V2UpdateClusterUISettingsForbidden {
	return &V2UpdateClusterUISettingsForbidden{}
}

/*
V2UpdateClusterUISettingsForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type V2UpdateClusterUISettingsForbidden struct {
	Payload *models.InfraError
}

// IsSuccess returns true when this v2 update cluster Ui settings forbidden response has a 2xx status code
func (o *V2UpdateClusterUISettingsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 update cluster Ui settings forbidden response has a 3xx status code
func (o *V2UpdateClusterUISettingsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 update cluster Ui settings forbidden response has a 4xx status code
func (o *V2UpdateClusterUISettingsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 update cluster Ui settings forbidden response has a 5xx status code
func (o *V2UpdateClusterUISettingsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 update cluster Ui settings forbidden response a status code equal to that given
func (o *V2UpdateClusterUISettingsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *V2UpdateClusterUISettingsForbidden) Error() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/ui-settings][%d] v2UpdateClusterUiSettingsForbidden  %+v", 403, o.Payload)
}

func (o *V2UpdateClusterUISettingsForbidden) String() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/ui-settings][%d] v2UpdateClusterUiSettingsForbidden  %+v", 403, o.Payload)
}

func (o *V2UpdateClusterUISettingsForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *V2UpdateClusterUISettingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateClusterUISettingsNotFound creates a V2UpdateClusterUISettingsNotFound with default headers values
func NewV2UpdateClusterUISettingsNotFound() *V2UpdateClusterUISettingsNotFound {
	return &V2UpdateClusterUISettingsNotFound{}
}

/*
V2UpdateClusterUISettingsNotFound describes a response with status code 404, with default header values.

Error.
*/
type V2UpdateClusterUISettingsNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 update cluster Ui settings not found response has a 2xx status code
func (o *V2UpdateClusterUISettingsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 update cluster Ui settings not found response has a 3xx status code
func (o *V2UpdateClusterUISettingsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 update cluster Ui settings not found response has a 4xx status code
func (o *V2UpdateClusterUISettingsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this v2 update cluster Ui settings not found response has a 5xx status code
func (o *V2UpdateClusterUISettingsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this v2 update cluster Ui settings not found response a status code equal to that given
func (o *V2UpdateClusterUISettingsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *V2UpdateClusterUISettingsNotFound) Error() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/ui-settings][%d] v2UpdateClusterUiSettingsNotFound  %+v", 404, o.Payload)
}

func (o *V2UpdateClusterUISettingsNotFound) String() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/ui-settings][%d] v2UpdateClusterUiSettingsNotFound  %+v", 404, o.Payload)
}

func (o *V2UpdateClusterUISettingsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UpdateClusterUISettingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewV2UpdateClusterUISettingsInternalServerError creates a V2UpdateClusterUISettingsInternalServerError with default headers values
func NewV2UpdateClusterUISettingsInternalServerError() *V2UpdateClusterUISettingsInternalServerError {
	return &V2UpdateClusterUISettingsInternalServerError{}
}

/*
V2UpdateClusterUISettingsInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type V2UpdateClusterUISettingsInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this v2 update cluster Ui settings internal server error response has a 2xx status code
func (o *V2UpdateClusterUISettingsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this v2 update cluster Ui settings internal server error response has a 3xx status code
func (o *V2UpdateClusterUISettingsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this v2 update cluster Ui settings internal server error response has a 4xx status code
func (o *V2UpdateClusterUISettingsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this v2 update cluster Ui settings internal server error response has a 5xx status code
func (o *V2UpdateClusterUISettingsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this v2 update cluster Ui settings internal server error response a status code equal to that given
func (o *V2UpdateClusterUISettingsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *V2UpdateClusterUISettingsInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/ui-settings][%d] v2UpdateClusterUiSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *V2UpdateClusterUISettingsInternalServerError) String() string {
	return fmt.Sprintf("[PUT /v2/clusters/{cluster_id}/ui-settings][%d] v2UpdateClusterUiSettingsInternalServerError  %+v", 500, o.Payload)
}

func (o *V2UpdateClusterUISettingsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *V2UpdateClusterUISettingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
