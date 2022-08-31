# TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | [**[]TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner**](TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner.md) | The IdP groups you want to connect to a GitHub team. When updating, the new &#x60;groups&#x60; object will replace the original one. You must include any existing groups that you don&#39;t want to remove. | 
**SyncedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewTeamsCreateOrUpdateIdpGroupConnectionsLegacyRequest

`func NewTeamsCreateOrUpdateIdpGroupConnectionsLegacyRequest(groups []TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner, ) *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequest`

NewTeamsCreateOrUpdateIdpGroupConnectionsLegacyRequest instantiates a new TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestWithDefaults

`func NewTeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestWithDefaults() *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequest`

NewTeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestWithDefaults instantiates a new TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequest) GetGroups() []TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequest) GetGroupsOk() (*[]TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequest) SetGroups(v []TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequestGroupsInner)`

SetGroups sets Groups field to given value.


### GetSyncedAt

`func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequest) GetSyncedAt() string`

GetSyncedAt returns the SyncedAt field if non-nil, zero value otherwise.

### GetSyncedAtOk

`func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequest) GetSyncedAtOk() (*string, bool)`

GetSyncedAtOk returns a tuple with the SyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedAt

`func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequest) SetSyncedAt(v string)`

SetSyncedAt sets SyncedAt field to given value.

### HasSyncedAt

`func (o *TeamsCreateOrUpdateIdpGroupConnectionsLegacyRequest) HasSyncedAt() bool`

HasSyncedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


