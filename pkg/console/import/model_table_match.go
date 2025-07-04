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

// checks if the TableMatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TableMatch{}

// TableMatch struct for TableMatch
type TableMatch struct {
	TargetTable *ConsoleTable `json:"target_table,omitempty"`
	// The warning message.
	WarningMessage *string `json:"warning_message,omitempty"`
}

// NewTableMatch instantiates a new TableMatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTableMatch() *TableMatch {
	this := TableMatch{}
	return &this
}

// NewTableMatchWithDefaults instantiates a new TableMatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTableMatchWithDefaults() *TableMatch {
	this := TableMatch{}
	return &this
}

// GetTargetTable returns the TargetTable field value if set, zero value otherwise.
func (o *TableMatch) GetTargetTable() ConsoleTable {
	if o == nil || IsNil(o.TargetTable) {
		var ret ConsoleTable
		return ret
	}
	return *o.TargetTable
}

// GetTargetTableOk returns a tuple with the TargetTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableMatch) GetTargetTableOk() (*ConsoleTable, bool) {
	if o == nil || IsNil(o.TargetTable) {
		return nil, false
	}
	return o.TargetTable, true
}

// HasTargetTable returns a boolean if a field has been set.
func (o *TableMatch) HasTargetTable() bool {
	if o != nil && !IsNil(o.TargetTable) {
		return true
	}

	return false
}

// SetTargetTable gets a reference to the given ConsoleTable and assigns it to the TargetTable field.
func (o *TableMatch) SetTargetTable(v ConsoleTable) {
	o.TargetTable = &v
}

// GetWarningMessage returns the WarningMessage field value if set, zero value otherwise.
func (o *TableMatch) GetWarningMessage() string {
	if o == nil || IsNil(o.WarningMessage) {
		var ret string
		return ret
	}
	return *o.WarningMessage
}

// GetWarningMessageOk returns a tuple with the WarningMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableMatch) GetWarningMessageOk() (*string, bool) {
	if o == nil || IsNil(o.WarningMessage) {
		return nil, false
	}
	return o.WarningMessage, true
}

// HasWarningMessage returns a boolean if a field has been set.
func (o *TableMatch) HasWarningMessage() bool {
	if o != nil && !IsNil(o.WarningMessage) {
		return true
	}

	return false
}

// SetWarningMessage gets a reference to the given string and assigns it to the WarningMessage field.
func (o *TableMatch) SetWarningMessage(v string) {
	o.WarningMessage = &v
}

func (o TableMatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TableMatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetTable) {
		toSerialize["target_table"] = o.TargetTable
	}
	if !IsNil(o.WarningMessage) {
		toSerialize["warning_message"] = o.WarningMessage
	}
	return toSerialize, nil
}

type NullableTableMatch struct {
	value *TableMatch
	isSet bool
}

func (v NullableTableMatch) Get() *TableMatch {
	return v.value
}

func (v *NullableTableMatch) Set(val *TableMatch) {
	v.value = val
	v.isSet = true
}

func (v NullableTableMatch) IsSet() bool {
	return v.isSet
}

func (v *NullableTableMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTableMatch(val *TableMatch) *NullableTableMatch {
	return &NullableTableMatch{value: val, isSet: true}
}

func (v NullableTableMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTableMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
