# IssuesUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to [**NullableIssuesUpdateRequestTitle**](IssuesUpdateRequestTitle.md) |  | [optional] 
**Body** | Pointer to **NullableString** | The contents of the issue. | [optional] 
**Assignee** | Pointer to **NullableString** | Login for the user that this issue should be assigned to. **This field is deprecated.** | [optional] 
**State** | Pointer to **string** | State of the issue. Either &#x60;open&#x60; or &#x60;closed&#x60;. | [optional] 
**StateReason** | Pointer to **NullableString** | The reason for the current state | [optional] 
**Milestone** | Pointer to [**NullableIssuesUpdateRequestMilestone**](IssuesUpdateRequestMilestone.md) |  | [optional] 
**Labels** | Pointer to [**[]IssuesCreateRequestLabelsInner**](IssuesCreateRequestLabelsInner.md) | Labels to associate with this issue. Pass one or more Labels to _replace_ the set of Labels on this Issue. Send an empty array (&#x60;[]&#x60;) to clear all Labels from the Issue. _NOTE: Only users with push access can set labels for issues. Labels are silently dropped otherwise._ | [optional] 
**Assignees** | Pointer to **[]string** | Logins for Users to assign to this issue. Pass one or more user logins to _replace_ the set of assignees on this Issue. Send an empty array (&#x60;[]&#x60;) to clear all assignees from the Issue. _NOTE: Only users with push access can set assignees for new issues. Assignees are silently dropped otherwise._ | [optional] 

## Methods

### NewIssuesUpdateRequest

`func NewIssuesUpdateRequest() *IssuesUpdateRequest`

NewIssuesUpdateRequest instantiates a new IssuesUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuesUpdateRequestWithDefaults

`func NewIssuesUpdateRequestWithDefaults() *IssuesUpdateRequest`

NewIssuesUpdateRequestWithDefaults instantiates a new IssuesUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *IssuesUpdateRequest) GetTitle() IssuesUpdateRequestTitle`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IssuesUpdateRequest) GetTitleOk() (*IssuesUpdateRequestTitle, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IssuesUpdateRequest) SetTitle(v IssuesUpdateRequestTitle)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *IssuesUpdateRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *IssuesUpdateRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *IssuesUpdateRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetBody

`func (o *IssuesUpdateRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *IssuesUpdateRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *IssuesUpdateRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *IssuesUpdateRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *IssuesUpdateRequest) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *IssuesUpdateRequest) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetAssignee

`func (o *IssuesUpdateRequest) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *IssuesUpdateRequest) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *IssuesUpdateRequest) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *IssuesUpdateRequest) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### SetAssigneeNil

`func (o *IssuesUpdateRequest) SetAssigneeNil(b bool)`

 SetAssigneeNil sets the value for Assignee to be an explicit nil

### UnsetAssignee
`func (o *IssuesUpdateRequest) UnsetAssignee()`

UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil
### GetState

`func (o *IssuesUpdateRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IssuesUpdateRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IssuesUpdateRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IssuesUpdateRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateReason

`func (o *IssuesUpdateRequest) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *IssuesUpdateRequest) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *IssuesUpdateRequest) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *IssuesUpdateRequest) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### SetStateReasonNil

`func (o *IssuesUpdateRequest) SetStateReasonNil(b bool)`

 SetStateReasonNil sets the value for StateReason to be an explicit nil

### UnsetStateReason
`func (o *IssuesUpdateRequest) UnsetStateReason()`

UnsetStateReason ensures that no value is present for StateReason, not even an explicit nil
### GetMilestone

`func (o *IssuesUpdateRequest) GetMilestone() IssuesUpdateRequestMilestone`

GetMilestone returns the Milestone field if non-nil, zero value otherwise.

### GetMilestoneOk

`func (o *IssuesUpdateRequest) GetMilestoneOk() (*IssuesUpdateRequestMilestone, bool)`

GetMilestoneOk returns a tuple with the Milestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestone

`func (o *IssuesUpdateRequest) SetMilestone(v IssuesUpdateRequestMilestone)`

SetMilestone sets Milestone field to given value.

### HasMilestone

`func (o *IssuesUpdateRequest) HasMilestone() bool`

HasMilestone returns a boolean if a field has been set.

### SetMilestoneNil

`func (o *IssuesUpdateRequest) SetMilestoneNil(b bool)`

 SetMilestoneNil sets the value for Milestone to be an explicit nil

### UnsetMilestone
`func (o *IssuesUpdateRequest) UnsetMilestone()`

UnsetMilestone ensures that no value is present for Milestone, not even an explicit nil
### GetLabels

`func (o *IssuesUpdateRequest) GetLabels() []IssuesCreateRequestLabelsInner`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *IssuesUpdateRequest) GetLabelsOk() (*[]IssuesCreateRequestLabelsInner, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *IssuesUpdateRequest) SetLabels(v []IssuesCreateRequestLabelsInner)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *IssuesUpdateRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAssignees

`func (o *IssuesUpdateRequest) GetAssignees() []string`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *IssuesUpdateRequest) GetAssigneesOk() (*[]string, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *IssuesUpdateRequest) SetAssignees(v []string)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *IssuesUpdateRequest) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


