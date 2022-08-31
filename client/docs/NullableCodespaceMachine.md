# NullableCodespaceMachine

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

### NewNullableCodespaceMachine

`func NewNullableCodespaceMachine(name string, displayName string, operatingSystem string, storageInBytes int32, memoryInBytes int32, cpus int32, prebuildAvailability NullableString, ) *NullableCodespaceMachine`

NewNullableCodespaceMachine instantiates a new NullableCodespaceMachine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNullableCodespaceMachineWithDefaults

`func NewNullableCodespaceMachineWithDefaults() *NullableCodespaceMachine`

NewNullableCodespaceMachineWithDefaults instantiates a new NullableCodespaceMachine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NullableCodespaceMachine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NullableCodespaceMachine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NullableCodespaceMachine) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *NullableCodespaceMachine) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NullableCodespaceMachine) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NullableCodespaceMachine) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetOperatingSystem

`func (o *NullableCodespaceMachine) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *NullableCodespaceMachine) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *NullableCodespaceMachine) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.


### GetStorageInBytes

`func (o *NullableCodespaceMachine) GetStorageInBytes() int32`

GetStorageInBytes returns the StorageInBytes field if non-nil, zero value otherwise.

### GetStorageInBytesOk

`func (o *NullableCodespaceMachine) GetStorageInBytesOk() (*int32, bool)`

GetStorageInBytesOk returns a tuple with the StorageInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInBytes

`func (o *NullableCodespaceMachine) SetStorageInBytes(v int32)`

SetStorageInBytes sets StorageInBytes field to given value.


### GetMemoryInBytes

`func (o *NullableCodespaceMachine) GetMemoryInBytes() int32`

GetMemoryInBytes returns the MemoryInBytes field if non-nil, zero value otherwise.

### GetMemoryInBytesOk

`func (o *NullableCodespaceMachine) GetMemoryInBytesOk() (*int32, bool)`

GetMemoryInBytesOk returns a tuple with the MemoryInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryInBytes

`func (o *NullableCodespaceMachine) SetMemoryInBytes(v int32)`

SetMemoryInBytes sets MemoryInBytes field to given value.


### GetCpus

`func (o *NullableCodespaceMachine) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *NullableCodespaceMachine) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *NullableCodespaceMachine) SetCpus(v int32)`

SetCpus sets Cpus field to given value.


### GetPrebuildAvailability

`func (o *NullableCodespaceMachine) GetPrebuildAvailability() string`

GetPrebuildAvailability returns the PrebuildAvailability field if non-nil, zero value otherwise.

### GetPrebuildAvailabilityOk

`func (o *NullableCodespaceMachine) GetPrebuildAvailabilityOk() (*string, bool)`

GetPrebuildAvailabilityOk returns a tuple with the PrebuildAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrebuildAvailability

`func (o *NullableCodespaceMachine) SetPrebuildAvailability(v string)`

SetPrebuildAvailability sets PrebuildAvailability field to given value.


### SetPrebuildAvailabilityNil

`func (o *NullableCodespaceMachine) SetPrebuildAvailabilityNil(b bool)`

 SetPrebuildAvailabilityNil sets the value for PrebuildAvailability to be an explicit nil

### UnsetPrebuildAvailability
`func (o *NullableCodespaceMachine) UnsetPrebuildAvailability()`

UnsetPrebuildAvailability ensures that no value is present for PrebuildAvailability, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


