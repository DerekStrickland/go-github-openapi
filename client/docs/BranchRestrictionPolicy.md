# BranchRestrictionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**UsersUrl** | **string** |  | 
**TeamsUrl** | **string** |  | 
**AppsUrl** | **string** |  | 
**Users** | [**[]RepositoryTemplateRepositoryOwner**](RepositoryTemplateRepositoryOwner.md) |  | 
**Teams** | [**[]BranchRestrictionPolicyTeamsInner**](BranchRestrictionPolicyTeamsInner.md) |  | 
**Apps** | [**[]BranchRestrictionPolicyAppsInner**](BranchRestrictionPolicyAppsInner.md) |  | 

## Methods

### NewBranchRestrictionPolicy

`func NewBranchRestrictionPolicy(url string, usersUrl string, teamsUrl string, appsUrl string, users []RepositoryTemplateRepositoryOwner, teams []BranchRestrictionPolicyTeamsInner, apps []BranchRestrictionPolicyAppsInner, ) *BranchRestrictionPolicy`

NewBranchRestrictionPolicy instantiates a new BranchRestrictionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchRestrictionPolicyWithDefaults

`func NewBranchRestrictionPolicyWithDefaults() *BranchRestrictionPolicy`

NewBranchRestrictionPolicyWithDefaults instantiates a new BranchRestrictionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *BranchRestrictionPolicy) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BranchRestrictionPolicy) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BranchRestrictionPolicy) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUsersUrl

`func (o *BranchRestrictionPolicy) GetUsersUrl() string`

GetUsersUrl returns the UsersUrl field if non-nil, zero value otherwise.

### GetUsersUrlOk

`func (o *BranchRestrictionPolicy) GetUsersUrlOk() (*string, bool)`

GetUsersUrlOk returns a tuple with the UsersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersUrl

`func (o *BranchRestrictionPolicy) SetUsersUrl(v string)`

SetUsersUrl sets UsersUrl field to given value.


### GetTeamsUrl

`func (o *BranchRestrictionPolicy) GetTeamsUrl() string`

GetTeamsUrl returns the TeamsUrl field if non-nil, zero value otherwise.

### GetTeamsUrlOk

`func (o *BranchRestrictionPolicy) GetTeamsUrlOk() (*string, bool)`

GetTeamsUrlOk returns a tuple with the TeamsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsUrl

`func (o *BranchRestrictionPolicy) SetTeamsUrl(v string)`

SetTeamsUrl sets TeamsUrl field to given value.


### GetAppsUrl

`func (o *BranchRestrictionPolicy) GetAppsUrl() string`

GetAppsUrl returns the AppsUrl field if non-nil, zero value otherwise.

### GetAppsUrlOk

`func (o *BranchRestrictionPolicy) GetAppsUrlOk() (*string, bool)`

GetAppsUrlOk returns a tuple with the AppsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppsUrl

`func (o *BranchRestrictionPolicy) SetAppsUrl(v string)`

SetAppsUrl sets AppsUrl field to given value.


### GetUsers

`func (o *BranchRestrictionPolicy) GetUsers() []RepositoryTemplateRepositoryOwner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *BranchRestrictionPolicy) GetUsersOk() (*[]RepositoryTemplateRepositoryOwner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *BranchRestrictionPolicy) SetUsers(v []RepositoryTemplateRepositoryOwner)`

SetUsers sets Users field to given value.


### GetTeams

`func (o *BranchRestrictionPolicy) GetTeams() []BranchRestrictionPolicyTeamsInner`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *BranchRestrictionPolicy) GetTeamsOk() (*[]BranchRestrictionPolicyTeamsInner, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *BranchRestrictionPolicy) SetTeams(v []BranchRestrictionPolicyTeamsInner)`

SetTeams sets Teams field to given value.


### GetApps

`func (o *BranchRestrictionPolicy) GetApps() []BranchRestrictionPolicyAppsInner`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *BranchRestrictionPolicy) GetAppsOk() (*[]BranchRestrictionPolicyAppsInner, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *BranchRestrictionPolicy) SetApps(v []BranchRestrictionPolicyAppsInner)`

SetApps sets Apps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


