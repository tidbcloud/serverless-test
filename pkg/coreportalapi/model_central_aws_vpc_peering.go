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

// checks if the CentralAwsVpcPeering type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CentralAwsVpcPeering{}

// CentralAwsVpcPeering struct for CentralAwsVpcPeering
type CentralAwsVpcPeering struct {
	Region          string `json:"region"`
	PrivateHostZone string `json:"private_host_zone"`
	VpcPeeringId    string `json:"vpc_peering_id"`
	VpcCidr         string `json:"vpc_cidr"`
	TenantVpcCidr   string `json:"tenant_vpc_cidr"`
	TenantVpcId     string `json:"tenant_vpc_id"`
	Status          string `json:"status"`
	Id              string `json:"id"`
	RegionName      string `json:"region_name"`
}

type _CentralAwsVpcPeering CentralAwsVpcPeering

// NewCentralAwsVpcPeering instantiates a new CentralAwsVpcPeering object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCentralAwsVpcPeering(region string, privateHostZone string, vpcPeeringId string, vpcCidr string, tenantVpcCidr string, tenantVpcId string, status string, id string, regionName string) *CentralAwsVpcPeering {
	this := CentralAwsVpcPeering{}
	this.Region = region
	this.PrivateHostZone = privateHostZone
	this.VpcPeeringId = vpcPeeringId
	this.VpcCidr = vpcCidr
	this.TenantVpcCidr = tenantVpcCidr
	this.TenantVpcId = tenantVpcId
	this.Status = status
	this.Id = id
	this.RegionName = regionName
	return &this
}

// NewCentralAwsVpcPeeringWithDefaults instantiates a new CentralAwsVpcPeering object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCentralAwsVpcPeeringWithDefaults() *CentralAwsVpcPeering {
	this := CentralAwsVpcPeering{}
	return &this
}

// GetRegion returns the Region field value
func (o *CentralAwsVpcPeering) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *CentralAwsVpcPeering) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *CentralAwsVpcPeering) SetRegion(v string) {
	o.Region = v
}

// GetPrivateHostZone returns the PrivateHostZone field value
func (o *CentralAwsVpcPeering) GetPrivateHostZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateHostZone
}

// GetPrivateHostZoneOk returns a tuple with the PrivateHostZone field value
// and a boolean to check if the value has been set.
func (o *CentralAwsVpcPeering) GetPrivateHostZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateHostZone, true
}

// SetPrivateHostZone sets field value
func (o *CentralAwsVpcPeering) SetPrivateHostZone(v string) {
	o.PrivateHostZone = v
}

// GetVpcPeeringId returns the VpcPeeringId field value
func (o *CentralAwsVpcPeering) GetVpcPeeringId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VpcPeeringId
}

// GetVpcPeeringIdOk returns a tuple with the VpcPeeringId field value
// and a boolean to check if the value has been set.
func (o *CentralAwsVpcPeering) GetVpcPeeringIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VpcPeeringId, true
}

// SetVpcPeeringId sets field value
func (o *CentralAwsVpcPeering) SetVpcPeeringId(v string) {
	o.VpcPeeringId = v
}

// GetVpcCidr returns the VpcCidr field value
func (o *CentralAwsVpcPeering) GetVpcCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VpcCidr
}

// GetVpcCidrOk returns a tuple with the VpcCidr field value
// and a boolean to check if the value has been set.
func (o *CentralAwsVpcPeering) GetVpcCidrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VpcCidr, true
}

// SetVpcCidr sets field value
func (o *CentralAwsVpcPeering) SetVpcCidr(v string) {
	o.VpcCidr = v
}

// GetTenantVpcCidr returns the TenantVpcCidr field value
func (o *CentralAwsVpcPeering) GetTenantVpcCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantVpcCidr
}

// GetTenantVpcCidrOk returns a tuple with the TenantVpcCidr field value
// and a boolean to check if the value has been set.
func (o *CentralAwsVpcPeering) GetTenantVpcCidrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantVpcCidr, true
}

// SetTenantVpcCidr sets field value
func (o *CentralAwsVpcPeering) SetTenantVpcCidr(v string) {
	o.TenantVpcCidr = v
}

// GetTenantVpcId returns the TenantVpcId field value
func (o *CentralAwsVpcPeering) GetTenantVpcId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantVpcId
}

// GetTenantVpcIdOk returns a tuple with the TenantVpcId field value
// and a boolean to check if the value has been set.
func (o *CentralAwsVpcPeering) GetTenantVpcIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantVpcId, true
}

// SetTenantVpcId sets field value
func (o *CentralAwsVpcPeering) SetTenantVpcId(v string) {
	o.TenantVpcId = v
}

// GetStatus returns the Status field value
func (o *CentralAwsVpcPeering) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CentralAwsVpcPeering) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CentralAwsVpcPeering) SetStatus(v string) {
	o.Status = v
}

// GetId returns the Id field value
func (o *CentralAwsVpcPeering) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CentralAwsVpcPeering) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CentralAwsVpcPeering) SetId(v string) {
	o.Id = v
}

// GetRegionName returns the RegionName field value
func (o *CentralAwsVpcPeering) GetRegionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value
// and a boolean to check if the value has been set.
func (o *CentralAwsVpcPeering) GetRegionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegionName, true
}

// SetRegionName sets field value
func (o *CentralAwsVpcPeering) SetRegionName(v string) {
	o.RegionName = v
}

func (o CentralAwsVpcPeering) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CentralAwsVpcPeering) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["region"] = o.Region
	toSerialize["private_host_zone"] = o.PrivateHostZone
	toSerialize["vpc_peering_id"] = o.VpcPeeringId
	toSerialize["vpc_cidr"] = o.VpcCidr
	toSerialize["tenant_vpc_cidr"] = o.TenantVpcCidr
	toSerialize["tenant_vpc_id"] = o.TenantVpcId
	toSerialize["status"] = o.Status
	toSerialize["id"] = o.Id
	toSerialize["region_name"] = o.RegionName
	return toSerialize, nil
}

func (o *CentralAwsVpcPeering) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"region",
		"private_host_zone",
		"vpc_peering_id",
		"vpc_cidr",
		"tenant_vpc_cidr",
		"tenant_vpc_id",
		"status",
		"id",
		"region_name",
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

	varCentralAwsVpcPeering := _CentralAwsVpcPeering{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCentralAwsVpcPeering)

	if err != nil {
		return err
	}

	*o = CentralAwsVpcPeering(varCentralAwsVpcPeering)

	return err
}

type NullableCentralAwsVpcPeering struct {
	value *CentralAwsVpcPeering
	isSet bool
}

func (v NullableCentralAwsVpcPeering) Get() *CentralAwsVpcPeering {
	return v.value
}

func (v *NullableCentralAwsVpcPeering) Set(val *CentralAwsVpcPeering) {
	v.value = val
	v.isSet = true
}

func (v NullableCentralAwsVpcPeering) IsSet() bool {
	return v.isSet
}

func (v *NullableCentralAwsVpcPeering) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCentralAwsVpcPeering(val *CentralAwsVpcPeering) *NullableCentralAwsVpcPeering {
	return &NullableCentralAwsVpcPeering{value: val, isSet: true}
}

func (v NullableCentralAwsVpcPeering) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCentralAwsVpcPeering) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
