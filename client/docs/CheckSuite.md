# CheckSuite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**HeadBranch** | **NullableString** |  | 
**HeadSha** | **string** | The SHA of the head commit that is being checked. | 
**Status** | **NullableString** |  | 
**Conclusion** | **NullableString** |  | 
**Url** | **NullableString** |  | 
**Before** | **NullableString** |  | 
**After** | **NullableString** |  | 
**PullRequests** | [**[]PullRequestMinimal**](PullRequestMinimal.md) |  | 
**App** | [**NullableNullableIntegration**](NullableIntegration.md) |  | 
**Repository** | [**MinimalRepository**](MinimalRepository.md) |  | 
**CreatedAt** | **NullableTime** |  | 
**UpdatedAt** | **NullableTime** |  | 
**HeadCommit** | [**SimpleCommit**](SimpleCommit.md) |  | 
**LatestCheckRunsCount** | **int32** |  | 
**CheckRunsUrl** | **string** |  | 
**Rerequestable** | Pointer to **bool** |  | [optional] 
**RunsRerequestable** | Pointer to **bool** |  | [optional] 

## Methods

### NewCheckSuite

`func NewCheckSuite(id int32, nodeId string, headBranch NullableString, headSha string, status NullableString, conclusion NullableString, url NullableString, before NullableString, after NullableString, pullRequests []PullRequestMinimal, app NullableNullableIntegration, repository MinimalRepository, createdAt NullableTime, updatedAt NullableTime, headCommit SimpleCommit, latestCheckRunsCount int32, checkRunsUrl string, ) *CheckSuite`

NewCheckSuite instantiates a new CheckSuite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckSuiteWithDefaults

`func NewCheckSuiteWithDefaults() *CheckSuite`

NewCheckSuiteWithDefaults instantiates a new CheckSuite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CheckSuite) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CheckSuite) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CheckSuite) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *CheckSuite) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *CheckSuite) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *CheckSuite) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetHeadBranch

`func (o *CheckSuite) GetHeadBranch() string`

GetHeadBranch returns the HeadBranch field if non-nil, zero value otherwise.

### GetHeadBranchOk

`func (o *CheckSuite) GetHeadBranchOk() (*string, bool)`

GetHeadBranchOk returns a tuple with the HeadBranch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadBranch

`func (o *CheckSuite) SetHeadBranch(v string)`

SetHeadBranch sets HeadBranch field to given value.


### SetHeadBranchNil

`func (o *CheckSuite) SetHeadBranchNil(b bool)`

 SetHeadBranchNil sets the value for HeadBranch to be an explicit nil

### UnsetHeadBranch
`func (o *CheckSuite) UnsetHeadBranch()`

UnsetHeadBranch ensures that no value is present for HeadBranch, not even an explicit nil
### GetHeadSha

`func (o *CheckSuite) GetHeadSha() string`

GetHeadSha returns the HeadSha field if non-nil, zero value otherwise.

### GetHeadShaOk

`func (o *CheckSuite) GetHeadShaOk() (*string, bool)`

GetHeadShaOk returns a tuple with the HeadSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadSha

`func (o *CheckSuite) SetHeadSha(v string)`

SetHeadSha sets HeadSha field to given value.


### GetStatus

`func (o *CheckSuite) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CheckSuite) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CheckSuite) SetStatus(v string)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *CheckSuite) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *CheckSuite) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetConclusion

`func (o *CheckSuite) GetConclusion() string`

GetConclusion returns the Conclusion field if non-nil, zero value otherwise.

### GetConclusionOk

`func (o *CheckSuite) GetConclusionOk() (*string, bool)`

GetConclusionOk returns a tuple with the Conclusion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConclusion

`func (o *CheckSuite) SetConclusion(v string)`

SetConclusion sets Conclusion field to given value.


### SetConclusionNil

`func (o *CheckSuite) SetConclusionNil(b bool)`

 SetConclusionNil sets the value for Conclusion to be an explicit nil

### UnsetConclusion
`func (o *CheckSuite) UnsetConclusion()`

UnsetConclusion ensures that no value is present for Conclusion, not even an explicit nil
### GetUrl

`func (o *CheckSuite) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CheckSuite) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CheckSuite) SetUrl(v string)`

SetUrl sets Url field to given value.


### SetUrlNil

`func (o *CheckSuite) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *CheckSuite) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetBefore

`func (o *CheckSuite) GetBefore() string`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *CheckSuite) GetBeforeOk() (*string, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *CheckSuite) SetBefore(v string)`

SetBefore sets Before field to given value.


### SetBeforeNil

`func (o *CheckSuite) SetBeforeNil(b bool)`

 SetBeforeNil sets the value for Before to be an explicit nil

### UnsetBefore
`func (o *CheckSuite) UnsetBefore()`

UnsetBefore ensures that no value is present for Before, not even an explicit nil
### GetAfter

`func (o *CheckSuite) GetAfter() string`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *CheckSuite) GetAfterOk() (*string, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *CheckSuite) SetAfter(v string)`

SetAfter sets After field to given value.


### SetAfterNil

`func (o *CheckSuite) SetAfterNil(b bool)`

 SetAfterNil sets the value for After to be an explicit nil

### UnsetAfter
`func (o *CheckSuite) UnsetAfter()`

UnsetAfter ensures that no value is present for After, not even an explicit nil
### GetPullRequests

`func (o *CheckSuite) GetPullRequests() []PullRequestMinimal`

GetPullRequests returns the PullRequests field if non-nil, zero value otherwise.

### GetPullRequestsOk

`func (o *CheckSuite) GetPullRequestsOk() (*[]PullRequestMinimal, bool)`

GetPullRequestsOk returns a tuple with the PullRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPullRequests

`func (o *CheckSuite) SetPullRequests(v []PullRequestMinimal)`

SetPullRequests sets PullRequests field to given value.


### SetPullRequestsNil

`func (o *CheckSuite) SetPullRequestsNil(b bool)`

 SetPullRequestsNil sets the value for PullRequests to be an explicit nil

### UnsetPullRequests
`func (o *CheckSuite) UnsetPullRequests()`

UnsetPullRequests ensures that no value is present for PullRequests, not even an explicit nil
### GetApp

`func (o *CheckSuite) GetApp() NullableIntegration`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *CheckSuite) GetAppOk() (*NullableIntegration, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *CheckSuite) SetApp(v NullableIntegration)`

SetApp sets App field to given value.


### SetAppNil

`func (o *CheckSuite) SetAppNil(b bool)`

 SetAppNil sets the value for App to be an explicit nil

### UnsetApp
`func (o *CheckSuite) UnsetApp()`

UnsetApp ensures that no value is present for App, not even an explicit nil
### GetRepository

`func (o *CheckSuite) GetRepository() MinimalRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *CheckSuite) GetRepositoryOk() (*MinimalRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *CheckSuite) SetRepository(v MinimalRepository)`

SetRepository sets Repository field to given value.


### GetCreatedAt

`func (o *CheckSuite) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CheckSuite) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CheckSuite) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *CheckSuite) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *CheckSuite) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *CheckSuite) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CheckSuite) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CheckSuite) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *CheckSuite) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *CheckSuite) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetHeadCommit

`func (o *CheckSuite) GetHeadCommit() SimpleCommit`

GetHeadCommit returns the HeadCommit field if non-nil, zero value otherwise.

### GetHeadCommitOk

`func (o *CheckSuite) GetHeadCommitOk() (*SimpleCommit, bool)`

GetHeadCommitOk returns a tuple with the HeadCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadCommit

`func (o *CheckSuite) SetHeadCommit(v SimpleCommit)`

SetHeadCommit sets HeadCommit field to given value.


### GetLatestCheckRunsCount

`func (o *CheckSuite) GetLatestCheckRunsCount() int32`

GetLatestCheckRunsCount returns the LatestCheckRunsCount field if non-nil, zero value otherwise.

### GetLatestCheckRunsCountOk

`func (o *CheckSuite) GetLatestCheckRunsCountOk() (*int32, bool)`

GetLatestCheckRunsCountOk returns a tuple with the LatestCheckRunsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestCheckRunsCount

`func (o *CheckSuite) SetLatestCheckRunsCount(v int32)`

SetLatestCheckRunsCount sets LatestCheckRunsCount field to given value.


### GetCheckRunsUrl

`func (o *CheckSuite) GetCheckRunsUrl() string`

GetCheckRunsUrl returns the CheckRunsUrl field if non-nil, zero value otherwise.

### GetCheckRunsUrlOk

`func (o *CheckSuite) GetCheckRunsUrlOk() (*string, bool)`

GetCheckRunsUrlOk returns a tuple with the CheckRunsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckRunsUrl

`func (o *CheckSuite) SetCheckRunsUrl(v string)`

SetCheckRunsUrl sets CheckRunsUrl field to given value.


### GetRerequestable

`func (o *CheckSuite) GetRerequestable() bool`

GetRerequestable returns the Rerequestable field if non-nil, zero value otherwise.

### GetRerequestableOk

`func (o *CheckSuite) GetRerequestableOk() (*bool, bool)`

GetRerequestableOk returns a tuple with the Rerequestable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerequestable

`func (o *CheckSuite) SetRerequestable(v bool)`

SetRerequestable sets Rerequestable field to given value.

### HasRerequestable

`func (o *CheckSuite) HasRerequestable() bool`

HasRerequestable returns a boolean if a field has been set.

### GetRunsRerequestable

`func (o *CheckSuite) GetRunsRerequestable() bool`

GetRunsRerequestable returns the RunsRerequestable field if non-nil, zero value otherwise.

### GetRunsRerequestableOk

`func (o *CheckSuite) GetRunsRerequestableOk() (*bool, bool)`

GetRunsRerequestableOk returns a tuple with the RunsRerequestable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunsRerequestable

`func (o *CheckSuite) SetRunsRerequestable(v bool)`

SetRunsRerequestable sets RunsRerequestable field to given value.

### HasRunsRerequestable

`func (o *CheckSuite) HasRunsRerequestable() bool`

HasRunsRerequestable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


