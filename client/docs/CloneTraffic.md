# CloneTraffic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Uniques** | **int32** |  | 
**Clones** | [**[]Traffic**](Traffic.md) |  | 

## Methods

### NewCloneTraffic

`func NewCloneTraffic(count int32, uniques int32, clones []Traffic, ) *CloneTraffic`

NewCloneTraffic instantiates a new CloneTraffic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloneTrafficWithDefaults

`func NewCloneTrafficWithDefaults() *CloneTraffic`

NewCloneTrafficWithDefaults instantiates a new CloneTraffic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CloneTraffic) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CloneTraffic) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CloneTraffic) SetCount(v int32)`

SetCount sets Count field to given value.


### GetUniques

`func (o *CloneTraffic) GetUniques() int32`

GetUniques returns the Uniques field if non-nil, zero value otherwise.

### GetUniquesOk

`func (o *CloneTraffic) GetUniquesOk() (*int32, bool)`

GetUniquesOk returns a tuple with the Uniques field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniques

`func (o *CloneTraffic) SetUniques(v int32)`

SetUniques sets Uniques field to given value.


### GetClones

`func (o *CloneTraffic) GetClones() []Traffic`

GetClones returns the Clones field if non-nil, zero value otherwise.

### GetClonesOk

`func (o *CloneTraffic) GetClonesOk() (*[]Traffic, bool)`

GetClonesOk returns a tuple with the Clones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClones

`func (o *CloneTraffic) SetClones(v []Traffic)`

SetClones sets Clones field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


