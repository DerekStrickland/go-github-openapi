# MigrationsStartImportRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VcsUrl** | **string** | The URL of the originating repository. | 
**Vcs** | Pointer to **string** | The originating VCS type. Without this parameter, the import job will take additional time to detect the VCS type before beginning the import. This detection step will be reflected in the response. | [optional] 
**VcsUsername** | Pointer to **string** | If authentication is required, the username to provide to &#x60;vcs_url&#x60;. | [optional] 
**VcsPassword** | Pointer to **string** | If authentication is required, the password to provide to &#x60;vcs_url&#x60;. | [optional] 
**TfvcProject** | Pointer to **string** | For a tfvc import, the name of the project that is being imported. | [optional] 

## Methods

### NewMigrationsStartImportRequest

`func NewMigrationsStartImportRequest(vcsUrl string, ) *MigrationsStartImportRequest`

NewMigrationsStartImportRequest instantiates a new MigrationsStartImportRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationsStartImportRequestWithDefaults

`func NewMigrationsStartImportRequestWithDefaults() *MigrationsStartImportRequest`

NewMigrationsStartImportRequestWithDefaults instantiates a new MigrationsStartImportRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVcsUrl

`func (o *MigrationsStartImportRequest) GetVcsUrl() string`

GetVcsUrl returns the VcsUrl field if non-nil, zero value otherwise.

### GetVcsUrlOk

`func (o *MigrationsStartImportRequest) GetVcsUrlOk() (*string, bool)`

GetVcsUrlOk returns a tuple with the VcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsUrl

`func (o *MigrationsStartImportRequest) SetVcsUrl(v string)`

SetVcsUrl sets VcsUrl field to given value.


### GetVcs

`func (o *MigrationsStartImportRequest) GetVcs() string`

GetVcs returns the Vcs field if non-nil, zero value otherwise.

### GetVcsOk

`func (o *MigrationsStartImportRequest) GetVcsOk() (*string, bool)`

GetVcsOk returns a tuple with the Vcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcs

`func (o *MigrationsStartImportRequest) SetVcs(v string)`

SetVcs sets Vcs field to given value.

### HasVcs

`func (o *MigrationsStartImportRequest) HasVcs() bool`

HasVcs returns a boolean if a field has been set.

### GetVcsUsername

`func (o *MigrationsStartImportRequest) GetVcsUsername() string`

GetVcsUsername returns the VcsUsername field if non-nil, zero value otherwise.

### GetVcsUsernameOk

`func (o *MigrationsStartImportRequest) GetVcsUsernameOk() (*string, bool)`

GetVcsUsernameOk returns a tuple with the VcsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsUsername

`func (o *MigrationsStartImportRequest) SetVcsUsername(v string)`

SetVcsUsername sets VcsUsername field to given value.

### HasVcsUsername

`func (o *MigrationsStartImportRequest) HasVcsUsername() bool`

HasVcsUsername returns a boolean if a field has been set.

### GetVcsPassword

`func (o *MigrationsStartImportRequest) GetVcsPassword() string`

GetVcsPassword returns the VcsPassword field if non-nil, zero value otherwise.

### GetVcsPasswordOk

`func (o *MigrationsStartImportRequest) GetVcsPasswordOk() (*string, bool)`

GetVcsPasswordOk returns a tuple with the VcsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsPassword

`func (o *MigrationsStartImportRequest) SetVcsPassword(v string)`

SetVcsPassword sets VcsPassword field to given value.

### HasVcsPassword

`func (o *MigrationsStartImportRequest) HasVcsPassword() bool`

HasVcsPassword returns a boolean if a field has been set.

### GetTfvcProject

`func (o *MigrationsStartImportRequest) GetTfvcProject() string`

GetTfvcProject returns the TfvcProject field if non-nil, zero value otherwise.

### GetTfvcProjectOk

`func (o *MigrationsStartImportRequest) GetTfvcProjectOk() (*string, bool)`

GetTfvcProjectOk returns a tuple with the TfvcProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfvcProject

`func (o *MigrationsStartImportRequest) SetTfvcProject(v string)`

SetTfvcProject sets TfvcProject field to given value.

### HasTfvcProject

`func (o *MigrationsStartImportRequest) HasTfvcProject() bool`

HasTfvcProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


