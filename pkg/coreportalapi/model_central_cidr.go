/*
gRPC for our DBaaS central service

notably uses gRPC-Gateway with OpenAPI

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package coreportalapi

import (
	"encoding/json"
)

// checks if the CentralCidr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CentralCidr{}

// CentralCidr Message for describing a CIDR.
type CentralCidr struct {
	// Output_only. The unique ID of the CIDR.
	Id *string `json:"id,omitempty"`
	// Required. The ID of the project.
	ProjectId *string `json:"project_id,omitempty"`
	// Required. The CIDR setting by user.
	Cidr     *string                          `json:"cidr,omitempty"`
	Region   *CentralCloudRegion              `json:"region,omitempty"`
	Provider *CentralCloudRegionCloudProvider `json:"provider,omitempty"`
	State    *CentralCidrState                `json:"state,omitempty"`
	// Output_only. The ID of the VPC that the CIDR belongs to.
	VpcId *string `json:"vpc_id,omitempty"`
}

// NewCentralCidr instantiates a new CentralCidr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCentralCidr() *CentralCidr {
	this := CentralCidr{}
	var provider CentralCloudRegionCloudProvider = CENTRALCLOUDREGIONCLOUDPROVIDER_CLOUD_PROVIDER_UNSPECIFIED
	this.Provider = &provider
	var state CentralCidrState = CENTRALCIDRSTATE_STATE_UNSPECIFIED
	this.State = &state
	return &this
}

// NewCentralCidrWithDefaults instantiates a new CentralCidr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCentralCidrWithDefaults() *CentralCidr {
	this := CentralCidr{}
	var provider CentralCloudRegionCloudProvider = CENTRALCLOUDREGIONCLOUDPROVIDER_CLOUD_PROVIDER_UNSPECIFIED
	this.Provider = &provider
	var state CentralCidrState = CENTRALCIDRSTATE_STATE_UNSPECIFIED
	this.State = &state
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CentralCidr) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CentralCidr) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CentralCidr) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CentralCidr) SetId(v string) {
	o.Id = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *CentralCidr) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CentralCidr) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *CentralCidr) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *CentralCidr) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetCidr returns the Cidr field value if set, zero value otherwise.
func (o *CentralCidr) GetCidr() string {
	if o == nil || IsNil(o.Cidr) {
		var ret string
		return ret
	}
	return *o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CentralCidr) GetCidrOk() (*string, bool) {
	if o == nil || IsNil(o.Cidr) {
		return nil, false
	}
	return o.Cidr, true
}

// HasCidr returns a boolean if a field has been set.
func (o *CentralCidr) HasCidr() bool {
	if o != nil && !IsNil(o.Cidr) {
		return true
	}

	return false
}

// SetCidr gets a reference to the given string and assigns it to the Cidr field.
func (o *CentralCidr) SetCidr(v string) {
	o.Cidr = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CentralCidr) GetRegion() CentralCloudRegion {
	if o == nil || IsNil(o.Region) {
		var ret CentralCloudRegion
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CentralCidr) GetRegionOk() (*CentralCloudRegion, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CentralCidr) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given CentralCloudRegion and assigns it to the Region field.
func (o *CentralCidr) SetRegion(v CentralCloudRegion) {
	o.Region = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *CentralCidr) GetProvider() CentralCloudRegionCloudProvider {
	if o == nil || IsNil(o.Provider) {
		var ret CentralCloudRegionCloudProvider
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CentralCidr) GetProviderOk() (*CentralCloudRegionCloudProvider, bool) {
	if o == nil || IsNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *CentralCidr) HasProvider() bool {
	if o != nil && !IsNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given CentralCloudRegionCloudProvider and assigns it to the Provider field.
func (o *CentralCidr) SetProvider(v CentralCloudRegionCloudProvider) {
	o.Provider = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CentralCidr) GetState() CentralCidrState {
	if o == nil || IsNil(o.State) {
		var ret CentralCidrState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CentralCidr) GetStateOk() (*CentralCidrState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CentralCidr) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given CentralCidrState and assigns it to the State field.
func (o *CentralCidr) SetState(v CentralCidrState) {
	o.State = &v
}

// GetVpcId returns the VpcId field value if set, zero value otherwise.
func (o *CentralCidr) GetVpcId() string {
	if o == nil || IsNil(o.VpcId) {
		var ret string
		return ret
	}
	return *o.VpcId
}

// GetVpcIdOk returns a tuple with the VpcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CentralCidr) GetVpcIdOk() (*string, bool) {
	if o == nil || IsNil(o.VpcId) {
		return nil, false
	}
	return o.VpcId, true
}

// HasVpcId returns a boolean if a field has been set.
func (o *CentralCidr) HasVpcId() bool {
	if o != nil && !IsNil(o.VpcId) {
		return true
	}

	return false
}

// SetVpcId gets a reference to the given string and assigns it to the VpcId field.
func (o *CentralCidr) SetVpcId(v string) {
	o.VpcId = &v
}

func (o CentralCidr) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CentralCidr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ProjectId) {
		toSerialize["project_id"] = o.ProjectId
	}
	if !IsNil(o.Cidr) {
		toSerialize["cidr"] = o.Cidr
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.VpcId) {
		toSerialize["vpc_id"] = o.VpcId
	}
	return toSerialize, nil
}

type NullableCentralCidr struct {
	value *CentralCidr
	isSet bool
}

func (v NullableCentralCidr) Get() *CentralCidr {
	return v.value
}

func (v *NullableCentralCidr) Set(val *CentralCidr) {
	v.value = val
	v.isSet = true
}

func (v NullableCentralCidr) IsSet() bool {
	return v.isSet
}

func (v *NullableCentralCidr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCentralCidr(val *CentralCidr) *NullableCentralCidr {
	return &NullableCentralCidr{value: val, isSet: true}
}

func (v NullableCentralCidr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCentralCidr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
