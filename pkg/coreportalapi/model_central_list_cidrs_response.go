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

// checks if the CentralListCidrsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CentralListCidrsResponse{}

// CentralListCidrsResponse struct for CentralListCidrsResponse
type CentralListCidrsResponse struct {
	// A list of CIDRs.
	Items    []CentralCidr   `json:"items"`
	Total    int32           `json:"total"`
	BaseResp CentralBaseResp `json:"base_resp"`
}

type _CentralListCidrsResponse CentralListCidrsResponse

// NewCentralListCidrsResponse instantiates a new CentralListCidrsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCentralListCidrsResponse(items []CentralCidr, total int32, baseResp CentralBaseResp) *CentralListCidrsResponse {
	this := CentralListCidrsResponse{}
	this.Items = items
	this.Total = total
	this.BaseResp = baseResp
	return &this
}

// NewCentralListCidrsResponseWithDefaults instantiates a new CentralListCidrsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCentralListCidrsResponseWithDefaults() *CentralListCidrsResponse {
	this := CentralListCidrsResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *CentralListCidrsResponse) GetItems() []CentralCidr {
	if o == nil {
		var ret []CentralCidr
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CentralListCidrsResponse) GetItemsOk() ([]CentralCidr, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *CentralListCidrsResponse) SetItems(v []CentralCidr) {
	o.Items = v
}

// GetTotal returns the Total field value
func (o *CentralListCidrsResponse) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *CentralListCidrsResponse) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *CentralListCidrsResponse) SetTotal(v int32) {
	o.Total = v
}

// GetBaseResp returns the BaseResp field value
func (o *CentralListCidrsResponse) GetBaseResp() CentralBaseResp {
	if o == nil {
		var ret CentralBaseResp
		return ret
	}

	return o.BaseResp
}

// GetBaseRespOk returns a tuple with the BaseResp field value
// and a boolean to check if the value has been set.
func (o *CentralListCidrsResponse) GetBaseRespOk() (*CentralBaseResp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseResp, true
}

// SetBaseResp sets field value
func (o *CentralListCidrsResponse) SetBaseResp(v CentralBaseResp) {
	o.BaseResp = v
}

func (o CentralListCidrsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CentralListCidrsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["total"] = o.Total
	toSerialize["base_resp"] = o.BaseResp
	return toSerialize, nil
}

func (o *CentralListCidrsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCentralListCidrsResponse := _CentralListCidrsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCentralListCidrsResponse)

	if err != nil {
		return err
	}

	*o = CentralListCidrsResponse(varCentralListCidrsResponse)

	return err
}

type NullableCentralListCidrsResponse struct {
	value *CentralListCidrsResponse
	isSet bool
}

func (v NullableCentralListCidrsResponse) Get() *CentralListCidrsResponse {
	return v.value
}

func (v *NullableCentralListCidrsResponse) Set(val *CentralListCidrsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCentralListCidrsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCentralListCidrsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCentralListCidrsResponse(val *CentralListCidrsResponse) *NullableCentralListCidrsResponse {
	return &NullableCentralListCidrsResponse{value: val, isSet: true}
}

func (v NullableCentralListCidrsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCentralListCidrsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
