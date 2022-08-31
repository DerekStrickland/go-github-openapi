# SimpleInstallation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The ID of the installation. | 
**NodeId** | **string** | The global node ID of the installation. | 

## Methods

### NewSimpleInstallation

`func NewSimpleInstallation(id int32, nodeId string, ) *SimpleInstallation`

NewSimpleInstallation instantiates a new SimpleInstallation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleInstallationWithDefaults

`func NewSimpleInstallationWithDefaults() *SimpleInstallation`

NewSimpleInstallationWithDefaults instantiates a new SimpleInstallation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimpleInstallation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimpleInstallation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimpleInstallation) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *SimpleInstallation) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *SimpleInstallation) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *SimpleInstallation) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


