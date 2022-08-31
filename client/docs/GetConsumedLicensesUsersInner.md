# GetConsumedLicensesUsersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GithubComLogin** | Pointer to **string** |  | [optional] 
**GithubComName** | Pointer to **NullableString** |  | [optional] 
**GithubComProfile** | Pointer to **NullableString** |  | [optional] 
**LicenseType** | Pointer to **string** |  | [optional] 
**GithubComMemberRoles** | Pointer to **[]string** |  | [optional] 
**GithubComEnterpriseRole** | Pointer to **NullableString** |  | [optional] 
**VisualStudioSubscriptionUser** | Pointer to **bool** |  | [optional] 
**GithubComVerifiedDomainEmails** | Pointer to **[]string** |  | [optional] 
**GithubComSamlNameId** | Pointer to **NullableString** |  | [optional] 
**EnterpriseServerUser** | Pointer to **NullableBool** |  | [optional] 
**EnterpriseServerEmails** | Pointer to **[]string** |  | [optional] 
**GithubComUser** | Pointer to **bool** |  | [optional] 
**TotalUserAccounts** | Pointer to **int32** |  | [optional] 
**EnterpriseServerUserIds** | Pointer to **[]string** |  | [optional] 
**GithubComOrgsWithPendingInvites** | Pointer to **[]string** |  | [optional] 
**VisualStudioSubscriptionEmail** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetConsumedLicensesUsersInner

`func NewGetConsumedLicensesUsersInner() *GetConsumedLicensesUsersInner`

NewGetConsumedLicensesUsersInner instantiates a new GetConsumedLicensesUsersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConsumedLicensesUsersInnerWithDefaults

`func NewGetConsumedLicensesUsersInnerWithDefaults() *GetConsumedLicensesUsersInner`

NewGetConsumedLicensesUsersInnerWithDefaults instantiates a new GetConsumedLicensesUsersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGithubComLogin

`func (o *GetConsumedLicensesUsersInner) GetGithubComLogin() string`

GetGithubComLogin returns the GithubComLogin field if non-nil, zero value otherwise.

### GetGithubComLoginOk

`func (o *GetConsumedLicensesUsersInner) GetGithubComLoginOk() (*string, bool)`

GetGithubComLoginOk returns a tuple with the GithubComLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubComLogin

`func (o *GetConsumedLicensesUsersInner) SetGithubComLogin(v string)`

SetGithubComLogin sets GithubComLogin field to given value.

### HasGithubComLogin

`func (o *GetConsumedLicensesUsersInner) HasGithubComLogin() bool`

HasGithubComLogin returns a boolean if a field has been set.

### GetGithubComName

`func (o *GetConsumedLicensesUsersInner) GetGithubComName() string`

GetGithubComName returns the GithubComName field if non-nil, zero value otherwise.

### GetGithubComNameOk

`func (o *GetConsumedLicensesUsersInner) GetGithubComNameOk() (*string, bool)`

GetGithubComNameOk returns a tuple with the GithubComName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubComName

`func (o *GetConsumedLicensesUsersInner) SetGithubComName(v string)`

SetGithubComName sets GithubComName field to given value.

### HasGithubComName

`func (o *GetConsumedLicensesUsersInner) HasGithubComName() bool`

HasGithubComName returns a boolean if a field has been set.

### SetGithubComNameNil

`func (o *GetConsumedLicensesUsersInner) SetGithubComNameNil(b bool)`

 SetGithubComNameNil sets the value for GithubComName to be an explicit nil

### UnsetGithubComName
`func (o *GetConsumedLicensesUsersInner) UnsetGithubComName()`

UnsetGithubComName ensures that no value is present for GithubComName, not even an explicit nil
### GetGithubComProfile

`func (o *GetConsumedLicensesUsersInner) GetGithubComProfile() string`

GetGithubComProfile returns the GithubComProfile field if non-nil, zero value otherwise.

### GetGithubComProfileOk

`func (o *GetConsumedLicensesUsersInner) GetGithubComProfileOk() (*string, bool)`

GetGithubComProfileOk returns a tuple with the GithubComProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubComProfile

`func (o *GetConsumedLicensesUsersInner) SetGithubComProfile(v string)`

SetGithubComProfile sets GithubComProfile field to given value.

### HasGithubComProfile

`func (o *GetConsumedLicensesUsersInner) HasGithubComProfile() bool`

HasGithubComProfile returns a boolean if a field has been set.

### SetGithubComProfileNil

`func (o *GetConsumedLicensesUsersInner) SetGithubComProfileNil(b bool)`

 SetGithubComProfileNil sets the value for GithubComProfile to be an explicit nil

### UnsetGithubComProfile
`func (o *GetConsumedLicensesUsersInner) UnsetGithubComProfile()`

UnsetGithubComProfile ensures that no value is present for GithubComProfile, not even an explicit nil
### GetLicenseType

`func (o *GetConsumedLicensesUsersInner) GetLicenseType() string`

GetLicenseType returns the LicenseType field if non-nil, zero value otherwise.

### GetLicenseTypeOk

`func (o *GetConsumedLicensesUsersInner) GetLicenseTypeOk() (*string, bool)`

GetLicenseTypeOk returns a tuple with the LicenseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseType

`func (o *GetConsumedLicensesUsersInner) SetLicenseType(v string)`

SetLicenseType sets LicenseType field to given value.

### HasLicenseType

`func (o *GetConsumedLicensesUsersInner) HasLicenseType() bool`

HasLicenseType returns a boolean if a field has been set.

### GetGithubComMemberRoles

`func (o *GetConsumedLicensesUsersInner) GetGithubComMemberRoles() []string`

GetGithubComMemberRoles returns the GithubComMemberRoles field if non-nil, zero value otherwise.

### GetGithubComMemberRolesOk

`func (o *GetConsumedLicensesUsersInner) GetGithubComMemberRolesOk() (*[]string, bool)`

GetGithubComMemberRolesOk returns a tuple with the GithubComMemberRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubComMemberRoles

`func (o *GetConsumedLicensesUsersInner) SetGithubComMemberRoles(v []string)`

SetGithubComMemberRoles sets GithubComMemberRoles field to given value.

### HasGithubComMemberRoles

`func (o *GetConsumedLicensesUsersInner) HasGithubComMemberRoles() bool`

HasGithubComMemberRoles returns a boolean if a field has been set.

### GetGithubComEnterpriseRole

`func (o *GetConsumedLicensesUsersInner) GetGithubComEnterpriseRole() string`

GetGithubComEnterpriseRole returns the GithubComEnterpriseRole field if non-nil, zero value otherwise.

### GetGithubComEnterpriseRoleOk

`func (o *GetConsumedLicensesUsersInner) GetGithubComEnterpriseRoleOk() (*string, bool)`

GetGithubComEnterpriseRoleOk returns a tuple with the GithubComEnterpriseRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubComEnterpriseRole

`func (o *GetConsumedLicensesUsersInner) SetGithubComEnterpriseRole(v string)`

SetGithubComEnterpriseRole sets GithubComEnterpriseRole field to given value.

### HasGithubComEnterpriseRole

`func (o *GetConsumedLicensesUsersInner) HasGithubComEnterpriseRole() bool`

HasGithubComEnterpriseRole returns a boolean if a field has been set.

### SetGithubComEnterpriseRoleNil

`func (o *GetConsumedLicensesUsersInner) SetGithubComEnterpriseRoleNil(b bool)`

 SetGithubComEnterpriseRoleNil sets the value for GithubComEnterpriseRole to be an explicit nil

### UnsetGithubComEnterpriseRole
`func (o *GetConsumedLicensesUsersInner) UnsetGithubComEnterpriseRole()`

UnsetGithubComEnterpriseRole ensures that no value is present for GithubComEnterpriseRole, not even an explicit nil
### GetVisualStudioSubscriptionUser

`func (o *GetConsumedLicensesUsersInner) GetVisualStudioSubscriptionUser() bool`

GetVisualStudioSubscriptionUser returns the VisualStudioSubscriptionUser field if non-nil, zero value otherwise.

### GetVisualStudioSubscriptionUserOk

`func (o *GetConsumedLicensesUsersInner) GetVisualStudioSubscriptionUserOk() (*bool, bool)`

GetVisualStudioSubscriptionUserOk returns a tuple with the VisualStudioSubscriptionUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualStudioSubscriptionUser

`func (o *GetConsumedLicensesUsersInner) SetVisualStudioSubscriptionUser(v bool)`

SetVisualStudioSubscriptionUser sets VisualStudioSubscriptionUser field to given value.

### HasVisualStudioSubscriptionUser

`func (o *GetConsumedLicensesUsersInner) HasVisualStudioSubscriptionUser() bool`

HasVisualStudioSubscriptionUser returns a boolean if a field has been set.

### GetGithubComVerifiedDomainEmails

`func (o *GetConsumedLicensesUsersInner) GetGithubComVerifiedDomainEmails() []string`

GetGithubComVerifiedDomainEmails returns the GithubComVerifiedDomainEmails field if non-nil, zero value otherwise.

### GetGithubComVerifiedDomainEmailsOk

`func (o *GetConsumedLicensesUsersInner) GetGithubComVerifiedDomainEmailsOk() (*[]string, bool)`

GetGithubComVerifiedDomainEmailsOk returns a tuple with the GithubComVerifiedDomainEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubComVerifiedDomainEmails

`func (o *GetConsumedLicensesUsersInner) SetGithubComVerifiedDomainEmails(v []string)`

SetGithubComVerifiedDomainEmails sets GithubComVerifiedDomainEmails field to given value.

### HasGithubComVerifiedDomainEmails

`func (o *GetConsumedLicensesUsersInner) HasGithubComVerifiedDomainEmails() bool`

HasGithubComVerifiedDomainEmails returns a boolean if a field has been set.

### GetGithubComSamlNameId

`func (o *GetConsumedLicensesUsersInner) GetGithubComSamlNameId() string`

GetGithubComSamlNameId returns the GithubComSamlNameId field if non-nil, zero value otherwise.

### GetGithubComSamlNameIdOk

`func (o *GetConsumedLicensesUsersInner) GetGithubComSamlNameIdOk() (*string, bool)`

GetGithubComSamlNameIdOk returns a tuple with the GithubComSamlNameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubComSamlNameId

`func (o *GetConsumedLicensesUsersInner) SetGithubComSamlNameId(v string)`

SetGithubComSamlNameId sets GithubComSamlNameId field to given value.

### HasGithubComSamlNameId

`func (o *GetConsumedLicensesUsersInner) HasGithubComSamlNameId() bool`

HasGithubComSamlNameId returns a boolean if a field has been set.

### SetGithubComSamlNameIdNil

`func (o *GetConsumedLicensesUsersInner) SetGithubComSamlNameIdNil(b bool)`

 SetGithubComSamlNameIdNil sets the value for GithubComSamlNameId to be an explicit nil

### UnsetGithubComSamlNameId
`func (o *GetConsumedLicensesUsersInner) UnsetGithubComSamlNameId()`

UnsetGithubComSamlNameId ensures that no value is present for GithubComSamlNameId, not even an explicit nil
### GetEnterpriseServerUser

`func (o *GetConsumedLicensesUsersInner) GetEnterpriseServerUser() bool`

GetEnterpriseServerUser returns the EnterpriseServerUser field if non-nil, zero value otherwise.

### GetEnterpriseServerUserOk

`func (o *GetConsumedLicensesUsersInner) GetEnterpriseServerUserOk() (*bool, bool)`

GetEnterpriseServerUserOk returns a tuple with the EnterpriseServerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseServerUser

`func (o *GetConsumedLicensesUsersInner) SetEnterpriseServerUser(v bool)`

SetEnterpriseServerUser sets EnterpriseServerUser field to given value.

### HasEnterpriseServerUser

`func (o *GetConsumedLicensesUsersInner) HasEnterpriseServerUser() bool`

HasEnterpriseServerUser returns a boolean if a field has been set.

### SetEnterpriseServerUserNil

`func (o *GetConsumedLicensesUsersInner) SetEnterpriseServerUserNil(b bool)`

 SetEnterpriseServerUserNil sets the value for EnterpriseServerUser to be an explicit nil

### UnsetEnterpriseServerUser
`func (o *GetConsumedLicensesUsersInner) UnsetEnterpriseServerUser()`

UnsetEnterpriseServerUser ensures that no value is present for EnterpriseServerUser, not even an explicit nil
### GetEnterpriseServerEmails

`func (o *GetConsumedLicensesUsersInner) GetEnterpriseServerEmails() []string`

GetEnterpriseServerEmails returns the EnterpriseServerEmails field if non-nil, zero value otherwise.

### GetEnterpriseServerEmailsOk

`func (o *GetConsumedLicensesUsersInner) GetEnterpriseServerEmailsOk() (*[]string, bool)`

GetEnterpriseServerEmailsOk returns a tuple with the EnterpriseServerEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseServerEmails

`func (o *GetConsumedLicensesUsersInner) SetEnterpriseServerEmails(v []string)`

SetEnterpriseServerEmails sets EnterpriseServerEmails field to given value.

### HasEnterpriseServerEmails

`func (o *GetConsumedLicensesUsersInner) HasEnterpriseServerEmails() bool`

HasEnterpriseServerEmails returns a boolean if a field has been set.

### GetGithubComUser

`func (o *GetConsumedLicensesUsersInner) GetGithubComUser() bool`

GetGithubComUser returns the GithubComUser field if non-nil, zero value otherwise.

### GetGithubComUserOk

`func (o *GetConsumedLicensesUsersInner) GetGithubComUserOk() (*bool, bool)`

GetGithubComUserOk returns a tuple with the GithubComUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubComUser

`func (o *GetConsumedLicensesUsersInner) SetGithubComUser(v bool)`

SetGithubComUser sets GithubComUser field to given value.

### HasGithubComUser

`func (o *GetConsumedLicensesUsersInner) HasGithubComUser() bool`

HasGithubComUser returns a boolean if a field has been set.

### GetTotalUserAccounts

`func (o *GetConsumedLicensesUsersInner) GetTotalUserAccounts() int32`

GetTotalUserAccounts returns the TotalUserAccounts field if non-nil, zero value otherwise.

### GetTotalUserAccountsOk

`func (o *GetConsumedLicensesUsersInner) GetTotalUserAccountsOk() (*int32, bool)`

GetTotalUserAccountsOk returns a tuple with the TotalUserAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalUserAccounts

`func (o *GetConsumedLicensesUsersInner) SetTotalUserAccounts(v int32)`

SetTotalUserAccounts sets TotalUserAccounts field to given value.

### HasTotalUserAccounts

`func (o *GetConsumedLicensesUsersInner) HasTotalUserAccounts() bool`

HasTotalUserAccounts returns a boolean if a field has been set.

### GetEnterpriseServerUserIds

`func (o *GetConsumedLicensesUsersInner) GetEnterpriseServerUserIds() []string`

GetEnterpriseServerUserIds returns the EnterpriseServerUserIds field if non-nil, zero value otherwise.

### GetEnterpriseServerUserIdsOk

`func (o *GetConsumedLicensesUsersInner) GetEnterpriseServerUserIdsOk() (*[]string, bool)`

GetEnterpriseServerUserIdsOk returns a tuple with the EnterpriseServerUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnterpriseServerUserIds

`func (o *GetConsumedLicensesUsersInner) SetEnterpriseServerUserIds(v []string)`

SetEnterpriseServerUserIds sets EnterpriseServerUserIds field to given value.

### HasEnterpriseServerUserIds

`func (o *GetConsumedLicensesUsersInner) HasEnterpriseServerUserIds() bool`

HasEnterpriseServerUserIds returns a boolean if a field has been set.

### GetGithubComOrgsWithPendingInvites

`func (o *GetConsumedLicensesUsersInner) GetGithubComOrgsWithPendingInvites() []string`

GetGithubComOrgsWithPendingInvites returns the GithubComOrgsWithPendingInvites field if non-nil, zero value otherwise.

### GetGithubComOrgsWithPendingInvitesOk

`func (o *GetConsumedLicensesUsersInner) GetGithubComOrgsWithPendingInvitesOk() (*[]string, bool)`

GetGithubComOrgsWithPendingInvitesOk returns a tuple with the GithubComOrgsWithPendingInvites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubComOrgsWithPendingInvites

`func (o *GetConsumedLicensesUsersInner) SetGithubComOrgsWithPendingInvites(v []string)`

SetGithubComOrgsWithPendingInvites sets GithubComOrgsWithPendingInvites field to given value.

### HasGithubComOrgsWithPendingInvites

`func (o *GetConsumedLicensesUsersInner) HasGithubComOrgsWithPendingInvites() bool`

HasGithubComOrgsWithPendingInvites returns a boolean if a field has been set.

### GetVisualStudioSubscriptionEmail

`func (o *GetConsumedLicensesUsersInner) GetVisualStudioSubscriptionEmail() string`

GetVisualStudioSubscriptionEmail returns the VisualStudioSubscriptionEmail field if non-nil, zero value otherwise.

### GetVisualStudioSubscriptionEmailOk

`func (o *GetConsumedLicensesUsersInner) GetVisualStudioSubscriptionEmailOk() (*string, bool)`

GetVisualStudioSubscriptionEmailOk returns a tuple with the VisualStudioSubscriptionEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualStudioSubscriptionEmail

`func (o *GetConsumedLicensesUsersInner) SetVisualStudioSubscriptionEmail(v string)`

SetVisualStudioSubscriptionEmail sets VisualStudioSubscriptionEmail field to given value.

### HasVisualStudioSubscriptionEmail

`func (o *GetConsumedLicensesUsersInner) HasVisualStudioSubscriptionEmail() bool`

HasVisualStudioSubscriptionEmail returns a boolean if a field has been set.

### SetVisualStudioSubscriptionEmailNil

`func (o *GetConsumedLicensesUsersInner) SetVisualStudioSubscriptionEmailNil(b bool)`

 SetVisualStudioSubscriptionEmailNil sets the value for VisualStudioSubscriptionEmail to be an explicit nil

### UnsetVisualStudioSubscriptionEmail
`func (o *GetConsumedLicensesUsersInner) UnsetVisualStudioSubscriptionEmail()`

UnsetVisualStudioSubscriptionEmail ensures that no value is present for VisualStudioSubscriptionEmail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


