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

// checks if the CentralTiFlashNodes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CentralTiFlashNodes{}

// CentralTiFlashNodes struct for CentralTiFlashNodes
type CentralTiFlashNodes struct {
	Total int32                `json:"total"`
	Items []CentralTiFlashNode `json:"items"`
}

type _CentralTiFlashNodes CentralTiFlashNodes

// NewCentralTiFlashNodes instantiates a new CentralTiFlashNodes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCentralTiFlashNodes(total int32, items []CentralTiFlashNode) *CentralTiFlashNodes {
	this := CentralTiFlashNodes{}
	this.Total = total
	this.Items = items
	return &this
}

// NewCentralTiFlashNodesWithDefaults instantiates a new CentralTiFlashNodes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCentralTiFlashNodesWithDefaults() *CentralTiFlashNodes {
	this := CentralTiFlashNodes{}
	return &this
}

// GetTotal returns the Total field value
func (o *CentralTiFlashNodes) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *CentralTiFlashNodes) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *CentralTiFlashNodes) SetTotal(v int32) {
	o.Total = v
}

// GetItems returns the Items field value
func (o *CentralTiFlashNodes) GetItems() []CentralTiFlashNode {
	if o == nil {
		var ret []CentralTiFlashNode
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CentralTiFlashNodes) GetItemsOk() ([]CentralTiFlashNode, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *CentralTiFlashNodes) SetItems(v []CentralTiFlashNode) {
	o.Items = v
}

func (o CentralTiFlashNodes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CentralTiFlashNodes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total"] = o.Total
	toSerialize["items"] = o.Items
	return toSerialize, nil
}

func (o *CentralTiFlashNodes) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"total",
		"items",
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

	varCentralTiFlashNodes := _CentralTiFlashNodes{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCentralTiFlashNodes)

	if err != nil {
		return err
	}

	*o = CentralTiFlashNodes(varCentralTiFlashNodes)

	return err
}

type NullableCentralTiFlashNodes struct {
	value *CentralTiFlashNodes
	isSet bool
}

func (v NullableCentralTiFlashNodes) Get() *CentralTiFlashNodes {
	return v.value
}

func (v *NullableCentralTiFlashNodes) Set(val *CentralTiFlashNodes) {
	v.value = val
	v.isSet = true
}

func (v NullableCentralTiFlashNodes) IsSet() bool {
	return v.isSet
}

func (v *NullableCentralTiFlashNodes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCentralTiFlashNodes(val *CentralTiFlashNodes) *NullableCentralTiFlashNodes {
	return &NullableCentralTiFlashNodes{value: val, isSet: true}
}

func (v NullableCentralTiFlashNodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCentralTiFlashNodes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
