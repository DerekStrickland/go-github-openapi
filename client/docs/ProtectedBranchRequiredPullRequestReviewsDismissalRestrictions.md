# ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**UsersUrl** | **string** |  | 
**TeamsUrl** | **string** |  | 
**Users** | [**[]SimpleUser**](SimpleUser.md) |  | 
**Teams** | [**[]Team**](Team.md) |  | 
**Apps** | Pointer to [**[]Integration**](Integration.md) |  | [optional] 

## Methods

### NewProtectedBranchRequiredPullRequestReviewsDismissalRestrictions

`func NewProtectedBranchRequiredPullRequestReviewsDismissalRestrictions(url string, usersUrl string, teamsUrl string, users []SimpleUser, teams []Team, ) *ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions`

NewProtectedBranchRequiredPullRequestReviewsDismissalRestrictions instantiates a new ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedBranchRequiredPullRequestReviewsDismissalRestrictionsWithDefaults

`func NewProtectedBranchRequiredPullRequestReviewsDismissalRestrictionsWithDefaults() *ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions`

NewProtectedBranchRequiredPullRequestReviewsDismissalRestrictionsWithDefaults instantiates a new ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUsersUrl

`func (o *ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions) GetUsersUrl() string`

GetUsersUrl returns the UsersUrl field if non-nil, zero value otherwise.

### GetUsersUrlOk

`func (o *ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions) GetUsersUrlOk() (*string, bool)`

GetUsersUrlOk returns a tuple with the UsersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersUrl

`func (o *ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions) SetUsersUrl(v string)`

SetUsersUrl sets UsersUrl field to given value.


### GetTeamsUrl

`func (o *ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions) GetTeamsUrl() string`

GetTeamsUrl returns the TeamsUrl field if non-nil, zero value otherwise.

### GetTeamsUrlOk

`func (o *ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions) GetTeamsUrlOk() (*string, bool)`

GetTeamsUrlOk returns a tuple with the TeamsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsUrl

`func (o *ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions) SetTeamsUrl(v string)`

SetTeamsUrl sets TeamsUrl field to given value.


### GetUsers

`func (o *ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions) GetUsers() []SimpleUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions) GetUsersOk() (*[]SimpleUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions) SetUsers(v []SimpleUser)`

SetUsers sets Users field to given value.


### GetTeams

`func (o *ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions) GetTeams() []Team`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions) GetTeamsOk() (*[]Team, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions) SetTeams(v []Team)`

SetTeams sets Teams field to given value.


### GetApps

`func (o *ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions) GetApps() []Integration`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions) GetAppsOk() (*[]Integration, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions) SetApps(v []Integration)`

SetApps sets Apps field to given value.

### HasApps

`func (o *ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions) HasApps() bool`

HasApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


