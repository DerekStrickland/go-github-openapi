# FullRepositoryPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | **bool** |  | 
**Maintain** | Pointer to **bool** |  | [optional] 
**Push** | **bool** |  | 
**Triage** | Pointer to **bool** |  | [optional] 
**Pull** | **bool** |  | 

## Methods

### NewFullRepositoryPermissions

`func NewFullRepositoryPermissions(admin bool, push bool, pull bool, ) *FullRepositoryPermissions`

NewFullRepositoryPermissions instantiates a new FullRepositoryPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullRepositoryPermissionsWithDefaults

`func NewFullRepositoryPermissionsWithDefaults() *FullRepositoryPermissions`

NewFullRepositoryPermissionsWithDefaults instantiates a new FullRepositoryPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *FullRepositoryPermissions) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *FullRepositoryPermissions) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *FullRepositoryPermissions) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.


### GetMaintain

`func (o *FullRepositoryPermissions) GetMaintain() bool`

GetMaintain returns the Maintain field if non-nil, zero value otherwise.

### GetMaintainOk

`func (o *FullRepositoryPermissions) GetMaintainOk() (*bool, bool)`

GetMaintainOk returns a tuple with the Maintain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintain

`func (o *FullRepositoryPermissions) SetMaintain(v bool)`

SetMaintain sets Maintain field to given value.

### HasMaintain

`func (o *FullRepositoryPermissions) HasMaintain() bool`

HasMaintain returns a boolean if a field has been set.

### GetPush

`func (o *FullRepositoryPermissions) GetPush() bool`

GetPush returns the Push field if non-nil, zero value otherwise.

### GetPushOk

`func (o *FullRepositoryPermissions) GetPushOk() (*bool, bool)`

GetPushOk returns a tuple with the Push field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPush

`func (o *FullRepositoryPermissions) SetPush(v bool)`

SetPush sets Push field to given value.


### GetTriage

`func (o *FullRepositoryPermissions) GetTriage() bool`

GetTriage returns the Triage field if non-nil, zero value otherwise.

### GetTriageOk

`func (o *FullRepositoryPermissions) GetTriageOk() (*bool, bool)`

GetTriageOk returns a tuple with the Triage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriage

`func (o *FullRepositoryPermissions) SetTriage(v bool)`

SetTriage sets Triage field to given value.

### HasTriage

`func (o *FullRepositoryPermissions) HasTriage() bool`

HasTriage returns a boolean if a field has been set.

### GetPull

`func (o *FullRepositoryPermissions) GetPull() bool`

GetPull returns the Pull field if non-nil, zero value otherwise.

### GetPullOk

`func (o *FullRepositoryPermissions) GetPullOk() (*bool, bool)`

GetPullOk returns a tuple with the Pull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPull

`func (o *FullRepositoryPermissions) SetPull(v bool)`

SetPull sets Pull field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


