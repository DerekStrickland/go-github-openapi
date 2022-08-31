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

// Actor Actor
type Actor struct {
	Id int32 `json:"id"`
	Login string `json:"login"`
	DisplayLogin *string `json:"display_login,omitempty"`
	GravatarId NullableString `json:"gravatar_id"`
	Url string `json:"url"`
	AvatarUrl string `json:"avatar_url"`
}

// NewActor instantiates a new Actor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActor(id int32, login string, gravatarId NullableString, url string, avatarUrl string) *Actor {
	this := Actor{}
	this.Id = id
	this.Login = login
	this.GravatarId = gravatarId
	this.Url = url
	this.AvatarUrl = avatarUrl
	return &this
}

// NewActorWithDefaults instantiates a new Actor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActorWithDefaults() *Actor {
	this := Actor{}
	return &this
}

// GetId returns the Id field value
func (o *Actor) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Actor) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Actor) SetId(v int32) {
	o.Id = v
}

// GetLogin returns the Login field value
func (o *Actor) GetLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Login
}

// GetLoginOk returns a tuple with the Login field value
// and a boolean to check if the value has been set.
func (o *Actor) GetLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Login, true
}

// SetLogin sets field value
func (o *Actor) SetLogin(v string) {
	o.Login = v
}

// GetDisplayLogin returns the DisplayLogin field value if set, zero value otherwise.
func (o *Actor) GetDisplayLogin() string {
	if o == nil || o.DisplayLogin == nil {
		var ret string
		return ret
	}
	return *o.DisplayLogin
}

// GetDisplayLoginOk returns a tuple with the DisplayLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actor) GetDisplayLoginOk() (*string, bool) {
	if o == nil || o.DisplayLogin == nil {
		return nil, false
	}
	return o.DisplayLogin, true
}

// HasDisplayLogin returns a boolean if a field has been set.
func (o *Actor) HasDisplayLogin() bool {
	if o != nil && o.DisplayLogin != nil {
		return true
	}

	return false
}

// SetDisplayLogin gets a reference to the given string and assigns it to the DisplayLogin field.
func (o *Actor) SetDisplayLogin(v string) {
	o.DisplayLogin = &v
}

// GetGravatarId returns the GravatarId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Actor) GetGravatarId() string {
	if o == nil || o.GravatarId.Get() == nil {
		var ret string
		return ret
	}

	return *o.GravatarId.Get()
}

// GetGravatarIdOk returns a tuple with the GravatarId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Actor) GetGravatarIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GravatarId.Get(), o.GravatarId.IsSet()
}

// SetGravatarId sets field value
func (o *Actor) SetGravatarId(v string) {
	o.GravatarId.Set(&v)
}

// GetUrl returns the Url field value
func (o *Actor) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Actor) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Actor) SetUrl(v string) {
	o.Url = v
}

// GetAvatarUrl returns the AvatarUrl field value
func (o *Actor) GetAvatarUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value
// and a boolean to check if the value has been set.
func (o *Actor) GetAvatarUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvatarUrl, true
}

// SetAvatarUrl sets field value
func (o *Actor) SetAvatarUrl(v string) {
	o.AvatarUrl = v
}

func (o Actor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["login"] = o.Login
	}
	if o.DisplayLogin != nil {
		toSerialize["display_login"] = o.DisplayLogin
	}
	if true {
		toSerialize["gravatar_id"] = o.GravatarId.Get()
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	return json.Marshal(toSerialize)
}

type NullableActor struct {
	value *Actor
	isSet bool
}

func (v NullableActor) Get() *Actor {
	return v.value
}

func (v *NullableActor) Set(val *Actor) {
	v.value = val
	v.isSet = true
}

func (v NullableActor) IsSet() bool {
	return v.isSet
}

func (v *NullableActor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActor(val *Actor) *NullableActor {
	return &NullableActor{value: val, isSet: true}
}

func (v NullableActor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


