/*
GitHub v3 REST API

GitHub's v3 REST API.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github

import (
	"encoding/json"
)

// EventRepo struct for EventRepo
type EventRepo struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	Url string `json:"url"`
}

// NewEventRepo instantiates a new EventRepo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventRepo(id int32, name string, url string) *EventRepo {
	this := EventRepo{}
	this.Id = id
	this.Name = name
	this.Url = url
	return &this
}

// NewEventRepoWithDefaults instantiates a new EventRepo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventRepoWithDefaults() *EventRepo {
	this := EventRepo{}
	return &this
}

// GetId returns the Id field value
func (o *EventRepo) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EventRepo) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EventRepo) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *EventRepo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EventRepo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EventRepo) SetName(v string) {
	o.Name = v
}

// GetUrl returns the Url field value
func (o *EventRepo) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *EventRepo) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *EventRepo) SetUrl(v string) {
	o.Url = v
}

func (o EventRepo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableEventRepo struct {
	value *EventRepo
	isSet bool
}

func (v NullableEventRepo) Get() *EventRepo {
	return v.value
}

func (v *NullableEventRepo) Set(val *EventRepo) {
	v.value = val
	v.isSet = true
}

func (v NullableEventRepo) IsSet() bool {
	return v.isSet
}

func (v *NullableEventRepo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventRepo(val *EventRepo) *NullableEventRepo {
	return &NullableEventRepo{value: val, isSet: true}
}

func (v NullableEventRepo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventRepo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


