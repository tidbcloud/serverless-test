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

// checks if the CalcNodeProfilesV2Request type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CalcNodeProfilesV2Request{}

// CalcNodeProfilesV2Request struct for CalcNodeProfilesV2Request
type CalcNodeProfilesV2Request struct {
	Provider  string `json:"provider"`
	Region    string `json:"region"`
	IsDevTier bool   `json:"is_dev_tier"`
}

type _CalcNodeProfilesV2Request CalcNodeProfilesV2Request

// NewCalcNodeProfilesV2Request instantiates a new CalcNodeProfilesV2Request object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCalcNodeProfilesV2Request(provider string, region string, isDevTier bool) *CalcNodeProfilesV2Request {
	this := CalcNodeProfilesV2Request{}
	this.Provider = provider
	this.Region = region
	this.IsDevTier = isDevTier
	return &this
}

// NewCalcNodeProfilesV2RequestWithDefaults instantiates a new CalcNodeProfilesV2Request object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCalcNodeProfilesV2RequestWithDefaults() *CalcNodeProfilesV2Request {
	this := CalcNodeProfilesV2Request{}
	return &this
}

// GetProvider returns the Provider field value
func (o *CalcNodeProfilesV2Request) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *CalcNodeProfilesV2Request) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *CalcNodeProfilesV2Request) SetProvider(v string) {
	o.Provider = v
}

// GetRegion returns the Region field value
func (o *CalcNodeProfilesV2Request) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *CalcNodeProfilesV2Request) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *CalcNodeProfilesV2Request) SetRegion(v string) {
	o.Region = v
}

// GetIsDevTier returns the IsDevTier field value
func (o *CalcNodeProfilesV2Request) GetIsDevTier() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDevTier
}

// GetIsDevTierOk returns a tuple with the IsDevTier field value
// and a boolean to check if the value has been set.
func (o *CalcNodeProfilesV2Request) GetIsDevTierOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDevTier, true
}

// SetIsDevTier sets field value
func (o *CalcNodeProfilesV2Request) SetIsDevTier(v bool) {
	o.IsDevTier = v
}

func (o CalcNodeProfilesV2Request) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CalcNodeProfilesV2Request) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["provider"] = o.Provider
	toSerialize["region"] = o.Region
	toSerialize["is_dev_tier"] = o.IsDevTier
	return toSerialize, nil
}

func (o *CalcNodeProfilesV2Request) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"provider",
		"region",
		"is_dev_tier",
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

	varCalcNodeProfilesV2Request := _CalcNodeProfilesV2Request{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCalcNodeProfilesV2Request)

	if err != nil {
		return err
	}

	*o = CalcNodeProfilesV2Request(varCalcNodeProfilesV2Request)

	return err
}

type NullableCalcNodeProfilesV2Request struct {
	value *CalcNodeProfilesV2Request
	isSet bool
}

func (v NullableCalcNodeProfilesV2Request) Get() *CalcNodeProfilesV2Request {
	return v.value
}

func (v *NullableCalcNodeProfilesV2Request) Set(val *CalcNodeProfilesV2Request) {
	v.value = val
	v.isSet = true
}

func (v NullableCalcNodeProfilesV2Request) IsSet() bool {
	return v.isSet
}

func (v *NullableCalcNodeProfilesV2Request) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCalcNodeProfilesV2Request(val *CalcNodeProfilesV2Request) *NullableCalcNodeProfilesV2Request {
	return &NullableCalcNodeProfilesV2Request{value: val, isSet: true}
}

func (v NullableCalcNodeProfilesV2Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCalcNodeProfilesV2Request) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
