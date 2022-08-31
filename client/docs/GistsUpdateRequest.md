# GistsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the gist | [optional] 
**Files** | Pointer to [**map[string]GistsUpdateRequestFilesValue**](GistsUpdateRequestFilesValue.md) | Names of files to be updated | [optional] 

## Methods

### NewGistsUpdateRequest

`func NewGistsUpdateRequest() *GistsUpdateRequest`

NewGistsUpdateRequest instantiates a new GistsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGistsUpdateRequestWithDefaults

`func NewGistsUpdateRequestWithDefaults() *GistsUpdateRequest`

NewGistsUpdateRequestWithDefaults instantiates a new GistsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GistsUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GistsUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GistsUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GistsUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFiles

`func (o *GistsUpdateRequest) GetFiles() map[string]GistsUpdateRequestFilesValue`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *GistsUpdateRequest) GetFilesOk() (*map[string]GistsUpdateRequestFilesValue, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *GistsUpdateRequest) SetFiles(v map[string]GistsUpdateRequestFilesValue)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *GistsUpdateRequest) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


