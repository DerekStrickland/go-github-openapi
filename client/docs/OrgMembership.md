# OrgMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**State** | **string** | The state of the member in the organization. The &#x60;pending&#x60; state indicates the user has not yet accepted an invitation. | 
**Role** | **string** | The user&#39;s membership type in the organization. | 
**OrganizationUrl** | **string** |  | 
**Organization** | [**OrganizationSimple**](OrganizationSimple.md) |  | 
**User** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 
**Permissions** | Pointer to [**OrgMembershipPermissions**](OrgMembershipPermissions.md) |  | [optional] 

## Methods

### NewOrgMembership

`func NewOrgMembership(url string, state string, role string, organizationUrl string, organization OrganizationSimple, user NullableNullableSimpleUser, ) *OrgMembership`

NewOrgMembership instantiates a new OrgMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgMembershipWithDefaults

`func NewOrgMembershipWithDefaults() *OrgMembership`

NewOrgMembershipWithDefaults instantiates a new OrgMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *OrgMembership) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OrgMembership) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OrgMembership) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetState

`func (o *OrgMembership) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OrgMembership) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OrgMembership) SetState(v string)`

SetState sets State field to given value.


### GetRole

`func (o *OrgMembership) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrgMembership) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrgMembership) SetRole(v string)`

SetRole sets Role field to given value.


### GetOrganizationUrl

`func (o *OrgMembership) GetOrganizationUrl() string`

GetOrganizationUrl returns the OrganizationUrl field if non-nil, zero value otherwise.

### GetOrganizationUrlOk

`func (o *OrgMembership) GetOrganizationUrlOk() (*string, bool)`

GetOrganizationUrlOk returns a tuple with the OrganizationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUrl

`func (o *OrgMembership) SetOrganizationUrl(v string)`

SetOrganizationUrl sets OrganizationUrl field to given value.


### GetOrganization

`func (o *OrgMembership) GetOrganization() OrganizationSimple`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *OrgMembership) GetOrganizationOk() (*OrganizationSimple, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *OrgMembership) SetOrganization(v OrganizationSimple)`

SetOrganization sets Organization field to given value.


### GetUser

`func (o *OrgMembership) GetUser() NullableSimpleUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OrgMembership) GetUserOk() (*NullableSimpleUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OrgMembership) SetUser(v NullableSimpleUser)`

SetUser sets User field to given value.


### SetUserNil

`func (o *OrgMembership) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *OrgMembership) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil
### GetPermissions

`func (o *OrgMembership) GetPermissions() OrgMembershipPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *OrgMembership) GetPermissionsOk() (*OrgMembershipPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *OrgMembership) SetPermissions(v OrgMembershipPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *OrgMembership) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


