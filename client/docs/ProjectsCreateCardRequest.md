# ProjectsCreateCardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Note** | **NullableString** | The project card&#39;s note | 
**ContentId** | **int32** | The unique identifier of the content associated with the card | 
**ContentType** | **string** | The piece of content associated with the card | 

## Methods

### NewProjectsCreateCardRequest

`func NewProjectsCreateCardRequest(note NullableString, contentId int32, contentType string, ) *ProjectsCreateCardRequest`

NewProjectsCreateCardRequest instantiates a new ProjectsCreateCardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsCreateCardRequestWithDefaults

`func NewProjectsCreateCardRequestWithDefaults() *ProjectsCreateCardRequest`

NewProjectsCreateCardRequestWithDefaults instantiates a new ProjectsCreateCardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNote

`func (o *ProjectsCreateCardRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ProjectsCreateCardRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ProjectsCreateCardRequest) SetNote(v string)`

SetNote sets Note field to given value.


### SetNoteNil

`func (o *ProjectsCreateCardRequest) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *ProjectsCreateCardRequest) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil
### GetContentId

`func (o *ProjectsCreateCardRequest) GetContentId() int32`

GetContentId returns the ContentId field if non-nil, zero value otherwise.

### GetContentIdOk

`func (o *ProjectsCreateCardRequest) GetContentIdOk() (*int32, bool)`

GetContentIdOk returns a tuple with the ContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentId

`func (o *ProjectsCreateCardRequest) SetContentId(v int32)`

SetContentId sets ContentId field to given value.


### GetContentType

`func (o *ProjectsCreateCardRequest) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ProjectsCreateCardRequest) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ProjectsCreateCardRequest) SetContentType(v string)`

SetContentType sets ContentType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


