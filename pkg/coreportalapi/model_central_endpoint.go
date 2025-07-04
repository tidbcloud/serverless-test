/*
gRPC for our DBaaS central service

notably uses gRPC-Gateway with OpenAPI

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package coreportalapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CentralEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CentralEndpoint{}

// CentralEndpoint struct for CentralEndpoint
type CentralEndpoint struct {
	Ready   bool    `json:"ready"`
	Address string  `json:"address"`
	Port    int32   `json:"port"`
	Token   *string `json:"token,omitempty"`
}

type _CentralEndpoint CentralEndpoint

// NewCentralEndpoint instantiates a new CentralEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCentralEndpoint(ready bool, address string, port int32) *CentralEndpoint {
	this := CentralEndpoint{}
	this.Ready = ready
	this.Address = address
	this.Port = port
	return &this
}

// NewCentralEndpointWithDefaults instantiates a new CentralEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCentralEndpointWithDefaults() *CentralEndpoint {
	this := CentralEndpoint{}
	return &this
}

// GetReady returns the Ready field value
func (o *CentralEndpoint) GetReady() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ready
}

// GetReadyOk returns a tuple with the Ready field value
// and a boolean to check if the value has been set.
func (o *CentralEndpoint) GetReadyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ready, true
}

// SetReady sets field value
func (o *CentralEndpoint) SetReady(v bool) {
	o.Ready = v
}

// GetAddress returns the Address field value
func (o *CentralEndpoint) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *CentralEndpoint) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *CentralEndpoint) SetAddress(v string) {
	o.Address = v
}

// GetPort returns the Port field value
func (o *CentralEndpoint) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *CentralEndpoint) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *CentralEndpoint) SetPort(v int32) {
	o.Port = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CentralEndpoint) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CentralEndpoint) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CentralEndpoint) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CentralEndpoint) SetToken(v string) {
	o.Token = &v
}

func (o CentralEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CentralEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ready"] = o.Ready
	toSerialize["address"] = o.Address
	toSerialize["port"] = o.Port
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

func (o *CentralEndpoint) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ready",
		"address",
		"port",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCentralEndpoint := _CentralEndpoint{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCentralEndpoint)

	if err != nil {
		return err
	}

	*o = CentralEndpoint(varCentralEndpoint)

	return err
}

type NullableCentralEndpoint struct {
	value *CentralEndpoint
	isSet bool
}

func (v NullableCentralEndpoint) Get() *CentralEndpoint {
	return v.value
}

func (v *NullableCentralEndpoint) Set(val *CentralEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableCentralEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableCentralEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCentralEndpoint(val *CentralEndpoint) *NullableCentralEndpoint {
	return &NullableCentralEndpoint{value: val, isSet: true}
}

func (v NullableCentralEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCentralEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
