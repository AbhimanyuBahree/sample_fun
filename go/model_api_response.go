/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ApiResponse - Describes the result of uploading an image resource
type ApiResponse struct {

	Code int32 `json:"code,omitempty"`

	Type string `json:"type,omitempty"`

	Message string `json:"message,omitempty"`
}

// AssertApiResponseRequired checks if the required fields are not zero-ed
func AssertApiResponseRequired(obj ApiResponse) error {
	return nil
}

// AssertRecurseApiResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ApiResponse (e.g. [][]ApiResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseApiResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aApiResponse, ok := obj.(ApiResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertApiResponseRequired(aApiResponse)
	})
}
