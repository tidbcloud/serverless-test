/*
gRPC for our DBaaS central service

notably uses gRPC-Gateway with OpenAPI

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package coreportalapi

import (
	"encoding/json"
)

// checks if the CentralGetAwsCmekResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CentralGetAwsCmekResp{}

// CentralGetAwsCmekResp struct for CentralGetAwsCmekResp
type CentralGetAwsCmekResp struct {
	Cmeks    []CentralAwsCmek `json:"cmeks,omitempty"`
	BaseResp *CentralBaseResp `json:"base_resp,omitempty"`
}

// NewCentralGetAwsCmekResp instantiates a new CentralGetAwsCmekResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCentralGetAwsCmekResp() *CentralGetAwsCmekResp {
	this := CentralGetAwsCmekResp{}
	return &this
}

// NewCentralGetAwsCmekRespWithDefaults instantiates a new CentralGetAwsCmekResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCentralGetAwsCmekRespWithDefaults() *CentralGetAwsCmekResp {
	this := CentralGetAwsCmekResp{}
	return &this
}

// GetCmeks returns the Cmeks field value if set, zero value otherwise.
func (o *CentralGetAwsCmekResp) GetCmeks() []CentralAwsCmek {
	if o == nil || IsNil(o.Cmeks) {
		var ret []CentralAwsCmek
		return ret
	}
	return o.Cmeks
}

// GetCmeksOk returns a tuple with the Cmeks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CentralGetAwsCmekResp) GetCmeksOk() ([]CentralAwsCmek, bool) {
	if o == nil || IsNil(o.Cmeks) {
		return nil, false
	}
	return o.Cmeks, true
}

// HasCmeks returns a boolean if a field has been set.
func (o *CentralGetAwsCmekResp) HasCmeks() bool {
	if o != nil && !IsNil(o.Cmeks) {
		return true
	}

	return false
}

// SetCmeks gets a reference to the given []CentralAwsCmek and assigns it to the Cmeks field.
func (o *CentralGetAwsCmekResp) SetCmeks(v []CentralAwsCmek) {
	o.Cmeks = v
}

// GetBaseResp returns the BaseResp field value if set, zero value otherwise.
func (o *CentralGetAwsCmekResp) GetBaseResp() CentralBaseResp {
	if o == nil || IsNil(o.BaseResp) {
		var ret CentralBaseResp
		return ret
	}
	return *o.BaseResp
}

// GetBaseRespOk returns a tuple with the BaseResp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CentralGetAwsCmekResp) GetBaseRespOk() (*CentralBaseResp, bool) {
	if o == nil || IsNil(o.BaseResp) {
		return nil, false
	}
	return o.BaseResp, true
}

// HasBaseResp returns a boolean if a field has been set.
func (o *CentralGetAwsCmekResp) HasBaseResp() bool {
	if o != nil && !IsNil(o.BaseResp) {
		return true
	}

	return false
}

// SetBaseResp gets a reference to the given CentralBaseResp and assigns it to the BaseResp field.
func (o *CentralGetAwsCmekResp) SetBaseResp(v CentralBaseResp) {
	o.BaseResp = &v
}

func (o CentralGetAwsCmekResp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CentralGetAwsCmekResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cmeks) {
		toSerialize["cmeks"] = o.Cmeks
	}
	if !IsNil(o.BaseResp) {
		toSerialize["base_resp"] = o.BaseResp
	}
	return toSerialize, nil
}

type NullableCentralGetAwsCmekResp struct {
	value *CentralGetAwsCmekResp
	isSet bool
}

func (v NullableCentralGetAwsCmekResp) Get() *CentralGetAwsCmekResp {
	return v.value
}

func (v *NullableCentralGetAwsCmekResp) Set(val *CentralGetAwsCmekResp) {
	v.value = val
	v.isSet = true
}

func (v NullableCentralGetAwsCmekResp) IsSet() bool {
	return v.isSet
}

func (v *NullableCentralGetAwsCmekResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCentralGetAwsCmekResp(val *CentralGetAwsCmekResp) *NullableCentralGetAwsCmekResp {
	return &NullableCentralGetAwsCmekResp{value: val, isSet: true}
}

func (v NullableCentralGetAwsCmekResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCentralGetAwsCmekResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
