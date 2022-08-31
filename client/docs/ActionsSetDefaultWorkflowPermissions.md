# ActionsSetDefaultWorkflowPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultWorkflowPermissions** | Pointer to [**ActionsDefaultWorkflowPermissions**](ActionsDefaultWorkflowPermissions.md) |  | [optional] 
**CanApprovePullRequestReviews** | Pointer to **bool** | Whether GitHub Actions can approve pull requests. Enabling this can be a security risk. | [optional] 

## Methods

### NewActionsSetDefaultWorkflowPermissions

`func NewActionsSetDefaultWorkflowPermissions() *ActionsSetDefaultWorkflowPermissions`

NewActionsSetDefaultWorkflowPermissions instantiates a new ActionsSetDefaultWorkflowPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsSetDefaultWorkflowPermissionsWithDefaults

`func NewActionsSetDefaultWorkflowPermissionsWithDefaults() *ActionsSetDefaultWorkflowPermissions`

NewActionsSetDefaultWorkflowPermissionsWithDefaults instantiates a new ActionsSetDefaultWorkflowPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultWorkflowPermissions

`func (o *ActionsSetDefaultWorkflowPermissions) GetDefaultWorkflowPermissions() ActionsDefaultWorkflowPermissions`

GetDefaultWorkflowPermissions returns the DefaultWorkflowPermissions field if non-nil, zero value otherwise.

### GetDefaultWorkflowPermissionsOk

`func (o *ActionsSetDefaultWorkflowPermissions) GetDefaultWorkflowPermissionsOk() (*ActionsDefaultWorkflowPermissions, bool)`

GetDefaultWorkflowPermissionsOk returns a tuple with the DefaultWorkflowPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWorkflowPermissions

`func (o *ActionsSetDefaultWorkflowPermissions) SetDefaultWorkflowPermissions(v ActionsDefaultWorkflowPermissions)`

SetDefaultWorkflowPermissions sets DefaultWorkflowPermissions field to given value.

### HasDefaultWorkflowPermissions

`func (o *ActionsSetDefaultWorkflowPermissions) HasDefaultWorkflowPermissions() bool`

HasDefaultWorkflowPermissions returns a boolean if a field has been set.

### GetCanApprovePullRequestReviews

`func (o *ActionsSetDefaultWorkflowPermissions) GetCanApprovePullRequestReviews() bool`

GetCanApprovePullRequestReviews returns the CanApprovePullRequestReviews field if non-nil, zero value otherwise.

### GetCanApprovePullRequestReviewsOk

`func (o *ActionsSetDefaultWorkflowPermissions) GetCanApprovePullRequestReviewsOk() (*bool, bool)`

GetCanApprovePullRequestReviewsOk returns a tuple with the CanApprovePullRequestReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanApprovePullRequestReviews

`func (o *ActionsSetDefaultWorkflowPermissions) SetCanApprovePullRequestReviews(v bool)`

SetCanApprovePullRequestReviews sets CanApprovePullRequestReviews field to given value.

### HasCanApprovePullRequestReviews

`func (o *ActionsSetDefaultWorkflowPermissions) HasCanApprovePullRequestReviews() bool`

HasCanApprovePullRequestReviews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


