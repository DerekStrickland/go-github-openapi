# CodeScanningUploadSarifRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitSha** | **string** | The SHA of the commit to which the analysis you are uploading relates. | 
**Ref** | **string** | The full Git reference, formatted as &#x60;refs/heads/&lt;branch name&gt;&#x60;, &#x60;refs/pull/&lt;number&gt;/merge&#x60;, or &#x60;refs/pull/&lt;number&gt;/head&#x60;. | 
**Sarif** | **string** | A Base64 string representing the SARIF file to upload. You must first compress your SARIF file using [&#x60;gzip&#x60;](http://www.gnu.org/software/gzip/manual/gzip.html) and then translate the contents of the file into a Base64 encoding string. For more information, see \&quot;[SARIF support for code scanning](https://docs.github.com/code-security/secure-coding/sarif-support-for-code-scanning).\&quot; | 
**CheckoutUri** | Pointer to **string** | The base directory used in the analysis, as it appears in the SARIF file. This property is used to convert file paths from absolute to relative, so that alerts can be mapped to their correct location in the repository. | [optional] 
**StartedAt** | Pointer to **time.Time** | The time that the analysis run began. This is a timestamp in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | [optional] 
**ToolName** | Pointer to **string** | The name of the tool used to generate the code scanning analysis. If this parameter is not used, the tool name defaults to \&quot;API\&quot;. If the uploaded SARIF contains a tool GUID, this will be available for filtering using the &#x60;tool_guid&#x60; parameter of operations such as &#x60;GET /repos/{owner}/{repo}/code-scanning/alerts&#x60;. | [optional] 

## Methods

### NewCodeScanningUploadSarifRequest

`func NewCodeScanningUploadSarifRequest(commitSha string, ref string, sarif string, ) *CodeScanningUploadSarifRequest`

NewCodeScanningUploadSarifRequest instantiates a new CodeScanningUploadSarifRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeScanningUploadSarifRequestWithDefaults

`func NewCodeScanningUploadSarifRequestWithDefaults() *CodeScanningUploadSarifRequest`

NewCodeScanningUploadSarifRequestWithDefaults instantiates a new CodeScanningUploadSarifRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitSha

`func (o *CodeScanningUploadSarifRequest) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *CodeScanningUploadSarifRequest) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *CodeScanningUploadSarifRequest) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.


### GetRef

`func (o *CodeScanningUploadSarifRequest) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *CodeScanningUploadSarifRequest) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *CodeScanningUploadSarifRequest) SetRef(v string)`

SetRef sets Ref field to given value.


### GetSarif

`func (o *CodeScanningUploadSarifRequest) GetSarif() string`

GetSarif returns the Sarif field if non-nil, zero value otherwise.

### GetSarifOk

`func (o *CodeScanningUploadSarifRequest) GetSarifOk() (*string, bool)`

GetSarifOk returns a tuple with the Sarif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSarif

`func (o *CodeScanningUploadSarifRequest) SetSarif(v string)`

SetSarif sets Sarif field to given value.


### GetCheckoutUri

`func (o *CodeScanningUploadSarifRequest) GetCheckoutUri() string`

GetCheckoutUri returns the CheckoutUri field if non-nil, zero value otherwise.

### GetCheckoutUriOk

`func (o *CodeScanningUploadSarifRequest) GetCheckoutUriOk() (*string, bool)`

GetCheckoutUriOk returns a tuple with the CheckoutUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutUri

`func (o *CodeScanningUploadSarifRequest) SetCheckoutUri(v string)`

SetCheckoutUri sets CheckoutUri field to given value.

### HasCheckoutUri

`func (o *CodeScanningUploadSarifRequest) HasCheckoutUri() bool`

HasCheckoutUri returns a boolean if a field has been set.

### GetStartedAt

`func (o *CodeScanningUploadSarifRequest) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CodeScanningUploadSarifRequest) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CodeScanningUploadSarifRequest) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *CodeScanningUploadSarifRequest) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetToolName

`func (o *CodeScanningUploadSarifRequest) GetToolName() string`

GetToolName returns the ToolName field if non-nil, zero value otherwise.

### GetToolNameOk

`func (o *CodeScanningUploadSarifRequest) GetToolNameOk() (*string, bool)`

GetToolNameOk returns a tuple with the ToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolName

`func (o *CodeScanningUploadSarifRequest) SetToolName(v string)`

SetToolName sets ToolName field to given value.

### HasToolName

`func (o *CodeScanningUploadSarifRequest) HasToolName() bool`

HasToolName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


