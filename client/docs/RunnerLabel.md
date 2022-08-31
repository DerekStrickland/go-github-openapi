# RunnerLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique identifier of the label. | [optional] 
**Name** | **string** | Name of the label. | 
**Type** | Pointer to **string** | The type of label. Read-only labels are applied automatically when the runner is configured. | [optional] 

## Methods

### NewRunnerLabel

`func NewRunnerLabel(name string, ) *RunnerLabel`

NewRunnerLabel instantiates a new RunnerLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerLabelWithDefaults

`func NewRunnerLabelWithDefaults() *RunnerLabel`

NewRunnerLabelWithDefaults instantiates a new RunnerLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RunnerLabel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RunnerLabel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RunnerLabel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *RunnerLabel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RunnerLabel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RunnerLabel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RunnerLabel) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *RunnerLabel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RunnerLabel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RunnerLabel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RunnerLabel) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


