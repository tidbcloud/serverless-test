/*
gRPC for our DBaaS central service

notably uses gRPC-Gateway with OpenAPI

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package coreportalapi

import (
	"encoding/json"
	"fmt"
)

// ListClustersFilterStatusesParameterInner the model 'ListClustersFilterStatusesParameterInner'
type ListClustersFilterStatusesParameterInner string

// List of ListClusters_filter_statuses_parameter_inner
const (
	LISTCLUSTERSFILTERSTATUSESPARAMETERINNER_NORMAL      ListClustersFilterStatusesParameterInner = "NORMAL"
	LISTCLUSTERSFILTERSTATUSESPARAMETERINNER_CREATING    ListClustersFilterStatusesParameterInner = "CREATING"
	LISTCLUSTERSFILTERSTATUSESPARAMETERINNER_SCALING     ListClustersFilterStatusesParameterInner = "SCALING"
	LISTCLUSTERSFILTERSTATUSESPARAMETERINNER_UPGRADING   ListClustersFilterStatusesParameterInner = "UPGRADING"
	LISTCLUSTERSFILTERSTATUSESPARAMETERINNER_UNAVAILABLE ListClustersFilterStatusesParameterInner = "UNAVAILABLE"
	LISTCLUSTERSFILTERSTATUSESPARAMETERINNER_RECOVERING  ListClustersFilterStatusesParameterInner = "RECOVERING"
	LISTCLUSTERSFILTERSTATUSESPARAMETERINNER_IMPORTING   ListClustersFilterStatusesParameterInner = "IMPORTING"
	LISTCLUSTERSFILTERSTATUSESPARAMETERINNER_PAUSED      ListClustersFilterStatusesParameterInner = "PAUSED"
	LISTCLUSTERSFILTERSTATUSESPARAMETERINNER_RESUMING    ListClustersFilterStatusesParameterInner = "RESUMING"
	LISTCLUSTERSFILTERSTATUSESPARAMETERINNER_DELETED     ListClustersFilterStatusesParameterInner = "DELETED"
	LISTCLUSTERSFILTERSTATUSESPARAMETERINNER_MAINTAINING ListClustersFilterStatusesParameterInner = "MAINTAINING"
	LISTCLUSTERSFILTERSTATUSESPARAMETERINNER_PAUSING     ListClustersFilterStatusesParameterInner = "PAUSING"
)

// All allowed values of ListClustersFilterStatusesParameterInner enum
var AllowedListClustersFilterStatusesParameterInnerEnumValues = []ListClustersFilterStatusesParameterInner{
	"NORMAL",
	"CREATING",
	"SCALING",
	"UPGRADING",
	"UNAVAILABLE",
	"RECOVERING",
	"IMPORTING",
	"PAUSED",
	"RESUMING",
	"DELETED",
	"MAINTAINING",
	"PAUSING",
}

func (v *ListClustersFilterStatusesParameterInner) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ListClustersFilterStatusesParameterInner(value)
	for _, existing := range AllowedListClustersFilterStatusesParameterInnerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ListClustersFilterStatusesParameterInner", value)
}

// NewListClustersFilterStatusesParameterInnerFromValue returns a pointer to a valid ListClustersFilterStatusesParameterInner
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewListClustersFilterStatusesParameterInnerFromValue(v string) (*ListClustersFilterStatusesParameterInner, error) {
	ev := ListClustersFilterStatusesParameterInner(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ListClustersFilterStatusesParameterInner: valid values are %v", v, AllowedListClustersFilterStatusesParameterInnerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ListClustersFilterStatusesParameterInner) IsValid() bool {
	for _, existing := range AllowedListClustersFilterStatusesParameterInnerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ListClusters_filter_statuses_parameter_inner value
func (v ListClustersFilterStatusesParameterInner) Ptr() *ListClustersFilterStatusesParameterInner {
	return &v
}

type NullableListClustersFilterStatusesParameterInner struct {
	value *ListClustersFilterStatusesParameterInner
	isSet bool
}

func (v NullableListClustersFilterStatusesParameterInner) Get() *ListClustersFilterStatusesParameterInner {
	return v.value
}

func (v *NullableListClustersFilterStatusesParameterInner) Set(val *ListClustersFilterStatusesParameterInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListClustersFilterStatusesParameterInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListClustersFilterStatusesParameterInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListClustersFilterStatusesParameterInner(val *ListClustersFilterStatusesParameterInner) *NullableListClustersFilterStatusesParameterInner {
	return &NullableListClustersFilterStatusesParameterInner{value: val, isSet: true}
}

func (v NullableListClustersFilterStatusesParameterInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListClustersFilterStatusesParameterInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
