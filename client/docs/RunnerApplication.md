# RunnerApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Os** | **string** |  | 
**Architecture** | **string** |  | 
**DownloadUrl** | **string** |  | 
**Filename** | **string** |  | 
**TempDownloadToken** | Pointer to **string** | A short lived bearer token used to download the runner, if needed. | [optional] 
**Sha256Checksum** | Pointer to **string** |  | [optional] 

## Methods

### NewRunnerApplication

`func NewRunnerApplication(os string, architecture string, downloadUrl string, filename string, ) *RunnerApplication`

NewRunnerApplication instantiates a new RunnerApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerApplicationWithDefaults

`func NewRunnerApplicationWithDefaults() *RunnerApplication`

NewRunnerApplicationWithDefaults instantiates a new RunnerApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOs

`func (o *RunnerApplication) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *RunnerApplication) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *RunnerApplication) SetOs(v string)`

SetOs sets Os field to given value.


### GetArchitecture

`func (o *RunnerApplication) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *RunnerApplication) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *RunnerApplication) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.


### GetDownloadUrl

`func (o *RunnerApplication) GetDownloadUrl() string`

GetDownloadUrl returns the DownloadUrl field if non-nil, zero value otherwise.

### GetDownloadUrlOk

`func (o *RunnerApplication) GetDownloadUrlOk() (*string, bool)`

GetDownloadUrlOk returns a tuple with the DownloadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrl

`func (o *RunnerApplication) SetDownloadUrl(v string)`

SetDownloadUrl sets DownloadUrl field to given value.


### GetFilename

`func (o *RunnerApplication) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *RunnerApplication) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *RunnerApplication) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetTempDownloadToken

`func (o *RunnerApplication) GetTempDownloadToken() string`

GetTempDownloadToken returns the TempDownloadToken field if non-nil, zero value otherwise.

### GetTempDownloadTokenOk

`func (o *RunnerApplication) GetTempDownloadTokenOk() (*string, bool)`

GetTempDownloadTokenOk returns a tuple with the TempDownloadToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempDownloadToken

`func (o *RunnerApplication) SetTempDownloadToken(v string)`

SetTempDownloadToken sets TempDownloadToken field to given value.

### HasTempDownloadToken

`func (o *RunnerApplication) HasTempDownloadToken() bool`

HasTempDownloadToken returns a boolean if a field has been set.

### GetSha256Checksum

`func (o *RunnerApplication) GetSha256Checksum() string`

GetSha256Checksum returns the Sha256Checksum field if non-nil, zero value otherwise.

### GetSha256ChecksumOk

`func (o *RunnerApplication) GetSha256ChecksumOk() (*string, bool)`

GetSha256ChecksumOk returns a tuple with the Sha256Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256Checksum

`func (o *RunnerApplication) SetSha256Checksum(v string)`

SetSha256Checksum sets Sha256Checksum field to given value.

### HasSha256Checksum

`func (o *RunnerApplication) HasSha256Checksum() bool`

HasSha256Checksum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


