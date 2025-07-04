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

// checks if the CentralAuditLogConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CentralAuditLogConfig{}

// CentralAuditLogConfig struct for CentralAuditLogConfig
type CentralAuditLogConfig struct {
	BucketUrl      string `json:"bucket_url"`
	Region         string `json:"region"`
	RoleArn        string `json:"role_arn"`
	DataSourceType int64  `json:"data_source_type"`
}

type _CentralAuditLogConfig CentralAuditLogConfig

// NewCentralAuditLogConfig instantiates a new CentralAuditLogConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCentralAuditLogConfig(bucketUrl string, region string, roleArn string, dataSourceType int64) *CentralAuditLogConfig {
	this := CentralAuditLogConfig{}
	this.BucketUrl = bucketUrl
	this.Region = region
	this.RoleArn = roleArn
	this.DataSourceType = dataSourceType
	return &this
}

// NewCentralAuditLogConfigWithDefaults instantiates a new CentralAuditLogConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCentralAuditLogConfigWithDefaults() *CentralAuditLogConfig {
	this := CentralAuditLogConfig{}
	return &this
}

// GetBucketUrl returns the BucketUrl field value
func (o *CentralAuditLogConfig) GetBucketUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BucketUrl
}

// GetBucketUrlOk returns a tuple with the BucketUrl field value
// and a boolean to check if the value has been set.
func (o *CentralAuditLogConfig) GetBucketUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BucketUrl, true
}

// SetBucketUrl sets field value
func (o *CentralAuditLogConfig) SetBucketUrl(v string) {
	o.BucketUrl = v
}

// GetRegion returns the Region field value
func (o *CentralAuditLogConfig) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *CentralAuditLogConfig) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *CentralAuditLogConfig) SetRegion(v string) {
	o.Region = v
}

// GetRoleArn returns the RoleArn field value
func (o *CentralAuditLogConfig) GetRoleArn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoleArn
}

// GetRoleArnOk returns a tuple with the RoleArn field value
// and a boolean to check if the value has been set.
func (o *CentralAuditLogConfig) GetRoleArnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleArn, true
}

// SetRoleArn sets field value
func (o *CentralAuditLogConfig) SetRoleArn(v string) {
	o.RoleArn = v
}

// GetDataSourceType returns the DataSourceType field value
func (o *CentralAuditLogConfig) GetDataSourceType() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DataSourceType
}

// GetDataSourceTypeOk returns a tuple with the DataSourceType field value
// and a boolean to check if the value has been set.
func (o *CentralAuditLogConfig) GetDataSourceTypeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSourceType, true
}

// SetDataSourceType sets field value
func (o *CentralAuditLogConfig) SetDataSourceType(v int64) {
	o.DataSourceType = v
}

func (o CentralAuditLogConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CentralAuditLogConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bucket_url"] = o.BucketUrl
	toSerialize["region"] = o.Region
	toSerialize["role_arn"] = o.RoleArn
	toSerialize["data_source_type"] = o.DataSourceType
	return toSerialize, nil
}

func (o *CentralAuditLogConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bucket_url",
		"region",
		"role_arn",
		"data_source_type",
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

	varCentralAuditLogConfig := _CentralAuditLogConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCentralAuditLogConfig)

	if err != nil {
		return err
	}

	*o = CentralAuditLogConfig(varCentralAuditLogConfig)

	return err
}

type NullableCentralAuditLogConfig struct {
	value *CentralAuditLogConfig
	isSet bool
}

func (v NullableCentralAuditLogConfig) Get() *CentralAuditLogConfig {
	return v.value
}

func (v *NullableCentralAuditLogConfig) Set(val *CentralAuditLogConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCentralAuditLogConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCentralAuditLogConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCentralAuditLogConfig(val *CentralAuditLogConfig) *NullableCentralAuditLogConfig {
	return &NullableCentralAuditLogConfig{value: val, isSet: true}
}

func (v NullableCentralAuditLogConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCentralAuditLogConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
