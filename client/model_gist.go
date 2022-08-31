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

// Gist Gist
type Gist struct {
	Url string `json:"url"`
	ForksUrl string `json:"forks_url"`
	CommitsUrl string `json:"commits_url"`
	Id string `json:"id"`
	NodeId string `json:"node_id"`
	GitPullUrl string `json:"git_pull_url"`
	GitPushUrl string `json:"git_push_url"`
	HtmlUrl string `json:"html_url"`
	Files map[string]BaseGistFilesValue `json:"files"`
	Public bool `json:"public"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Description NullableString `json:"description"`
	Comments int32 `json:"comments"`
	User NullableNullableSimpleUser `json:"user"`
	CommentsUrl string `json:"comments_url"`
	Owner NullableNullableSimpleUser `json:"owner,omitempty"`
	Truncated *bool `json:"truncated,omitempty"`
	Forks []interface{} `json:"forks,omitempty"`
	History []interface{} `json:"history,omitempty"`
}

// NewGist instantiates a new Gist object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGist(url string, forksUrl string, commitsUrl string, id string, nodeId string, gitPullUrl string, gitPushUrl string, htmlUrl string, files map[string]BaseGistFilesValue, public bool, createdAt time.Time, updatedAt time.Time, description NullableString, comments int32, user NullableNullableSimpleUser, commentsUrl string) *Gist {
	this := Gist{}
	this.Url = url
	this.ForksUrl = forksUrl
	this.CommitsUrl = commitsUrl
	this.Id = id
	this.NodeId = nodeId
	this.GitPullUrl = gitPullUrl
	this.GitPushUrl = gitPushUrl
	this.HtmlUrl = htmlUrl
	this.Files = files
	this.Public = public
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Description = description
	this.Comments = comments
	this.User = user
	this.CommentsUrl = commentsUrl
	return &this
}

// NewGistWithDefaults instantiates a new Gist object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGistWithDefaults() *Gist {
	this := Gist{}
	return &this
}

// GetUrl returns the Url field value
func (o *Gist) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Gist) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Gist) SetUrl(v string) {
	o.Url = v
}

// GetForksUrl returns the ForksUrl field value
func (o *Gist) GetForksUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ForksUrl
}

// GetForksUrlOk returns a tuple with the ForksUrl field value
// and a boolean to check if the value has been set.
func (o *Gist) GetForksUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForksUrl, true
}

// SetForksUrl sets field value
func (o *Gist) SetForksUrl(v string) {
	o.ForksUrl = v
}

// GetCommitsUrl returns the CommitsUrl field value
func (o *Gist) GetCommitsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommitsUrl
}

// GetCommitsUrlOk returns a tuple with the CommitsUrl field value
// and a boolean to check if the value has been set.
func (o *Gist) GetCommitsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommitsUrl, true
}

// SetCommitsUrl sets field value
func (o *Gist) SetCommitsUrl(v string) {
	o.CommitsUrl = v
}

// GetId returns the Id field value
func (o *Gist) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Gist) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Gist) SetId(v string) {
	o.Id = v
}

// GetNodeId returns the NodeId field value
func (o *Gist) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *Gist) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *Gist) SetNodeId(v string) {
	o.NodeId = v
}

// GetGitPullUrl returns the GitPullUrl field value
func (o *Gist) GetGitPullUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GitPullUrl
}

// GetGitPullUrlOk returns a tuple with the GitPullUrl field value
// and a boolean to check if the value has been set.
func (o *Gist) GetGitPullUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GitPullUrl, true
}

// SetGitPullUrl sets field value
func (o *Gist) SetGitPullUrl(v string) {
	o.GitPullUrl = v
}

// GetGitPushUrl returns the GitPushUrl field value
func (o *Gist) GetGitPushUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GitPushUrl
}

// GetGitPushUrlOk returns a tuple with the GitPushUrl field value
// and a boolean to check if the value has been set.
func (o *Gist) GetGitPushUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GitPushUrl, true
}

// SetGitPushUrl sets field value
func (o *Gist) SetGitPushUrl(v string) {
	o.GitPushUrl = v
}

// GetHtmlUrl returns the HtmlUrl field value
func (o *Gist) GetHtmlUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HtmlUrl
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value
// and a boolean to check if the value has been set.
func (o *Gist) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlUrl, true
}

// SetHtmlUrl sets field value
func (o *Gist) SetHtmlUrl(v string) {
	o.HtmlUrl = v
}

// GetFiles returns the Files field value
func (o *Gist) GetFiles() map[string]BaseGistFilesValue {
	if o == nil {
		var ret map[string]BaseGistFilesValue
		return ret
	}

	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *Gist) GetFilesOk() (*map[string]BaseGistFilesValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Files, true
}

// SetFiles sets field value
func (o *Gist) SetFiles(v map[string]BaseGistFilesValue) {
	o.Files = v
}

// GetPublic returns the Public field value
func (o *Gist) GetPublic() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Public
}

// GetPublicOk returns a tuple with the Public field value
// and a boolean to check if the value has been set.
func (o *Gist) GetPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Public, true
}

// SetPublic sets field value
func (o *Gist) SetPublic(v bool) {
	o.Public = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Gist) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Gist) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Gist) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Gist) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Gist) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Gist) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Gist) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Gist) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *Gist) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetComments returns the Comments field value
func (o *Gist) GetComments() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value
// and a boolean to check if the value has been set.
func (o *Gist) GetCommentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comments, true
}

// SetComments sets field value
func (o *Gist) SetComments(v int32) {
	o.Comments = v
}

// GetUser returns the User field value
// If the value is explicit nil, the zero value for NullableSimpleUser will be returned
func (o *Gist) GetUser() NullableSimpleUser {
	if o == nil || o.User.Get() == nil {
		var ret NullableSimpleUser
		return ret
	}

	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Gist) GetUserOk() (*NullableSimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// SetUser sets field value
func (o *Gist) SetUser(v NullableSimpleUser) {
	o.User.Set(&v)
}

// GetCommentsUrl returns the CommentsUrl field value
func (o *Gist) GetCommentsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommentsUrl
}

// GetCommentsUrlOk returns a tuple with the CommentsUrl field value
// and a boolean to check if the value has been set.
func (o *Gist) GetCommentsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommentsUrl, true
}

// SetCommentsUrl sets field value
func (o *Gist) SetCommentsUrl(v string) {
	o.CommentsUrl = v
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Gist) GetOwner() NullableSimpleUser {
	if o == nil || o.Owner.Get() == nil {
		var ret NullableSimpleUser
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Gist) GetOwnerOk() (*NullableSimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *Gist) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableNullableSimpleUser and assigns it to the Owner field.
func (o *Gist) SetOwner(v NullableSimpleUser) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *Gist) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *Gist) UnsetOwner() {
	o.Owner.Unset()
}

// GetTruncated returns the Truncated field value if set, zero value otherwise.
func (o *Gist) GetTruncated() bool {
	if o == nil || o.Truncated == nil {
		var ret bool
		return ret
	}
	return *o.Truncated
}

// GetTruncatedOk returns a tuple with the Truncated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gist) GetTruncatedOk() (*bool, bool) {
	if o == nil || o.Truncated == nil {
		return nil, false
	}
	return o.Truncated, true
}

// HasTruncated returns a boolean if a field has been set.
func (o *Gist) HasTruncated() bool {
	if o != nil && o.Truncated != nil {
		return true
	}

	return false
}

// SetTruncated gets a reference to the given bool and assigns it to the Truncated field.
func (o *Gist) SetTruncated(v bool) {
	o.Truncated = &v
}

// GetForks returns the Forks field value if set, zero value otherwise.
func (o *Gist) GetForks() []interface{} {
	if o == nil || o.Forks == nil {
		var ret []interface{}
		return ret
	}
	return o.Forks
}

// GetForksOk returns a tuple with the Forks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gist) GetForksOk() ([]interface{}, bool) {
	if o == nil || o.Forks == nil {
		return nil, false
	}
	return o.Forks, true
}

// HasForks returns a boolean if a field has been set.
func (o *Gist) HasForks() bool {
	if o != nil && o.Forks != nil {
		return true
	}

	return false
}

// SetForks gets a reference to the given []interface{} and assigns it to the Forks field.
func (o *Gist) SetForks(v []interface{}) {
	o.Forks = v
}

// GetHistory returns the History field value if set, zero value otherwise.
func (o *Gist) GetHistory() []interface{} {
	if o == nil || o.History == nil {
		var ret []interface{}
		return ret
	}
	return o.History
}

// GetHistoryOk returns a tuple with the History field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gist) GetHistoryOk() ([]interface{}, bool) {
	if o == nil || o.History == nil {
		return nil, false
	}
	return o.History, true
}

// HasHistory returns a boolean if a field has been set.
func (o *Gist) HasHistory() bool {
	if o != nil && o.History != nil {
		return true
	}

	return false
}

// SetHistory gets a reference to the given []interface{} and assigns it to the History field.
func (o *Gist) SetHistory(v []interface{}) {
	o.History = v
}

func (o Gist) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["forks_url"] = o.ForksUrl
	}
	if true {
		toSerialize["commits_url"] = o.CommitsUrl
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["node_id"] = o.NodeId
	}
	if true {
		toSerialize["git_pull_url"] = o.GitPullUrl
	}
	if true {
		toSerialize["git_push_url"] = o.GitPushUrl
	}
	if true {
		toSerialize["html_url"] = o.HtmlUrl
	}
	if true {
		toSerialize["files"] = o.Files
	}
	if true {
		toSerialize["public"] = o.Public
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["description"] = o.Description.Get()
	}
	if true {
		toSerialize["comments"] = o.Comments
	}
	if true {
		toSerialize["user"] = o.User.Get()
	}
	if true {
		toSerialize["comments_url"] = o.CommentsUrl
	}
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	if o.Truncated != nil {
		toSerialize["truncated"] = o.Truncated
	}
	if o.Forks != nil {
		toSerialize["forks"] = o.Forks
	}
	if o.History != nil {
		toSerialize["history"] = o.History
	}
	return json.Marshal(toSerialize)
}

type NullableGist struct {
	value *Gist
	isSet bool
}

func (v NullableGist) Get() *Gist {
	return v.value
}

func (v *NullableGist) Set(val *Gist) {
	v.value = val
	v.isSet = true
}

func (v NullableGist) IsSet() bool {
	return v.isSet
}

func (v *NullableGist) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGist(val *Gist) *NullableGist {
	return &NullableGist{value: val, isSet: true}
}

func (v NullableGist) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGist) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


