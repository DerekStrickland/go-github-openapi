# CodeScanningAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | **string** | The full Git reference, formatted as &#x60;refs/heads/&lt;branch name&gt;&#x60;, &#x60;refs/pull/&lt;number&gt;/merge&#x60;, or &#x60;refs/pull/&lt;number&gt;/head&#x60;. | 
**CommitSha** | **string** | The SHA of the commit to which the analysis you are uploading relates. | 
**AnalysisKey** | **string** | Identifies the configuration under which the analysis was executed. For example, in GitHub Actions this includes the workflow filename and job name. | 
**Environment** | **string** | Identifies the variable values associated with the environment in which this analysis was performed. | 
**Category** | Pointer to **string** | Identifies the configuration under which the analysis was executed. Used to distinguish between multiple analyses for the same tool and commit, but performed on different languages or different parts of the code. | [optional] 
**Error** | **string** |  | 
**CreatedAt** | **time.Time** | The time that the analysis was created in ISO 8601 format: &#x60;YYYY-MM-DDTHH:MM:SSZ&#x60;. | [readonly] 
**ResultsCount** | **int32** | The total number of results in the analysis. | 
**RulesCount** | **int32** | The total number of rules used in the analysis. | 
**Id** | **int32** | Unique identifier for this analysis. | 
**Url** | **string** | The REST API URL of the analysis resource. | [readonly] 
**SarifId** | **string** | An identifier for the upload. | 
**Tool** | [**CodeScanningAnalysisTool**](CodeScanningAnalysisTool.md) |  | 
**Deletable** | **bool** |  | 
**Warning** | **string** | Warning generated when processing the analysis | 

## Methods

### NewCodeScanningAnalysis

`func NewCodeScanningAnalysis(ref string, commitSha string, analysisKey string, environment string, error_ string, createdAt time.Time, resultsCount int32, rulesCount int32, id int32, url string, sarifId string, tool CodeScanningAnalysisTool, deletable bool, warning string, ) *CodeScanningAnalysis`

NewCodeScanningAnalysis instantiates a new CodeScanningAnalysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeScanningAnalysisWithDefaults

`func NewCodeScanningAnalysisWithDefaults() *CodeScanningAnalysis`

NewCodeScanningAnalysisWithDefaults instantiates a new CodeScanningAnalysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *CodeScanningAnalysis) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *CodeScanningAnalysis) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *CodeScanningAnalysis) SetRef(v string)`

SetRef sets Ref field to given value.


### GetCommitSha

`func (o *CodeScanningAnalysis) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *CodeScanningAnalysis) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *CodeScanningAnalysis) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.


### GetAnalysisKey

`func (o *CodeScanningAnalysis) GetAnalysisKey() string`

GetAnalysisKey returns the AnalysisKey field if non-nil, zero value otherwise.

### GetAnalysisKeyOk

`func (o *CodeScanningAnalysis) GetAnalysisKeyOk() (*string, bool)`

GetAnalysisKeyOk returns a tuple with the AnalysisKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisKey

`func (o *CodeScanningAnalysis) SetAnalysisKey(v string)`

SetAnalysisKey sets AnalysisKey field to given value.


### GetEnvironment

`func (o *CodeScanningAnalysis) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CodeScanningAnalysis) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CodeScanningAnalysis) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.


### GetCategory

`func (o *CodeScanningAnalysis) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CodeScanningAnalysis) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CodeScanningAnalysis) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CodeScanningAnalysis) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetError

`func (o *CodeScanningAnalysis) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CodeScanningAnalysis) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CodeScanningAnalysis) SetError(v string)`

SetError sets Error field to given value.


### GetCreatedAt

`func (o *CodeScanningAnalysis) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CodeScanningAnalysis) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CodeScanningAnalysis) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetResultsCount

`func (o *CodeScanningAnalysis) GetResultsCount() int32`

GetResultsCount returns the ResultsCount field if non-nil, zero value otherwise.

### GetResultsCountOk

`func (o *CodeScanningAnalysis) GetResultsCountOk() (*int32, bool)`

GetResultsCountOk returns a tuple with the ResultsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultsCount

`func (o *CodeScanningAnalysis) SetResultsCount(v int32)`

SetResultsCount sets ResultsCount field to given value.


### GetRulesCount

`func (o *CodeScanningAnalysis) GetRulesCount() int32`

GetRulesCount returns the RulesCount field if non-nil, zero value otherwise.

### GetRulesCountOk

`func (o *CodeScanningAnalysis) GetRulesCountOk() (*int32, bool)`

GetRulesCountOk returns a tuple with the RulesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesCount

`func (o *CodeScanningAnalysis) SetRulesCount(v int32)`

SetRulesCount sets RulesCount field to given value.


### GetId

`func (o *CodeScanningAnalysis) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CodeScanningAnalysis) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CodeScanningAnalysis) SetId(v int32)`

SetId sets Id field to given value.


### GetUrl

`func (o *CodeScanningAnalysis) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CodeScanningAnalysis) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CodeScanningAnalysis) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSarifId

`func (o *CodeScanningAnalysis) GetSarifId() string`

GetSarifId returns the SarifId field if non-nil, zero value otherwise.

### GetSarifIdOk

`func (o *CodeScanningAnalysis) GetSarifIdOk() (*string, bool)`

GetSarifIdOk returns a tuple with the SarifId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSarifId

`func (o *CodeScanningAnalysis) SetSarifId(v string)`

SetSarifId sets SarifId field to given value.


### GetTool

`func (o *CodeScanningAnalysis) GetTool() CodeScanningAnalysisTool`

GetTool returns the Tool field if non-nil, zero value otherwise.

### GetToolOk

`func (o *CodeScanningAnalysis) GetToolOk() (*CodeScanningAnalysisTool, bool)`

GetToolOk returns a tuple with the Tool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTool

`func (o *CodeScanningAnalysis) SetTool(v CodeScanningAnalysisTool)`

SetTool sets Tool field to given value.


### GetDeletable

`func (o *CodeScanningAnalysis) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *CodeScanningAnalysis) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *CodeScanningAnalysis) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.


### GetWarning

`func (o *CodeScanningAnalysis) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *CodeScanningAnalysis) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *CodeScanningAnalysis) SetWarning(v string)`

SetWarning sets Warning field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


