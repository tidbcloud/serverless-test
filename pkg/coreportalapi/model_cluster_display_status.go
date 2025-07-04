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

// ClusterDisplayStatus the model 'ClusterDisplayStatus'
type ClusterDisplayStatus string

// List of ClusterDisplayStatus
const (
	CLUSTERDISPLAYSTATUS_DISPLAY_UNSPECIFIED ClusterDisplayStatus = "DISPLAY_UNSPECIFIED"
	CLUSTERDISPLAYSTATUS_DISPLAY_CREATING    ClusterDisplayStatus = "DISPLAY_CREATING"
	CLUSTERDISPLAYSTATUS_DISPLAY_RESTORING   ClusterDisplayStatus = "DISPLAY_RESTORING"
	CLUSTERDISPLAYSTATUS_DISPLAY_AVAILABLE   ClusterDisplayStatus = "DISPLAY_AVAILABLE"
	CLUSTERDISPLAYSTATUS_DISPLAY_IMPORTING   ClusterDisplayStatus = "DISPLAY_IMPORTING"
	CLUSTERDISPLAYSTATUS_DISPLAY_MODIFYING   ClusterDisplayStatus = "DISPLAY_MODIFYING"
	CLUSTERDISPLAYSTATUS_DISPLAY_PAUSED      ClusterDisplayStatus = "DISPLAY_PAUSED"
	CLUSTERDISPLAYSTATUS_DISPLAY_RESUMING    ClusterDisplayStatus = "DISPLAY_RESUMING"
	CLUSTERDISPLAYSTATUS_DISPLAY_UNAVAILABLE ClusterDisplayStatus = "DISPLAY_UNAVAILABLE"
	CLUSTERDISPLAYSTATUS_DISPLAY_MAINTAINING ClusterDisplayStatus = "DISPLAY_MAINTAINING"
	CLUSTERDISPLAYSTATUS_DISPLAY_PAUSING     ClusterDisplayStatus = "DISPLAY_PAUSING"
)

// All allowed values of ClusterDisplayStatus enum
var AllowedClusterDisplayStatusEnumValues = []ClusterDisplayStatus{
	"DISPLAY_UNSPECIFIED",
	"DISPLAY_CREATING",
	"DISPLAY_RESTORING",
	"DISPLAY_AVAILABLE",
	"DISPLAY_IMPORTING",
	"DISPLAY_MODIFYING",
	"DISPLAY_PAUSED",
	"DISPLAY_RESUMING",
	"DISPLAY_UNAVAILABLE",
	"DISPLAY_MAINTAINING",
	"DISPLAY_PAUSING",
}

func (v *ClusterDisplayStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ClusterDisplayStatus(value)
	for _, existing := range AllowedClusterDisplayStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ClusterDisplayStatus", value)
}

// NewClusterDisplayStatusFromValue returns a pointer to a valid ClusterDisplayStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewClusterDisplayStatusFromValue(v string) (*ClusterDisplayStatus, error) {
	ev := ClusterDisplayStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ClusterDisplayStatus: valid values are %v", v, AllowedClusterDisplayStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ClusterDisplayStatus) IsValid() bool {
	for _, existing := range AllowedClusterDisplayStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ClusterDisplayStatus value
func (v ClusterDisplayStatus) Ptr() *ClusterDisplayStatus {
	return &v
}

type NullableClusterDisplayStatus struct {
	value *ClusterDisplayStatus
	isSet bool
}

func (v NullableClusterDisplayStatus) Get() *ClusterDisplayStatus {
	return v.value
}

func (v *NullableClusterDisplayStatus) Set(val *ClusterDisplayStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterDisplayStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterDisplayStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterDisplayStatus(val *ClusterDisplayStatus) *NullableClusterDisplayStatus {
	return &NullableClusterDisplayStatus{value: val, isSet: true}
}

func (v NullableClusterDisplayStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterDisplayStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
