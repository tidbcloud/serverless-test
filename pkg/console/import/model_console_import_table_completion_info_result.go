/*
Import APIs for TiDB Cloud Serverless

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: alpha
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package consoleimportapi

import (
	"encoding/json"
	"fmt"
)

// ConsoleImportTableCompletionInfoResult the model 'ConsoleImportTableCompletionInfoResult'
type ConsoleImportTableCompletionInfoResult string

// List of console.ImportTableCompletionInfo.Result
const (
	CONSOLEIMPORTTABLECOMPLETIONINFORESULT_SUCCESS ConsoleImportTableCompletionInfoResult = "SUCCESS"
	CONSOLEIMPORTTABLECOMPLETIONINFORESULT_WARNING ConsoleImportTableCompletionInfoResult = "WARNING"
	CONSOLEIMPORTTABLECOMPLETIONINFORESULT_ERROR   ConsoleImportTableCompletionInfoResult = "ERROR"
)

// All allowed values of ConsoleImportTableCompletionInfoResult enum
var AllowedConsoleImportTableCompletionInfoResultEnumValues = []ConsoleImportTableCompletionInfoResult{
	"SUCCESS",
	"WARNING",
	"ERROR",
}

func (v *ConsoleImportTableCompletionInfoResult) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConsoleImportTableCompletionInfoResult(value)
	for _, existing := range AllowedConsoleImportTableCompletionInfoResultEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConsoleImportTableCompletionInfoResult", value)
}

// NewConsoleImportTableCompletionInfoResultFromValue returns a pointer to a valid ConsoleImportTableCompletionInfoResult
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConsoleImportTableCompletionInfoResultFromValue(v string) (*ConsoleImportTableCompletionInfoResult, error) {
	ev := ConsoleImportTableCompletionInfoResult(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConsoleImportTableCompletionInfoResult: valid values are %v", v, AllowedConsoleImportTableCompletionInfoResultEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConsoleImportTableCompletionInfoResult) IsValid() bool {
	for _, existing := range AllowedConsoleImportTableCompletionInfoResultEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to console.ImportTableCompletionInfo.Result value
func (v ConsoleImportTableCompletionInfoResult) Ptr() *ConsoleImportTableCompletionInfoResult {
	return &v
}

type NullableConsoleImportTableCompletionInfoResult struct {
	value *ConsoleImportTableCompletionInfoResult
	isSet bool
}

func (v NullableConsoleImportTableCompletionInfoResult) Get() *ConsoleImportTableCompletionInfoResult {
	return v.value
}

func (v *NullableConsoleImportTableCompletionInfoResult) Set(val *ConsoleImportTableCompletionInfoResult) {
	v.value = val
	v.isSet = true
}

func (v NullableConsoleImportTableCompletionInfoResult) IsSet() bool {
	return v.isSet
}

func (v *NullableConsoleImportTableCompletionInfoResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsoleImportTableCompletionInfoResult(val *ConsoleImportTableCompletionInfoResult) *NullableConsoleImportTableCompletionInfoResult {
	return &NullableConsoleImportTableCompletionInfoResult{value: val, isSet: true}
}

func (v NullableConsoleImportTableCompletionInfoResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsoleImportTableCompletionInfoResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
