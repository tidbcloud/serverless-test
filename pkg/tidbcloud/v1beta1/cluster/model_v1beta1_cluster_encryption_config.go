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

// checks if the V1beta1ClusterEncryptionConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1beta1ClusterEncryptionConfig{}

// V1beta1ClusterEncryptionConfig Message for encryption settings for a cluster.
type V1beta1ClusterEncryptionConfig struct {
	// Optional. Whether enhanced encryption for cluster data is enabled.
	EnhancedEncryptionEnabled *bool `json:"enhancedEncryptionEnabled,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _V1beta1ClusterEncryptionConfig V1beta1ClusterEncryptionConfig

// NewV1beta1ClusterEncryptionConfig instantiates a new V1beta1ClusterEncryptionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1beta1ClusterEncryptionConfig() *V1beta1ClusterEncryptionConfig {
	this := V1beta1ClusterEncryptionConfig{}
	return &this
}

// NewV1beta1ClusterEncryptionConfigWithDefaults instantiates a new V1beta1ClusterEncryptionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1beta1ClusterEncryptionConfigWithDefaults() *V1beta1ClusterEncryptionConfig {
	this := V1beta1ClusterEncryptionConfig{}
	return &this
}

// GetEnhancedEncryptionEnabled returns the EnhancedEncryptionEnabled field value if set, zero value otherwise.
func (o *V1beta1ClusterEncryptionConfig) GetEnhancedEncryptionEnabled() bool {
	if o == nil || IsNil(o.EnhancedEncryptionEnabled) {
		var ret bool
		return ret
	}
	return *o.EnhancedEncryptionEnabled
}

// GetEnhancedEncryptionEnabledOk returns a tuple with the EnhancedEncryptionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1beta1ClusterEncryptionConfig) GetEnhancedEncryptionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EnhancedEncryptionEnabled) {
		return nil, false
	}
	return o.EnhancedEncryptionEnabled, true
}

// HasEnhancedEncryptionEnabled returns a boolean if a field has been set.
func (o *V1beta1ClusterEncryptionConfig) HasEnhancedEncryptionEnabled() bool {
	if o != nil && !IsNil(o.EnhancedEncryptionEnabled) {
		return true
	}

	return false
}

// SetEnhancedEncryptionEnabled gets a reference to the given bool and assigns it to the EnhancedEncryptionEnabled field.
func (o *V1beta1ClusterEncryptionConfig) SetEnhancedEncryptionEnabled(v bool) {
	o.EnhancedEncryptionEnabled = &v
}

func (o V1beta1ClusterEncryptionConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1beta1ClusterEncryptionConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EnhancedEncryptionEnabled) {
		toSerialize["enhancedEncryptionEnabled"] = o.EnhancedEncryptionEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V1beta1ClusterEncryptionConfig) UnmarshalJSON(data []byte) (err error) {
	varV1beta1ClusterEncryptionConfig := _V1beta1ClusterEncryptionConfig{}

	err = json.Unmarshal(data, &varV1beta1ClusterEncryptionConfig)

	if err != nil {
		return err
	}

	*o = V1beta1ClusterEncryptionConfig(varV1beta1ClusterEncryptionConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "enhancedEncryptionEnabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV1beta1ClusterEncryptionConfig struct {
	value *V1beta1ClusterEncryptionConfig
	isSet bool
}

func (v NullableV1beta1ClusterEncryptionConfig) Get() *V1beta1ClusterEncryptionConfig {
	return v.value
}

func (v *NullableV1beta1ClusterEncryptionConfig) Set(val *V1beta1ClusterEncryptionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableV1beta1ClusterEncryptionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableV1beta1ClusterEncryptionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1beta1ClusterEncryptionConfig(val *V1beta1ClusterEncryptionConfig) *NullableV1beta1ClusterEncryptionConfig {
	return &NullableV1beta1ClusterEncryptionConfig{value: val, isSet: true}
}

func (v NullableV1beta1ClusterEncryptionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1beta1ClusterEncryptionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
