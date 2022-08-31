# TeamMembership

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Role** | **string** | The role of the user in the team. | [default to "member"]
**State** | **string** | The state of the user&#39;s membership in the team. | 

## Methods

### NewTeamMembership

`func NewTeamMembership(url string, role string, state string, ) *TeamMembership`

NewTeamMembership instantiates a new TeamMembership object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamMembershipWithDefaults

`func NewTeamMembershipWithDefaults() *TeamMembership`

NewTeamMembershipWithDefaults instantiates a new TeamMembership object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *TeamMembership) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TeamMembership) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TeamMembership) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRole

`func (o *TeamMembership) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *TeamMembership) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *TeamMembership) SetRole(v string)`

SetRole sets Role field to given value.


### GetState

`func (o *TeamMembership) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TeamMembership) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TeamMembership) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


