# RenamedIssueEvent

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
**Rename** | [**RenamedIssueEventRename**](RenamedIssueEventRename.md) |  | 

## Methods

### NewRenamedIssueEvent

`func NewRenamedIssueEvent(id int32, nodeId string, url string, actor SimpleUser, event string, commitId NullableString, commitUrl NullableString, createdAt string, performedViaGithubApp NullableNullableIntegration, rename RenamedIssueEventRename, ) *RenamedIssueEvent`

NewRenamedIssueEvent instantiates a new RenamedIssueEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRenamedIssueEventWithDefaults

`func NewRenamedIssueEventWithDefaults() *RenamedIssueEvent`

NewRenamedIssueEventWithDefaults instantiates a new RenamedIssueEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RenamedIssueEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RenamedIssueEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RenamedIssueEvent) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *RenamedIssueEvent) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *RenamedIssueEvent) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *RenamedIssueEvent) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *RenamedIssueEvent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RenamedIssueEvent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RenamedIssueEvent) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetActor

`func (o *RenamedIssueEvent) GetActor() SimpleUser`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *RenamedIssueEvent) GetActorOk() (*SimpleUser, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *RenamedIssueEvent) SetActor(v SimpleUser)`

SetActor sets Actor field to given value.


### GetEvent

`func (o *RenamedIssueEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *RenamedIssueEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *RenamedIssueEvent) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetCommitId

`func (o *RenamedIssueEvent) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *RenamedIssueEvent) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *RenamedIssueEvent) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.


### SetCommitIdNil

`func (o *RenamedIssueEvent) SetCommitIdNil(b bool)`

 SetCommitIdNil sets the value for CommitId to be an explicit nil

### UnsetCommitId
`func (o *RenamedIssueEvent) UnsetCommitId()`

UnsetCommitId ensures that no value is present for CommitId, not even an explicit nil
### GetCommitUrl

`func (o *RenamedIssueEvent) GetCommitUrl() string`

GetCommitUrl returns the CommitUrl field if non-nil, zero value otherwise.

### GetCommitUrlOk

`func (o *RenamedIssueEvent) GetCommitUrlOk() (*string, bool)`

GetCommitUrlOk returns a tuple with the CommitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitUrl

`func (o *RenamedIssueEvent) SetCommitUrl(v string)`

SetCommitUrl sets CommitUrl field to given value.


### SetCommitUrlNil

`func (o *RenamedIssueEvent) SetCommitUrlNil(b bool)`

 SetCommitUrlNil sets the value for CommitUrl to be an explicit nil

### UnsetCommitUrl
`func (o *RenamedIssueEvent) UnsetCommitUrl()`

UnsetCommitUrl ensures that no value is present for CommitUrl, not even an explicit nil
### GetCreatedAt

`func (o *RenamedIssueEvent) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RenamedIssueEvent) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RenamedIssueEvent) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetPerformedViaGithubApp

`func (o *RenamedIssueEvent) GetPerformedViaGithubApp() NullableIntegration`

GetPerformedViaGithubApp returns the PerformedViaGithubApp field if non-nil, zero value otherwise.

### GetPerformedViaGithubAppOk

`func (o *RenamedIssueEvent) GetPerformedViaGithubAppOk() (*NullableIntegration, bool)`

GetPerformedViaGithubAppOk returns a tuple with the PerformedViaGithubApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformedViaGithubApp

`func (o *RenamedIssueEvent) SetPerformedViaGithubApp(v NullableIntegration)`

SetPerformedViaGithubApp sets PerformedViaGithubApp field to given value.


### SetPerformedViaGithubAppNil

`func (o *RenamedIssueEvent) SetPerformedViaGithubAppNil(b bool)`

 SetPerformedViaGithubAppNil sets the value for PerformedViaGithubApp to be an explicit nil

### UnsetPerformedViaGithubApp
`func (o *RenamedIssueEvent) UnsetPerformedViaGithubApp()`

UnsetPerformedViaGithubApp ensures that no value is present for PerformedViaGithubApp, not even an explicit nil
### GetRename

`func (o *RenamedIssueEvent) GetRename() RenamedIssueEventRename`

GetRename returns the Rename field if non-nil, zero value otherwise.

### GetRenameOk

`func (o *RenamedIssueEvent) GetRenameOk() (*RenamedIssueEventRename, bool)`

GetRenameOk returns a tuple with the Rename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRename

`func (o *RenamedIssueEvent) SetRename(v RenamedIssueEventRename)`

SetRename sets Rename field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


