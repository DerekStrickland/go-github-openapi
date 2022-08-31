# ProtectedBranchPullRequestReviewDismissalRestrictions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]SimpleUser**](SimpleUser.md) | The list of users with review dismissal access. | [optional] 
**Teams** | Pointer to [**[]Team**](Team.md) | The list of teams with review dismissal access. | [optional] 
**Apps** | Pointer to [**[]Integration**](Integration.md) | The list of apps with review dismissal access. | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**UsersUrl** | Pointer to **string** |  | [optional] 
**TeamsUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewProtectedBranchPullRequestReviewDismissalRestrictions

`func NewProtectedBranchPullRequestReviewDismissalRestrictions() *ProtectedBranchPullRequestReviewDismissalRestrictions`

NewProtectedBranchPullRequestReviewDismissalRestrictions instantiates a new ProtectedBranchPullRequestReviewDismissalRestrictions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedBranchPullRequestReviewDismissalRestrictionsWithDefaults

`func NewProtectedBranchPullRequestReviewDismissalRestrictionsWithDefaults() *ProtectedBranchPullRequestReviewDismissalRestrictions`

NewProtectedBranchPullRequestReviewDismissalRestrictionsWithDefaults instantiates a new ProtectedBranchPullRequestReviewDismissalRestrictions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *ProtectedBranchPullRequestReviewDismissalRestrictions) GetUsers() []SimpleUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ProtectedBranchPullRequestReviewDismissalRestrictions) GetUsersOk() (*[]SimpleUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ProtectedBranchPullRequestReviewDismissalRestrictions) SetUsers(v []SimpleUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ProtectedBranchPullRequestReviewDismissalRestrictions) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetTeams

`func (o *ProtectedBranchPullRequestReviewDismissalRestrictions) GetTeams() []Team`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *ProtectedBranchPullRequestReviewDismissalRestrictions) GetTeamsOk() (*[]Team, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *ProtectedBranchPullRequestReviewDismissalRestrictions) SetTeams(v []Team)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *ProtectedBranchPullRequestReviewDismissalRestrictions) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetApps

`func (o *ProtectedBranchPullRequestReviewDismissalRestrictions) GetApps() []Integration`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *ProtectedBranchPullRequestReviewDismissalRestrictions) GetAppsOk() (*[]Integration, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *ProtectedBranchPullRequestReviewDismissalRestrictions) SetApps(v []Integration)`

SetApps sets Apps field to given value.

### HasApps

`func (o *ProtectedBranchPullRequestReviewDismissalRestrictions) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetUrl

`func (o *ProtectedBranchPullRequestReviewDismissalRestrictions) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProtectedBranchPullRequestReviewDismissalRestrictions) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProtectedBranchPullRequestReviewDismissalRestrictions) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProtectedBranchPullRequestReviewDismissalRestrictions) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsersUrl

`func (o *ProtectedBranchPullRequestReviewDismissalRestrictions) GetUsersUrl() string`

GetUsersUrl returns the UsersUrl field if non-nil, zero value otherwise.

### GetUsersUrlOk

`func (o *ProtectedBranchPullRequestReviewDismissalRestrictions) GetUsersUrlOk() (*string, bool)`

GetUsersUrlOk returns a tuple with the UsersUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsersUrl

`func (o *ProtectedBranchPullRequestReviewDismissalRestrictions) SetUsersUrl(v string)`

SetUsersUrl sets UsersUrl field to given value.

### HasUsersUrl

`func (o *ProtectedBranchPullRequestReviewDismissalRestrictions) HasUsersUrl() bool`

HasUsersUrl returns a boolean if a field has been set.

### GetTeamsUrl

`func (o *ProtectedBranchPullRequestReviewDismissalRestrictions) GetTeamsUrl() string`

GetTeamsUrl returns the TeamsUrl field if non-nil, zero value otherwise.

### GetTeamsUrlOk

`func (o *ProtectedBranchPullRequestReviewDismissalRestrictions) GetTeamsUrlOk() (*string, bool)`

GetTeamsUrlOk returns a tuple with the TeamsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsUrl

`func (o *ProtectedBranchPullRequestReviewDismissalRestrictions) SetTeamsUrl(v string)`

SetTeamsUrl sets TeamsUrl field to given value.

### HasTeamsUrl

`func (o *ProtectedBranchPullRequestReviewDismissalRestrictions) HasTeamsUrl() bool`

HasTeamsUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


