# Stargazer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StarredAt** | **time.Time** |  | 
**User** | [**NullableNullableSimpleUser**](NullableSimpleUser.md) |  | 

## Methods

### NewStargazer

`func NewStargazer(starredAt time.Time, user NullableNullableSimpleUser, ) *Stargazer`

NewStargazer instantiates a new Stargazer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStargazerWithDefaults

`func NewStargazerWithDefaults() *Stargazer`

NewStargazerWithDefaults instantiates a new Stargazer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStarredAt

`func (o *Stargazer) GetStarredAt() time.Time`

GetStarredAt returns the StarredAt field if non-nil, zero value otherwise.

### GetStarredAtOk

`func (o *Stargazer) GetStarredAtOk() (*time.Time, bool)`

GetStarredAtOk returns a tuple with the StarredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarredAt

`func (o *Stargazer) SetStarredAt(v time.Time)`

SetStarredAt sets StarredAt field to given value.


### GetUser

`func (o *Stargazer) GetUser() NullableSimpleUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Stargazer) GetUserOk() (*NullableSimpleUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Stargazer) SetUser(v NullableSimpleUser)`

SetUser sets User field to given value.


### SetUserNil

`func (o *Stargazer) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *Stargazer) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


