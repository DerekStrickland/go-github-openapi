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

// SimpleUser Simple User
type SimpleUser struct {
	Name NullableString `json:"name,omitempty"`
	Email NullableString `json:"email,omitempty"`
	Login string `json:"login"`
	Id int32 `json:"id"`
	NodeId string `json:"node_id"`
	AvatarUrl string `json:"avatar_url"`
	GravatarId NullableString `json:"gravatar_id"`
	Url string `json:"url"`
	HtmlUrl string `json:"html_url"`
	FollowersUrl string `json:"followers_url"`
	FollowingUrl string `json:"following_url"`
	GistsUrl string `json:"gists_url"`
	StarredUrl string `json:"starred_url"`
	SubscriptionsUrl string `json:"subscriptions_url"`
	OrganizationsUrl string `json:"organizations_url"`
	ReposUrl string `json:"repos_url"`
	EventsUrl string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type string `json:"type"`
	SiteAdmin bool `json:"site_admin"`
	StarredAt *string `json:"starred_at,omitempty"`
}

// NewSimpleUser instantiates a new SimpleUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleUser(login string, id int32, nodeId string, avatarUrl string, gravatarId NullableString, url string, htmlUrl string, followersUrl string, followingUrl string, gistsUrl string, starredUrl string, subscriptionsUrl string, organizationsUrl string, reposUrl string, eventsUrl string, receivedEventsUrl string, type_ string, siteAdmin bool) *SimpleUser {
	this := SimpleUser{}
	this.Login = login
	this.Id = id
	this.NodeId = nodeId
	this.AvatarUrl = avatarUrl
	this.GravatarId = gravatarId
	this.Url = url
	this.HtmlUrl = htmlUrl
	this.FollowersUrl = followersUrl
	this.FollowingUrl = followingUrl
	this.GistsUrl = gistsUrl
	this.StarredUrl = starredUrl
	this.SubscriptionsUrl = subscriptionsUrl
	this.OrganizationsUrl = organizationsUrl
	this.ReposUrl = reposUrl
	this.EventsUrl = eventsUrl
	this.ReceivedEventsUrl = receivedEventsUrl
	this.Type = type_
	this.SiteAdmin = siteAdmin
	return &this
}

// NewSimpleUserWithDefaults instantiates a new SimpleUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleUserWithDefaults() *SimpleUser {
	this := SimpleUser{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimpleUser) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimpleUser) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *SimpleUser) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *SimpleUser) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *SimpleUser) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *SimpleUser) UnsetName() {
	o.Name.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SimpleUser) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimpleUser) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *SimpleUser) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *SimpleUser) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *SimpleUser) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *SimpleUser) UnsetEmail() {
	o.Email.Unset()
}

// GetLogin returns the Login field value
func (o *SimpleUser) GetLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Login
}

// GetLoginOk returns a tuple with the Login field value
// and a boolean to check if the value has been set.
func (o *SimpleUser) GetLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Login, true
}

// SetLogin sets field value
func (o *SimpleUser) SetLogin(v string) {
	o.Login = v
}

// GetId returns the Id field value
func (o *SimpleUser) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SimpleUser) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SimpleUser) SetId(v int32) {
	o.Id = v
}

// GetNodeId returns the NodeId field value
func (o *SimpleUser) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *SimpleUser) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *SimpleUser) SetNodeId(v string) {
	o.NodeId = v
}

// GetAvatarUrl returns the AvatarUrl field value
func (o *SimpleUser) GetAvatarUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value
// and a boolean to check if the value has been set.
func (o *SimpleUser) GetAvatarUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvatarUrl, true
}

// SetAvatarUrl sets field value
func (o *SimpleUser) SetAvatarUrl(v string) {
	o.AvatarUrl = v
}

// GetGravatarId returns the GravatarId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *SimpleUser) GetGravatarId() string {
	if o == nil || o.GravatarId.Get() == nil {
		var ret string
		return ret
	}

	return *o.GravatarId.Get()
}

// GetGravatarIdOk returns a tuple with the GravatarId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SimpleUser) GetGravatarIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GravatarId.Get(), o.GravatarId.IsSet()
}

// SetGravatarId sets field value
func (o *SimpleUser) SetGravatarId(v string) {
	o.GravatarId.Set(&v)
}

// GetUrl returns the Url field value
func (o *SimpleUser) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *SimpleUser) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *SimpleUser) SetUrl(v string) {
	o.Url = v
}

// GetHtmlUrl returns the HtmlUrl field value
func (o *SimpleUser) GetHtmlUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value
// and a boolean to check if the value has been set.
func (o *SimpleUser) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlUrl, true
}

// SetHtmlUrl sets field value
func (o *SimpleUser) SetHtmlUrl(v string) {
	o.HtmlUrl = v
}

// GetFollowersUrl returns the FollowersUrl field value
func (o *SimpleUser) GetFollowersUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FollowersUrl
}

// GetFollowersUrlOk returns a tuple with the FollowersUrl field value
// and a boolean to check if the value has been set.
func (o *SimpleUser) GetFollowersUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FollowersUrl, true
}

// SetFollowersUrl sets field value
func (o *SimpleUser) SetFollowersUrl(v string) {
	o.FollowersUrl = v
}

// GetFollowingUrl returns the FollowingUrl field value
func (o *SimpleUser) GetFollowingUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FollowingUrl
}

// GetFollowingUrlOk returns a tuple with the FollowingUrl field value
// and a boolean to check if the value has been set.
func (o *SimpleUser) GetFollowingUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FollowingUrl, true
}

// SetFollowingUrl sets field value
func (o *SimpleUser) SetFollowingUrl(v string) {
	o.FollowingUrl = v
}

// GetGistsUrl returns the GistsUrl field value
func (o *SimpleUser) GetGistsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GistsUrl
}

// GetGistsUrlOk returns a tuple with the GistsUrl field value
// and a boolean to check if the value has been set.
func (o *SimpleUser) GetGistsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GistsUrl, true
}

// SetGistsUrl sets field value
func (o *SimpleUser) SetGistsUrl(v string) {
	o.GistsUrl = v
}

// GetStarredUrl returns the StarredUrl field value
func (o *SimpleUser) GetStarredUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StarredUrl
}

// GetStarredUrlOk returns a tuple with the StarredUrl field value
// and a boolean to check if the value has been set.
func (o *SimpleUser) GetStarredUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StarredUrl, true
}

// SetStarredUrl sets field value
func (o *SimpleUser) SetStarredUrl(v string) {
	o.StarredUrl = v
}

// GetSubscriptionsUrl returns the SubscriptionsUrl field value
func (o *SimpleUser) GetSubscriptionsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubscriptionsUrl
}

// GetSubscriptionsUrlOk returns a tuple with the SubscriptionsUrl field value
// and a boolean to check if the value has been set.
func (o *SimpleUser) GetSubscriptionsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubscriptionsUrl, true
}

// SetSubscriptionsUrl sets field value
func (o *SimpleUser) SetSubscriptionsUrl(v string) {
	o.SubscriptionsUrl = v
}

// GetOrganizationsUrl returns the OrganizationsUrl field value
func (o *SimpleUser) GetOrganizationsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationsUrl
}

// GetOrganizationsUrlOk returns a tuple with the OrganizationsUrl field value
// and a boolean to check if the value has been set.
func (o *SimpleUser) GetOrganizationsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationsUrl, true
}

// SetOrganizationsUrl sets field value
func (o *SimpleUser) SetOrganizationsUrl(v string) {
	o.OrganizationsUrl = v
}

// GetReposUrl returns the ReposUrl field value
func (o *SimpleUser) GetReposUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReposUrl
}

// GetReposUrlOk returns a tuple with the ReposUrl field value
// and a boolean to check if the value has been set.
func (o *SimpleUser) GetReposUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReposUrl, true
}

// SetReposUrl sets field value
func (o *SimpleUser) SetReposUrl(v string) {
	o.ReposUrl = v
}

// GetEventsUrl returns the EventsUrl field value
func (o *SimpleUser) GetEventsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventsUrl
}

// GetEventsUrlOk returns a tuple with the EventsUrl field value
// and a boolean to check if the value has been set.
func (o *SimpleUser) GetEventsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventsUrl, true
}

// SetEventsUrl sets field value
func (o *SimpleUser) SetEventsUrl(v string) {
	o.EventsUrl = v
}

// GetReceivedEventsUrl returns the ReceivedEventsUrl field value
func (o *SimpleUser) GetReceivedEventsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReceivedEventsUrl
}

// GetReceivedEventsUrlOk returns a tuple with the ReceivedEventsUrl field value
// and a boolean to check if the value has been set.
func (o *SimpleUser) GetReceivedEventsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceivedEventsUrl, true
}

// SetReceivedEventsUrl sets field value
func (o *SimpleUser) SetReceivedEventsUrl(v string) {
	o.ReceivedEventsUrl = v
}

// GetType returns the Type field value
func (o *SimpleUser) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SimpleUser) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SimpleUser) SetType(v string) {
	o.Type = v
}

// GetSiteAdmin returns the SiteAdmin field value
func (o *SimpleUser) GetSiteAdmin() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SiteAdmin
}

// GetSiteAdminOk returns a tuple with the SiteAdmin field value
// and a boolean to check if the value has been set.
func (o *SimpleUser) GetSiteAdminOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SiteAdmin, true
}

// SetSiteAdmin sets field value
func (o *SimpleUser) SetSiteAdmin(v bool) {
	o.SiteAdmin = v
}

// GetStarredAt returns the StarredAt field value if set, zero value otherwise.
func (o *SimpleUser) GetStarredAt() string {
	if o == nil || o.StarredAt == nil {
		var ret string
		return ret
	}
	return *o.StarredAt
}

// GetStarredAtOk returns a tuple with the StarredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleUser) GetStarredAtOk() (*string, bool) {
	if o == nil || o.StarredAt == nil {
		return nil, false
	}
	return o.StarredAt, true
}

// HasStarredAt returns a boolean if a field has been set.
func (o *SimpleUser) HasStarredAt() bool {
	if o != nil && o.StarredAt != nil {
		return true
	}

	return false
}

// SetStarredAt gets a reference to the given string and assigns it to the StarredAt field.
func (o *SimpleUser) SetStarredAt(v string) {
	o.StarredAt = &v
}

func (o SimpleUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if true {
		toSerialize["login"] = o.Login
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["node_id"] = o.NodeId
	}
	if true {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	if true {
		toSerialize["gravatar_id"] = o.GravatarId.Get()
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["html_url"] = o.HtmlUrl
	}
	if true {
		toSerialize["followers_url"] = o.FollowersUrl
	}
	if true {
		toSerialize["following_url"] = o.FollowingUrl
	}
	if true {
		toSerialize["gists_url"] = o.GistsUrl
	}
	if true {
		toSerialize["starred_url"] = o.StarredUrl
	}
	if true {
		toSerialize["subscriptions_url"] = o.SubscriptionsUrl
	}
	if true {
		toSerialize["organizations_url"] = o.OrganizationsUrl
	}
	if true {
		toSerialize["repos_url"] = o.ReposUrl
	}
	if true {
		toSerialize["events_url"] = o.EventsUrl
	}
	if true {
		toSerialize["received_events_url"] = o.ReceivedEventsUrl
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["site_admin"] = o.SiteAdmin
	}
	if o.StarredAt != nil {
		toSerialize["starred_at"] = o.StarredAt
	}
	return json.Marshal(toSerialize)
}

type NullableSimpleUser struct {
	value *SimpleUser
	isSet bool
}

func (v NullableSimpleUser) Get() *SimpleUser {
	return v.value
}

func (v *NullableSimpleUser) Set(val *SimpleUser) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleUser) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleUser(val *SimpleUser) *NullableSimpleUser {
	return &NullableSimpleUser{value: val, isSet: true}
}

func (v NullableSimpleUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

