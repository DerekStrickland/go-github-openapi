# CodespaceMachine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the machine. | 
**DisplayName** | **string** | The display name of the machine includes cores, memory, and storage. | 
**OperatingSystem** | **string** | The operating system of the machine. | 
**StorageInBytes** | **int32** | How much storage is available to the codespace. | 
**MemoryInBytes** | **int32** | How much memory is available to the codespace. | 
**Cpus** | **int32** | How many cores are available to the codespace. | 
**PrebuildAvailability** | **NullableString** | Whether a prebuild is currently available when creating a codespace for this machine and repository. If a branch was not specified as a ref, the default branch will be assumed. Value will be \&quot;null\&quot; if prebuilds are not supported or prebuild availability could not be determined. Value will be \&quot;none\&quot; if no prebuild is available. Latest values \&quot;ready\&quot; and \&quot;in_progress\&quot; indicate the prebuild availability status. | 

## Methods

### NewCodespaceMachine

`func NewCodespaceMachine(name string, displayName string, operatingSystem string, storageInBytes int32, memoryInBytes int32, cpus int32, prebuildAvailability NullableString, ) *CodespaceMachine`

NewCodespaceMachine instantiates a new CodespaceMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodespaceMachineWithDefaults

`func NewCodespaceMachineWithDefaults() *CodespaceMachine`

NewCodespaceMachineWithDefaults instantiates a new CodespaceMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CodespaceMachine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CodespaceMachine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CodespaceMachine) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *CodespaceMachine) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CodespaceMachine) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CodespaceMachine) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetOperatingSystem

`func (o *CodespaceMachine) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *CodespaceMachine) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *CodespaceMachine) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.


### GetStorageInBytes

`func (o *CodespaceMachine) GetStorageInBytes() int32`

GetStorageInBytes returns the StorageInBytes field if non-nil, zero value otherwise.

### GetStorageInBytesOk

`func (o *CodespaceMachine) GetStorageInBytesOk() (*int32, bool)`

GetStorageInBytesOk returns a tuple with the StorageInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInBytes

`func (o *CodespaceMachine) SetStorageInBytes(v int32)`

SetStorageInBytes sets StorageInBytes field to given value.


### GetMemoryInBytes

`func (o *CodespaceMachine) GetMemoryInBytes() int32`

GetMemoryInBytes returns the MemoryInBytes field if non-nil, zero value otherwise.

### GetMemoryInBytesOk

`func (o *CodespaceMachine) GetMemoryInBytesOk() (*int32, bool)`

GetMemoryInBytesOk returns a tuple with the MemoryInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryInBytes

`func (o *CodespaceMachine) SetMemoryInBytes(v int32)`

SetMemoryInBytes sets MemoryInBytes field to given value.


### GetCpus

`func (o *CodespaceMachine) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *CodespaceMachine) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *CodespaceMachine) SetCpus(v int32)`

SetCpus sets Cpus field to given value.


### GetPrebuildAvailability

`func (o *CodespaceMachine) GetPrebuildAvailability() string`

GetPrebuildAvailability returns the PrebuildAvailability field if non-nil, zero value otherwise.

### GetPrebuildAvailabilityOk

`func (o *CodespaceMachine) GetPrebuildAvailabilityOk() (*string, bool)`

GetPrebuildAvailabilityOk returns a tuple with the PrebuildAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrebuildAvailability

`func (o *CodespaceMachine) SetPrebuildAvailability(v string)`

SetPrebuildAvailability sets PrebuildAvailability field to given value.


### SetPrebuildAvailabilityNil

`func (o *CodespaceMachine) SetPrebuildAvailabilityNil(b bool)`

 SetPrebuildAvailabilityNil sets the value for PrebuildAvailability to be an explicit nil

### UnsetPrebuildAvailability
`func (o *CodespaceMachine) UnsetPrebuildAvailability()`

UnsetPrebuildAvailability ensures that no value is present for PrebuildAvailability, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


