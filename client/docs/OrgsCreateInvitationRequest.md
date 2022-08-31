# OrgsCreateInvitationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InviteeId** | Pointer to **int32** | **Required unless you provide &#x60;email&#x60;**. GitHub user ID for the person you are inviting. | [optional] 
**Email** | Pointer to **string** | **Required unless you provide &#x60;invitee_id&#x60;**. Email address of the person you are inviting, which can be an existing GitHub user. | [optional] 
**Role** | Pointer to **string** | The role for the new member.  \\* &#x60;admin&#x60; - Organization owners with full administrative rights to the organization and complete access to all repositories and teams.   \\* &#x60;direct_member&#x60; - Non-owner organization members with ability to see other members and join teams by invitation.   \\* &#x60;billing_manager&#x60; - Non-owner organization members with ability to manage the billing settings of your organization. | [optional] [default to "direct_member"]
**TeamIds** | Pointer to **[]int32** | Specify IDs for the teams you want to invite new members to. | [optional] 

## Methods

### NewOrgsCreateInvitationRequest

`func NewOrgsCreateInvitationRequest() *OrgsCreateInvitationRequest`

NewOrgsCreateInvitationRequest instantiates a new OrgsCreateInvitationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsCreateInvitationRequestWithDefaults

`func NewOrgsCreateInvitationRequestWithDefaults() *OrgsCreateInvitationRequest`

NewOrgsCreateInvitationRequestWithDefaults instantiates a new OrgsCreateInvitationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInviteeId

`func (o *OrgsCreateInvitationRequest) GetInviteeId() int32`

GetInviteeId returns the InviteeId field if non-nil, zero value otherwise.

### GetInviteeIdOk

`func (o *OrgsCreateInvitationRequest) GetInviteeIdOk() (*int32, bool)`

GetInviteeIdOk returns a tuple with the InviteeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteeId

`func (o *OrgsCreateInvitationRequest) SetInviteeId(v int32)`

SetInviteeId sets InviteeId field to given value.

### HasInviteeId

`func (o *OrgsCreateInvitationRequest) HasInviteeId() bool`

HasInviteeId returns a boolean if a field has been set.

### GetEmail

`func (o *OrgsCreateInvitationRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrgsCreateInvitationRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrgsCreateInvitationRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OrgsCreateInvitationRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetRole

`func (o *OrgsCreateInvitationRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrgsCreateInvitationRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrgsCreateInvitationRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrgsCreateInvitationRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTeamIds

`func (o *OrgsCreateInvitationRequest) GetTeamIds() []int32`

GetTeamIds returns the TeamIds field if non-nil, zero value otherwise.

### GetTeamIdsOk

`func (o *OrgsCreateInvitationRequest) GetTeamIdsOk() (*[]int32, bool)`

GetTeamIdsOk returns a tuple with the TeamIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamIds

`func (o *OrgsCreateInvitationRequest) SetTeamIds(v []int32)`

SetTeamIds sets TeamIds field to given value.

### HasTeamIds

`func (o *OrgsCreateInvitationRequest) HasTeamIds() bool`

HasTeamIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


