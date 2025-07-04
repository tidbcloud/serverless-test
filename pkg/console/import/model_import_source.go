/*
Import APIs for TiDB Cloud Serverless

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package consoleimportapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the ImportSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportSource{}

// ImportSource struct for ImportSource
type ImportSource struct {
	Type         ImportSourceTypeEnum `json:"type"`
	Local        *LocalSource         `json:"local,omitempty"`
	S3           *S3Source            `json:"s3,omitempty"`
	Gcs          *GCSSource           `json:"gcs,omitempty"`
	AzureBlob    *AzureBlobSource     `json:"azure_blob,omitempty"`
	S3Compatible *S3CompatibleSource  `json:"s3_compatible,omitempty"`
	Oss          *OSSSource           `json:"oss,omitempty"`
}

type _ImportSource ImportSource

// NewImportSource instantiates a new ImportSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportSource(type_ ImportSourceTypeEnum) *ImportSource {
	this := ImportSource{}
	this.Type = type_
	return &this
}

// NewImportSourceWithDefaults instantiates a new ImportSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportSourceWithDefaults() *ImportSource {
	this := ImportSource{}
	return &this
}

// GetType returns the Type field value
func (o *ImportSource) GetType() ImportSourceTypeEnum {
	if o == nil {
		var ret ImportSourceTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ImportSource) GetTypeOk() (*ImportSourceTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ImportSource) SetType(v ImportSourceTypeEnum) {
	o.Type = v
}

// GetLocal returns the Local field value if set, zero value otherwise.
func (o *ImportSource) GetLocal() LocalSource {
	if o == nil || IsNil(o.Local) {
		var ret LocalSource
		return ret
	}
	return *o.Local
}

// GetLocalOk returns a tuple with the Local field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSource) GetLocalOk() (*LocalSource, bool) {
	if o == nil || IsNil(o.Local) {
		return nil, false
	}
	return o.Local, true
}

// HasLocal returns a boolean if a field has been set.
func (o *ImportSource) HasLocal() bool {
	if o != nil && !IsNil(o.Local) {
		return true
	}

	return false
}

// SetLocal gets a reference to the given LocalSource and assigns it to the Local field.
func (o *ImportSource) SetLocal(v LocalSource) {
	o.Local = &v
}

// GetS3 returns the S3 field value if set, zero value otherwise.
func (o *ImportSource) GetS3() S3Source {
	if o == nil || IsNil(o.S3) {
		var ret S3Source
		return ret
	}
	return *o.S3
}

// GetS3Ok returns a tuple with the S3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSource) GetS3Ok() (*S3Source, bool) {
	if o == nil || IsNil(o.S3) {
		return nil, false
	}
	return o.S3, true
}

// HasS3 returns a boolean if a field has been set.
func (o *ImportSource) HasS3() bool {
	if o != nil && !IsNil(o.S3) {
		return true
	}

	return false
}

// SetS3 gets a reference to the given S3Source and assigns it to the S3 field.
func (o *ImportSource) SetS3(v S3Source) {
	o.S3 = &v
}

// GetGcs returns the Gcs field value if set, zero value otherwise.
func (o *ImportSource) GetGcs() GCSSource {
	if o == nil || IsNil(o.Gcs) {
		var ret GCSSource
		return ret
	}
	return *o.Gcs
}

// GetGcsOk returns a tuple with the Gcs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSource) GetGcsOk() (*GCSSource, bool) {
	if o == nil || IsNil(o.Gcs) {
		return nil, false
	}
	return o.Gcs, true
}

// HasGcs returns a boolean if a field has been set.
func (o *ImportSource) HasGcs() bool {
	if o != nil && !IsNil(o.Gcs) {
		return true
	}

	return false
}

// SetGcs gets a reference to the given GCSSource and assigns it to the Gcs field.
func (o *ImportSource) SetGcs(v GCSSource) {
	o.Gcs = &v
}

// GetAzureBlob returns the AzureBlob field value if set, zero value otherwise.
func (o *ImportSource) GetAzureBlob() AzureBlobSource {
	if o == nil || IsNil(o.AzureBlob) {
		var ret AzureBlobSource
		return ret
	}
	return *o.AzureBlob
}

// GetAzureBlobOk returns a tuple with the AzureBlob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSource) GetAzureBlobOk() (*AzureBlobSource, bool) {
	if o == nil || IsNil(o.AzureBlob) {
		return nil, false
	}
	return o.AzureBlob, true
}

// HasAzureBlob returns a boolean if a field has been set.
func (o *ImportSource) HasAzureBlob() bool {
	if o != nil && !IsNil(o.AzureBlob) {
		return true
	}

	return false
}

// SetAzureBlob gets a reference to the given AzureBlobSource and assigns it to the AzureBlob field.
func (o *ImportSource) SetAzureBlob(v AzureBlobSource) {
	o.AzureBlob = &v
}

// GetS3Compatible returns the S3Compatible field value if set, zero value otherwise.
func (o *ImportSource) GetS3Compatible() S3CompatibleSource {
	if o == nil || IsNil(o.S3Compatible) {
		var ret S3CompatibleSource
		return ret
	}
	return *o.S3Compatible
}

// GetS3CompatibleOk returns a tuple with the S3Compatible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSource) GetS3CompatibleOk() (*S3CompatibleSource, bool) {
	if o == nil || IsNil(o.S3Compatible) {
		return nil, false
	}
	return o.S3Compatible, true
}

// HasS3Compatible returns a boolean if a field has been set.
func (o *ImportSource) HasS3Compatible() bool {
	if o != nil && !IsNil(o.S3Compatible) {
		return true
	}

	return false
}

// SetS3Compatible gets a reference to the given S3CompatibleSource and assigns it to the S3Compatible field.
func (o *ImportSource) SetS3Compatible(v S3CompatibleSource) {
	o.S3Compatible = &v
}

// GetOss returns the Oss field value if set, zero value otherwise.
func (o *ImportSource) GetOss() OSSSource {
	if o == nil || IsNil(o.Oss) {
		var ret OSSSource
		return ret
	}
	return *o.Oss
}

// GetOssOk returns a tuple with the Oss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportSource) GetOssOk() (*OSSSource, bool) {
	if o == nil || IsNil(o.Oss) {
		return nil, false
	}
	return o.Oss, true
}

// HasOss returns a boolean if a field has been set.
func (o *ImportSource) HasOss() bool {
	if o != nil && !IsNil(o.Oss) {
		return true
	}

	return false
}

// SetOss gets a reference to the given OSSSource and assigns it to the Oss field.
func (o *ImportSource) SetOss(v OSSSource) {
	o.Oss = &v
}

func (o ImportSource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Local) {
		toSerialize["local"] = o.Local
	}
	if !IsNil(o.S3) {
		toSerialize["s3"] = o.S3
	}
	if !IsNil(o.Gcs) {
		toSerialize["gcs"] = o.Gcs
	}
	if !IsNil(o.AzureBlob) {
		toSerialize["azure_blob"] = o.AzureBlob
	}
	if !IsNil(o.S3Compatible) {
		toSerialize["s3_compatible"] = o.S3Compatible
	}
	if !IsNil(o.Oss) {
		toSerialize["oss"] = o.Oss
	}
	return toSerialize, nil
}

func (o *ImportSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varImportSource := _ImportSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varImportSource)

	if err != nil {
		return err
	}

	*o = ImportSource(varImportSource)

	return err
}

type NullableImportSource struct {
	value *ImportSource
	isSet bool
}

func (v NullableImportSource) Get() *ImportSource {
	return v.value
}

func (v *NullableImportSource) Set(val *ImportSource) {
	v.value = val
	v.isSet = true
}

func (v NullableImportSource) IsSet() bool {
	return v.isSet
}

func (v *NullableImportSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportSource(val *ImportSource) *NullableImportSource {
	return &NullableImportSource{value: val, isSet: true}
}

func (v NullableImportSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
