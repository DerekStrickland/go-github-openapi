# ActionsWorkflowAccessToRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessLevel** | **string** | Defines the level of access that workflows outside of the repository have to actions and reusable workflows within the repository. &#x60;none&#x60; means access is only possible from workflows in this repository. | 

## Methods

### NewActionsWorkflowAccessToRepository

`func NewActionsWorkflowAccessToRepository(accessLevel string, ) *ActionsWorkflowAccessToRepository`

NewActionsWorkflowAccessToRepository instantiates a new ActionsWorkflowAccessToRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsWorkflowAccessToRepositoryWithDefaults

`func NewActionsWorkflowAccessToRepositoryWithDefaults() *ActionsWorkflowAccessToRepository`

NewActionsWorkflowAccessToRepositoryWithDefaults instantiates a new ActionsWorkflowAccessToRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessLevel

`func (o *ActionsWorkflowAccessToRepository) GetAccessLevel() string`

GetAccessLevel returns the AccessLevel field if non-nil, zero value otherwise.

### GetAccessLevelOk

`func (o *ActionsWorkflowAccessToRepository) GetAccessLevelOk() (*string, bool)`

GetAccessLevelOk returns a tuple with the AccessLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessLevel

`func (o *ActionsWorkflowAccessToRepository) SetAccessLevel(v string)`

SetAccessLevel sets AccessLevel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


