# CheckRun

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The id of the check. | 
**HeadSha** | **string** | The SHA of the commit that is being checked. | 
**NodeId** | **string** |  | 
**ExternalId** | **NullableString** |  | 
**Url** | **string** |  | 
**HtmlUrl** | **NullableString** |  | 
**DetailsUrl** | **NullableString** |  | 
**Status** | **string** | The phase of the lifecycle that the check is currently in. | 
**Conclusion** | **NullableString** |  | 
**StartedAt** | **NullableTime** |  | 
**CompletedAt** | **NullableTime** |  | 
**Output** | [**CheckRunOutput**](CheckRunOutput.md) |  | 
**Name** | **string** | The name of the check. | 
**CheckSuite** | [**NullableCheckRunCheckSuite**](CheckRunCheckSuite.md) |  | 
**App** | [**NullableNullableIntegration**](NullableIntegration.md) |  | 
**PullRequests** | [**[]PullRequestMinimal**](PullRequestMinimal.md) |  | 
**Deployment** | Pointer to [**DeploymentSimple**](DeploymentSimple.md) |  | [optional] 

## Methods

### NewCheckRun

`func NewCheckRun(id int32, headSha string, nodeId string, externalId NullableString, url string, htmlUrl NullableString, detailsUrl NullableString, status string, conclusion NullableString, startedAt NullableTime, completedAt NullableTime, output CheckRunOutput, name string, checkSuite NullableCheckRunCheckSuite, app NullableNullableIntegration, pullRequests []PullRequestMinimal, ) *CheckRun`

NewCheckRun instantiates a new CheckRun object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckRunWithDefaults

`func NewCheckRunWithDefaults() *CheckRun`

NewCheckRunWithDefaults instantiates a new CheckRun object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CheckRun) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckRun) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckRun) SetId(v int32)`

SetId sets Id field to given value.


### GetHeadSha

`func (o *CheckRun) GetHeadSha() string`

GetHeadSha returns the HeadSha field if non-nil, zero value otherwise.

### GetHeadShaOk

`func (o *CheckRun) GetHeadShaOk() (*string, bool)`

GetHeadShaOk returns a tuple with the HeadSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadSha

`func (o *CheckRun) SetHeadSha(v string)`

SetHeadSha sets HeadSha field to given value.


### GetNodeId

`func (o *CheckRun) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *CheckRun) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *CheckRun) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetExternalId

`func (o *CheckRun) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CheckRun) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CheckRun) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### SetExternalIdNil

`func (o *CheckRun) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *CheckRun) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetUrl

`func (o *CheckRun) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CheckRun) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CheckRun) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *CheckRun) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *CheckRun) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *CheckRun) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### SetHtmlUrlNil

`func (o *CheckRun) SetHtmlUrlNil(b bool)`

 SetHtmlUrlNil sets the value for HtmlUrl to be an explicit nil

### UnsetHtmlUrl
`func (o *CheckRun) UnsetHtmlUrl()`

UnsetHtmlUrl ensures that no value is present for HtmlUrl, not even an explicit nil
### GetDetailsUrl

`func (o *CheckRun) GetDetailsUrl() string`

GetDetailsUrl returns the DetailsUrl field if non-nil, zero value otherwise.

### GetDetailsUrlOk

`func (o *CheckRun) GetDetailsUrlOk() (*string, bool)`

GetDetailsUrlOk returns a tuple with the DetailsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailsUrl

`func (o *CheckRun) SetDetailsUrl(v string)`

SetDetailsUrl sets DetailsUrl field to given value.


### SetDetailsUrlNil

`func (o *CheckRun) SetDetailsUrlNil(b bool)`

 SetDetailsUrlNil sets the value for DetailsUrl to be an explicit nil

### UnsetDetailsUrl
`func (o *CheckRun) UnsetDetailsUrl()`

UnsetDetailsUrl ensures that no value is present for DetailsUrl, not even an explicit nil
### GetStatus

`func (o *CheckRun) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CheckRun) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CheckRun) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetConclusion

`func (o *CheckRun) GetConclusion() string`

GetConclusion returns the Conclusion field if non-nil, zero value otherwise.

### GetConclusionOk

`func (o *CheckRun) GetConclusionOk() (*string, bool)`

GetConclusionOk returns a tuple with the Conclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusion

`func (o *CheckRun) SetConclusion(v string)`

SetConclusion sets Conclusion field to given value.


### SetConclusionNil

`func (o *CheckRun) SetConclusionNil(b bool)`

 SetConclusionNil sets the value for Conclusion to be an explicit nil

### UnsetConclusion
`func (o *CheckRun) UnsetConclusion()`

UnsetConclusion ensures that no value is present for Conclusion, not even an explicit nil
### GetStartedAt

`func (o *CheckRun) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CheckRun) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CheckRun) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### SetStartedAtNil

`func (o *CheckRun) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *CheckRun) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *CheckRun) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *CheckRun) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *CheckRun) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.


### SetCompletedAtNil

`func (o *CheckRun) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *CheckRun) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetOutput

`func (o *CheckRun) GetOutput() CheckRunOutput`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *CheckRun) GetOutputOk() (*CheckRunOutput, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *CheckRun) SetOutput(v CheckRunOutput)`

SetOutput sets Output field to given value.


### GetName

`func (o *CheckRun) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CheckRun) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CheckRun) SetName(v string)`

SetName sets Name field to given value.


### GetCheckSuite

`func (o *CheckRun) GetCheckSuite() CheckRunCheckSuite`

GetCheckSuite returns the CheckSuite field if non-nil, zero value otherwise.

### GetCheckSuiteOk

`func (o *CheckRun) GetCheckSuiteOk() (*CheckRunCheckSuite, bool)`

GetCheckSuiteOk returns a tuple with the CheckSuite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckSuite

`func (o *CheckRun) SetCheckSuite(v CheckRunCheckSuite)`

SetCheckSuite sets CheckSuite field to given value.


### SetCheckSuiteNil

`func (o *CheckRun) SetCheckSuiteNil(b bool)`

 SetCheckSuiteNil sets the value for CheckSuite to be an explicit nil

### UnsetCheckSuite
`func (o *CheckRun) UnsetCheckSuite()`

UnsetCheckSuite ensures that no value is present for CheckSuite, not even an explicit nil
### GetApp

`func (o *CheckRun) GetApp() NullableIntegration`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *CheckRun) GetAppOk() (*NullableIntegration, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *CheckRun) SetApp(v NullableIntegration)`

SetApp sets App field to given value.


### SetAppNil

`func (o *CheckRun) SetAppNil(b bool)`

 SetAppNil sets the value for App to be an explicit nil

### UnsetApp
`func (o *CheckRun) UnsetApp()`

UnsetApp ensures that no value is present for App, not even an explicit nil
### GetPullRequests

`func (o *CheckRun) GetPullRequests() []PullRequestMinimal`

GetPullRequests returns the PullRequests field if non-nil, zero value otherwise.

### GetPullRequestsOk

`func (o *CheckRun) GetPullRequestsOk() (*[]PullRequestMinimal, bool)`

GetPullRequestsOk returns a tuple with the PullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequests

`func (o *CheckRun) SetPullRequests(v []PullRequestMinimal)`

SetPullRequests sets PullRequests field to given value.


### GetDeployment

`func (o *CheckRun) GetDeployment() DeploymentSimple`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *CheckRun) GetDeploymentOk() (*DeploymentSimple, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *CheckRun) SetDeployment(v DeploymentSimple)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *CheckRun) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


