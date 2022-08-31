# ExternalGroupMembersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MemberId** | **int32** | The internal user ID of the identity | 
**MemberLogin** | **string** | The handle/login for the user | 
**MemberName** | **string** | The user display name/profile name | 
**MemberEmail** | **string** | An email attached to a user | 

## Methods

### NewExternalGroupMembersInner

`func NewExternalGroupMembersInner(memberId int32, memberLogin string, memberName string, memberEmail string, ) *ExternalGroupMembersInner`

NewExternalGroupMembersInner instantiates a new ExternalGroupMembersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalGroupMembersInnerWithDefaults

`func NewExternalGroupMembersInnerWithDefaults() *ExternalGroupMembersInner`

NewExternalGroupMembersInnerWithDefaults instantiates a new ExternalGroupMembersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemberId

`func (o *ExternalGroupMembersInner) GetMemberId() int32`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *ExternalGroupMembersInner) GetMemberIdOk() (*int32, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *ExternalGroupMembersInner) SetMemberId(v int32)`

SetMemberId sets MemberId field to given value.


### GetMemberLogin

`func (o *ExternalGroupMembersInner) GetMemberLogin() string`

GetMemberLogin returns the MemberLogin field if non-nil, zero value otherwise.

### GetMemberLoginOk

`func (o *ExternalGroupMembersInner) GetMemberLoginOk() (*string, bool)`

GetMemberLoginOk returns a tuple with the MemberLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberLogin

`func (o *ExternalGroupMembersInner) SetMemberLogin(v string)`

SetMemberLogin sets MemberLogin field to given value.


### GetMemberName

`func (o *ExternalGroupMembersInner) GetMemberName() string`

GetMemberName returns the MemberName field if non-nil, zero value otherwise.

### GetMemberNameOk

`func (o *ExternalGroupMembersInner) GetMemberNameOk() (*string, bool)`

GetMemberNameOk returns a tuple with the MemberName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberName

`func (o *ExternalGroupMembersInner) SetMemberName(v string)`

SetMemberName sets MemberName field to given value.


### GetMemberEmail

`func (o *ExternalGroupMembersInner) GetMemberEmail() string`

GetMemberEmail returns the MemberEmail field if non-nil, zero value otherwise.

### GetMemberEmailOk

`func (o *ExternalGroupMembersInner) GetMemberEmailOk() (*string, bool)`

GetMemberEmailOk returns a tuple with the MemberEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberEmail

`func (o *ExternalGroupMembersInner) SetMemberEmail(v string)`

SetMemberEmail sets MemberEmail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


