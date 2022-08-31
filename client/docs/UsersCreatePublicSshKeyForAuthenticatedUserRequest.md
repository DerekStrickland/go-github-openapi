# UsersCreatePublicSshKeyForAuthenticatedUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | A descriptive name for the new key. | [optional] 
**Key** | **string** | The public SSH key to add to your GitHub account. | 

## Methods

### NewUsersCreatePublicSshKeyForAuthenticatedUserRequest

`func NewUsersCreatePublicSshKeyForAuthenticatedUserRequest(key string, ) *UsersCreatePublicSshKeyForAuthenticatedUserRequest`

NewUsersCreatePublicSshKeyForAuthenticatedUserRequest instantiates a new UsersCreatePublicSshKeyForAuthenticatedUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersCreatePublicSshKeyForAuthenticatedUserRequestWithDefaults

`func NewUsersCreatePublicSshKeyForAuthenticatedUserRequestWithDefaults() *UsersCreatePublicSshKeyForAuthenticatedUserRequest`

NewUsersCreatePublicSshKeyForAuthenticatedUserRequestWithDefaults instantiates a new UsersCreatePublicSshKeyForAuthenticatedUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UsersCreatePublicSshKeyForAuthenticatedUserRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UsersCreatePublicSshKeyForAuthenticatedUserRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UsersCreatePublicSshKeyForAuthenticatedUserRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UsersCreatePublicSshKeyForAuthenticatedUserRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetKey

`func (o *UsersCreatePublicSshKeyForAuthenticatedUserRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UsersCreatePublicSshKeyForAuthenticatedUserRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UsersCreatePublicSshKeyForAuthenticatedUserRequest) SetKey(v string)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


