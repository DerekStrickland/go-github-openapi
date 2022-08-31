# UsersUpdateAuthenticatedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The new name of the user. | [optional] 
**Email** | Pointer to **string** | The publicly visible email address of the user. | [optional] 
**Blog** | Pointer to **string** | The new blog URL of the user. | [optional] 
**TwitterUsername** | Pointer to **NullableString** | The new Twitter username of the user. | [optional] 
**Company** | Pointer to **string** | The new company of the user. | [optional] 
**Location** | Pointer to **string** | The new location of the user. | [optional] 
**Hireable** | Pointer to **bool** | The new hiring availability of the user. | [optional] 
**Bio** | Pointer to **string** | The new short biography of the user. | [optional] 

## Methods

### NewUsersUpdateAuthenticatedRequest

`func NewUsersUpdateAuthenticatedRequest() *UsersUpdateAuthenticatedRequest`

NewUsersUpdateAuthenticatedRequest instantiates a new UsersUpdateAuthenticatedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsersUpdateAuthenticatedRequestWithDefaults

`func NewUsersUpdateAuthenticatedRequestWithDefaults() *UsersUpdateAuthenticatedRequest`

NewUsersUpdateAuthenticatedRequestWithDefaults instantiates a new UsersUpdateAuthenticatedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UsersUpdateAuthenticatedRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UsersUpdateAuthenticatedRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UsersUpdateAuthenticatedRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UsersUpdateAuthenticatedRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *UsersUpdateAuthenticatedRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UsersUpdateAuthenticatedRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UsersUpdateAuthenticatedRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UsersUpdateAuthenticatedRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetBlog

`func (o *UsersUpdateAuthenticatedRequest) GetBlog() string`

GetBlog returns the Blog field if non-nil, zero value otherwise.

### GetBlogOk

`func (o *UsersUpdateAuthenticatedRequest) GetBlogOk() (*string, bool)`

GetBlogOk returns a tuple with the Blog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlog

`func (o *UsersUpdateAuthenticatedRequest) SetBlog(v string)`

SetBlog sets Blog field to given value.

### HasBlog

`func (o *UsersUpdateAuthenticatedRequest) HasBlog() bool`

HasBlog returns a boolean if a field has been set.

### GetTwitterUsername

`func (o *UsersUpdateAuthenticatedRequest) GetTwitterUsername() string`

GetTwitterUsername returns the TwitterUsername field if non-nil, zero value otherwise.

### GetTwitterUsernameOk

`func (o *UsersUpdateAuthenticatedRequest) GetTwitterUsernameOk() (*string, bool)`

GetTwitterUsernameOk returns a tuple with the TwitterUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterUsername

`func (o *UsersUpdateAuthenticatedRequest) SetTwitterUsername(v string)`

SetTwitterUsername sets TwitterUsername field to given value.

### HasTwitterUsername

`func (o *UsersUpdateAuthenticatedRequest) HasTwitterUsername() bool`

HasTwitterUsername returns a boolean if a field has been set.

### SetTwitterUsernameNil

`func (o *UsersUpdateAuthenticatedRequest) SetTwitterUsernameNil(b bool)`

 SetTwitterUsernameNil sets the value for TwitterUsername to be an explicit nil

### UnsetTwitterUsername
`func (o *UsersUpdateAuthenticatedRequest) UnsetTwitterUsername()`

UnsetTwitterUsername ensures that no value is present for TwitterUsername, not even an explicit nil
### GetCompany

`func (o *UsersUpdateAuthenticatedRequest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UsersUpdateAuthenticatedRequest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UsersUpdateAuthenticatedRequest) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *UsersUpdateAuthenticatedRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetLocation

`func (o *UsersUpdateAuthenticatedRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UsersUpdateAuthenticatedRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UsersUpdateAuthenticatedRequest) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *UsersUpdateAuthenticatedRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetHireable

`func (o *UsersUpdateAuthenticatedRequest) GetHireable() bool`

GetHireable returns the Hireable field if non-nil, zero value otherwise.

### GetHireableOk

`func (o *UsersUpdateAuthenticatedRequest) GetHireableOk() (*bool, bool)`

GetHireableOk returns a tuple with the Hireable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHireable

`func (o *UsersUpdateAuthenticatedRequest) SetHireable(v bool)`

SetHireable sets Hireable field to given value.

### HasHireable

`func (o *UsersUpdateAuthenticatedRequest) HasHireable() bool`

HasHireable returns a boolean if a field has been set.

### GetBio

`func (o *UsersUpdateAuthenticatedRequest) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *UsersUpdateAuthenticatedRequest) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *UsersUpdateAuthenticatedRequest) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *UsersUpdateAuthenticatedRequest) HasBio() bool`

HasBio returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


