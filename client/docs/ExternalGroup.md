# ExternalGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **int32** | The internal ID of the group | 
**GroupName** | **string** | The display name for the group | 
**UpdatedAt** | Pointer to **string** | The date when the group was last updated_at | [optional] 
**Teams** | [**[]ExternalGroupTeamsInner**](ExternalGroupTeamsInner.md) | An array of teams linked to this group | 
**Members** | [**[]ExternalGroupMembersInner**](ExternalGroupMembersInner.md) | An array of external members linked to this group | 

## Methods

### NewExternalGroup

`func NewExternalGroup(groupId int32, groupName string, teams []ExternalGroupTeamsInner, members []ExternalGroupMembersInner, ) *ExternalGroup`

NewExternalGroup instantiates a new ExternalGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalGroupWithDefaults

`func NewExternalGroupWithDefaults() *ExternalGroup`

NewExternalGroupWithDefaults instantiates a new ExternalGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *ExternalGroup) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ExternalGroup) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ExternalGroup) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.


### GetGroupName

`func (o *ExternalGroup) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ExternalGroup) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ExternalGroup) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetUpdatedAt

`func (o *ExternalGroup) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ExternalGroup) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ExternalGroup) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ExternalGroup) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetTeams

`func (o *ExternalGroup) GetTeams() []ExternalGroupTeamsInner`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *ExternalGroup) GetTeamsOk() (*[]ExternalGroupTeamsInner, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *ExternalGroup) SetTeams(v []ExternalGroupTeamsInner)`

SetTeams sets Teams field to given value.


### GetMembers

`func (o *ExternalGroup) GetMembers() []ExternalGroupMembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ExternalGroup) GetMembersOk() (*[]ExternalGroupMembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ExternalGroup) SetMembers(v []ExternalGroupMembersInner)`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


