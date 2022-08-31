# ConvertedNoteToIssueIssueEvent

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
**PerformedViaGithubApp** | [**Integration**](Integration.md) |  | 
**ProjectCard** | Pointer to [**AddedToProjectIssueEventProjectCard**](AddedToProjectIssueEventProjectCard.md) |  | [optional] 

## Methods

### NewConvertedNoteToIssueIssueEvent

`func NewConvertedNoteToIssueIssueEvent(id int32, nodeId string, url string, actor SimpleUser, event string, commitId NullableString, commitUrl NullableString, createdAt string, performedViaGithubApp Integration, ) *ConvertedNoteToIssueIssueEvent`

NewConvertedNoteToIssueIssueEvent instantiates a new ConvertedNoteToIssueIssueEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConvertedNoteToIssueIssueEventWithDefaults

`func NewConvertedNoteToIssueIssueEventWithDefaults() *ConvertedNoteToIssueIssueEvent`

NewConvertedNoteToIssueIssueEventWithDefaults instantiates a new ConvertedNoteToIssueIssueEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConvertedNoteToIssueIssueEvent) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConvertedNoteToIssueIssueEvent) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConvertedNoteToIssueIssueEvent) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *ConvertedNoteToIssueIssueEvent) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ConvertedNoteToIssueIssueEvent) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ConvertedNoteToIssueIssueEvent) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *ConvertedNoteToIssueIssueEvent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ConvertedNoteToIssueIssueEvent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ConvertedNoteToIssueIssueEvent) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetActor

`func (o *ConvertedNoteToIssueIssueEvent) GetActor() SimpleUser`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *ConvertedNoteToIssueIssueEvent) GetActorOk() (*SimpleUser, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *ConvertedNoteToIssueIssueEvent) SetActor(v SimpleUser)`

SetActor sets Actor field to given value.


### GetEvent

`func (o *ConvertedNoteToIssueIssueEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *ConvertedNoteToIssueIssueEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *ConvertedNoteToIssueIssueEvent) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetCommitId

`func (o *ConvertedNoteToIssueIssueEvent) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *ConvertedNoteToIssueIssueEvent) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *ConvertedNoteToIssueIssueEvent) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.


### SetCommitIdNil

`func (o *ConvertedNoteToIssueIssueEvent) SetCommitIdNil(b bool)`

 SetCommitIdNil sets the value for CommitId to be an explicit nil

### UnsetCommitId
`func (o *ConvertedNoteToIssueIssueEvent) UnsetCommitId()`

UnsetCommitId ensures that no value is present for CommitId, not even an explicit nil
### GetCommitUrl

`func (o *ConvertedNoteToIssueIssueEvent) GetCommitUrl() string`

GetCommitUrl returns the CommitUrl field if non-nil, zero value otherwise.

### GetCommitUrlOk

`func (o *ConvertedNoteToIssueIssueEvent) GetCommitUrlOk() (*string, bool)`

GetCommitUrlOk returns a tuple with the CommitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitUrl

`func (o *ConvertedNoteToIssueIssueEvent) SetCommitUrl(v string)`

SetCommitUrl sets CommitUrl field to given value.


### SetCommitUrlNil

`func (o *ConvertedNoteToIssueIssueEvent) SetCommitUrlNil(b bool)`

 SetCommitUrlNil sets the value for CommitUrl to be an explicit nil

### UnsetCommitUrl
`func (o *ConvertedNoteToIssueIssueEvent) UnsetCommitUrl()`

UnsetCommitUrl ensures that no value is present for CommitUrl, not even an explicit nil
### GetCreatedAt

`func (o *ConvertedNoteToIssueIssueEvent) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConvertedNoteToIssueIssueEvent) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConvertedNoteToIssueIssueEvent) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetPerformedViaGithubApp

`func (o *ConvertedNoteToIssueIssueEvent) GetPerformedViaGithubApp() Integration`

GetPerformedViaGithubApp returns the PerformedViaGithubApp field if non-nil, zero value otherwise.

### GetPerformedViaGithubAppOk

`func (o *ConvertedNoteToIssueIssueEvent) GetPerformedViaGithubAppOk() (*Integration, bool)`

GetPerformedViaGithubAppOk returns a tuple with the PerformedViaGithubApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformedViaGithubApp

`func (o *ConvertedNoteToIssueIssueEvent) SetPerformedViaGithubApp(v Integration)`

SetPerformedViaGithubApp sets PerformedViaGithubApp field to given value.


### GetProjectCard

`func (o *ConvertedNoteToIssueIssueEvent) GetProjectCard() AddedToProjectIssueEventProjectCard`

GetProjectCard returns the ProjectCard field if non-nil, zero value otherwise.

### GetProjectCardOk

`func (o *ConvertedNoteToIssueIssueEvent) GetProjectCardOk() (*AddedToProjectIssueEventProjectCard, bool)`

GetProjectCardOk returns a tuple with the ProjectCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectCard

`func (o *ConvertedNoteToIssueIssueEvent) SetProjectCard(v AddedToProjectIssueEventProjectCard)`

SetProjectCard sets ProjectCard field to given value.

### HasProjectCard

`func (o *ConvertedNoteToIssueIssueEvent) HasProjectCard() bool`

HasProjectCard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


