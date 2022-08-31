# EventPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Issue** | Pointer to [**Issue**](Issue.md) |  | [optional] 
**Comment** | Pointer to [**IssueComment**](IssueComment.md) |  | [optional] 
**Pages** | Pointer to [**[]EventPayloadPagesInner**](EventPayloadPagesInner.md) |  | [optional] 

## Methods

### NewEventPayload

`func NewEventPayload() *EventPayload`

NewEventPayload instantiates a new EventPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventPayloadWithDefaults

`func NewEventPayloadWithDefaults() *EventPayload`

NewEventPayloadWithDefaults instantiates a new EventPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *EventPayload) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *EventPayload) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *EventPayload) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *EventPayload) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetIssue

`func (o *EventPayload) GetIssue() Issue`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *EventPayload) GetIssueOk() (*Issue, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *EventPayload) SetIssue(v Issue)`

SetIssue sets Issue field to given value.

### HasIssue

`func (o *EventPayload) HasIssue() bool`

HasIssue returns a boolean if a field has been set.

### GetComment

`func (o *EventPayload) GetComment() IssueComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *EventPayload) GetCommentOk() (*IssueComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *EventPayload) SetComment(v IssueComment)`

SetComment sets Comment field to given value.

### HasComment

`func (o *EventPayload) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetPages

`func (o *EventPayload) GetPages() []EventPayloadPagesInner`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *EventPayload) GetPagesOk() (*[]EventPayloadPagesInner, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *EventPayload) SetPages(v []EventPayloadPagesInner)`

SetPages sets Pages field to given value.

### HasPages

`func (o *EventPayload) HasPages() bool`

HasPages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


