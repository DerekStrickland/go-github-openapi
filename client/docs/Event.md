# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **NullableString** |  | 
**Actor** | [**Actor**](Actor.md) |  | 
**Repo** | [**EventRepo**](EventRepo.md) |  | 
**Org** | Pointer to [**Actor**](Actor.md) |  | [optional] 
**Payload** | [**EventPayload**](EventPayload.md) |  | 
**Public** | **bool** |  | 
**CreatedAt** | **NullableTime** |  | 

## Methods

### NewEvent

`func NewEvent(id string, type_ NullableString, actor Actor, repo EventRepo, payload EventPayload, public bool, createdAt NullableTime, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Event) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Event) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Event) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *Event) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Event) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Event) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *Event) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Event) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetActor

`func (o *Event) GetActor() Actor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *Event) GetActorOk() (*Actor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *Event) SetActor(v Actor)`

SetActor sets Actor field to given value.


### GetRepo

`func (o *Event) GetRepo() EventRepo`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *Event) GetRepoOk() (*EventRepo, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *Event) SetRepo(v EventRepo)`

SetRepo sets Repo field to given value.


### GetOrg

`func (o *Event) GetOrg() Actor`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *Event) GetOrgOk() (*Actor, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *Event) SetOrg(v Actor)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *Event) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPayload

`func (o *Event) GetPayload() EventPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Event) GetPayloadOk() (*EventPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Event) SetPayload(v EventPayload)`

SetPayload sets Payload field to given value.


### GetPublic

`func (o *Event) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *Event) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *Event) SetPublic(v bool)`

SetPublic sets Public field to given value.


### GetCreatedAt

`func (o *Event) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Event) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Event) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *Event) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Event) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


