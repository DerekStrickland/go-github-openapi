# TimelineUnassignedIssueEvent

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
**Assignee** | [**SimpleUser**](SimpleUser.md) |  | 

## Methods

### NewTimelineUnassignedIssueEvent

`func NewTimelineUnassignedIssueEvent(id int32, nodeId string, url string, actor SimpleUser, event string, commitId NullableString, commitUrl NullableString, createdAt string, performedViaGithubApp NullableNullableIntegration, assignee SimpleUser, ) *TimelineUnassignedIssueEvent`

NewTimelineUnassignedIssueEvent instantiates a new TimelineUnassignedIssueEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineUnassignedIssueEventWithDefaults

`func NewTimelineUnassignedIssueEventWithDefaults() *TimelineUnassignedIssueEvent`

NewTimelineUnassignedIssueEventWithDefaults instantiates a new TimelineUnassignedIssueEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimelineUnassignedIssueEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimelineUnassignedIssueEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimelineUnassignedIssueEvent) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *TimelineUnassignedIssueEvent) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *TimelineUnassignedIssueEvent) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *TimelineUnassignedIssueEvent) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *TimelineUnassignedIssueEvent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TimelineUnassignedIssueEvent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TimelineUnassignedIssueEvent) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetActor

`func (o *TimelineUnassignedIssueEvent) GetActor() SimpleUser`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *TimelineUnassignedIssueEvent) GetActorOk() (*SimpleUser, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *TimelineUnassignedIssueEvent) SetActor(v SimpleUser)`

SetActor sets Actor field to given value.


### GetEvent

`func (o *TimelineUnassignedIssueEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *TimelineUnassignedIssueEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *TimelineUnassignedIssueEvent) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetCommitId

`func (o *TimelineUnassignedIssueEvent) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *TimelineUnassignedIssueEvent) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *TimelineUnassignedIssueEvent) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.


### SetCommitIdNil

`func (o *TimelineUnassignedIssueEvent) SetCommitIdNil(b bool)`

 SetCommitIdNil sets the value for CommitId to be an explicit nil

### UnsetCommitId
`func (o *TimelineUnassignedIssueEvent) UnsetCommitId()`

UnsetCommitId ensures that no value is present for CommitId, not even an explicit nil
### GetCommitUrl

`func (o *TimelineUnassignedIssueEvent) GetCommitUrl() string`

GetCommitUrl returns the CommitUrl field if non-nil, zero value otherwise.

### GetCommitUrlOk

`func (o *TimelineUnassignedIssueEvent) GetCommitUrlOk() (*string, bool)`

GetCommitUrlOk returns a tuple with the CommitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitUrl

`func (o *TimelineUnassignedIssueEvent) SetCommitUrl(v string)`

SetCommitUrl sets CommitUrl field to given value.


### SetCommitUrlNil

`func (o *TimelineUnassignedIssueEvent) SetCommitUrlNil(b bool)`

 SetCommitUrlNil sets the value for CommitUrl to be an explicit nil

### UnsetCommitUrl
`func (o *TimelineUnassignedIssueEvent) UnsetCommitUrl()`

UnsetCommitUrl ensures that no value is present for CommitUrl, not even an explicit nil
### GetCreatedAt

`func (o *TimelineUnassignedIssueEvent) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TimelineUnassignedIssueEvent) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TimelineUnassignedIssueEvent) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetPerformedViaGithubApp

`func (o *TimelineUnassignedIssueEvent) GetPerformedViaGithubApp() NullableIntegration`

GetPerformedViaGithubApp returns the PerformedViaGithubApp field if non-nil, zero value otherwise.

### GetPerformedViaGithubAppOk

`func (o *TimelineUnassignedIssueEvent) GetPerformedViaGithubAppOk() (*NullableIntegration, bool)`

GetPerformedViaGithubAppOk returns a tuple with the PerformedViaGithubApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformedViaGithubApp

`func (o *TimelineUnassignedIssueEvent) SetPerformedViaGithubApp(v NullableIntegration)`

SetPerformedViaGithubApp sets PerformedViaGithubApp field to given value.


### SetPerformedViaGithubAppNil

`func (o *TimelineUnassignedIssueEvent) SetPerformedViaGithubAppNil(b bool)`

 SetPerformedViaGithubAppNil sets the value for PerformedViaGithubApp to be an explicit nil

### UnsetPerformedViaGithubApp
`func (o *TimelineUnassignedIssueEvent) UnsetPerformedViaGithubApp()`

UnsetPerformedViaGithubApp ensures that no value is present for PerformedViaGithubApp, not even an explicit nil
### GetAssignee

`func (o *TimelineUnassignedIssueEvent) GetAssignee() SimpleUser`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *TimelineUnassignedIssueEvent) GetAssigneeOk() (*SimpleUser, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *TimelineUnassignedIssueEvent) SetAssignee(v SimpleUser)`

SetAssignee sets Assignee field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


