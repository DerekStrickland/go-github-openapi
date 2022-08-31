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

// GetLicenseSyncStatusServerInstancesInnerLastSync struct for GetLicenseSyncStatusServerInstancesInnerLastSync
type GetLicenseSyncStatusServerInstancesInnerLastSync struct {
	Date *string `json:"date,omitempty"`
	Status *string `json:"status,omitempty"`
	Error *string `json:"error,omitempty"`
}

// NewGetLicenseSyncStatusServerInstancesInnerLastSync instantiates a new GetLicenseSyncStatusServerInstancesInnerLastSync object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLicenseSyncStatusServerInstancesInnerLastSync() *GetLicenseSyncStatusServerInstancesInnerLastSync {
	this := GetLicenseSyncStatusServerInstancesInnerLastSync{}
	return &this
}

// NewGetLicenseSyncStatusServerInstancesInnerLastSyncWithDefaults instantiates a new GetLicenseSyncStatusServerInstancesInnerLastSync object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLicenseSyncStatusServerInstancesInnerLastSyncWithDefaults() *GetLicenseSyncStatusServerInstancesInnerLastSync {
	this := GetLicenseSyncStatusServerInstancesInnerLastSync{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *GetLicenseSyncStatusServerInstancesInnerLastSync) GetDate() string {
	if o == nil || o.Date == nil {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLicenseSyncStatusServerInstancesInnerLastSync) GetDateOk() (*string, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *GetLicenseSyncStatusServerInstancesInnerLastSync) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *GetLicenseSyncStatusServerInstancesInnerLastSync) SetDate(v string) {
	o.Date = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetLicenseSyncStatusServerInstancesInnerLastSync) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLicenseSyncStatusServerInstancesInnerLastSync) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetLicenseSyncStatusServerInstancesInnerLastSync) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetLicenseSyncStatusServerInstancesInnerLastSync) SetStatus(v string) {
	o.Status = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *GetLicenseSyncStatusServerInstancesInnerLastSync) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLicenseSyncStatusServerInstancesInnerLastSync) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *GetLicenseSyncStatusServerInstancesInnerLastSync) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *GetLicenseSyncStatusServerInstancesInnerLastSync) SetError(v string) {
	o.Error = &v
}

func (o GetLicenseSyncStatusServerInstancesInnerLastSync) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableGetLicenseSyncStatusServerInstancesInnerLastSync struct {
	value *GetLicenseSyncStatusServerInstancesInnerLastSync
	isSet bool
}

func (v NullableGetLicenseSyncStatusServerInstancesInnerLastSync) Get() *GetLicenseSyncStatusServerInstancesInnerLastSync {
	return v.value
}

func (v *NullableGetLicenseSyncStatusServerInstancesInnerLastSync) Set(val *GetLicenseSyncStatusServerInstancesInnerLastSync) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLicenseSyncStatusServerInstancesInnerLastSync) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLicenseSyncStatusServerInstancesInnerLastSync) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLicenseSyncStatusServerInstancesInnerLastSync(val *GetLicenseSyncStatusServerInstancesInnerLastSync) *NullableGetLicenseSyncStatusServerInstancesInnerLastSync {
	return &NullableGetLicenseSyncStatusServerInstancesInnerLastSync{value: val, isSet: true}
}

func (v NullableGetLicenseSyncStatusServerInstancesInnerLastSync) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLicenseSyncStatusServerInstancesInnerLastSync) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


