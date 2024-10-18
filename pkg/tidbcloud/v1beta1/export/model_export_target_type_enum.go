/*
TiDB Cloud Serverless Export Open API

TiDB Cloud Serverless Export Open API

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package export

import (
	"encoding/json"
	"fmt"
)

// ExportTargetTypeEnum the model 'ExportTargetTypeEnum'
type ExportTargetTypeEnum string

// List of ExportTargetType.Enum
const (
	EXPORTTARGETTYPEENUM_LOCAL      ExportTargetTypeEnum = "LOCAL"
	EXPORTTARGETTYPEENUM_S3         ExportTargetTypeEnum = "S3"
	EXPORTTARGETTYPEENUM_GCS        ExportTargetTypeEnum = "GCS"
	EXPORTTARGETTYPEENUM_AZURE_BLOB ExportTargetTypeEnum = "AZURE_BLOB"
)

// All allowed values of ExportTargetTypeEnum enum
var AllowedExportTargetTypeEnumEnumValues = []ExportTargetTypeEnum{
	"LOCAL",
	"S3",
	"GCS",
	"AZURE_BLOB",
}

func (v *ExportTargetTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExportTargetTypeEnum(value)
	for _, existing := range AllowedExportTargetTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExportTargetTypeEnum", value)
}

// NewExportTargetTypeEnumFromValue returns a pointer to a valid ExportTargetTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExportTargetTypeEnumFromValue(v string) (*ExportTargetTypeEnum, error) {
	ev := ExportTargetTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExportTargetTypeEnum: valid values are %v", v, AllowedExportTargetTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExportTargetTypeEnum) IsValid() bool {
	for _, existing := range AllowedExportTargetTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExportTargetType.Enum value
func (v ExportTargetTypeEnum) Ptr() *ExportTargetTypeEnum {
	return &v
}

type NullableExportTargetTypeEnum struct {
	value *ExportTargetTypeEnum
	isSet bool
}

func (v NullableExportTargetTypeEnum) Get() *ExportTargetTypeEnum {
	return v.value
}

func (v *NullableExportTargetTypeEnum) Set(val *ExportTargetTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableExportTargetTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableExportTargetTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportTargetTypeEnum(val *ExportTargetTypeEnum) *NullableExportTargetTypeEnum {
	return &NullableExportTargetTypeEnum{value: val, isSet: true}
}

func (v NullableExportTargetTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportTargetTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
