# Email

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Primary** | **bool** |  | 
**Verified** | **bool** |  | 
**Visibility** | **NullableString** |  | 

## Methods

### NewEmail

`func NewEmail(email string, primary bool, verified bool, visibility NullableString, ) *Email`

NewEmail instantiates a new Email object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailWithDefaults

`func NewEmailWithDefaults() *Email`

NewEmailWithDefaults instantiates a new Email object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *Email) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Email) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Email) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPrimary

`func (o *Email) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *Email) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *Email) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.


### GetVerified

`func (o *Email) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *Email) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *Email) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetVisibility

`func (o *Email) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Email) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Email) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### SetVisibilityNil

`func (o *Email) SetVisibilityNil(b bool)`

 SetVisibilityNil sets the value for Visibility to be an explicit nil

### UnsetVisibility
`func (o *Email) UnsetVisibility()`

UnsetVisibility ensures that no value is present for Visibility, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


