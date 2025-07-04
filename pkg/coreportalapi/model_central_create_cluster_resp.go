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

// checks if the CentralCreateClusterResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CentralCreateClusterResp{}

// CentralCreateClusterResp struct for CentralCreateClusterResp
type CentralCreateClusterResp struct {
	Id       string          `json:"id"`
	BaseResp CentralBaseResp `json:"base_resp"`
}

type _CentralCreateClusterResp CentralCreateClusterResp

// NewCentralCreateClusterResp instantiates a new CentralCreateClusterResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCentralCreateClusterResp(id string, baseResp CentralBaseResp) *CentralCreateClusterResp {
	this := CentralCreateClusterResp{}
	this.Id = id
	this.BaseResp = baseResp
	return &this
}

// NewCentralCreateClusterRespWithDefaults instantiates a new CentralCreateClusterResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCentralCreateClusterRespWithDefaults() *CentralCreateClusterResp {
	this := CentralCreateClusterResp{}
	return &this
}

// GetId returns the Id field value
func (o *CentralCreateClusterResp) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CentralCreateClusterResp) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CentralCreateClusterResp) SetId(v string) {
	o.Id = v
}

// GetBaseResp returns the BaseResp field value
func (o *CentralCreateClusterResp) GetBaseResp() CentralBaseResp {
	if o == nil {
		var ret CentralBaseResp
		return ret
	}

	return o.BaseResp
}

// GetBaseRespOk returns a tuple with the BaseResp field value
// and a boolean to check if the value has been set.
func (o *CentralCreateClusterResp) GetBaseRespOk() (*CentralBaseResp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseResp, true
}

// SetBaseResp sets field value
func (o *CentralCreateClusterResp) SetBaseResp(v CentralBaseResp) {
	o.BaseResp = v
}

func (o CentralCreateClusterResp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CentralCreateClusterResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["base_resp"] = o.BaseResp
	return toSerialize, nil
}

func (o *CentralCreateClusterResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"base_resp",
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

	varCentralCreateClusterResp := _CentralCreateClusterResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCentralCreateClusterResp)

	if err != nil {
		return err
	}

	*o = CentralCreateClusterResp(varCentralCreateClusterResp)

	return err
}

type NullableCentralCreateClusterResp struct {
	value *CentralCreateClusterResp
	isSet bool
}

func (v NullableCentralCreateClusterResp) Get() *CentralCreateClusterResp {
	return v.value
}

func (v *NullableCentralCreateClusterResp) Set(val *CentralCreateClusterResp) {
	v.value = val
	v.isSet = true
}

func (v NullableCentralCreateClusterResp) IsSet() bool {
	return v.isSet
}

func (v *NullableCentralCreateClusterResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCentralCreateClusterResp(val *CentralCreateClusterResp) *NullableCentralCreateClusterResp {
	return &NullableCentralCreateClusterResp{value: val, isSet: true}
}

func (v NullableCentralCreateClusterResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCentralCreateClusterResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
