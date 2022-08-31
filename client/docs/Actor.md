# Actor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Login** | **string** |  | 
**DisplayLogin** | Pointer to **string** |  | [optional] 
**GravatarId** | **NullableString** |  | 
**Url** | **string** |  | 
**AvatarUrl** | **string** |  | 

## Methods

### NewActor

`func NewActor(id int32, login string, gravatarId NullableString, url string, avatarUrl string, ) *Actor`

NewActor instantiates a new Actor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActorWithDefaults

`func NewActorWithDefaults() *Actor`

NewActorWithDefaults instantiates a new Actor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Actor) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Actor) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Actor) SetId(v int32)`

SetId sets Id field to given value.


### GetLogin

`func (o *Actor) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *Actor) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *Actor) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetDisplayLogin

`func (o *Actor) GetDisplayLogin() string`

GetDisplayLogin returns the DisplayLogin field if non-nil, zero value otherwise.

### GetDisplayLoginOk

`func (o *Actor) GetDisplayLoginOk() (*string, bool)`

GetDisplayLoginOk returns a tuple with the DisplayLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLogin

`func (o *Actor) SetDisplayLogin(v string)`

SetDisplayLogin sets DisplayLogin field to given value.

### HasDisplayLogin

`func (o *Actor) HasDisplayLogin() bool`

HasDisplayLogin returns a boolean if a field has been set.

### GetGravatarId

`func (o *Actor) GetGravatarId() string`

GetGravatarId returns the GravatarId field if non-nil, zero value otherwise.

### GetGravatarIdOk

`func (o *Actor) GetGravatarIdOk() (*string, bool)`

GetGravatarIdOk returns a tuple with the GravatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGravatarId

`func (o *Actor) SetGravatarId(v string)`

SetGravatarId sets GravatarId field to given value.


### SetGravatarIdNil

`func (o *Actor) SetGravatarIdNil(b bool)`

 SetGravatarIdNil sets the value for GravatarId to be an explicit nil

### UnsetGravatarId
`func (o *Actor) UnsetGravatarId()`

UnsetGravatarId ensures that no value is present for GravatarId, not even an explicit nil
### GetUrl

`func (o *Actor) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Actor) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Actor) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAvatarUrl

`func (o *Actor) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *Actor) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *Actor) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


