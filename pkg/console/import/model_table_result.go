/*
Import APIs for TiDB Cloud Serverless

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package consoleimportapi

import (
	"encoding/json"
)

// checks if the TableResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TableResult{}

// TableResult struct for TableResult
type TableResult struct {
	TargetTable *ConsoleTable `json:"target_table,omitempty"`
	// The error message.
	ErrorMessage *string `json:"error_message,omitempty"`
	// The matched data files.
	MatchedDataFiles []string `json:"matched_data_files,omitempty"`
	// The matched schema file.
	MatchedSchemaFile *string `json:"matched_schema_file,omitempty"`
	// whether the matched schema file is used in import.
	UseSchemaFile *bool `json:"use_schema_file,omitempty"`
}

// NewTableResult instantiates a new TableResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTableResult() *TableResult {
	this := TableResult{}
	return &this
}

// NewTableResultWithDefaults instantiates a new TableResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTableResultWithDefaults() *TableResult {
	this := TableResult{}
	return &this
}

// GetTargetTable returns the TargetTable field value if set, zero value otherwise.
func (o *TableResult) GetTargetTable() ConsoleTable {
	if o == nil || IsNil(o.TargetTable) {
		var ret ConsoleTable
		return ret
	}
	return *o.TargetTable
}

// GetTargetTableOk returns a tuple with the TargetTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResult) GetTargetTableOk() (*ConsoleTable, bool) {
	if o == nil || IsNil(o.TargetTable) {
		return nil, false
	}
	return o.TargetTable, true
}

// HasTargetTable returns a boolean if a field has been set.
func (o *TableResult) HasTargetTable() bool {
	if o != nil && !IsNil(o.TargetTable) {
		return true
	}

	return false
}

// SetTargetTable gets a reference to the given ConsoleTable and assigns it to the TargetTable field.
func (o *TableResult) SetTargetTable(v ConsoleTable) {
	o.TargetTable = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *TableResult) GetErrorMessage() string {
	if o == nil || IsNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResult) GetErrorMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorMessage) {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *TableResult) HasErrorMessage() bool {
	if o != nil && !IsNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *TableResult) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetMatchedDataFiles returns the MatchedDataFiles field value if set, zero value otherwise.
func (o *TableResult) GetMatchedDataFiles() []string {
	if o == nil || IsNil(o.MatchedDataFiles) {
		var ret []string
		return ret
	}
	return o.MatchedDataFiles
}

// GetMatchedDataFilesOk returns a tuple with the MatchedDataFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResult) GetMatchedDataFilesOk() ([]string, bool) {
	if o == nil || IsNil(o.MatchedDataFiles) {
		return nil, false
	}
	return o.MatchedDataFiles, true
}

// HasMatchedDataFiles returns a boolean if a field has been set.
func (o *TableResult) HasMatchedDataFiles() bool {
	if o != nil && !IsNil(o.MatchedDataFiles) {
		return true
	}

	return false
}

// SetMatchedDataFiles gets a reference to the given []string and assigns it to the MatchedDataFiles field.
func (o *TableResult) SetMatchedDataFiles(v []string) {
	o.MatchedDataFiles = v
}

// GetMatchedSchemaFile returns the MatchedSchemaFile field value if set, zero value otherwise.
func (o *TableResult) GetMatchedSchemaFile() string {
	if o == nil || IsNil(o.MatchedSchemaFile) {
		var ret string
		return ret
	}
	return *o.MatchedSchemaFile
}

// GetMatchedSchemaFileOk returns a tuple with the MatchedSchemaFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResult) GetMatchedSchemaFileOk() (*string, bool) {
	if o == nil || IsNil(o.MatchedSchemaFile) {
		return nil, false
	}
	return o.MatchedSchemaFile, true
}

// HasMatchedSchemaFile returns a boolean if a field has been set.
func (o *TableResult) HasMatchedSchemaFile() bool {
	if o != nil && !IsNil(o.MatchedSchemaFile) {
		return true
	}

	return false
}

// SetMatchedSchemaFile gets a reference to the given string and assigns it to the MatchedSchemaFile field.
func (o *TableResult) SetMatchedSchemaFile(v string) {
	o.MatchedSchemaFile = &v
}

// GetUseSchemaFile returns the UseSchemaFile field value if set, zero value otherwise.
func (o *TableResult) GetUseSchemaFile() bool {
	if o == nil || IsNil(o.UseSchemaFile) {
		var ret bool
		return ret
	}
	return *o.UseSchemaFile
}

// GetUseSchemaFileOk returns a tuple with the UseSchemaFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableResult) GetUseSchemaFileOk() (*bool, bool) {
	if o == nil || IsNil(o.UseSchemaFile) {
		return nil, false
	}
	return o.UseSchemaFile, true
}

// HasUseSchemaFile returns a boolean if a field has been set.
func (o *TableResult) HasUseSchemaFile() bool {
	if o != nil && !IsNil(o.UseSchemaFile) {
		return true
	}

	return false
}

// SetUseSchemaFile gets a reference to the given bool and assigns it to the UseSchemaFile field.
func (o *TableResult) SetUseSchemaFile(v bool) {
	o.UseSchemaFile = &v
}

func (o TableResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TableResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetTable) {
		toSerialize["target_table"] = o.TargetTable
	}
	if !IsNil(o.ErrorMessage) {
		toSerialize["error_message"] = o.ErrorMessage
	}
	if !IsNil(o.MatchedDataFiles) {
		toSerialize["matched_data_files"] = o.MatchedDataFiles
	}
	if !IsNil(o.MatchedSchemaFile) {
		toSerialize["matched_schema_file"] = o.MatchedSchemaFile
	}
	if !IsNil(o.UseSchemaFile) {
		toSerialize["use_schema_file"] = o.UseSchemaFile
	}
	return toSerialize, nil
}

type NullableTableResult struct {
	value *TableResult
	isSet bool
}

func (v NullableTableResult) Get() *TableResult {
	return v.value
}

func (v *NullableTableResult) Set(val *TableResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTableResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTableResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTableResult(val *TableResult) *NullableTableResult {
	return &NullableTableResult{value: val, isSet: true}
}

func (v NullableTableResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTableResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
