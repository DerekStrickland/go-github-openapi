# SecretScanningLocationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | The file path in the repository | 
**StartLine** | **float32** | Line number at which the secret starts in the file | 
**EndLine** | **float32** | Line number at which the secret ends in the file | 
**StartColumn** | **float32** | The column at which the secret starts within the start line when the file is interpreted as 8BIT ASCII | 
**EndColumn** | **float32** | The column at which the secret ends within the end line when the file is interpreted as 8BIT ASCII | 
**BlobSha** | **string** | SHA-1 hash ID of the associated blob | 
**BlobUrl** | **string** | The API URL to get the associated blob resource | 
**CommitSha** | **string** | SHA-1 hash ID of the associated commit | 
**CommitUrl** | **string** | The API URL to get the associated commit resource | 

## Methods

### NewSecretScanningLocationDetails

`func NewSecretScanningLocationDetails(path string, startLine float32, endLine float32, startColumn float32, endColumn float32, blobSha string, blobUrl string, commitSha string, commitUrl string, ) *SecretScanningLocationDetails`

NewSecretScanningLocationDetails instantiates a new SecretScanningLocationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretScanningLocationDetailsWithDefaults

`func NewSecretScanningLocationDetailsWithDefaults() *SecretScanningLocationDetails`

NewSecretScanningLocationDetailsWithDefaults instantiates a new SecretScanningLocationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *SecretScanningLocationDetails) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SecretScanningLocationDetails) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SecretScanningLocationDetails) SetPath(v string)`

SetPath sets Path field to given value.


### GetStartLine

`func (o *SecretScanningLocationDetails) GetStartLine() float32`

GetStartLine returns the StartLine field if non-nil, zero value otherwise.

### GetStartLineOk

`func (o *SecretScanningLocationDetails) GetStartLineOk() (*float32, bool)`

GetStartLineOk returns a tuple with the StartLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLine

`func (o *SecretScanningLocationDetails) SetStartLine(v float32)`

SetStartLine sets StartLine field to given value.


### GetEndLine

`func (o *SecretScanningLocationDetails) GetEndLine() float32`

GetEndLine returns the EndLine field if non-nil, zero value otherwise.

### GetEndLineOk

`func (o *SecretScanningLocationDetails) GetEndLineOk() (*float32, bool)`

GetEndLineOk returns a tuple with the EndLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLine

`func (o *SecretScanningLocationDetails) SetEndLine(v float32)`

SetEndLine sets EndLine field to given value.


### GetStartColumn

`func (o *SecretScanningLocationDetails) GetStartColumn() float32`

GetStartColumn returns the StartColumn field if non-nil, zero value otherwise.

### GetStartColumnOk

`func (o *SecretScanningLocationDetails) GetStartColumnOk() (*float32, bool)`

GetStartColumnOk returns a tuple with the StartColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartColumn

`func (o *SecretScanningLocationDetails) SetStartColumn(v float32)`

SetStartColumn sets StartColumn field to given value.


### GetEndColumn

`func (o *SecretScanningLocationDetails) GetEndColumn() float32`

GetEndColumn returns the EndColumn field if non-nil, zero value otherwise.

### GetEndColumnOk

`func (o *SecretScanningLocationDetails) GetEndColumnOk() (*float32, bool)`

GetEndColumnOk returns a tuple with the EndColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndColumn

`func (o *SecretScanningLocationDetails) SetEndColumn(v float32)`

SetEndColumn sets EndColumn field to given value.


### GetBlobSha

`func (o *SecretScanningLocationDetails) GetBlobSha() string`

GetBlobSha returns the BlobSha field if non-nil, zero value otherwise.

### GetBlobShaOk

`func (o *SecretScanningLocationDetails) GetBlobShaOk() (*string, bool)`

GetBlobShaOk returns a tuple with the BlobSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobSha

`func (o *SecretScanningLocationDetails) SetBlobSha(v string)`

SetBlobSha sets BlobSha field to given value.


### GetBlobUrl

`func (o *SecretScanningLocationDetails) GetBlobUrl() string`

GetBlobUrl returns the BlobUrl field if non-nil, zero value otherwise.

### GetBlobUrlOk

`func (o *SecretScanningLocationDetails) GetBlobUrlOk() (*string, bool)`

GetBlobUrlOk returns a tuple with the BlobUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobUrl

`func (o *SecretScanningLocationDetails) SetBlobUrl(v string)`

SetBlobUrl sets BlobUrl field to given value.


### GetCommitSha

`func (o *SecretScanningLocationDetails) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *SecretScanningLocationDetails) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *SecretScanningLocationDetails) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.


### GetCommitUrl

`func (o *SecretScanningLocationDetails) GetCommitUrl() string`

GetCommitUrl returns the CommitUrl field if non-nil, zero value otherwise.

### GetCommitUrlOk

`func (o *SecretScanningLocationDetails) GetCommitUrlOk() (*string, bool)`

GetCommitUrlOk returns a tuple with the CommitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitUrl

`func (o *SecretScanningLocationDetails) SetCommitUrl(v string)`

SetCommitUrl sets CommitUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


