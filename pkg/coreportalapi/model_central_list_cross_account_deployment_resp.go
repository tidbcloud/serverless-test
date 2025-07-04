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

// checks if the CentralListCrossAccountDeploymentResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CentralListCrossAccountDeploymentResp{}

// CentralListCrossAccountDeploymentResp struct for CentralListCrossAccountDeploymentResp
type CentralListCrossAccountDeploymentResp struct {
	Items    []CentralCrossAccountDeployment `json:"items"`
	Total    int32                           `json:"total"`
	BaseResp CentralBaseResp                 `json:"base_resp"`
}

type _CentralListCrossAccountDeploymentResp CentralListCrossAccountDeploymentResp

// NewCentralListCrossAccountDeploymentResp instantiates a new CentralListCrossAccountDeploymentResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCentralListCrossAccountDeploymentResp(items []CentralCrossAccountDeployment, total int32, baseResp CentralBaseResp) *CentralListCrossAccountDeploymentResp {
	this := CentralListCrossAccountDeploymentResp{}
	this.Items = items
	this.Total = total
	this.BaseResp = baseResp
	return &this
}

// NewCentralListCrossAccountDeploymentRespWithDefaults instantiates a new CentralListCrossAccountDeploymentResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCentralListCrossAccountDeploymentRespWithDefaults() *CentralListCrossAccountDeploymentResp {
	this := CentralListCrossAccountDeploymentResp{}
	return &this
}

// GetItems returns the Items field value
func (o *CentralListCrossAccountDeploymentResp) GetItems() []CentralCrossAccountDeployment {
	if o == nil {
		var ret []CentralCrossAccountDeployment
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CentralListCrossAccountDeploymentResp) GetItemsOk() ([]CentralCrossAccountDeployment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *CentralListCrossAccountDeploymentResp) SetItems(v []CentralCrossAccountDeployment) {
	o.Items = v
}

// GetTotal returns the Total field value
func (o *CentralListCrossAccountDeploymentResp) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *CentralListCrossAccountDeploymentResp) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *CentralListCrossAccountDeploymentResp) SetTotal(v int32) {
	o.Total = v
}

// GetBaseResp returns the BaseResp field value
func (o *CentralListCrossAccountDeploymentResp) GetBaseResp() CentralBaseResp {
	if o == nil {
		var ret CentralBaseResp
		return ret
	}

	return o.BaseResp
}

// GetBaseRespOk returns a tuple with the BaseResp field value
// and a boolean to check if the value has been set.
func (o *CentralListCrossAccountDeploymentResp) GetBaseRespOk() (*CentralBaseResp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseResp, true
}

// SetBaseResp sets field value
func (o *CentralListCrossAccountDeploymentResp) SetBaseResp(v CentralBaseResp) {
	o.BaseResp = v
}

func (o CentralListCrossAccountDeploymentResp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CentralListCrossAccountDeploymentResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["total"] = o.Total
	toSerialize["base_resp"] = o.BaseResp
	return toSerialize, nil
}

func (o *CentralListCrossAccountDeploymentResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"items",
		"total",
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

	varCentralListCrossAccountDeploymentResp := _CentralListCrossAccountDeploymentResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCentralListCrossAccountDeploymentResp)

	if err != nil {
		return err
	}

	*o = CentralListCrossAccountDeploymentResp(varCentralListCrossAccountDeploymentResp)

	return err
}

type NullableCentralListCrossAccountDeploymentResp struct {
	value *CentralListCrossAccountDeploymentResp
	isSet bool
}

func (v NullableCentralListCrossAccountDeploymentResp) Get() *CentralListCrossAccountDeploymentResp {
	return v.value
}

func (v *NullableCentralListCrossAccountDeploymentResp) Set(val *CentralListCrossAccountDeploymentResp) {
	v.value = val
	v.isSet = true
}

func (v NullableCentralListCrossAccountDeploymentResp) IsSet() bool {
	return v.isSet
}

func (v *NullableCentralListCrossAccountDeploymentResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCentralListCrossAccountDeploymentResp(val *CentralListCrossAccountDeploymentResp) *NullableCentralListCrossAccountDeploymentResp {
	return &NullableCentralListCrossAccountDeploymentResp{value: val, isSet: true}
}

func (v NullableCentralListCrossAccountDeploymentResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCentralListCrossAccountDeploymentResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
