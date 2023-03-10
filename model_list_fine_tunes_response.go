/*
OpenAI API

APIs for sampling from and fine-tuning language models

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ListFineTunesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListFineTunesResponse{}

// ListFineTunesResponse struct for ListFineTunesResponse
type ListFineTunesResponse struct {
	Object string `json:"object"`
	Data []FineTune `json:"data"`
}

// NewListFineTunesResponse instantiates a new ListFineTunesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFineTunesResponse(object string, data []FineTune) *ListFineTunesResponse {
	this := ListFineTunesResponse{}
	this.Object = object
	this.Data = data
	return &this
}

// NewListFineTunesResponseWithDefaults instantiates a new ListFineTunesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFineTunesResponseWithDefaults() *ListFineTunesResponse {
	this := ListFineTunesResponse{}
	return &this
}

// GetObject returns the Object field value
func (o *ListFineTunesResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *ListFineTunesResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *ListFineTunesResponse) SetObject(v string) {
	o.Object = v
}

// GetData returns the Data field value
func (o *ListFineTunesResponse) GetData() []FineTune {
	if o == nil {
		var ret []FineTune
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListFineTunesResponse) GetDataOk() ([]FineTune, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListFineTunesResponse) SetData(v []FineTune) {
	o.Data = v
}

func (o ListFineTunesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListFineTunesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object"] = o.Object
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableListFineTunesResponse struct {
	value *ListFineTunesResponse
	isSet bool
}

func (v NullableListFineTunesResponse) Get() *ListFineTunesResponse {
	return v.value
}

func (v *NullableListFineTunesResponse) Set(val *ListFineTunesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListFineTunesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListFineTunesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFineTunesResponse(val *ListFineTunesResponse) *NullableListFineTunesResponse {
	return &NullableListFineTunesResponse{value: val, isSet: true}
}

func (v NullableListFineTunesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFineTunesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


