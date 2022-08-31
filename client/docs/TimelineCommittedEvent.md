# TimelineCommittedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Event** | Pointer to **string** |  | [optional] 
**Sha** | **string** | SHA for the commit | 
**NodeId** | **string** |  | 
**Url** | **string** |  | 
**Author** | [**GitCommitAuthor**](GitCommitAuthor.md) |  | 
**Committer** | [**GitCommitAuthor**](GitCommitAuthor.md) |  | 
**Message** | **string** | Message describing the purpose of the commit | 
**Tree** | [**GitCommitTree**](GitCommitTree.md) |  | 
**Parents** | [**[]GitCommitParentsInner**](GitCommitParentsInner.md) |  | 
**Verification** | [**GitCommitVerification**](GitCommitVerification.md) |  | 
**HtmlUrl** | **string** |  | 

## Methods

### NewTimelineCommittedEvent

`func NewTimelineCommittedEvent(sha string, nodeId string, url string, author GitCommitAuthor, committer GitCommitAuthor, message string, tree GitCommitTree, parents []GitCommitParentsInner, verification GitCommitVerification, htmlUrl string, ) *TimelineCommittedEvent`

NewTimelineCommittedEvent instantiates a new TimelineCommittedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineCommittedEventWithDefaults

`func NewTimelineCommittedEventWithDefaults() *TimelineCommittedEvent`

NewTimelineCommittedEventWithDefaults instantiates a new TimelineCommittedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvent

`func (o *TimelineCommittedEvent) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *TimelineCommittedEvent) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *TimelineCommittedEvent) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *TimelineCommittedEvent) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetSha

`func (o *TimelineCommittedEvent) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *TimelineCommittedEvent) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *TimelineCommittedEvent) SetSha(v string)`

SetSha sets Sha field to given value.


### GetNodeId

`func (o *TimelineCommittedEvent) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *TimelineCommittedEvent) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *TimelineCommittedEvent) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *TimelineCommittedEvent) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TimelineCommittedEvent) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TimelineCommittedEvent) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAuthor

`func (o *TimelineCommittedEvent) GetAuthor() GitCommitAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *TimelineCommittedEvent) GetAuthorOk() (*GitCommitAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *TimelineCommittedEvent) SetAuthor(v GitCommitAuthor)`

SetAuthor sets Author field to given value.


### GetCommitter

`func (o *TimelineCommittedEvent) GetCommitter() GitCommitAuthor`

GetCommitter returns the Committer field if non-nil, zero value otherwise.

### GetCommitterOk

`func (o *TimelineCommittedEvent) GetCommitterOk() (*GitCommitAuthor, bool)`

GetCommitterOk returns a tuple with the Committer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitter

`func (o *TimelineCommittedEvent) SetCommitter(v GitCommitAuthor)`

SetCommitter sets Committer field to given value.


### GetMessage

`func (o *TimelineCommittedEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TimelineCommittedEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TimelineCommittedEvent) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTree

`func (o *TimelineCommittedEvent) GetTree() GitCommitTree`

GetTree returns the Tree field if non-nil, zero value otherwise.

### GetTreeOk

`func (o *TimelineCommittedEvent) GetTreeOk() (*GitCommitTree, bool)`

GetTreeOk returns a tuple with the Tree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTree

`func (o *TimelineCommittedEvent) SetTree(v GitCommitTree)`

SetTree sets Tree field to given value.


### GetParents

`func (o *TimelineCommittedEvent) GetParents() []GitCommitParentsInner`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *TimelineCommittedEvent) GetParentsOk() (*[]GitCommitParentsInner, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *TimelineCommittedEvent) SetParents(v []GitCommitParentsInner)`

SetParents sets Parents field to given value.


### GetVerification

`func (o *TimelineCommittedEvent) GetVerification() GitCommitVerification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *TimelineCommittedEvent) GetVerificationOk() (*GitCommitVerification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *TimelineCommittedEvent) SetVerification(v GitCommitVerification)`

SetVerification sets Verification field to given value.


### GetHtmlUrl

`func (o *TimelineCommittedEvent) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *TimelineCommittedEvent) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *TimelineCommittedEvent) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


