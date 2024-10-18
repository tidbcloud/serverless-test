/*
TiDB Cloud Serverless Open API

TiDB Cloud Serverless Open API

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cluster

import (
	"encoding/json"
)

// checks if the GooglerpcStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GooglerpcStatus{}

// GooglerpcStatus The `Status` type defines a logical error model that is suitable for different programming environments, including REST APIs and RPC APIs. It is used by [gRPC](https://github.com/grpc). Each `Status` message contains three pieces of data: error code, error message, and error details.  You can find out more about this error model and how to work with it in the [API Design Guide](https://cloud.google.com/apis/design/errors).
type GooglerpcStatus struct {
	// The status code, which should be an enum value of [google.rpc.Code][google.rpc.Code].
	Code *int32 `json:"code,omitempty"`
	// A developer-facing error message, which should be in English. Any user-facing error message should be localized and sent in the [google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client.
	Message *string `json:"message,omitempty"`
	// A list of messages that carry the error details.  There is a common set of message types for APIs to use.
	Details              []ProtobufAny `json:"details,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GooglerpcStatus GooglerpcStatus

// NewGooglerpcStatus instantiates a new GooglerpcStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGooglerpcStatus() *GooglerpcStatus {
	this := GooglerpcStatus{}
	return &this
}

// NewGooglerpcStatusWithDefaults instantiates a new GooglerpcStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGooglerpcStatusWithDefaults() *GooglerpcStatus {
	this := GooglerpcStatus{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GooglerpcStatus) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GooglerpcStatus) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GooglerpcStatus) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *GooglerpcStatus) SetCode(v int32) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GooglerpcStatus) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GooglerpcStatus) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GooglerpcStatus) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GooglerpcStatus) SetMessage(v string) {
	o.Message = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *GooglerpcStatus) GetDetails() []ProtobufAny {
	if o == nil || IsNil(o.Details) {
		var ret []ProtobufAny
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GooglerpcStatus) GetDetailsOk() ([]ProtobufAny, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *GooglerpcStatus) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []ProtobufAny and assigns it to the Details field.
func (o *GooglerpcStatus) SetDetails(v []ProtobufAny) {
	o.Details = v
}

func (o GooglerpcStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GooglerpcStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GooglerpcStatus) UnmarshalJSON(data []byte) (err error) {
	varGooglerpcStatus := _GooglerpcStatus{}

	err = json.Unmarshal(data, &varGooglerpcStatus)

	if err != nil {
		return err
	}

	*o = GooglerpcStatus(varGooglerpcStatus)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "message")
		delete(additionalProperties, "details")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGooglerpcStatus struct {
	value *GooglerpcStatus
	isSet bool
}

func (v NullableGooglerpcStatus) Get() *GooglerpcStatus {
	return v.value
}

func (v *NullableGooglerpcStatus) Set(val *GooglerpcStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableGooglerpcStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableGooglerpcStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGooglerpcStatus(val *GooglerpcStatus) *NullableGooglerpcStatus {
	return &NullableGooglerpcStatus{value: val, isSet: true}
}

func (v NullableGooglerpcStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGooglerpcStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
