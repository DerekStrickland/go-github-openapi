# OrgsSetMembershipForUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** | The role to give the user in the organization. Can be one of:   \\* &#x60;admin&#x60; - The user will become an owner of the organization.   \\* &#x60;member&#x60; - The user will become a non-owner member of the organization. | [optional] [default to "member"]

## Methods

### NewOrgsSetMembershipForUserRequest

`func NewOrgsSetMembershipForUserRequest() *OrgsSetMembershipForUserRequest`

NewOrgsSetMembershipForUserRequest instantiates a new OrgsSetMembershipForUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsSetMembershipForUserRequestWithDefaults

`func NewOrgsSetMembershipForUserRequestWithDefaults() *OrgsSetMembershipForUserRequest`

NewOrgsSetMembershipForUserRequestWithDefaults instantiates a new OrgsSetMembershipForUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *OrgsSetMembershipForUserRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *OrgsSetMembershipForUserRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *OrgsSetMembershipForUserRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *OrgsSetMembershipForUserRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


