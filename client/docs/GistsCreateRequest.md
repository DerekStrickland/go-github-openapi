# GistsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the gist | [optional] 
**Files** | [**map[string]GistsCreateRequestFilesValue**](GistsCreateRequestFilesValue.md) | Names and content for the files that make up the gist | 
**Public** | Pointer to [**GistsCreateRequestPublic**](GistsCreateRequestPublic.md) |  | [optional] 

## Methods

### NewGistsCreateRequest

`func NewGistsCreateRequest(files map[string]GistsCreateRequestFilesValue, ) *GistsCreateRequest`

NewGistsCreateRequest instantiates a new GistsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGistsCreateRequestWithDefaults

`func NewGistsCreateRequestWithDefaults() *GistsCreateRequest`

NewGistsCreateRequestWithDefaults instantiates a new GistsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GistsCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GistsCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GistsCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GistsCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFiles

`func (o *GistsCreateRequest) GetFiles() map[string]GistsCreateRequestFilesValue`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *GistsCreateRequest) GetFilesOk() (*map[string]GistsCreateRequestFilesValue, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *GistsCreateRequest) SetFiles(v map[string]GistsCreateRequestFilesValue)`

SetFiles sets Files field to given value.


### GetPublic

`func (o *GistsCreateRequest) GetPublic() GistsCreateRequestPublic`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *GistsCreateRequest) GetPublicOk() (*GistsCreateRequestPublic, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *GistsCreateRequest) SetPublic(v GistsCreateRequestPublic)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *GistsCreateRequest) HasPublic() bool`

HasPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


