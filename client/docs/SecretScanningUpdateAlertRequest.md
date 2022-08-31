# SecretScanningUpdateAlertRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**SecretScanningAlertState**](SecretScanningAlertState.md) |  | 
**Resolution** | Pointer to [**NullableSecretScanningAlertResolution**](SecretScanningAlertResolution.md) |  | [optional] 

## Methods

### NewSecretScanningUpdateAlertRequest

`func NewSecretScanningUpdateAlertRequest(state SecretScanningAlertState, ) *SecretScanningUpdateAlertRequest`

NewSecretScanningUpdateAlertRequest instantiates a new SecretScanningUpdateAlertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretScanningUpdateAlertRequestWithDefaults

`func NewSecretScanningUpdateAlertRequestWithDefaults() *SecretScanningUpdateAlertRequest`

NewSecretScanningUpdateAlertRequestWithDefaults instantiates a new SecretScanningUpdateAlertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *SecretScanningUpdateAlertRequest) GetState() SecretScanningAlertState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SecretScanningUpdateAlertRequest) GetStateOk() (*SecretScanningAlertState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SecretScanningUpdateAlertRequest) SetState(v SecretScanningAlertState)`

SetState sets State field to given value.


### GetResolution

`func (o *SecretScanningUpdateAlertRequest) GetResolution() SecretScanningAlertResolution`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *SecretScanningUpdateAlertRequest) GetResolutionOk() (*SecretScanningAlertResolution, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *SecretScanningUpdateAlertRequest) SetResolution(v SecretScanningAlertResolution)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *SecretScanningUpdateAlertRequest) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### SetResolutionNil

`func (o *SecretScanningUpdateAlertRequest) SetResolutionNil(b bool)`

 SetResolutionNil sets the value for Resolution to be an explicit nil

### UnsetResolution
`func (o *SecretScanningUpdateAlertRequest) UnsetResolution()`

UnsetResolution ensures that no value is present for Resolution, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


