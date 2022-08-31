# RepositoryPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | **bool** |  | 
**Pull** | **bool** |  | 
**Triage** | Pointer to **bool** |  | [optional] 
**Push** | **bool** |  | 
**Maintain** | Pointer to **bool** |  | [optional] 

## Methods

### NewRepositoryPermissions

`func NewRepositoryPermissions(admin bool, pull bool, push bool, ) *RepositoryPermissions`

NewRepositoryPermissions instantiates a new RepositoryPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRepositoryPermissionsWithDefaults

`func NewRepositoryPermissionsWithDefaults() *RepositoryPermissions`

NewRepositoryPermissionsWithDefaults instantiates a new RepositoryPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *RepositoryPermissions) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *RepositoryPermissions) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *RepositoryPermissions) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.


### GetPull

`func (o *RepositoryPermissions) GetPull() bool`

GetPull returns the Pull field if non-nil, zero value otherwise.

### GetPullOk

`func (o *RepositoryPermissions) GetPullOk() (*bool, bool)`

GetPullOk returns a tuple with the Pull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPull

`func (o *RepositoryPermissions) SetPull(v bool)`

SetPull sets Pull field to given value.


### GetTriage

`func (o *RepositoryPermissions) GetTriage() bool`

GetTriage returns the Triage field if non-nil, zero value otherwise.

### GetTriageOk

`func (o *RepositoryPermissions) GetTriageOk() (*bool, bool)`

GetTriageOk returns a tuple with the Triage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriage

`func (o *RepositoryPermissions) SetTriage(v bool)`

SetTriage sets Triage field to given value.

### HasTriage

`func (o *RepositoryPermissions) HasTriage() bool`

HasTriage returns a boolean if a field has been set.

### GetPush

`func (o *RepositoryPermissions) GetPush() bool`

GetPush returns the Push field if non-nil, zero value otherwise.

### GetPushOk

`func (o *RepositoryPermissions) GetPushOk() (*bool, bool)`

GetPushOk returns a tuple with the Push field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPush

`func (o *RepositoryPermissions) SetPush(v bool)`

SetPush sets Push field to given value.


### GetMaintain

`func (o *RepositoryPermissions) GetMaintain() bool`

GetMaintain returns the Maintain field if non-nil, zero value otherwise.

### GetMaintainOk

`func (o *RepositoryPermissions) GetMaintainOk() (*bool, bool)`

GetMaintainOk returns a tuple with the Maintain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintain

`func (o *RepositoryPermissions) SetMaintain(v bool)`

SetMaintain sets Maintain field to given value.

### HasMaintain

`func (o *RepositoryPermissions) HasMaintain() bool`

HasMaintain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


