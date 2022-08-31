# OrganizationInvitation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Login** | **NullableString** |  | 
**Email** | **NullableString** |  | 
**Role** | **string** |  | 
**CreatedAt** | **string** |  | 
**FailedAt** | Pointer to **NullableString** |  | [optional] 
**FailedReason** | Pointer to **NullableString** |  | [optional] 
**Inviter** | [**SimpleUser**](SimpleUser.md) |  | 
**TeamCount** | **int32** |  | 
**NodeId** | **string** |  | 
**InvitationTeamsUrl** | **string** |  | 

## Methods

### NewOrganizationInvitation

`func NewOrganizationInvitation(id int32, login NullableString, email NullableString, role string, createdAt string, inviter SimpleUser, teamCount int32, nodeId string, invitationTeamsUrl string, ) *OrganizationInvitation`

NewOrganizationInvitation instantiates a new OrganizationInvitation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationInvitationWithDefaults

`func NewOrganizationInvitationWithDefaults() *OrganizationInvitation`

NewOrganizationInvitationWithDefaults instantiates a new OrganizationInvitation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrganizationInvitation) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationInvitation) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationInvitation) SetId(v int32)`

SetId sets Id field to given value.


### GetLogin

`func (o *OrganizationInvitation) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *OrganizationInvitation) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *OrganizationInvitation) SetLogin(v string)`

SetLogin sets Login field to given value.


### SetLoginNil

`func (o *OrganizationInvitation) SetLoginNil(b bool)`

 SetLoginNil sets the value for Login to be an explicit nil

### UnsetLogin
`func (o *OrganizationInvitation) UnsetLogin()`

UnsetLogin ensures that no value is present for Login, not even an explicit nil
### GetEmail

`func (o *OrganizationInvitation) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationInvitation) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationInvitation) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *OrganizationInvitation) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *OrganizationInvitation) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetRole

`func (o *OrganizationInvitation) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrganizationInvitation) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrganizationInvitation) SetRole(v string)`

SetRole sets Role field to given value.


### GetCreatedAt

`func (o *OrganizationInvitation) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationInvitation) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationInvitation) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetFailedAt

`func (o *OrganizationInvitation) GetFailedAt() string`

GetFailedAt returns the FailedAt field if non-nil, zero value otherwise.

### GetFailedAtOk

`func (o *OrganizationInvitation) GetFailedAtOk() (*string, bool)`

GetFailedAtOk returns a tuple with the FailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAt

`func (o *OrganizationInvitation) SetFailedAt(v string)`

SetFailedAt sets FailedAt field to given value.

### HasFailedAt

`func (o *OrganizationInvitation) HasFailedAt() bool`

HasFailedAt returns a boolean if a field has been set.

### SetFailedAtNil

`func (o *OrganizationInvitation) SetFailedAtNil(b bool)`

 SetFailedAtNil sets the value for FailedAt to be an explicit nil

### UnsetFailedAt
`func (o *OrganizationInvitation) UnsetFailedAt()`

UnsetFailedAt ensures that no value is present for FailedAt, not even an explicit nil
### GetFailedReason

`func (o *OrganizationInvitation) GetFailedReason() string`

GetFailedReason returns the FailedReason field if non-nil, zero value otherwise.

### GetFailedReasonOk

`func (o *OrganizationInvitation) GetFailedReasonOk() (*string, bool)`

GetFailedReasonOk returns a tuple with the FailedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedReason

`func (o *OrganizationInvitation) SetFailedReason(v string)`

SetFailedReason sets FailedReason field to given value.

### HasFailedReason

`func (o *OrganizationInvitation) HasFailedReason() bool`

HasFailedReason returns a boolean if a field has been set.

### SetFailedReasonNil

`func (o *OrganizationInvitation) SetFailedReasonNil(b bool)`

 SetFailedReasonNil sets the value for FailedReason to be an explicit nil

### UnsetFailedReason
`func (o *OrganizationInvitation) UnsetFailedReason()`

UnsetFailedReason ensures that no value is present for FailedReason, not even an explicit nil
### GetInviter

`func (o *OrganizationInvitation) GetInviter() SimpleUser`

GetInviter returns the Inviter field if non-nil, zero value otherwise.

### GetInviterOk

`func (o *OrganizationInvitation) GetInviterOk() (*SimpleUser, bool)`

GetInviterOk returns a tuple with the Inviter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviter

`func (o *OrganizationInvitation) SetInviter(v SimpleUser)`

SetInviter sets Inviter field to given value.


### GetTeamCount

`func (o *OrganizationInvitation) GetTeamCount() int32`

GetTeamCount returns the TeamCount field if non-nil, zero value otherwise.

### GetTeamCountOk

`func (o *OrganizationInvitation) GetTeamCountOk() (*int32, bool)`

GetTeamCountOk returns a tuple with the TeamCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamCount

`func (o *OrganizationInvitation) SetTeamCount(v int32)`

SetTeamCount sets TeamCount field to given value.


### GetNodeId

`func (o *OrganizationInvitation) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *OrganizationInvitation) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *OrganizationInvitation) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetInvitationTeamsUrl

`func (o *OrganizationInvitation) GetInvitationTeamsUrl() string`

GetInvitationTeamsUrl returns the InvitationTeamsUrl field if non-nil, zero value otherwise.

### GetInvitationTeamsUrlOk

`func (o *OrganizationInvitation) GetInvitationTeamsUrlOk() (*string, bool)`

GetInvitationTeamsUrlOk returns a tuple with the InvitationTeamsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationTeamsUrl

`func (o *OrganizationInvitation) SetInvitationTeamsUrl(v string)`

SetInvitationTeamsUrl sets InvitationTeamsUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


