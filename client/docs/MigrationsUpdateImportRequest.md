# MigrationsUpdateImportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VcsUsername** | Pointer to **string** | The username to provide to the originating repository. | [optional] 
**VcsPassword** | Pointer to **string** | The password to provide to the originating repository. | [optional] 
**Vcs** | Pointer to **string** | The type of version control system you are migrating from. | [optional] 
**TfvcProject** | Pointer to **string** | For a tfvc import, the name of the project that is being imported. | [optional] 

## Methods

### NewMigrationsUpdateImportRequest

`func NewMigrationsUpdateImportRequest() *MigrationsUpdateImportRequest`

NewMigrationsUpdateImportRequest instantiates a new MigrationsUpdateImportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationsUpdateImportRequestWithDefaults

`func NewMigrationsUpdateImportRequestWithDefaults() *MigrationsUpdateImportRequest`

NewMigrationsUpdateImportRequestWithDefaults instantiates a new MigrationsUpdateImportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVcsUsername

`func (o *MigrationsUpdateImportRequest) GetVcsUsername() string`

GetVcsUsername returns the VcsUsername field if non-nil, zero value otherwise.

### GetVcsUsernameOk

`func (o *MigrationsUpdateImportRequest) GetVcsUsernameOk() (*string, bool)`

GetVcsUsernameOk returns a tuple with the VcsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsUsername

`func (o *MigrationsUpdateImportRequest) SetVcsUsername(v string)`

SetVcsUsername sets VcsUsername field to given value.

### HasVcsUsername

`func (o *MigrationsUpdateImportRequest) HasVcsUsername() bool`

HasVcsUsername returns a boolean if a field has been set.

### GetVcsPassword

`func (o *MigrationsUpdateImportRequest) GetVcsPassword() string`

GetVcsPassword returns the VcsPassword field if non-nil, zero value otherwise.

### GetVcsPasswordOk

`func (o *MigrationsUpdateImportRequest) GetVcsPasswordOk() (*string, bool)`

GetVcsPasswordOk returns a tuple with the VcsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsPassword

`func (o *MigrationsUpdateImportRequest) SetVcsPassword(v string)`

SetVcsPassword sets VcsPassword field to given value.

### HasVcsPassword

`func (o *MigrationsUpdateImportRequest) HasVcsPassword() bool`

HasVcsPassword returns a boolean if a field has been set.

### GetVcs

`func (o *MigrationsUpdateImportRequest) GetVcs() string`

GetVcs returns the Vcs field if non-nil, zero value otherwise.

### GetVcsOk

`func (o *MigrationsUpdateImportRequest) GetVcsOk() (*string, bool)`

GetVcsOk returns a tuple with the Vcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcs

`func (o *MigrationsUpdateImportRequest) SetVcs(v string)`

SetVcs sets Vcs field to given value.

### HasVcs

`func (o *MigrationsUpdateImportRequest) HasVcs() bool`

HasVcs returns a boolean if a field has been set.

### GetTfvcProject

`func (o *MigrationsUpdateImportRequest) GetTfvcProject() string`

GetTfvcProject returns the TfvcProject field if non-nil, zero value otherwise.

### GetTfvcProjectOk

`func (o *MigrationsUpdateImportRequest) GetTfvcProjectOk() (*string, bool)`

GetTfvcProjectOk returns a tuple with the TfvcProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfvcProject

`func (o *MigrationsUpdateImportRequest) SetTfvcProject(v string)`

SetTfvcProject sets TfvcProject field to given value.

### HasTfvcProject

`func (o *MigrationsUpdateImportRequest) HasTfvcProject() bool`

HasTfvcProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


