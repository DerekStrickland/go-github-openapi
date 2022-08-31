# TeamPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pull** | **bool** |  | 
**Triage** | **bool** |  | 
**Push** | **bool** |  | 
**Maintain** | **bool** |  | 
**Admin** | **bool** |  | 

## Methods

### NewTeamPermissions

`func NewTeamPermissions(pull bool, triage bool, push bool, maintain bool, admin bool, ) *TeamPermissions`

NewTeamPermissions instantiates a new TeamPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamPermissionsWithDefaults

`func NewTeamPermissionsWithDefaults() *TeamPermissions`

NewTeamPermissionsWithDefaults instantiates a new TeamPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPull

`func (o *TeamPermissions) GetPull() bool`

GetPull returns the Pull field if non-nil, zero value otherwise.

### GetPullOk

`func (o *TeamPermissions) GetPullOk() (*bool, bool)`

GetPullOk returns a tuple with the Pull field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPull

`func (o *TeamPermissions) SetPull(v bool)`

SetPull sets Pull field to given value.


### GetTriage

`func (o *TeamPermissions) GetTriage() bool`

GetTriage returns the Triage field if non-nil, zero value otherwise.

### GetTriageOk

`func (o *TeamPermissions) GetTriageOk() (*bool, bool)`

GetTriageOk returns a tuple with the Triage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriage

`func (o *TeamPermissions) SetTriage(v bool)`

SetTriage sets Triage field to given value.


### GetPush

`func (o *TeamPermissions) GetPush() bool`

GetPush returns the Push field if non-nil, zero value otherwise.

### GetPushOk

`func (o *TeamPermissions) GetPushOk() (*bool, bool)`

GetPushOk returns a tuple with the Push field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPush

`func (o *TeamPermissions) SetPush(v bool)`

SetPush sets Push field to given value.


### GetMaintain

`func (o *TeamPermissions) GetMaintain() bool`

GetMaintain returns the Maintain field if non-nil, zero value otherwise.

### GetMaintainOk

`func (o *TeamPermissions) GetMaintainOk() (*bool, bool)`

GetMaintainOk returns a tuple with the Maintain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintain

`func (o *TeamPermissions) SetMaintain(v bool)`

SetMaintain sets Maintain field to given value.


### GetAdmin

`func (o *TeamPermissions) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *TeamPermissions) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *TeamPermissions) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


