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

// checks if the CentralRegionCIDR type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CentralRegionCIDR{}

// CentralRegionCIDR struct for CentralRegionCIDR
type CentralRegionCIDR struct {
	Region     string `json:"region"`
	RegionName string `json:"region_name"`
	Cidr       string `json:"cidr"`
	Active     bool   `json:"active"`
}

type _CentralRegionCIDR CentralRegionCIDR

// NewCentralRegionCIDR instantiates a new CentralRegionCIDR object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCentralRegionCIDR(region string, regionName string, cidr string, active bool) *CentralRegionCIDR {
	this := CentralRegionCIDR{}
	this.Region = region
	this.RegionName = regionName
	this.Cidr = cidr
	this.Active = active
	return &this
}

// NewCentralRegionCIDRWithDefaults instantiates a new CentralRegionCIDR object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCentralRegionCIDRWithDefaults() *CentralRegionCIDR {
	this := CentralRegionCIDR{}
	return &this
}

// GetRegion returns the Region field value
func (o *CentralRegionCIDR) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *CentralRegionCIDR) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *CentralRegionCIDR) SetRegion(v string) {
	o.Region = v
}

// GetRegionName returns the RegionName field value
func (o *CentralRegionCIDR) GetRegionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value
// and a boolean to check if the value has been set.
func (o *CentralRegionCIDR) GetRegionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionName, true
}

// SetRegionName sets field value
func (o *CentralRegionCIDR) SetRegionName(v string) {
	o.RegionName = v
}

// GetCidr returns the Cidr field value
func (o *CentralRegionCIDR) GetCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value
// and a boolean to check if the value has been set.
func (o *CentralRegionCIDR) GetCidrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cidr, true
}

// SetCidr sets field value
func (o *CentralRegionCIDR) SetCidr(v string) {
	o.Cidr = v
}

// GetActive returns the Active field value
func (o *CentralRegionCIDR) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *CentralRegionCIDR) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *CentralRegionCIDR) SetActive(v bool) {
	o.Active = v
}

func (o CentralRegionCIDR) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CentralRegionCIDR) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["region"] = o.Region
	toSerialize["region_name"] = o.RegionName
	toSerialize["cidr"] = o.Cidr
	toSerialize["active"] = o.Active
	return toSerialize, nil
}

func (o *CentralRegionCIDR) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"region",
		"region_name",
		"cidr",
		"active",
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

	varCentralRegionCIDR := _CentralRegionCIDR{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCentralRegionCIDR)

	if err != nil {
		return err
	}

	*o = CentralRegionCIDR(varCentralRegionCIDR)

	return err
}

type NullableCentralRegionCIDR struct {
	value *CentralRegionCIDR
	isSet bool
}

func (v NullableCentralRegionCIDR) Get() *CentralRegionCIDR {
	return v.value
}

func (v *NullableCentralRegionCIDR) Set(val *CentralRegionCIDR) {
	v.value = val
	v.isSet = true
}

func (v NullableCentralRegionCIDR) IsSet() bool {
	return v.isSet
}

func (v *NullableCentralRegionCIDR) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCentralRegionCIDR(val *CentralRegionCIDR) *NullableCentralRegionCIDR {
	return &NullableCentralRegionCIDR{value: val, isSet: true}
}

func (v NullableCentralRegionCIDR) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCentralRegionCIDR) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
