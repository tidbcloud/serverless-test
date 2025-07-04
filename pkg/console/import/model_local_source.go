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

// checks if the LocalSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocalSource{}

// LocalSource struct for LocalSource
type LocalSource struct {
	// The upload id of import source file.
	UploadId string `json:"upload_id"`
	// The target database of import.
	TargetDatabase string `json:"target_database"`
	// The target table of import.
	TargetTable string `json:"target_table"`
	// The file name of import source file.
	FileName *string `json:"file_name,omitempty"`
	// INPUT_ONLY. The columns definition of the target table.
	Columns []ColumnInfo `json:"columns"`
	// INPUT_ONLY. The primary key columns of the target table.
	PkColumns []string `json:"pk_columns"`
}

type _LocalSource LocalSource

// NewLocalSource instantiates a new LocalSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalSource(uploadId string, targetDatabase string, targetTable string, columns []ColumnInfo, pkColumns []string) *LocalSource {
	this := LocalSource{}
	this.UploadId = uploadId
	this.TargetDatabase = targetDatabase
	this.TargetTable = targetTable
	this.Columns = columns
	this.PkColumns = pkColumns
	return &this
}

// NewLocalSourceWithDefaults instantiates a new LocalSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalSourceWithDefaults() *LocalSource {
	this := LocalSource{}
	return &this
}

// GetUploadId returns the UploadId field value
func (o *LocalSource) GetUploadId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UploadId
}

// GetUploadIdOk returns a tuple with the UploadId field value
// and a boolean to check if the value has been set.
func (o *LocalSource) GetUploadIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploadId, true
}

// SetUploadId sets field value
func (o *LocalSource) SetUploadId(v string) {
	o.UploadId = v
}

// GetTargetDatabase returns the TargetDatabase field value
func (o *LocalSource) GetTargetDatabase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetDatabase
}

// GetTargetDatabaseOk returns a tuple with the TargetDatabase field value
// and a boolean to check if the value has been set.
func (o *LocalSource) GetTargetDatabaseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetDatabase, true
}

// SetTargetDatabase sets field value
func (o *LocalSource) SetTargetDatabase(v string) {
	o.TargetDatabase = v
}

// GetTargetTable returns the TargetTable field value
func (o *LocalSource) GetTargetTable() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetTable
}

// GetTargetTableOk returns a tuple with the TargetTable field value
// and a boolean to check if the value has been set.
func (o *LocalSource) GetTargetTableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetTable, true
}

// SetTargetTable sets field value
func (o *LocalSource) SetTargetTable(v string) {
	o.TargetTable = v
}

// GetFileName returns the FileName field value if set, zero value otherwise.
func (o *LocalSource) GetFileName() string {
	if o == nil || IsNil(o.FileName) {
		var ret string
		return ret
	}
	return *o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalSource) GetFileNameOk() (*string, bool) {
	if o == nil || IsNil(o.FileName) {
		return nil, false
	}
	return o.FileName, true
}

// HasFileName returns a boolean if a field has been set.
func (o *LocalSource) HasFileName() bool {
	if o != nil && !IsNil(o.FileName) {
		return true
	}

	return false
}

// SetFileName gets a reference to the given string and assigns it to the FileName field.
func (o *LocalSource) SetFileName(v string) {
	o.FileName = &v
}

// GetColumns returns the Columns field value
func (o *LocalSource) GetColumns() []ColumnInfo {
	if o == nil {
		var ret []ColumnInfo
		return ret
	}

	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *LocalSource) GetColumnsOk() ([]ColumnInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Columns, true
}

// SetColumns sets field value
func (o *LocalSource) SetColumns(v []ColumnInfo) {
	o.Columns = v
}

// GetPkColumns returns the PkColumns field value
func (o *LocalSource) GetPkColumns() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PkColumns
}

// GetPkColumnsOk returns a tuple with the PkColumns field value
// and a boolean to check if the value has been set.
func (o *LocalSource) GetPkColumnsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PkColumns, true
}

// SetPkColumns sets field value
func (o *LocalSource) SetPkColumns(v []string) {
	o.PkColumns = v
}

func (o LocalSource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocalSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["upload_id"] = o.UploadId
	toSerialize["target_database"] = o.TargetDatabase
	toSerialize["target_table"] = o.TargetTable
	if !IsNil(o.FileName) {
		toSerialize["file_name"] = o.FileName
	}
	toSerialize["columns"] = o.Columns
	toSerialize["pk_columns"] = o.PkColumns
	return toSerialize, nil
}

func (o *LocalSource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"upload_id",
		"target_database",
		"target_table",
		"columns",
		"pk_columns",
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

	varLocalSource := _LocalSource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLocalSource)

	if err != nil {
		return err
	}

	*o = LocalSource(varLocalSource)

	return err
}

type NullableLocalSource struct {
	value *LocalSource
	isSet bool
}

func (v NullableLocalSource) Get() *LocalSource {
	return v.value
}

func (v *NullableLocalSource) Set(val *LocalSource) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalSource) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalSource(val *LocalSource) *NullableLocalSource {
	return &NullableLocalSource{value: val, isSet: true}
}

func (v NullableLocalSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
