# ProjectsUpdateCardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Note** | Pointer to **NullableString** | The project card&#39;s note | [optional] 
**Archived** | Pointer to **bool** | Whether or not the card is archived | [optional] 

## Methods

### NewProjectsUpdateCardRequest

`func NewProjectsUpdateCardRequest() *ProjectsUpdateCardRequest`

NewProjectsUpdateCardRequest instantiates a new ProjectsUpdateCardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsUpdateCardRequestWithDefaults

`func NewProjectsUpdateCardRequestWithDefaults() *ProjectsUpdateCardRequest`

NewProjectsUpdateCardRequestWithDefaults instantiates a new ProjectsUpdateCardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNote

`func (o *ProjectsUpdateCardRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ProjectsUpdateCardRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ProjectsUpdateCardRequest) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ProjectsUpdateCardRequest) HasNote() bool`

HasNote returns a boolean if a field has been set.

### SetNoteNil

`func (o *ProjectsUpdateCardRequest) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *ProjectsUpdateCardRequest) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil
### GetArchived

`func (o *ProjectsUpdateCardRequest) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ProjectsUpdateCardRequest) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ProjectsUpdateCardRequest) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *ProjectsUpdateCardRequest) HasArchived() bool`

HasArchived returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


