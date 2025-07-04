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

// checks if the CentralTiKVNode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CentralTiKVNode{}

// CentralTiKVNode struct for CentralTiKVNode
type CentralTiKVNode struct {
	Name      string                `json:"name"`
	Status    CentralTiKVNodeStatus `json:"status"`
	Az        string                `json:"az"`
	Compute   CentralCompute        `json:"compute"`
	StorageGi int64                 `json:"storage_gi"`
	Type      CentralTiKVNodeType   `json:"type"`
}

type _CentralTiKVNode CentralTiKVNode

// NewCentralTiKVNode instantiates a new CentralTiKVNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCentralTiKVNode(name string, status CentralTiKVNodeStatus, az string, compute CentralCompute, storageGi int64, type_ CentralTiKVNodeType) *CentralTiKVNode {
	this := CentralTiKVNode{}
	this.Name = name
	this.Status = status
	this.Az = az
	this.Compute = compute
	this.StorageGi = storageGi
	this.Type = type_
	return &this
}

// NewCentralTiKVNodeWithDefaults instantiates a new CentralTiKVNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCentralTiKVNodeWithDefaults() *CentralTiKVNode {
	this := CentralTiKVNode{}
	var status CentralTiKVNodeStatus = CENTRALTIKVNODESTATUS_AVAILABLE
	this.Status = status
	var type_ CentralTiKVNodeType = CENTRALTIKVNODETYPE_STORAGE
	this.Type = type_
	return &this
}

// GetName returns the Name field value
func (o *CentralTiKVNode) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CentralTiKVNode) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CentralTiKVNode) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *CentralTiKVNode) GetStatus() CentralTiKVNodeStatus {
	if o == nil {
		var ret CentralTiKVNodeStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CentralTiKVNode) GetStatusOk() (*CentralTiKVNodeStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CentralTiKVNode) SetStatus(v CentralTiKVNodeStatus) {
	o.Status = v
}

// GetAz returns the Az field value
func (o *CentralTiKVNode) GetAz() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Az
}

// GetAzOk returns a tuple with the Az field value
// and a boolean to check if the value has been set.
func (o *CentralTiKVNode) GetAzOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Az, true
}

// SetAz sets field value
func (o *CentralTiKVNode) SetAz(v string) {
	o.Az = v
}

// GetCompute returns the Compute field value
func (o *CentralTiKVNode) GetCompute() CentralCompute {
	if o == nil {
		var ret CentralCompute
		return ret
	}

	return o.Compute
}

// GetComputeOk returns a tuple with the Compute field value
// and a boolean to check if the value has been set.
func (o *CentralTiKVNode) GetComputeOk() (*CentralCompute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compute, true
}

// SetCompute sets field value
func (o *CentralTiKVNode) SetCompute(v CentralCompute) {
	o.Compute = v
}

// GetStorageGi returns the StorageGi field value
func (o *CentralTiKVNode) GetStorageGi() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StorageGi
}

// GetStorageGiOk returns a tuple with the StorageGi field value
// and a boolean to check if the value has been set.
func (o *CentralTiKVNode) GetStorageGiOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageGi, true
}

// SetStorageGi sets field value
func (o *CentralTiKVNode) SetStorageGi(v int64) {
	o.StorageGi = v
}

// GetType returns the Type field value
func (o *CentralTiKVNode) GetType() CentralTiKVNodeType {
	if o == nil {
		var ret CentralTiKVNodeType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CentralTiKVNode) GetTypeOk() (*CentralTiKVNodeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CentralTiKVNode) SetType(v CentralTiKVNodeType) {
	o.Type = v
}

func (o CentralTiKVNode) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CentralTiKVNode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	toSerialize["az"] = o.Az
	toSerialize["compute"] = o.Compute
	toSerialize["storage_gi"] = o.StorageGi
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *CentralTiKVNode) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"status",
		"az",
		"compute",
		"storage_gi",
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

	varCentralTiKVNode := _CentralTiKVNode{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCentralTiKVNode)

	if err != nil {
		return err
	}

	*o = CentralTiKVNode(varCentralTiKVNode)

	return err
}

type NullableCentralTiKVNode struct {
	value *CentralTiKVNode
	isSet bool
}

func (v NullableCentralTiKVNode) Get() *CentralTiKVNode {
	return v.value
}

func (v *NullableCentralTiKVNode) Set(val *CentralTiKVNode) {
	v.value = val
	v.isSet = true
}

func (v NullableCentralTiKVNode) IsSet() bool {
	return v.isSet
}

func (v *NullableCentralTiKVNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCentralTiKVNode(val *CentralTiKVNode) *NullableCentralTiKVNode {
	return &NullableCentralTiKVNode{value: val, isSet: true}
}

func (v NullableCentralTiKVNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCentralTiKVNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
