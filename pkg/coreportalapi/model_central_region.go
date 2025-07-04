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

// checks if the CentralRegion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CentralRegion{}

// CentralRegion struct for CentralRegion
type CentralRegion struct {
	Name                     string  `json:"name"`
	DisplayName              string  `json:"display_name"`
	CrossAccountDeploymentId *string `json:"cross_account_deployment_id,omitempty"`
	Area                     *string `json:"area,omitempty"`
	CountryCode              *string `json:"country_code,omitempty"`
}

type _CentralRegion CentralRegion

// NewCentralRegion instantiates a new CentralRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCentralRegion(name string, displayName string) *CentralRegion {
	this := CentralRegion{}
	this.Name = name
	this.DisplayName = displayName
	return &this
}

// NewCentralRegionWithDefaults instantiates a new CentralRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCentralRegionWithDefaults() *CentralRegion {
	this := CentralRegion{}
	return &this
}

// GetName returns the Name field value
func (o *CentralRegion) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CentralRegion) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CentralRegion) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value
func (o *CentralRegion) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *CentralRegion) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *CentralRegion) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetCrossAccountDeploymentId returns the CrossAccountDeploymentId field value if set, zero value otherwise.
func (o *CentralRegion) GetCrossAccountDeploymentId() string {
	if o == nil || IsNil(o.CrossAccountDeploymentId) {
		var ret string
		return ret
	}
	return *o.CrossAccountDeploymentId
}

// GetCrossAccountDeploymentIdOk returns a tuple with the CrossAccountDeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CentralRegion) GetCrossAccountDeploymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.CrossAccountDeploymentId) {
		return nil, false
	}
	return o.CrossAccountDeploymentId, true
}

// HasCrossAccountDeploymentId returns a boolean if a field has been set.
func (o *CentralRegion) HasCrossAccountDeploymentId() bool {
	if o != nil && !IsNil(o.CrossAccountDeploymentId) {
		return true
	}

	return false
}

// SetCrossAccountDeploymentId gets a reference to the given string and assigns it to the CrossAccountDeploymentId field.
func (o *CentralRegion) SetCrossAccountDeploymentId(v string) {
	o.CrossAccountDeploymentId = &v
}

// GetArea returns the Area field value if set, zero value otherwise.
func (o *CentralRegion) GetArea() string {
	if o == nil || IsNil(o.Area) {
		var ret string
		return ret
	}
	return *o.Area
}

// GetAreaOk returns a tuple with the Area field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CentralRegion) GetAreaOk() (*string, bool) {
	if o == nil || IsNil(o.Area) {
		return nil, false
	}
	return o.Area, true
}

// HasArea returns a boolean if a field has been set.
func (o *CentralRegion) HasArea() bool {
	if o != nil && !IsNil(o.Area) {
		return true
	}

	return false
}

// SetArea gets a reference to the given string and assigns it to the Area field.
func (o *CentralRegion) SetArea(v string) {
	o.Area = &v
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *CentralRegion) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CentralRegion) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *CentralRegion) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *CentralRegion) SetCountryCode(v string) {
	o.CountryCode = &v
}

func (o CentralRegion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CentralRegion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["display_name"] = o.DisplayName
	if !IsNil(o.CrossAccountDeploymentId) {
		toSerialize["cross_account_deployment_id"] = o.CrossAccountDeploymentId
	}
	if !IsNil(o.Area) {
		toSerialize["area"] = o.Area
	}
	if !IsNil(o.CountryCode) {
		toSerialize["country_code"] = o.CountryCode
	}
	return toSerialize, nil
}

func (o *CentralRegion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"display_name",
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

	varCentralRegion := _CentralRegion{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCentralRegion)

	if err != nil {
		return err
	}

	*o = CentralRegion(varCentralRegion)

	return err
}

type NullableCentralRegion struct {
	value *CentralRegion
	isSet bool
}

func (v NullableCentralRegion) Get() *CentralRegion {
	return v.value
}

func (v *NullableCentralRegion) Set(val *CentralRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableCentralRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableCentralRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCentralRegion(val *CentralRegion) *NullableCentralRegion {
	return &NullableCentralRegion{value: val, isSet: true}
}

func (v NullableCentralRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCentralRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
