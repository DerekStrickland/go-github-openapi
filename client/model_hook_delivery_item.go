/*
GitHub v3 REST API

GitHub's v3 REST API.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github

import (
	"encoding/json"
	"time"
)

// HookDeliveryItem Delivery made by a webhook, without request and response information.
type HookDeliveryItem struct {
	// Unique identifier of the webhook delivery.
	Id int32 `json:"id"`
	// Unique identifier for the event (shared with all deliveries for all webhooks that subscribe to this event).
	Guid string `json:"guid"`
	// Time when the webhook delivery occurred.
	DeliveredAt time.Time `json:"delivered_at"`
	// Whether the webhook delivery is a redelivery.
	Redelivery bool `json:"redelivery"`
	// Time spent delivering.
	Duration float32 `json:"duration"`
	// Describes the response returned after attempting the delivery.
	Status string `json:"status"`
	// Status code received when delivery was made.
	StatusCode int32 `json:"status_code"`
	// The event that triggered the delivery.
	Event string `json:"event"`
	// The type of activity for the event that triggered the delivery.
	Action NullableString `json:"action"`
	// The id of the GitHub App installation associated with this event.
	InstallationId NullableInt32 `json:"installation_id"`
	// The id of the repository associated with this event.
	RepositoryId NullableInt32 `json:"repository_id"`
}

// NewHookDeliveryItem instantiates a new HookDeliveryItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHookDeliveryItem(id int32, guid string, deliveredAt time.Time, redelivery bool, duration float32, status string, statusCode int32, event string, action NullableString, installationId NullableInt32, repositoryId NullableInt32) *HookDeliveryItem {
	this := HookDeliveryItem{}
	this.Id = id
	this.Guid = guid
	this.DeliveredAt = deliveredAt
	this.Redelivery = redelivery
	this.Duration = duration
	this.Status = status
	this.StatusCode = statusCode
	this.Event = event
	this.Action = action
	this.InstallationId = installationId
	this.RepositoryId = repositoryId
	return &this
}

// NewHookDeliveryItemWithDefaults instantiates a new HookDeliveryItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHookDeliveryItemWithDefaults() *HookDeliveryItem {
	this := HookDeliveryItem{}
	return &this
}

// GetId returns the Id field value
func (o *HookDeliveryItem) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *HookDeliveryItem) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *HookDeliveryItem) SetId(v int32) {
	o.Id = v
}

// GetGuid returns the Guid field value
func (o *HookDeliveryItem) GetGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Guid
}

// GetGuidOk returns a tuple with the Guid field value
// and a boolean to check if the value has been set.
func (o *HookDeliveryItem) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Guid, true
}

// SetGuid sets field value
func (o *HookDeliveryItem) SetGuid(v string) {
	o.Guid = v
}

// GetDeliveredAt returns the DeliveredAt field value
func (o *HookDeliveryItem) GetDeliveredAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DeliveredAt
}

// GetDeliveredAtOk returns a tuple with the DeliveredAt field value
// and a boolean to check if the value has been set.
func (o *HookDeliveryItem) GetDeliveredAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeliveredAt, true
}

// SetDeliveredAt sets field value
func (o *HookDeliveryItem) SetDeliveredAt(v time.Time) {
	o.DeliveredAt = v
}

// GetRedelivery returns the Redelivery field value
func (o *HookDeliveryItem) GetRedelivery() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Redelivery
}

// GetRedeliveryOk returns a tuple with the Redelivery field value
// and a boolean to check if the value has been set.
func (o *HookDeliveryItem) GetRedeliveryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Redelivery, true
}

// SetRedelivery sets field value
func (o *HookDeliveryItem) SetRedelivery(v bool) {
	o.Redelivery = v
}

// GetDuration returns the Duration field value
func (o *HookDeliveryItem) GetDuration() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *HookDeliveryItem) GetDurationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *HookDeliveryItem) SetDuration(v float32) {
	o.Duration = v
}

// GetStatus returns the Status field value
func (o *HookDeliveryItem) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *HookDeliveryItem) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *HookDeliveryItem) SetStatus(v string) {
	o.Status = v
}

// GetStatusCode returns the StatusCode field value
func (o *HookDeliveryItem) GetStatusCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *HookDeliveryItem) GetStatusCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *HookDeliveryItem) SetStatusCode(v int32) {
	o.StatusCode = v
}

// GetEvent returns the Event field value
func (o *HookDeliveryItem) GetEvent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *HookDeliveryItem) GetEventOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *HookDeliveryItem) SetEvent(v string) {
	o.Event = v
}

// GetAction returns the Action field value
// If the value is explicit nil, the zero value for string will be returned
func (o *HookDeliveryItem) GetAction() string {
	if o == nil || o.Action.Get() == nil {
		var ret string
		return ret
	}

	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HookDeliveryItem) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// SetAction sets field value
func (o *HookDeliveryItem) SetAction(v string) {
	o.Action.Set(&v)
}

// GetInstallationId returns the InstallationId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *HookDeliveryItem) GetInstallationId() int32 {
	if o == nil || o.InstallationId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.InstallationId.Get()
}

// GetInstallationIdOk returns a tuple with the InstallationId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HookDeliveryItem) GetInstallationIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.InstallationId.Get(), o.InstallationId.IsSet()
}

// SetInstallationId sets field value
func (o *HookDeliveryItem) SetInstallationId(v int32) {
	o.InstallationId.Set(&v)
}

// GetRepositoryId returns the RepositoryId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *HookDeliveryItem) GetRepositoryId() int32 {
	if o == nil || o.RepositoryId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.RepositoryId.Get()
}

// GetRepositoryIdOk returns a tuple with the RepositoryId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HookDeliveryItem) GetRepositoryIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RepositoryId.Get(), o.RepositoryId.IsSet()
}

// SetRepositoryId sets field value
func (o *HookDeliveryItem) SetRepositoryId(v int32) {
	o.RepositoryId.Set(&v)
}

func (o HookDeliveryItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["guid"] = o.Guid
	}
	if true {
		toSerialize["delivered_at"] = o.DeliveredAt
	}
	if true {
		toSerialize["redelivery"] = o.Redelivery
	}
	if true {
		toSerialize["duration"] = o.Duration
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["status_code"] = o.StatusCode
	}
	if true {
		toSerialize["event"] = o.Event
	}
	if true {
		toSerialize["action"] = o.Action.Get()
	}
	if true {
		toSerialize["installation_id"] = o.InstallationId.Get()
	}
	if true {
		toSerialize["repository_id"] = o.RepositoryId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableHookDeliveryItem struct {
	value *HookDeliveryItem
	isSet bool
}

func (v NullableHookDeliveryItem) Get() *HookDeliveryItem {
	return v.value
}

func (v *NullableHookDeliveryItem) Set(val *HookDeliveryItem) {
	v.value = val
	v.isSet = true
}

func (v NullableHookDeliveryItem) IsSet() bool {
	return v.isSet
}

func (v *NullableHookDeliveryItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHookDeliveryItem(val *HookDeliveryItem) *NullableHookDeliveryItem {
	return &NullableHookDeliveryItem{value: val, isSet: true}
}

func (v NullableHookDeliveryItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHookDeliveryItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

