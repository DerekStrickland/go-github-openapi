# CollaboratorPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pull** | **bool** |  | 
**Triage** | Pointer to **bool** |  | [optional] 
**Push** | **bool** |  | 
**Maintain** | Pointer to **bool** |  | [optional] 
**Admin** | **bool** |  | 

## Methods

### NewCollaboratorPermissions

`func NewCollaboratorPermissions(pull bool, push bool, admin bool, ) *CollaboratorPermissions`

NewCollaboratorPermissions instantiates a new CollaboratorPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollaboratorPermissionsWithDefaults

`func NewCollaboratorPermissionsWithDefaults() *CollaboratorPermissions`

NewCollaboratorPermissionsWithDefaults instantiates a new CollaboratorPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPull

`func (o *CollaboratorPermissions) GetPull() bool`

GetPull returns the Pull field if non-nil, zero value otherwise.

### GetPullOk

`func (o *CollaboratorPermissions) GetPullOk() (*bool, bool)`

GetPullOk returns a tuple with the Pull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPull

`func (o *CollaboratorPermissions) SetPull(v bool)`

SetPull sets Pull field to given value.


### GetTriage

`func (o *CollaboratorPermissions) GetTriage() bool`

GetTriage returns the Triage field if non-nil, zero value otherwise.

### GetTriageOk

`func (o *CollaboratorPermissions) GetTriageOk() (*bool, bool)`

GetTriageOk returns a tuple with the Triage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriage

`func (o *CollaboratorPermissions) SetTriage(v bool)`

SetTriage sets Triage field to given value.

### HasTriage

`func (o *CollaboratorPermissions) HasTriage() bool`

HasTriage returns a boolean if a field has been set.

### GetPush

`func (o *CollaboratorPermissions) GetPush() bool`

GetPush returns the Push field if non-nil, zero value otherwise.

### GetPushOk

`func (o *CollaboratorPermissions) GetPushOk() (*bool, bool)`

GetPushOk returns a tuple with the Push field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPush

`func (o *CollaboratorPermissions) SetPush(v bool)`

SetPush sets Push field to given value.


### GetMaintain

`func (o *CollaboratorPermissions) GetMaintain() bool`

GetMaintain returns the Maintain field if non-nil, zero value otherwise.

### GetMaintainOk

`func (o *CollaboratorPermissions) GetMaintainOk() (*bool, bool)`

GetMaintainOk returns a tuple with the Maintain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintain

`func (o *CollaboratorPermissions) SetMaintain(v bool)`

SetMaintain sets Maintain field to given value.

### HasMaintain

`func (o *CollaboratorPermissions) HasMaintain() bool`

HasMaintain returns a boolean if a field has been set.

### GetAdmin

`func (o *CollaboratorPermissions) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *CollaboratorPermissions) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *CollaboratorPermissions) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


