# ProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowances

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]SimpleUser**](SimpleUser.md) |  | 
**Teams** | [**[]Team**](Team.md) |  | 
**Apps** | Pointer to [**[]Integration**](Integration.md) |  | [optional] 

## Methods

### NewProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowances

`func NewProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowances(users []SimpleUser, teams []Team, ) *ProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowances`

NewProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowances instantiates a new ProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowances object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowancesWithDefaults

`func NewProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowancesWithDefaults() *ProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowances`

NewProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowancesWithDefaults instantiates a new ProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowances object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *ProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowances) GetUsers() []SimpleUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowances) GetUsersOk() (*[]SimpleUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowances) SetUsers(v []SimpleUser)`

SetUsers sets Users field to given value.


### GetTeams

`func (o *ProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowances) GetTeams() []Team`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *ProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowances) GetTeamsOk() (*[]Team, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *ProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowances) SetTeams(v []Team)`

SetTeams sets Teams field to given value.


### GetApps

`func (o *ProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowances) GetApps() []Integration`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *ProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowances) GetAppsOk() (*[]Integration, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *ProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowances) SetApps(v []Integration)`

SetApps sets Apps field to given value.

### HasApps

`func (o *ProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowances) HasApps() bool`

HasApps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


