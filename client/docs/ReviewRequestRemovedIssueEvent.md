# ReviewRequestRemovedIssueEvent

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

### NewReviewRequestRemovedIssueEvent

`func NewReviewRequestRemovedIssueEvent(id int32, nodeId string, url string, actor SimpleUser, event string, commitId NullableString, commitUrl NullableString, createdAt string, performedViaGithubApp NullableNullableIntegration, reviewRequester SimpleUser, ) *ReviewRequestRemovedIssueEvent`

NewReviewRequestRemovedIssueEvent instantiates a new ReviewRequestRemovedIssueEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewRequestRemovedIssueEventWithDefaults

`func NewReviewRequestRemovedIssueEventWithDefaults() *ReviewRequestRemovedIssueEvent`

NewReviewRequestRemovedIssueEventWithDefaults instantiates a new ReviewRequestRemovedIssueEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReviewRequestRemovedIssueEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReviewRequestRemovedIssueEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReviewRequestRemovedIssueEvent) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *ReviewRequestRemovedIssueEvent) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ReviewRequestRemovedIssueEvent) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ReviewRequestRemovedIssueEvent) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *ReviewRequestRemovedIssueEvent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ReviewRequestRemovedIssueEvent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ReviewRequestRemovedIssueEvent) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetActor

`func (o *ReviewRequestRemovedIssueEvent) GetActor() SimpleUser`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *ReviewRequestRemovedIssueEvent) GetActorOk() (*SimpleUser, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *ReviewRequestRemovedIssueEvent) SetActor(v SimpleUser)`

SetActor sets Actor field to given value.


### GetEvent

`func (o *ReviewRequestRemovedIssueEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *ReviewRequestRemovedIssueEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *ReviewRequestRemovedIssueEvent) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetCommitId

`func (o *ReviewRequestRemovedIssueEvent) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *ReviewRequestRemovedIssueEvent) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *ReviewRequestRemovedIssueEvent) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.


### SetCommitIdNil

`func (o *ReviewRequestRemovedIssueEvent) SetCommitIdNil(b bool)`

 SetCommitIdNil sets the value for CommitId to be an explicit nil

### UnsetCommitId
`func (o *ReviewRequestRemovedIssueEvent) UnsetCommitId()`

UnsetCommitId ensures that no value is present for CommitId, not even an explicit nil
### GetCommitUrl

`func (o *ReviewRequestRemovedIssueEvent) GetCommitUrl() string`

GetCommitUrl returns the CommitUrl field if non-nil, zero value otherwise.

### GetCommitUrlOk

`func (o *ReviewRequestRemovedIssueEvent) GetCommitUrlOk() (*string, bool)`

GetCommitUrlOk returns a tuple with the CommitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitUrl

`func (o *ReviewRequestRemovedIssueEvent) SetCommitUrl(v string)`

SetCommitUrl sets CommitUrl field to given value.


### SetCommitUrlNil

`func (o *ReviewRequestRemovedIssueEvent) SetCommitUrlNil(b bool)`

 SetCommitUrlNil sets the value for CommitUrl to be an explicit nil

### UnsetCommitUrl
`func (o *ReviewRequestRemovedIssueEvent) UnsetCommitUrl()`

UnsetCommitUrl ensures that no value is present for CommitUrl, not even an explicit nil
### GetCreatedAt

`func (o *ReviewRequestRemovedIssueEvent) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReviewRequestRemovedIssueEvent) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReviewRequestRemovedIssueEvent) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetPerformedViaGithubApp

`func (o *ReviewRequestRemovedIssueEvent) GetPerformedViaGithubApp() NullableIntegration`

GetPerformedViaGithubApp returns the PerformedViaGithubApp field if non-nil, zero value otherwise.

### GetPerformedViaGithubAppOk

`func (o *ReviewRequestRemovedIssueEvent) GetPerformedViaGithubAppOk() (*NullableIntegration, bool)`

GetPerformedViaGithubAppOk returns a tuple with the PerformedViaGithubApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformedViaGithubApp

`func (o *ReviewRequestRemovedIssueEvent) SetPerformedViaGithubApp(v NullableIntegration)`

SetPerformedViaGithubApp sets PerformedViaGithubApp field to given value.


### SetPerformedViaGithubAppNil

`func (o *ReviewRequestRemovedIssueEvent) SetPerformedViaGithubAppNil(b bool)`

 SetPerformedViaGithubAppNil sets the value for PerformedViaGithubApp to be an explicit nil

### UnsetPerformedViaGithubApp
`func (o *ReviewRequestRemovedIssueEvent) UnsetPerformedViaGithubApp()`

UnsetPerformedViaGithubApp ensures that no value is present for PerformedViaGithubApp, not even an explicit nil
### GetReviewRequester

`func (o *ReviewRequestRemovedIssueEvent) GetReviewRequester() SimpleUser`

GetReviewRequester returns the ReviewRequester field if non-nil, zero value otherwise.

### GetReviewRequesterOk

`func (o *ReviewRequestRemovedIssueEvent) GetReviewRequesterOk() (*SimpleUser, bool)`

GetReviewRequesterOk returns a tuple with the ReviewRequester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewRequester

`func (o *ReviewRequestRemovedIssueEvent) SetReviewRequester(v SimpleUser)`

SetReviewRequester sets ReviewRequester field to given value.


### GetRequestedTeam

`func (o *ReviewRequestRemovedIssueEvent) GetRequestedTeam() Team`

GetRequestedTeam returns the RequestedTeam field if non-nil, zero value otherwise.

### GetRequestedTeamOk

`func (o *ReviewRequestRemovedIssueEvent) GetRequestedTeamOk() (*Team, bool)`

GetRequestedTeamOk returns a tuple with the RequestedTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedTeam

`func (o *ReviewRequestRemovedIssueEvent) SetRequestedTeam(v Team)`

SetRequestedTeam sets RequestedTeam field to given value.

### HasRequestedTeam

`func (o *ReviewRequestRemovedIssueEvent) HasRequestedTeam() bool`

HasRequestedTeam returns a boolean if a field has been set.

### GetRequestedReviewer

`func (o *ReviewRequestRemovedIssueEvent) GetRequestedReviewer() SimpleUser`

GetRequestedReviewer returns the RequestedReviewer field if non-nil, zero value otherwise.

### GetRequestedReviewerOk

`func (o *ReviewRequestRemovedIssueEvent) GetRequestedReviewerOk() (*SimpleUser, bool)`

GetRequestedReviewerOk returns a tuple with the RequestedReviewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedReviewer

`func (o *ReviewRequestRemovedIssueEvent) SetRequestedReviewer(v SimpleUser)`

SetRequestedReviewer sets RequestedReviewer field to given value.

### HasRequestedReviewer

`func (o *ReviewRequestRemovedIssueEvent) HasRequestedReviewer() bool`

HasRequestedReviewer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


