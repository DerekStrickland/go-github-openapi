# GroupMappingGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **string** | The ID of the group | 
**GroupName** | **string** | The name of the group | 
**GroupDescription** | **string** | a description of the group | 
**Status** | Pointer to **string** | synchronization status for this group mapping | [optional] 
**SyncedAt** | Pointer to **NullableString** | the time of the last sync for this group-mapping | [optional] 

## Methods

### NewGroupMappingGroupsInner

`func NewGroupMappingGroupsInner(groupId string, groupName string, groupDescription string, ) *GroupMappingGroupsInner`

NewGroupMappingGroupsInner instantiates a new GroupMappingGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMappingGroupsInnerWithDefaults

`func NewGroupMappingGroupsInnerWithDefaults() *GroupMappingGroupsInner`

NewGroupMappingGroupsInnerWithDefaults instantiates a new GroupMappingGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *GroupMappingGroupsInner) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *GroupMappingGroupsInner) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *GroupMappingGroupsInner) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetGroupName

`func (o *GroupMappingGroupsInner) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *GroupMappingGroupsInner) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *GroupMappingGroupsInner) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetGroupDescription

`func (o *GroupMappingGroupsInner) GetGroupDescription() string`

GetGroupDescription returns the GroupDescription field if non-nil, zero value otherwise.

### GetGroupDescriptionOk

`func (o *GroupMappingGroupsInner) GetGroupDescriptionOk() (*string, bool)`

GetGroupDescriptionOk returns a tuple with the GroupDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDescription

`func (o *GroupMappingGroupsInner) SetGroupDescription(v string)`

SetGroupDescription sets GroupDescription field to given value.


### GetStatus

`func (o *GroupMappingGroupsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GroupMappingGroupsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GroupMappingGroupsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GroupMappingGroupsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSyncedAt

`func (o *GroupMappingGroupsInner) GetSyncedAt() string`

GetSyncedAt returns the SyncedAt field if non-nil, zero value otherwise.

### GetSyncedAtOk

`func (o *GroupMappingGroupsInner) GetSyncedAtOk() (*string, bool)`

GetSyncedAtOk returns a tuple with the SyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedAt

`func (o *GroupMappingGroupsInner) SetSyncedAt(v string)`

SetSyncedAt sets SyncedAt field to given value.

### HasSyncedAt

`func (o *GroupMappingGroupsInner) HasSyncedAt() bool`

HasSyncedAt returns a boolean if a field has been set.

### SetSyncedAtNil

`func (o *GroupMappingGroupsInner) SetSyncedAtNil(b bool)`

 SetSyncedAtNil sets the value for SyncedAt to be an explicit nil

### UnsetSyncedAt
`func (o *GroupMappingGroupsInner) UnsetSyncedAt()`

UnsetSyncedAt ensures that no value is present for SyncedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


