# ReposUpdateBranchProtectionRequestRestrictions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | **[]string** | The list of user &#x60;login&#x60;s with push access | 
**Teams** | **[]string** | The list of team &#x60;slug&#x60;s with push access | 
**Apps** | Pointer to **[]string** | The list of app &#x60;slug&#x60;s with push access | [optional] 

## Methods

### NewReposUpdateBranchProtectionRequestRestrictions

`func NewReposUpdateBranchProtectionRequestRestrictions(users []string, teams []string, ) *ReposUpdateBranchProtectionRequestRestrictions`

NewReposUpdateBranchProtectionRequestRestrictions instantiates a new ReposUpdateBranchProtectionRequestRestrictions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposUpdateBranchProtectionRequestRestrictionsWithDefaults

`func NewReposUpdateBranchProtectionRequestRestrictionsWithDefaults() *ReposUpdateBranchProtectionRequestRestrictions`

NewReposUpdateBranchProtectionRequestRestrictionsWithDefaults instantiates a new ReposUpdateBranchProtectionRequestRestrictions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *ReposUpdateBranchProtectionRequestRestrictions) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ReposUpdateBranchProtectionRequestRestrictions) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ReposUpdateBranchProtectionRequestRestrictions) SetUsers(v []string)`

SetUsers sets Users field to given value.


### GetTeams

`func (o *ReposUpdateBranchProtectionRequestRestrictions) GetTeams() []string`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *ReposUpdateBranchProtectionRequestRestrictions) GetTeamsOk() (*[]string, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *ReposUpdateBranchProtectionRequestRestrictions) SetTeams(v []string)`

SetTeams sets Teams field to given value.


### GetApps

`func (o *ReposUpdateBranchProtectionRequestRestrictions) GetApps() []string`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *ReposUpdateBranchProtectionRequestRestrictions) GetAppsOk() (*[]string, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *ReposUpdateBranchProtectionRequestRestrictions) SetApps(v []string)`

SetApps sets Apps field to given value.

### HasApps

`func (o *ReposUpdateBranchProtectionRequestRestrictions) HasApps() bool`

HasApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


