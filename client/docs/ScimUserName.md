# ScimUserName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GivenName** | **NullableString** |  | 
**FamilyName** | **NullableString** |  | 
**Formatted** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewScimUserName

`func NewScimUserName(givenName NullableString, familyName NullableString, ) *ScimUserName`

NewScimUserName instantiates a new ScimUserName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimUserNameWithDefaults

`func NewScimUserNameWithDefaults() *ScimUserName`

NewScimUserNameWithDefaults instantiates a new ScimUserName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGivenName

`func (o *ScimUserName) GetGivenName() string`

GetGivenName returns the GivenName field if non-nil, zero value otherwise.

### GetGivenNameOk

`func (o *ScimUserName) GetGivenNameOk() (*string, bool)`

GetGivenNameOk returns a tuple with the GivenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenName

`func (o *ScimUserName) SetGivenName(v string)`

SetGivenName sets GivenName field to given value.


### SetGivenNameNil

`func (o *ScimUserName) SetGivenNameNil(b bool)`

 SetGivenNameNil sets the value for GivenName to be an explicit nil

### UnsetGivenName
`func (o *ScimUserName) UnsetGivenName()`

UnsetGivenName ensures that no value is present for GivenName, not even an explicit nil
### GetFamilyName

`func (o *ScimUserName) GetFamilyName() string`

GetFamilyName returns the FamilyName field if non-nil, zero value otherwise.

### GetFamilyNameOk

`func (o *ScimUserName) GetFamilyNameOk() (*string, bool)`

GetFamilyNameOk returns a tuple with the FamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamilyName

`func (o *ScimUserName) SetFamilyName(v string)`

SetFamilyName sets FamilyName field to given value.


### SetFamilyNameNil

`func (o *ScimUserName) SetFamilyNameNil(b bool)`

 SetFamilyNameNil sets the value for FamilyName to be an explicit nil

### UnsetFamilyName
`func (o *ScimUserName) UnsetFamilyName()`

UnsetFamilyName ensures that no value is present for FamilyName, not even an explicit nil
### GetFormatted

`func (o *ScimUserName) GetFormatted() string`

GetFormatted returns the Formatted field if non-nil, zero value otherwise.

### GetFormattedOk

`func (o *ScimUserName) GetFormattedOk() (*string, bool)`

GetFormattedOk returns a tuple with the Formatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatted

`func (o *ScimUserName) SetFormatted(v string)`

SetFormatted sets Formatted field to given value.

### HasFormatted

`func (o *ScimUserName) HasFormatted() bool`

HasFormatted returns a boolean if a field has been set.

### SetFormattedNil

`func (o *ScimUserName) SetFormattedNil(b bool)`

 SetFormattedNil sets the value for Formatted to be an explicit nil

### UnsetFormatted
`func (o *ScimUserName) UnsetFormatted()`

UnsetFormatted ensures that no value is present for Formatted, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


