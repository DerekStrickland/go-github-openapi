# IssuesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | [**IssuesCreateRequestTitle**](IssuesCreateRequestTitle.md) |  | 
**Body** | Pointer to **string** | The contents of the issue. | [optional] 
**Assignee** | Pointer to **NullableString** | Login for the user that this issue should be assigned to. _NOTE: Only users with push access can set the assignee for new issues. The assignee is silently dropped otherwise. **This field is deprecated.**_ | [optional] 
**Milestone** | Pointer to [**NullableIssuesCreateRequestMilestone**](IssuesCreateRequestMilestone.md) |  | [optional] 
**Labels** | Pointer to [**[]IssuesCreateRequestLabelsInner**](IssuesCreateRequestLabelsInner.md) | Labels to associate with this issue. _NOTE: Only users with push access can set labels for new issues. Labels are silently dropped otherwise._ | [optional] 
**Assignees** | Pointer to **[]string** | Logins for Users to assign to this issue. _NOTE: Only users with push access can set assignees for new issues. Assignees are silently dropped otherwise._ | [optional] 

## Methods

### NewIssuesCreateRequest

`func NewIssuesCreateRequest(title IssuesCreateRequestTitle, ) *IssuesCreateRequest`

NewIssuesCreateRequest instantiates a new IssuesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuesCreateRequestWithDefaults

`func NewIssuesCreateRequestWithDefaults() *IssuesCreateRequest`

NewIssuesCreateRequestWithDefaults instantiates a new IssuesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *IssuesCreateRequest) GetTitle() IssuesCreateRequestTitle`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *IssuesCreateRequest) GetTitleOk() (*IssuesCreateRequestTitle, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *IssuesCreateRequest) SetTitle(v IssuesCreateRequestTitle)`

SetTitle sets Title field to given value.


### GetBody

`func (o *IssuesCreateRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *IssuesCreateRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *IssuesCreateRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *IssuesCreateRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetAssignee

`func (o *IssuesCreateRequest) GetAssignee() string`

GetAssignee returns the Assignee field if non-nil, zero value otherwise.

### GetAssigneeOk

`func (o *IssuesCreateRequest) GetAssigneeOk() (*string, bool)`

GetAssigneeOk returns a tuple with the Assignee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignee

`func (o *IssuesCreateRequest) SetAssignee(v string)`

SetAssignee sets Assignee field to given value.

### HasAssignee

`func (o *IssuesCreateRequest) HasAssignee() bool`

HasAssignee returns a boolean if a field has been set.

### SetAssigneeNil

`func (o *IssuesCreateRequest) SetAssigneeNil(b bool)`

 SetAssigneeNil sets the value for Assignee to be an explicit nil

### UnsetAssignee
`func (o *IssuesCreateRequest) UnsetAssignee()`

UnsetAssignee ensures that no value is present for Assignee, not even an explicit nil
### GetMilestone

`func (o *IssuesCreateRequest) GetMilestone() IssuesCreateRequestMilestone`

GetMilestone returns the Milestone field if non-nil, zero value otherwise.

### GetMilestoneOk

`func (o *IssuesCreateRequest) GetMilestoneOk() (*IssuesCreateRequestMilestone, bool)`

GetMilestoneOk returns a tuple with the Milestone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMilestone

`func (o *IssuesCreateRequest) SetMilestone(v IssuesCreateRequestMilestone)`

SetMilestone sets Milestone field to given value.

### HasMilestone

`func (o *IssuesCreateRequest) HasMilestone() bool`

HasMilestone returns a boolean if a field has been set.

### SetMilestoneNil

`func (o *IssuesCreateRequest) SetMilestoneNil(b bool)`

 SetMilestoneNil sets the value for Milestone to be an explicit nil

### UnsetMilestone
`func (o *IssuesCreateRequest) UnsetMilestone()`

UnsetMilestone ensures that no value is present for Milestone, not even an explicit nil
### GetLabels

`func (o *IssuesCreateRequest) GetLabels() []IssuesCreateRequestLabelsInner`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *IssuesCreateRequest) GetLabelsOk() (*[]IssuesCreateRequestLabelsInner, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *IssuesCreateRequest) SetLabels(v []IssuesCreateRequestLabelsInner)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *IssuesCreateRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetAssignees

`func (o *IssuesCreateRequest) GetAssignees() []string`

GetAssignees returns the Assignees field if non-nil, zero value otherwise.

### GetAssigneesOk

`func (o *IssuesCreateRequest) GetAssigneesOk() (*[]string, bool)`

GetAssigneesOk returns a tuple with the Assignees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignees

`func (o *IssuesCreateRequest) SetAssignees(v []string)`

SetAssignees sets Assignees field to given value.

### HasAssignees

`func (o *IssuesCreateRequest) HasAssignees() bool`

HasAssignees returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


