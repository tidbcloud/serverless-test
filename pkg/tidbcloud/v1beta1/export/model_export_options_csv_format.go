/*
TiDB Cloud Serverless Export Open API

TiDB Cloud Serverless Export Open API

API version: v1beta1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package export

import (
	"encoding/json"
)

// checks if the ExportOptionsCSVFormat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportOptionsCSVFormat{}

// ExportOptionsCSVFormat struct for ExportOptionsCSVFormat
type ExportOptionsCSVFormat struct {
	// Separator of each value in CSV files. It is recommended to use '|+|' or other uncommon character combinations. Default is ','.
	Separator *string `json:"separator,omitempty"`
	// Delimiter of string type variables in CSV files. Default is '\"'.
	Delimiter NullableString `json:"delimiter,omitempty"`
	// Representation of null values in CSV files. Default is \"\\N\".
	NullValue NullableString `json:"nullValue,omitempty"`
	// Export CSV files of the tables without header. Default is false.
	SkipHeader           *bool `json:"skipHeader,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExportOptionsCSVFormat ExportOptionsCSVFormat

// NewExportOptionsCSVFormat instantiates a new ExportOptionsCSVFormat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportOptionsCSVFormat() *ExportOptionsCSVFormat {
	this := ExportOptionsCSVFormat{}
	return &this
}

// NewExportOptionsCSVFormatWithDefaults instantiates a new ExportOptionsCSVFormat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportOptionsCSVFormatWithDefaults() *ExportOptionsCSVFormat {
	this := ExportOptionsCSVFormat{}
	return &this
}

// GetSeparator returns the Separator field value if set, zero value otherwise.
func (o *ExportOptionsCSVFormat) GetSeparator() string {
	if o == nil || IsNil(o.Separator) {
		var ret string
		return ret
	}
	return *o.Separator
}

// GetSeparatorOk returns a tuple with the Separator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportOptionsCSVFormat) GetSeparatorOk() (*string, bool) {
	if o == nil || IsNil(o.Separator) {
		return nil, false
	}
	return o.Separator, true
}

// HasSeparator returns a boolean if a field has been set.
func (o *ExportOptionsCSVFormat) HasSeparator() bool {
	if o != nil && !IsNil(o.Separator) {
		return true
	}

	return false
}

// SetSeparator gets a reference to the given string and assigns it to the Separator field.
func (o *ExportOptionsCSVFormat) SetSeparator(v string) {
	o.Separator = &v
}

// GetDelimiter returns the Delimiter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportOptionsCSVFormat) GetDelimiter() string {
	if o == nil || IsNil(o.Delimiter.Get()) {
		var ret string
		return ret
	}
	return *o.Delimiter.Get()
}

// GetDelimiterOk returns a tuple with the Delimiter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportOptionsCSVFormat) GetDelimiterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Delimiter.Get(), o.Delimiter.IsSet()
}

// HasDelimiter returns a boolean if a field has been set.
func (o *ExportOptionsCSVFormat) HasDelimiter() bool {
	if o != nil && o.Delimiter.IsSet() {
		return true
	}

	return false
}

// SetDelimiter gets a reference to the given NullableString and assigns it to the Delimiter field.
func (o *ExportOptionsCSVFormat) SetDelimiter(v string) {
	o.Delimiter.Set(&v)
}

// SetDelimiterNil sets the value for Delimiter to be an explicit nil
func (o *ExportOptionsCSVFormat) SetDelimiterNil() {
	o.Delimiter.Set(nil)
}

// UnsetDelimiter ensures that no value is present for Delimiter, not even an explicit nil
func (o *ExportOptionsCSVFormat) UnsetDelimiter() {
	o.Delimiter.Unset()
}

// GetNullValue returns the NullValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExportOptionsCSVFormat) GetNullValue() string {
	if o == nil || IsNil(o.NullValue.Get()) {
		var ret string
		return ret
	}
	return *o.NullValue.Get()
}

// GetNullValueOk returns a tuple with the NullValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExportOptionsCSVFormat) GetNullValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NullValue.Get(), o.NullValue.IsSet()
}

// HasNullValue returns a boolean if a field has been set.
func (o *ExportOptionsCSVFormat) HasNullValue() bool {
	if o != nil && o.NullValue.IsSet() {
		return true
	}

	return false
}

// SetNullValue gets a reference to the given NullableString and assigns it to the NullValue field.
func (o *ExportOptionsCSVFormat) SetNullValue(v string) {
	o.NullValue.Set(&v)
}

// SetNullValueNil sets the value for NullValue to be an explicit nil
func (o *ExportOptionsCSVFormat) SetNullValueNil() {
	o.NullValue.Set(nil)
}

// UnsetNullValue ensures that no value is present for NullValue, not even an explicit nil
func (o *ExportOptionsCSVFormat) UnsetNullValue() {
	o.NullValue.Unset()
}

// GetSkipHeader returns the SkipHeader field value if set, zero value otherwise.
func (o *ExportOptionsCSVFormat) GetSkipHeader() bool {
	if o == nil || IsNil(o.SkipHeader) {
		var ret bool
		return ret
	}
	return *o.SkipHeader
}

// GetSkipHeaderOk returns a tuple with the SkipHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportOptionsCSVFormat) GetSkipHeaderOk() (*bool, bool) {
	if o == nil || IsNil(o.SkipHeader) {
		return nil, false
	}
	return o.SkipHeader, true
}

// HasSkipHeader returns a boolean if a field has been set.
func (o *ExportOptionsCSVFormat) HasSkipHeader() bool {
	if o != nil && !IsNil(o.SkipHeader) {
		return true
	}

	return false
}

// SetSkipHeader gets a reference to the given bool and assigns it to the SkipHeader field.
func (o *ExportOptionsCSVFormat) SetSkipHeader(v bool) {
	o.SkipHeader = &v
}

func (o ExportOptionsCSVFormat) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportOptionsCSVFormat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Separator) {
		toSerialize["separator"] = o.Separator
	}
	if o.Delimiter.IsSet() {
		toSerialize["delimiter"] = o.Delimiter.Get()
	}
	if o.NullValue.IsSet() {
		toSerialize["nullValue"] = o.NullValue.Get()
	}
	if !IsNil(o.SkipHeader) {
		toSerialize["skipHeader"] = o.SkipHeader
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExportOptionsCSVFormat) UnmarshalJSON(data []byte) (err error) {
	varExportOptionsCSVFormat := _ExportOptionsCSVFormat{}

	err = json.Unmarshal(data, &varExportOptionsCSVFormat)

	if err != nil {
		return err
	}

	*o = ExportOptionsCSVFormat(varExportOptionsCSVFormat)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "separator")
		delete(additionalProperties, "delimiter")
		delete(additionalProperties, "nullValue")
		delete(additionalProperties, "skipHeader")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExportOptionsCSVFormat struct {
	value *ExportOptionsCSVFormat
	isSet bool
}

func (v NullableExportOptionsCSVFormat) Get() *ExportOptionsCSVFormat {
	return v.value
}

func (v *NullableExportOptionsCSVFormat) Set(val *ExportOptionsCSVFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableExportOptionsCSVFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableExportOptionsCSVFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportOptionsCSVFormat(val *ExportOptionsCSVFormat) *NullableExportOptionsCSVFormat {
	return &NullableExportOptionsCSVFormat{value: val, isSet: true}
}

func (v NullableExportOptionsCSVFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportOptionsCSVFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
