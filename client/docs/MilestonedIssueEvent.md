# MilestonedIssueEvent

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
**Milestone** | [**MilestonedIssueEventMilestone**](MilestonedIssueEventMilestone.md) |  | 

## Methods

### NewMilestonedIssueEvent

`func NewMilestonedIssueEvent(id int32, nodeId string, url string, actor SimpleUser, event string, commitId NullableString, commitUrl NullableString, createdAt string, performedViaGithubApp NullableNullableIntegration, milestone MilestonedIssueEventMilestone, ) *MilestonedIssueEvent`

NewMilestonedIssueEvent instantiates a new MilestonedIssueEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMilestonedIssueEventWithDefaults

`func NewMilestonedIssueEventWithDefaults() *MilestonedIssueEvent`

NewMilestonedIssueEventWithDefaults instantiates a new MilestonedIssueEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MilestonedIssueEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MilestonedIssueEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MilestonedIssueEvent) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *MilestonedIssueEvent) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *MilestonedIssueEvent) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *MilestonedIssueEvent) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *MilestonedIssueEvent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MilestonedIssueEvent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MilestonedIssueEvent) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetActor

`func (o *MilestonedIssueEvent) GetActor() SimpleUser`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *MilestonedIssueEvent) GetActorOk() (*SimpleUser, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *MilestonedIssueEvent) SetActor(v SimpleUser)`

SetActor sets Actor field to given value.


### GetEvent

`func (o *MilestonedIssueEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *MilestonedIssueEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *MilestonedIssueEvent) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetCommitId

`func (o *MilestonedIssueEvent) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *MilestonedIssueEvent) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *MilestonedIssueEvent) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.


### SetCommitIdNil

`func (o *MilestonedIssueEvent) SetCommitIdNil(b bool)`

 SetCommitIdNil sets the value for CommitId to be an explicit nil

### UnsetCommitId
`func (o *MilestonedIssueEvent) UnsetCommitId()`

UnsetCommitId ensures that no value is present for CommitId, not even an explicit nil
### GetCommitUrl

`func (o *MilestonedIssueEvent) GetCommitUrl() string`

GetCommitUrl returns the CommitUrl field if non-nil, zero value otherwise.

### GetCommitUrlOk

`func (o *MilestonedIssueEvent) GetCommitUrlOk() (*string, bool)`

GetCommitUrlOk returns a tuple with the CommitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitUrl

`func (o *MilestonedIssueEvent) SetCommitUrl(v string)`

SetCommitUrl sets CommitUrl field to given value.


### SetCommitUrlNil

`func (o *MilestonedIssueEvent) SetCommitUrlNil(b bool)`

 SetCommitUrlNil sets the value for CommitUrl to be an explicit nil

### UnsetCommitUrl
`func (o *MilestonedIssueEvent) UnsetCommitUrl()`

UnsetCommitUrl ensures that no value is present for CommitUrl, not even an explicit nil
### GetCreatedAt

`func (o *MilestonedIssueEvent) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MilestonedIssueEvent) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MilestonedIssueEvent) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetPerformedViaGithubApp

`func (o *MilestonedIssueEvent) GetPerformedViaGithubApp() NullableIntegration`

GetPerformedViaGithubApp returns the PerformedViaGithubApp field if non-nil, zero value otherwise.

### GetPerformedViaGithubAppOk

`func (o *MilestonedIssueEvent) GetPerformedViaGithubAppOk() (*NullableIntegration, bool)`

GetPerformedViaGithubAppOk returns a tuple with the PerformedViaGithubApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformedViaGithubApp

`func (o *MilestonedIssueEvent) SetPerformedViaGithubApp(v NullableIntegration)`

SetPerformedViaGithubApp sets PerformedViaGithubApp field to given value.


### SetPerformedViaGithubAppNil

`func (o *MilestonedIssueEvent) SetPerformedViaGithubAppNil(b bool)`

 SetPerformedViaGithubAppNil sets the value for PerformedViaGithubApp to be an explicit nil

### UnsetPerformedViaGithubApp
`func (o *MilestonedIssueEvent) UnsetPerformedViaGithubApp()`

UnsetPerformedViaGithubApp ensures that no value is present for PerformedViaGithubApp, not even an explicit nil
### GetMilestone

`func (o *MilestonedIssueEvent) GetMilestone() MilestonedIssueEventMilestone`

GetMilestone returns the Milestone field if non-nil, zero value otherwise.

### GetMilestoneOk

`func (o *MilestonedIssueEvent) GetMilestoneOk() (*MilestonedIssueEventMilestone, bool)`

GetMilestoneOk returns a tuple with the Milestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestone

`func (o *MilestonedIssueEvent) SetMilestone(v MilestonedIssueEventMilestone)`

SetMilestone sets Milestone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


