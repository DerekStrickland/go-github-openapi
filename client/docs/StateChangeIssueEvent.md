# StateChangeIssueEvent

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
**StateReason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewStateChangeIssueEvent

`func NewStateChangeIssueEvent(id int32, nodeId string, url string, actor SimpleUser, event string, commitId NullableString, commitUrl NullableString, createdAt string, performedViaGithubApp NullableNullableIntegration, ) *StateChangeIssueEvent`

NewStateChangeIssueEvent instantiates a new StateChangeIssueEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateChangeIssueEventWithDefaults

`func NewStateChangeIssueEventWithDefaults() *StateChangeIssueEvent`

NewStateChangeIssueEventWithDefaults instantiates a new StateChangeIssueEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StateChangeIssueEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StateChangeIssueEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StateChangeIssueEvent) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *StateChangeIssueEvent) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *StateChangeIssueEvent) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *StateChangeIssueEvent) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *StateChangeIssueEvent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StateChangeIssueEvent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StateChangeIssueEvent) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetActor

`func (o *StateChangeIssueEvent) GetActor() SimpleUser`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *StateChangeIssueEvent) GetActorOk() (*SimpleUser, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *StateChangeIssueEvent) SetActor(v SimpleUser)`

SetActor sets Actor field to given value.


### GetEvent

`func (o *StateChangeIssueEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *StateChangeIssueEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *StateChangeIssueEvent) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetCommitId

`func (o *StateChangeIssueEvent) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *StateChangeIssueEvent) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *StateChangeIssueEvent) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.


### SetCommitIdNil

`func (o *StateChangeIssueEvent) SetCommitIdNil(b bool)`

 SetCommitIdNil sets the value for CommitId to be an explicit nil

### UnsetCommitId
`func (o *StateChangeIssueEvent) UnsetCommitId()`

UnsetCommitId ensures that no value is present for CommitId, not even an explicit nil
### GetCommitUrl

`func (o *StateChangeIssueEvent) GetCommitUrl() string`

GetCommitUrl returns the CommitUrl field if non-nil, zero value otherwise.

### GetCommitUrlOk

`func (o *StateChangeIssueEvent) GetCommitUrlOk() (*string, bool)`

GetCommitUrlOk returns a tuple with the CommitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitUrl

`func (o *StateChangeIssueEvent) SetCommitUrl(v string)`

SetCommitUrl sets CommitUrl field to given value.


### SetCommitUrlNil

`func (o *StateChangeIssueEvent) SetCommitUrlNil(b bool)`

 SetCommitUrlNil sets the value for CommitUrl to be an explicit nil

### UnsetCommitUrl
`func (o *StateChangeIssueEvent) UnsetCommitUrl()`

UnsetCommitUrl ensures that no value is present for CommitUrl, not even an explicit nil
### GetCreatedAt

`func (o *StateChangeIssueEvent) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StateChangeIssueEvent) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StateChangeIssueEvent) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetPerformedViaGithubApp

`func (o *StateChangeIssueEvent) GetPerformedViaGithubApp() NullableIntegration`

GetPerformedViaGithubApp returns the PerformedViaGithubApp field if non-nil, zero value otherwise.

### GetPerformedViaGithubAppOk

`func (o *StateChangeIssueEvent) GetPerformedViaGithubAppOk() (*NullableIntegration, bool)`

GetPerformedViaGithubAppOk returns a tuple with the PerformedViaGithubApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformedViaGithubApp

`func (o *StateChangeIssueEvent) SetPerformedViaGithubApp(v NullableIntegration)`

SetPerformedViaGithubApp sets PerformedViaGithubApp field to given value.


### SetPerformedViaGithubAppNil

`func (o *StateChangeIssueEvent) SetPerformedViaGithubAppNil(b bool)`

 SetPerformedViaGithubAppNil sets the value for PerformedViaGithubApp to be an explicit nil

### UnsetPerformedViaGithubApp
`func (o *StateChangeIssueEvent) UnsetPerformedViaGithubApp()`

UnsetPerformedViaGithubApp ensures that no value is present for PerformedViaGithubApp, not even an explicit nil
### GetStateReason

`func (o *StateChangeIssueEvent) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *StateChangeIssueEvent) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *StateChangeIssueEvent) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *StateChangeIssueEvent) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### SetStateReasonNil

`func (o *StateChangeIssueEvent) SetStateReasonNil(b bool)`

 SetStateReasonNil sets the value for StateReason to be an explicit nil

### UnsetStateReason
`func (o *StateChangeIssueEvent) UnsetStateReason()`

UnsetStateReason ensures that no value is present for StateReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


