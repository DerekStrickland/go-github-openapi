# UsersAddEmailForAuthenticatedUserRequestOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | **[]string** | Adds one or more email addresses to your GitHub account. Must contain at least one email address. **Note:** Alternatively, you can pass a single email address or an &#x60;array&#x60; of emails addresses directly, but we recommend that you pass an object using the &#x60;emails&#x60; key. | 

## Methods

### NewUsersAddEmailForAuthenticatedUserRequestOneOf

`func NewUsersAddEmailForAuthenticatedUserRequestOneOf(emails []string, ) *UsersAddEmailForAuthenticatedUserRequestOneOf`

NewUsersAddEmailForAuthenticatedUserRequestOneOf instantiates a new UsersAddEmailForAuthenticatedUserRequestOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersAddEmailForAuthenticatedUserRequestOneOfWithDefaults

`func NewUsersAddEmailForAuthenticatedUserRequestOneOfWithDefaults() *UsersAddEmailForAuthenticatedUserRequestOneOf`

NewUsersAddEmailForAuthenticatedUserRequestOneOfWithDefaults instantiates a new UsersAddEmailForAuthenticatedUserRequestOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *UsersAddEmailForAuthenticatedUserRequestOneOf) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *UsersAddEmailForAuthenticatedUserRequestOneOf) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *UsersAddEmailForAuthenticatedUserRequestOneOf) SetEmails(v []string)`

SetEmails sets Emails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


