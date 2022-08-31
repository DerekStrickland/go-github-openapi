# ActionsGetDefaultWorkflowPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultWorkflowPermissions** | [**ActionsDefaultWorkflowPermissions**](ActionsDefaultWorkflowPermissions.md) |  | 
**CanApprovePullRequestReviews** | **bool** | Whether GitHub Actions can approve pull requests. Enabling this can be a security risk. | 

## Methods

### NewActionsGetDefaultWorkflowPermissions

`func NewActionsGetDefaultWorkflowPermissions(defaultWorkflowPermissions ActionsDefaultWorkflowPermissions, canApprovePullRequestReviews bool, ) *ActionsGetDefaultWorkflowPermissions`

NewActionsGetDefaultWorkflowPermissions instantiates a new ActionsGetDefaultWorkflowPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsGetDefaultWorkflowPermissionsWithDefaults

`func NewActionsGetDefaultWorkflowPermissionsWithDefaults() *ActionsGetDefaultWorkflowPermissions`

NewActionsGetDefaultWorkflowPermissionsWithDefaults instantiates a new ActionsGetDefaultWorkflowPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultWorkflowPermissions

`func (o *ActionsGetDefaultWorkflowPermissions) GetDefaultWorkflowPermissions() ActionsDefaultWorkflowPermissions`

GetDefaultWorkflowPermissions returns the DefaultWorkflowPermissions field if non-nil, zero value otherwise.

### GetDefaultWorkflowPermissionsOk

`func (o *ActionsGetDefaultWorkflowPermissions) GetDefaultWorkflowPermissionsOk() (*ActionsDefaultWorkflowPermissions, bool)`

GetDefaultWorkflowPermissionsOk returns a tuple with the DefaultWorkflowPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWorkflowPermissions

`func (o *ActionsGetDefaultWorkflowPermissions) SetDefaultWorkflowPermissions(v ActionsDefaultWorkflowPermissions)`

SetDefaultWorkflowPermissions sets DefaultWorkflowPermissions field to given value.


### GetCanApprovePullRequestReviews

`func (o *ActionsGetDefaultWorkflowPermissions) GetCanApprovePullRequestReviews() bool`

GetCanApprovePullRequestReviews returns the CanApprovePullRequestReviews field if non-nil, zero value otherwise.

### GetCanApprovePullRequestReviewsOk

`func (o *ActionsGetDefaultWorkflowPermissions) GetCanApprovePullRequestReviewsOk() (*bool, bool)`

GetCanApprovePullRequestReviewsOk returns a tuple with the CanApprovePullRequestReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanApprovePullRequestReviews

`func (o *ActionsGetDefaultWorkflowPermissions) SetCanApprovePullRequestReviews(v bool)`

SetCanApprovePullRequestReviews sets CanApprovePullRequestReviews field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


