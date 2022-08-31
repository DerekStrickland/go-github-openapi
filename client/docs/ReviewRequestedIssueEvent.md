# ReviewRequestedIssueEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**Url** | **string** |  | 
**Actor** | [**SimpleUser**](SimpleUser.md) |  | 
**Event** | **string** |  | 
**CommitId** | **NullableString** |  | 
**CommitUrl** | **NullableString** |  | 
**CreatedAt** | **string** |  | 
**PerformedViaGithubApp** | [**NullableNullableIntegration**](NullableIntegration.md) |  | 
**ReviewRequester** | [**SimpleUser**](SimpleUser.md) |  | 
**RequestedTeam** | Pointer to [**Team**](Team.md) |  | [optional] 
**RequestedReviewer** | Pointer to [**SimpleUser**](SimpleUser.md) |  | [optional] 

## Methods

### NewReviewRequestedIssueEvent

`func NewReviewRequestedIssueEvent(id int32, nodeId string, url string, actor SimpleUser, event string, commitId NullableString, commitUrl NullableString, createdAt string, performedViaGithubApp NullableNullableIntegration, reviewRequester SimpleUser, ) *ReviewRequestedIssueEvent`

NewReviewRequestedIssueEvent instantiates a new ReviewRequestedIssueEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewRequestedIssueEventWithDefaults

`func NewReviewRequestedIssueEventWithDefaults() *ReviewRequestedIssueEvent`

NewReviewRequestedIssueEventWithDefaults instantiates a new ReviewRequestedIssueEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReviewRequestedIssueEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReviewRequestedIssueEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReviewRequestedIssueEvent) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *ReviewRequestedIssueEvent) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ReviewRequestedIssueEvent) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ReviewRequestedIssueEvent) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *ReviewRequestedIssueEvent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ReviewRequestedIssueEvent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ReviewRequestedIssueEvent) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetActor

`func (o *ReviewRequestedIssueEvent) GetActor() SimpleUser`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *ReviewRequestedIssueEvent) GetActorOk() (*SimpleUser, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *ReviewRequestedIssueEvent) SetActor(v SimpleUser)`

SetActor sets Actor field to given value.


### GetEvent

`func (o *ReviewRequestedIssueEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *ReviewRequestedIssueEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *ReviewRequestedIssueEvent) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetCommitId

`func (o *ReviewRequestedIssueEvent) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *ReviewRequestedIssueEvent) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *ReviewRequestedIssueEvent) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.


### SetCommitIdNil

`func (o *ReviewRequestedIssueEvent) SetCommitIdNil(b bool)`

 SetCommitIdNil sets the value for CommitId to be an explicit nil

### UnsetCommitId
`func (o *ReviewRequestedIssueEvent) UnsetCommitId()`

UnsetCommitId ensures that no value is present for CommitId, not even an explicit nil
### GetCommitUrl

`func (o *ReviewRequestedIssueEvent) GetCommitUrl() string`

GetCommitUrl returns the CommitUrl field if non-nil, zero value otherwise.

### GetCommitUrlOk

`func (o *ReviewRequestedIssueEvent) GetCommitUrlOk() (*string, bool)`

GetCommitUrlOk returns a tuple with the CommitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitUrl

`func (o *ReviewRequestedIssueEvent) SetCommitUrl(v string)`

SetCommitUrl sets CommitUrl field to given value.


### SetCommitUrlNil

`func (o *ReviewRequestedIssueEvent) SetCommitUrlNil(b bool)`

 SetCommitUrlNil sets the value for CommitUrl to be an explicit nil

### UnsetCommitUrl
`func (o *ReviewRequestedIssueEvent) UnsetCommitUrl()`

UnsetCommitUrl ensures that no value is present for CommitUrl, not even an explicit nil
### GetCreatedAt

`func (o *ReviewRequestedIssueEvent) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReviewRequestedIssueEvent) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReviewRequestedIssueEvent) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetPerformedViaGithubApp

`func (o *ReviewRequestedIssueEvent) GetPerformedViaGithubApp() NullableIntegration`

GetPerformedViaGithubApp returns the PerformedViaGithubApp field if non-nil, zero value otherwise.

### GetPerformedViaGithubAppOk

`func (o *ReviewRequestedIssueEvent) GetPerformedViaGithubAppOk() (*NullableIntegration, bool)`

GetPerformedViaGithubAppOk returns a tuple with the PerformedViaGithubApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformedViaGithubApp

`func (o *ReviewRequestedIssueEvent) SetPerformedViaGithubApp(v NullableIntegration)`

SetPerformedViaGithubApp sets PerformedViaGithubApp field to given value.


### SetPerformedViaGithubAppNil

`func (o *ReviewRequestedIssueEvent) SetPerformedViaGithubAppNil(b bool)`

 SetPerformedViaGithubAppNil sets the value for PerformedViaGithubApp to be an explicit nil

### UnsetPerformedViaGithubApp
`func (o *ReviewRequestedIssueEvent) UnsetPerformedViaGithubApp()`

UnsetPerformedViaGithubApp ensures that no value is present for PerformedViaGithubApp, not even an explicit nil
### GetReviewRequester

`func (o *ReviewRequestedIssueEvent) GetReviewRequester() SimpleUser`

GetReviewRequester returns the ReviewRequester field if non-nil, zero value otherwise.

### GetReviewRequesterOk

`func (o *ReviewRequestedIssueEvent) GetReviewRequesterOk() (*SimpleUser, bool)`

GetReviewRequesterOk returns a tuple with the ReviewRequester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewRequester

`func (o *ReviewRequestedIssueEvent) SetReviewRequester(v SimpleUser)`

SetReviewRequester sets ReviewRequester field to given value.


### GetRequestedTeam

`func (o *ReviewRequestedIssueEvent) GetRequestedTeam() Team`

GetRequestedTeam returns the RequestedTeam field if non-nil, zero value otherwise.

### GetRequestedTeamOk

`func (o *ReviewRequestedIssueEvent) GetRequestedTeamOk() (*Team, bool)`

GetRequestedTeamOk returns a tuple with the RequestedTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTeam

`func (o *ReviewRequestedIssueEvent) SetRequestedTeam(v Team)`

SetRequestedTeam sets RequestedTeam field to given value.

### HasRequestedTeam

`func (o *ReviewRequestedIssueEvent) HasRequestedTeam() bool`

HasRequestedTeam returns a boolean if a field has been set.

### GetRequestedReviewer

`func (o *ReviewRequestedIssueEvent) GetRequestedReviewer() SimpleUser`

GetRequestedReviewer returns the RequestedReviewer field if non-nil, zero value otherwise.

### GetRequestedReviewerOk

`func (o *ReviewRequestedIssueEvent) GetRequestedReviewerOk() (*SimpleUser, bool)`

GetRequestedReviewerOk returns a tuple with the RequestedReviewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedReviewer

`func (o *ReviewRequestedIssueEvent) SetRequestedReviewer(v SimpleUser)`

SetRequestedReviewer sets RequestedReviewer field to given value.

### HasRequestedReviewer

`func (o *ReviewRequestedIssueEvent) HasRequestedReviewer() bool`

HasRequestedReviewer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


