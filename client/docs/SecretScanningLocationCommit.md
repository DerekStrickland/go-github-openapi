# SecretScanningLocationCommit

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

### NewSecretScanningLocationCommit

`func NewSecretScanningLocationCommit(path string, startLine float32, endLine float32, startColumn float32, endColumn float32, blobSha string, blobUrl string, commitSha string, commitUrl string, ) *SecretScanningLocationCommit`

NewSecretScanningLocationCommit instantiates a new SecretScanningLocationCommit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretScanningLocationCommitWithDefaults

`func NewSecretScanningLocationCommitWithDefaults() *SecretScanningLocationCommit`

NewSecretScanningLocationCommitWithDefaults instantiates a new SecretScanningLocationCommit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *SecretScanningLocationCommit) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *SecretScanningLocationCommit) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *SecretScanningLocationCommit) SetPath(v string)`

SetPath sets Path field to given value.


### GetStartLine

`func (o *SecretScanningLocationCommit) GetStartLine() float32`

GetStartLine returns the StartLine field if non-nil, zero value otherwise.

### GetStartLineOk

`func (o *SecretScanningLocationCommit) GetStartLineOk() (*float32, bool)`

GetStartLineOk returns a tuple with the StartLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLine

`func (o *SecretScanningLocationCommit) SetStartLine(v float32)`

SetStartLine sets StartLine field to given value.


### GetEndLine

`func (o *SecretScanningLocationCommit) GetEndLine() float32`

GetEndLine returns the EndLine field if non-nil, zero value otherwise.

### GetEndLineOk

`func (o *SecretScanningLocationCommit) GetEndLineOk() (*float32, bool)`

GetEndLineOk returns a tuple with the EndLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLine

`func (o *SecretScanningLocationCommit) SetEndLine(v float32)`

SetEndLine sets EndLine field to given value.


### GetStartColumn

`func (o *SecretScanningLocationCommit) GetStartColumn() float32`

GetStartColumn returns the StartColumn field if non-nil, zero value otherwise.

### GetStartColumnOk

`func (o *SecretScanningLocationCommit) GetStartColumnOk() (*float32, bool)`

GetStartColumnOk returns a tuple with the StartColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartColumn

`func (o *SecretScanningLocationCommit) SetStartColumn(v float32)`

SetStartColumn sets StartColumn field to given value.


### GetEndColumn

`func (o *SecretScanningLocationCommit) GetEndColumn() float32`

GetEndColumn returns the EndColumn field if non-nil, zero value otherwise.

### GetEndColumnOk

`func (o *SecretScanningLocationCommit) GetEndColumnOk() (*float32, bool)`

GetEndColumnOk returns a tuple with the EndColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndColumn

`func (o *SecretScanningLocationCommit) SetEndColumn(v float32)`

SetEndColumn sets EndColumn field to given value.


### GetBlobSha

`func (o *SecretScanningLocationCommit) GetBlobSha() string`

GetBlobSha returns the BlobSha field if non-nil, zero value otherwise.

### GetBlobShaOk

`func (o *SecretScanningLocationCommit) GetBlobShaOk() (*string, bool)`

GetBlobShaOk returns a tuple with the BlobSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobSha

`func (o *SecretScanningLocationCommit) SetBlobSha(v string)`

SetBlobSha sets BlobSha field to given value.


### GetBlobUrl

`func (o *SecretScanningLocationCommit) GetBlobUrl() string`

GetBlobUrl returns the BlobUrl field if non-nil, zero value otherwise.

### GetBlobUrlOk

`func (o *SecretScanningLocationCommit) GetBlobUrlOk() (*string, bool)`

GetBlobUrlOk returns a tuple with the BlobUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobUrl

`func (o *SecretScanningLocationCommit) SetBlobUrl(v string)`

SetBlobUrl sets BlobUrl field to given value.


### GetCommitSha

`func (o *SecretScanningLocationCommit) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *SecretScanningLocationCommit) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *SecretScanningLocationCommit) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.


### GetCommitUrl

`func (o *SecretScanningLocationCommit) GetCommitUrl() string`

GetCommitUrl returns the CommitUrl field if non-nil, zero value otherwise.

### GetCommitUrlOk

`func (o *SecretScanningLocationCommit) GetCommitUrlOk() (*string, bool)`

GetCommitUrlOk returns a tuple with the CommitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitUrl

`func (o *SecretScanningLocationCommit) SetCommitUrl(v string)`

SetCommitUrl sets CommitUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


