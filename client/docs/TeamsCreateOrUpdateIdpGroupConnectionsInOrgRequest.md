# TeamsCreateOrUpdateIdpGroupConnectionsInOrgRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]TeamsCreateOrUpdateIdpGroupConnectionsInOrgRequestGroupsInner**](TeamsCreateOrUpdateIdpGroupConnectionsInOrgRequestGroupsInner.md) | The IdP groups you want to connect to a GitHub team. When updating, the new &#x60;groups&#x60; object will replace the original one. You must include any existing groups that you don&#39;t want to remove. | [optional] 

## Methods

### NewTeamsCreateOrUpdateIdpGroupConnectionsInOrgRequest

`func NewTeamsCreateOrUpdateIdpGroupConnectionsInOrgRequest() *TeamsCreateOrUpdateIdpGroupConnectionsInOrgRequest`

NewTeamsCreateOrUpdateIdpGroupConnectionsInOrgRequest instantiates a new TeamsCreateOrUpdateIdpGroupConnectionsInOrgRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsCreateOrUpdateIdpGroupConnectionsInOrgRequestWithDefaults

`func NewTeamsCreateOrUpdateIdpGroupConnectionsInOrgRequestWithDefaults() *TeamsCreateOrUpdateIdpGroupConnectionsInOrgRequest`

NewTeamsCreateOrUpdateIdpGroupConnectionsInOrgRequestWithDefaults instantiates a new TeamsCreateOrUpdateIdpGroupConnectionsInOrgRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *TeamsCreateOrUpdateIdpGroupConnectionsInOrgRequest) GetGroups() []TeamsCreateOrUpdateIdpGroupConnectionsInOrgRequestGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *TeamsCreateOrUpdateIdpGroupConnectionsInOrgRequest) GetGroupsOk() (*[]TeamsCreateOrUpdateIdpGroupConnectionsInOrgRequestGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *TeamsCreateOrUpdateIdpGroupConnectionsInOrgRequest) SetGroups(v []TeamsCreateOrUpdateIdpGroupConnectionsInOrgRequestGroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *TeamsCreateOrUpdateIdpGroupConnectionsInOrgRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


