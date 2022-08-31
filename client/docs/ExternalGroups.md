# ExternalGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]ExternalGroupsGroupsInner**](ExternalGroupsGroupsInner.md) | An array of external groups available to be mapped to a team | [optional] 

## Methods

### NewExternalGroups

`func NewExternalGroups() *ExternalGroups`

NewExternalGroups instantiates a new ExternalGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalGroupsWithDefaults

`func NewExternalGroupsWithDefaults() *ExternalGroups`

NewExternalGroupsWithDefaults instantiates a new ExternalGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *ExternalGroups) GetGroups() []ExternalGroupsGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ExternalGroups) GetGroupsOk() (*[]ExternalGroupsGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ExternalGroups) SetGroups(v []ExternalGroupsGroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *ExternalGroups) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


