# ActionsCreateWorkflowDispatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | **string** | The git reference for the workflow. The reference can be a branch or tag name. | 
**Inputs** | Pointer to **map[string]string** | Input keys and values configured in the workflow file. The maximum number of properties is 10. Any default properties configured in the workflow file will be used when &#x60;inputs&#x60; are omitted. | [optional] 

## Methods

### NewActionsCreateWorkflowDispatchRequest

`func NewActionsCreateWorkflowDispatchRequest(ref string, ) *ActionsCreateWorkflowDispatchRequest`

NewActionsCreateWorkflowDispatchRequest instantiates a new ActionsCreateWorkflowDispatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsCreateWorkflowDispatchRequestWithDefaults

`func NewActionsCreateWorkflowDispatchRequestWithDefaults() *ActionsCreateWorkflowDispatchRequest`

NewActionsCreateWorkflowDispatchRequestWithDefaults instantiates a new ActionsCreateWorkflowDispatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *ActionsCreateWorkflowDispatchRequest) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ActionsCreateWorkflowDispatchRequest) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ActionsCreateWorkflowDispatchRequest) SetRef(v string)`

SetRef sets Ref field to given value.


### GetInputs

`func (o *ActionsCreateWorkflowDispatchRequest) GetInputs() map[string]string`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *ActionsCreateWorkflowDispatchRequest) GetInputsOk() (*map[string]string, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *ActionsCreateWorkflowDispatchRequest) SetInputs(v map[string]string)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *ActionsCreateWorkflowDispatchRequest) HasInputs() bool`

HasInputs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


