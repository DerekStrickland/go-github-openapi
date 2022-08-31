# CodeScanningAlertInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The full Git reference, formatted as &#x60;refs/heads/&lt;branch name&gt;&#x60;, &#x60;refs/pull/&lt;number&gt;/merge&#x60;, or &#x60;refs/pull/&lt;number&gt;/head&#x60;. | [optional] 
**AnalysisKey** | Pointer to **string** | Identifies the configuration under which the analysis was executed. For example, in GitHub Actions this includes the workflow filename and job name. | [optional] 
**Environment** | Pointer to **string** | Identifies the variable values associated with the environment in which the analysis that generated this alert instance was performed, such as the language that was analyzed. | [optional] 
**Category** | Pointer to **string** | Identifies the configuration under which the analysis was executed. Used to distinguish between multiple analyses for the same tool and commit, but performed on different languages or different parts of the code. | [optional] 
**State** | Pointer to [**CodeScanningAlertState**](CodeScanningAlertState.md) |  | [optional] 
**CommitSha** | Pointer to **string** |  | [optional] 
**Message** | Pointer to [**CodeScanningAlertInstanceMessage**](CodeScanningAlertInstanceMessage.md) |  | [optional] 
**Location** | Pointer to [**CodeScanningAlertLocation**](CodeScanningAlertLocation.md) |  | [optional] 
**HtmlUrl** | Pointer to **string** |  | [optional] 
**Classifications** | Pointer to [**[]CodeScanningAlertClassification**](CodeScanningAlertClassification.md) | Classifications that have been applied to the file that triggered the alert. For example identifying it as documentation, or a generated file. | [optional] 

## Methods

### NewCodeScanningAlertInstance

`func NewCodeScanningAlertInstance() *CodeScanningAlertInstance`

NewCodeScanningAlertInstance instantiates a new CodeScanningAlertInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeScanningAlertInstanceWithDefaults

`func NewCodeScanningAlertInstanceWithDefaults() *CodeScanningAlertInstance`

NewCodeScanningAlertInstanceWithDefaults instantiates a new CodeScanningAlertInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *CodeScanningAlertInstance) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *CodeScanningAlertInstance) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *CodeScanningAlertInstance) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *CodeScanningAlertInstance) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAnalysisKey

`func (o *CodeScanningAlertInstance) GetAnalysisKey() string`

GetAnalysisKey returns the AnalysisKey field if non-nil, zero value otherwise.

### GetAnalysisKeyOk

`func (o *CodeScanningAlertInstance) GetAnalysisKeyOk() (*string, bool)`

GetAnalysisKeyOk returns a tuple with the AnalysisKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysisKey

`func (o *CodeScanningAlertInstance) SetAnalysisKey(v string)`

SetAnalysisKey sets AnalysisKey field to given value.

### HasAnalysisKey

`func (o *CodeScanningAlertInstance) HasAnalysisKey() bool`

HasAnalysisKey returns a boolean if a field has been set.

### GetEnvironment

`func (o *CodeScanningAlertInstance) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CodeScanningAlertInstance) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CodeScanningAlertInstance) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CodeScanningAlertInstance) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetCategory

`func (o *CodeScanningAlertInstance) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CodeScanningAlertInstance) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CodeScanningAlertInstance) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CodeScanningAlertInstance) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetState

`func (o *CodeScanningAlertInstance) GetState() CodeScanningAlertState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CodeScanningAlertInstance) GetStateOk() (*CodeScanningAlertState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CodeScanningAlertInstance) SetState(v CodeScanningAlertState)`

SetState sets State field to given value.

### HasState

`func (o *CodeScanningAlertInstance) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCommitSha

`func (o *CodeScanningAlertInstance) GetCommitSha() string`

GetCommitSha returns the CommitSha field if non-nil, zero value otherwise.

### GetCommitShaOk

`func (o *CodeScanningAlertInstance) GetCommitShaOk() (*string, bool)`

GetCommitShaOk returns a tuple with the CommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitSha

`func (o *CodeScanningAlertInstance) SetCommitSha(v string)`

SetCommitSha sets CommitSha field to given value.

### HasCommitSha

`func (o *CodeScanningAlertInstance) HasCommitSha() bool`

HasCommitSha returns a boolean if a field has been set.

### GetMessage

`func (o *CodeScanningAlertInstance) GetMessage() CodeScanningAlertInstanceMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CodeScanningAlertInstance) GetMessageOk() (*CodeScanningAlertInstanceMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CodeScanningAlertInstance) SetMessage(v CodeScanningAlertInstanceMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CodeScanningAlertInstance) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetLocation

`func (o *CodeScanningAlertInstance) GetLocation() CodeScanningAlertLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CodeScanningAlertInstance) GetLocationOk() (*CodeScanningAlertLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CodeScanningAlertInstance) SetLocation(v CodeScanningAlertLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CodeScanningAlertInstance) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *CodeScanningAlertInstance) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *CodeScanningAlertInstance) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *CodeScanningAlertInstance) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *CodeScanningAlertInstance) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.

### GetClassifications

`func (o *CodeScanningAlertInstance) GetClassifications() []CodeScanningAlertClassification`

GetClassifications returns the Classifications field if non-nil, zero value otherwise.

### GetClassificationsOk

`func (o *CodeScanningAlertInstance) GetClassificationsOk() (*[]CodeScanningAlertClassification, bool)`

GetClassificationsOk returns a tuple with the Classifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifications

`func (o *CodeScanningAlertInstance) SetClassifications(v []CodeScanningAlertClassification)`

SetClassifications sets Classifications field to given value.

### HasClassifications

`func (o *CodeScanningAlertInstance) HasClassifications() bool`

HasClassifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


