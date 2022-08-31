# ProjectCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Id** | **int32** | The project card&#39;s ID | 
**NodeId** | **string** |  | 
**Note** | **NullableString** |  | 
**Creator** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Archived** | Pointer to **bool** | Whether or not the card is archived | [optional] 
**ColumnName** | Pointer to **string** |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**ColumnUrl** | **string** |  | 
**ContentUrl** | Pointer to **string** |  | [optional] 
**ProjectUrl** | **string** |  | 

## Methods

### NewProjectCard

`func NewProjectCard(url string, id int32, nodeId string, note NullableString, creator NullableNullableSimpleUser, createdAt time.Time, updatedAt time.Time, columnUrl string, projectUrl string, ) *ProjectCard`

NewProjectCard instantiates a new ProjectCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCardWithDefaults

`func NewProjectCardWithDefaults() *ProjectCard`

NewProjectCardWithDefaults instantiates a new ProjectCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ProjectCard) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProjectCard) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProjectCard) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetId

`func (o *ProjectCard) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectCard) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectCard) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *ProjectCard) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ProjectCard) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ProjectCard) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetNote

`func (o *ProjectCard) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ProjectCard) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ProjectCard) SetNote(v string)`

SetNote sets Note field to given value.


### SetNoteNil

`func (o *ProjectCard) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *ProjectCard) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil
### GetCreator

`func (o *ProjectCard) GetCreator() NullableSimpleUser`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *ProjectCard) GetCreatorOk() (*NullableSimpleUser, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *ProjectCard) SetCreator(v NullableSimpleUser)`

SetCreator sets Creator field to given value.


### SetCreatorNil

`func (o *ProjectCard) SetCreatorNil(b bool)`

 SetCreatorNil sets the value for Creator to be an explicit nil

### UnsetCreator
`func (o *ProjectCard) UnsetCreator()`

UnsetCreator ensures that no value is present for Creator, not even an explicit nil
### GetCreatedAt

`func (o *ProjectCard) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProjectCard) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProjectCard) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ProjectCard) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProjectCard) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProjectCard) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetArchived

`func (o *ProjectCard) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ProjectCard) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ProjectCard) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *ProjectCard) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetColumnName

`func (o *ProjectCard) GetColumnName() string`

GetColumnName returns the ColumnName field if non-nil, zero value otherwise.

### GetColumnNameOk

`func (o *ProjectCard) GetColumnNameOk() (*string, bool)`

GetColumnNameOk returns a tuple with the ColumnName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnName

`func (o *ProjectCard) SetColumnName(v string)`

SetColumnName sets ColumnName field to given value.

### HasColumnName

`func (o *ProjectCard) HasColumnName() bool`

HasColumnName returns a boolean if a field has been set.

### GetProjectId

`func (o *ProjectCard) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ProjectCard) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ProjectCard) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ProjectCard) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetColumnUrl

`func (o *ProjectCard) GetColumnUrl() string`

GetColumnUrl returns the ColumnUrl field if non-nil, zero value otherwise.

### GetColumnUrlOk

`func (o *ProjectCard) GetColumnUrlOk() (*string, bool)`

GetColumnUrlOk returns a tuple with the ColumnUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnUrl

`func (o *ProjectCard) SetColumnUrl(v string)`

SetColumnUrl sets ColumnUrl field to given value.


### GetContentUrl

`func (o *ProjectCard) GetContentUrl() string`

GetContentUrl returns the ContentUrl field if non-nil, zero value otherwise.

### GetContentUrlOk

`func (o *ProjectCard) GetContentUrlOk() (*string, bool)`

GetContentUrlOk returns a tuple with the ContentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentUrl

`func (o *ProjectCard) SetContentUrl(v string)`

SetContentUrl sets ContentUrl field to given value.

### HasContentUrl

`func (o *ProjectCard) HasContentUrl() bool`

HasContentUrl returns a boolean if a field has been set.

### GetProjectUrl

`func (o *ProjectCard) GetProjectUrl() string`

GetProjectUrl returns the ProjectUrl field if non-nil, zero value otherwise.

### GetProjectUrlOk

`func (o *ProjectCard) GetProjectUrlOk() (*string, bool)`

GetProjectUrlOk returns a tuple with the ProjectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectUrl

`func (o *ProjectCard) SetProjectUrl(v string)`

SetProjectUrl sets ProjectUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


