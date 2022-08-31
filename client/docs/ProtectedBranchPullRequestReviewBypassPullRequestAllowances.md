# ProtectedBranchPullRequestReviewBypassPullRequestAllowances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | Pointer to [**[]SimpleUser**](SimpleUser.md) | The list of users allowed to bypass pull request requirements. | [optional] 
**Teams** | Pointer to [**[]Team**](Team.md) | The list of teams allowed to bypass pull request requirements. | [optional] 
**Apps** | Pointer to [**[]Integration**](Integration.md) | The list of apps allowed to bypass pull request requirements. | [optional] 

## Methods

### NewProtectedBranchPullRequestReviewBypassPullRequestAllowances

`func NewProtectedBranchPullRequestReviewBypassPullRequestAllowances() *ProtectedBranchPullRequestReviewBypassPullRequestAllowances`

NewProtectedBranchPullRequestReviewBypassPullRequestAllowances instantiates a new ProtectedBranchPullRequestReviewBypassPullRequestAllowances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedBranchPullRequestReviewBypassPullRequestAllowancesWithDefaults

`func NewProtectedBranchPullRequestReviewBypassPullRequestAllowancesWithDefaults() *ProtectedBranchPullRequestReviewBypassPullRequestAllowances`

NewProtectedBranchPullRequestReviewBypassPullRequestAllowancesWithDefaults instantiates a new ProtectedBranchPullRequestReviewBypassPullRequestAllowances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) GetUsers() []SimpleUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) GetUsersOk() (*[]SimpleUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) SetUsers(v []SimpleUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetTeams

`func (o *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) GetTeams() []Team`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) GetTeamsOk() (*[]Team, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) SetTeams(v []Team)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetApps

`func (o *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) GetApps() []Integration`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) GetAppsOk() (*[]Integration, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) SetApps(v []Integration)`

SetApps sets Apps field to given value.

### HasApps

`func (o *ProtectedBranchPullRequestReviewBypassPullRequestAllowances) HasApps() bool`

HasApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


