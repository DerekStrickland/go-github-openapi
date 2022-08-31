# Runner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The id of the runner. | 
**Name** | **string** | The name of the runner. | 
**Os** | **string** | The Operating System of the runner. | 
**Status** | **string** | The status of the runner. | 
**Busy** | **bool** |  | 
**Labels** | [**[]RunnerLabel**](RunnerLabel.md) |  | 

## Methods

### NewRunner

`func NewRunner(id int32, name string, os string, status string, busy bool, labels []RunnerLabel, ) *Runner`

NewRunner instantiates a new Runner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunnerWithDefaults

`func NewRunnerWithDefaults() *Runner`

NewRunnerWithDefaults instantiates a new Runner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Runner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Runner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Runner) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *Runner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Runner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Runner) SetName(v string)`

SetName sets Name field to given value.


### GetOs

`func (o *Runner) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *Runner) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *Runner) SetOs(v string)`

SetOs sets Os field to given value.


### GetStatus

`func (o *Runner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Runner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Runner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetBusy

`func (o *Runner) GetBusy() bool`

GetBusy returns the Busy field if non-nil, zero value otherwise.

### GetBusyOk

`func (o *Runner) GetBusyOk() (*bool, bool)`

GetBusyOk returns a tuple with the Busy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusy

`func (o *Runner) SetBusy(v bool)`

SetBusy sets Busy field to given value.


### GetLabels

`func (o *Runner) GetLabels() []RunnerLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Runner) GetLabelsOk() (*[]RunnerLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Runner) SetLabels(v []RunnerLabel)`

SetLabels sets Labels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


