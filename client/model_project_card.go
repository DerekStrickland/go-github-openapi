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

// ProjectCard Project cards represent a scope of work.
type ProjectCard struct {
	Url string `json:"url"`
	// The project card's ID
	Id int32 `json:"id"`
	NodeId string `json:"node_id"`
	Note NullableString `json:"note"`
	Creator NullableNullableSimpleUser `json:"creator"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Whether or not the card is archived
	Archived *bool `json:"archived,omitempty"`
	ColumnName *string `json:"column_name,omitempty"`
	ProjectId *string `json:"project_id,omitempty"`
	ColumnUrl string `json:"column_url"`
	ContentUrl *string `json:"content_url,omitempty"`
	ProjectUrl string `json:"project_url"`
}

// NewProjectCard instantiates a new ProjectCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectCard(url string, id int32, nodeId string, note NullableString, creator NullableNullableSimpleUser, createdAt time.Time, updatedAt time.Time, columnUrl string, projectUrl string) *ProjectCard {
	this := ProjectCard{}
	this.Url = url
	this.Id = id
	this.NodeId = nodeId
	this.Note = note
	this.Creator = creator
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.ColumnUrl = columnUrl
	this.ProjectUrl = projectUrl
	return &this
}

// NewProjectCardWithDefaults instantiates a new ProjectCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectCardWithDefaults() *ProjectCard {
	this := ProjectCard{}
	return &this
}

// GetUrl returns the Url field value
func (o *ProjectCard) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ProjectCard) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ProjectCard) SetUrl(v string) {
	o.Url = v
}

// GetId returns the Id field value
func (o *ProjectCard) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectCard) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectCard) SetId(v int32) {
	o.Id = v
}

// GetNodeId returns the NodeId field value
func (o *ProjectCard) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ProjectCard) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ProjectCard) SetNodeId(v string) {
	o.NodeId = v
}

// GetNote returns the Note field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProjectCard) GetNote() string {
	if o == nil || o.Note.Get() == nil {
		var ret string
		return ret
	}

	return *o.Note.Get()
}

// GetNoteOk returns a tuple with the Note field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectCard) GetNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Note.Get(), o.Note.IsSet()
}

// SetNote sets field value
func (o *ProjectCard) SetNote(v string) {
	o.Note.Set(&v)
}

// GetCreator returns the Creator field value
// If the value is explicit nil, the zero value for NullableSimpleUser will be returned
func (o *ProjectCard) GetCreator() NullableSimpleUser {
	if o == nil || o.Creator.Get() == nil {
		var ret NullableSimpleUser
		return ret
	}

	return *o.Creator.Get()
}

// GetCreatorOk returns a tuple with the Creator field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProjectCard) GetCreatorOk() (*NullableSimpleUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Creator.Get(), o.Creator.IsSet()
}

// SetCreator sets field value
func (o *ProjectCard) SetCreator(v NullableSimpleUser) {
	o.Creator.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *ProjectCard) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ProjectCard) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ProjectCard) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ProjectCard) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ProjectCard) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ProjectCard) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *ProjectCard) GetArchived() bool {
	if o == nil || o.Archived == nil {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCard) GetArchivedOk() (*bool, bool) {
	if o == nil || o.Archived == nil {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *ProjectCard) HasArchived() bool {
	if o != nil && o.Archived != nil {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *ProjectCard) SetArchived(v bool) {
	o.Archived = &v
}

// GetColumnName returns the ColumnName field value if set, zero value otherwise.
func (o *ProjectCard) GetColumnName() string {
	if o == nil || o.ColumnName == nil {
		var ret string
		return ret
	}
	return *o.ColumnName
}

// GetColumnNameOk returns a tuple with the ColumnName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCard) GetColumnNameOk() (*string, bool) {
	if o == nil || o.ColumnName == nil {
		return nil, false
	}
	return o.ColumnName, true
}

// HasColumnName returns a boolean if a field has been set.
func (o *ProjectCard) HasColumnName() bool {
	if o != nil && o.ColumnName != nil {
		return true
	}

	return false
}

// SetColumnName gets a reference to the given string and assigns it to the ColumnName field.
func (o *ProjectCard) SetColumnName(v string) {
	o.ColumnName = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *ProjectCard) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCard) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *ProjectCard) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *ProjectCard) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetColumnUrl returns the ColumnUrl field value
func (o *ProjectCard) GetColumnUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ColumnUrl
}

// GetColumnUrlOk returns a tuple with the ColumnUrl field value
// and a boolean to check if the value has been set.
func (o *ProjectCard) GetColumnUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ColumnUrl, true
}

// SetColumnUrl sets field value
func (o *ProjectCard) SetColumnUrl(v string) {
	o.ColumnUrl = v
}

// GetContentUrl returns the ContentUrl field value if set, zero value otherwise.
func (o *ProjectCard) GetContentUrl() string {
	if o == nil || o.ContentUrl == nil {
		var ret string
		return ret
	}
	return *o.ContentUrl
}

// GetContentUrlOk returns a tuple with the ContentUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectCard) GetContentUrlOk() (*string, bool) {
	if o == nil || o.ContentUrl == nil {
		return nil, false
	}
	return o.ContentUrl, true
}

// HasContentUrl returns a boolean if a field has been set.
func (o *ProjectCard) HasContentUrl() bool {
	if o != nil && o.ContentUrl != nil {
		return true
	}

	return false
}

// SetContentUrl gets a reference to the given string and assigns it to the ContentUrl field.
func (o *ProjectCard) SetContentUrl(v string) {
	o.ContentUrl = &v
}

// GetProjectUrl returns the ProjectUrl field value
func (o *ProjectCard) GetProjectUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectUrl
}

// GetProjectUrlOk returns a tuple with the ProjectUrl field value
// and a boolean to check if the value has been set.
func (o *ProjectCard) GetProjectUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectUrl, true
}

// SetProjectUrl sets field value
func (o *ProjectCard) SetProjectUrl(v string) {
	o.ProjectUrl = v
}

func (o ProjectCard) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["node_id"] = o.NodeId
	}
	if true {
		toSerialize["note"] = o.Note.Get()
	}
	if true {
		toSerialize["creator"] = o.Creator.Get()
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.Archived != nil {
		toSerialize["archived"] = o.Archived
	}
	if o.ColumnName != nil {
		toSerialize["column_name"] = o.ColumnName
	}
	if o.ProjectId != nil {
		toSerialize["project_id"] = o.ProjectId
	}
	if true {
		toSerialize["column_url"] = o.ColumnUrl
	}
	if o.ContentUrl != nil {
		toSerialize["content_url"] = o.ContentUrl
	}
	if true {
		toSerialize["project_url"] = o.ProjectUrl
	}
	return json.Marshal(toSerialize)
}

type NullableProjectCard struct {
	value *ProjectCard
	isSet bool
}

func (v NullableProjectCard) Get() *ProjectCard {
	return v.value
}

func (v *NullableProjectCard) Set(val *ProjectCard) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectCard) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectCard(val *ProjectCard) *NullableProjectCard {
	return &NullableProjectCard{value: val, isSet: true}
}

func (v NullableProjectCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

