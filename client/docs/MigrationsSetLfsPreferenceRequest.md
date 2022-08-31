# MigrationsSetLfsPreferenceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseLfs** | **string** | Whether to store large files during the import. &#x60;opt_in&#x60; means large files will be stored using Git LFS. &#x60;opt_out&#x60; means large files will be removed during the import. | 

## Methods

### NewMigrationsSetLfsPreferenceRequest

`func NewMigrationsSetLfsPreferenceRequest(useLfs string, ) *MigrationsSetLfsPreferenceRequest`

NewMigrationsSetLfsPreferenceRequest instantiates a new MigrationsSetLfsPreferenceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationsSetLfsPreferenceRequestWithDefaults

`func NewMigrationsSetLfsPreferenceRequestWithDefaults() *MigrationsSetLfsPreferenceRequest`

NewMigrationsSetLfsPreferenceRequestWithDefaults instantiates a new MigrationsSetLfsPreferenceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseLfs

`func (o *MigrationsSetLfsPreferenceRequest) GetUseLfs() string`

GetUseLfs returns the UseLfs field if non-nil, zero value otherwise.

### GetUseLfsOk

`func (o *MigrationsSetLfsPreferenceRequest) GetUseLfsOk() (*string, bool)`

GetUseLfsOk returns a tuple with the UseLfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLfs

`func (o *MigrationsSetLfsPreferenceRequest) SetUseLfs(v string)`

SetUseLfs sets UseLfs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


