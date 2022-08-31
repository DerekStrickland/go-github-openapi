# EnvironmentApprovals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environments** | [**[]EnvironmentApprovalsEnvironmentsInner**](EnvironmentApprovalsEnvironmentsInner.md) | The list of environments that were approved or rejected | 
**State** | **string** | Whether deployment to the environment(s) was approved or rejected | 
**User** | [**SimpleUser**](SimpleUser.md) |  | 
**Comment** | **string** | The comment submitted with the deployment review | 

## Methods

### NewEnvironmentApprovals

`func NewEnvironmentApprovals(environments []EnvironmentApprovalsEnvironmentsInner, state string, user SimpleUser, comment string, ) *EnvironmentApprovals`

NewEnvironmentApprovals instantiates a new EnvironmentApprovals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentApprovalsWithDefaults

`func NewEnvironmentApprovalsWithDefaults() *EnvironmentApprovals`

NewEnvironmentApprovalsWithDefaults instantiates a new EnvironmentApprovals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironments

`func (o *EnvironmentApprovals) GetEnvironments() []EnvironmentApprovalsEnvironmentsInner`

GetEnvironments returns the Environments field if non-nil, zero value otherwise.

### GetEnvironmentsOk

`func (o *EnvironmentApprovals) GetEnvironmentsOk() (*[]EnvironmentApprovalsEnvironmentsInner, bool)`

GetEnvironmentsOk returns a tuple with the Environments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironments

`func (o *EnvironmentApprovals) SetEnvironments(v []EnvironmentApprovalsEnvironmentsInner)`

SetEnvironments sets Environments field to given value.


### GetState

`func (o *EnvironmentApprovals) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EnvironmentApprovals) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EnvironmentApprovals) SetState(v string)`

SetState sets State field to given value.


### GetUser

`func (o *EnvironmentApprovals) GetUser() SimpleUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *EnvironmentApprovals) GetUserOk() (*SimpleUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *EnvironmentApprovals) SetUser(v SimpleUser)`

SetUser sets User field to given value.


### GetComment

`func (o *EnvironmentApprovals) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *EnvironmentApprovals) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *EnvironmentApprovals) SetComment(v string)`

SetComment sets Comment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


