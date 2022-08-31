# ReposUpdateBranchProtectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequiredStatusChecks** | [**NullableReposUpdateBranchProtectionRequestRequiredStatusChecks**](ReposUpdateBranchProtectionRequestRequiredStatusChecks.md) |  | 
**EnforceAdmins** | **NullableBool** | Enforce all configured restrictions for administrators. Set to &#x60;true&#x60; to enforce required status checks for repository administrators. Set to &#x60;null&#x60; to disable. | 
**RequiredPullRequestReviews** | [**NullableReposUpdateBranchProtectionRequestRequiredPullRequestReviews**](ReposUpdateBranchProtectionRequestRequiredPullRequestReviews.md) |  | 
**Restrictions** | [**NullableReposUpdateBranchProtectionRequestRestrictions**](ReposUpdateBranchProtectionRequestRestrictions.md) |  | 
**RequiredLinearHistory** | Pointer to **bool** | Enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch. Set to &#x60;true&#x60; to enforce a linear commit history. Set to &#x60;false&#x60; to disable a linear commit Git history. Your repository must allow squash merging or rebase merging before you can enable a linear commit history. Default: &#x60;false&#x60;. For more information, see \&quot;[Requiring a linear commit history](https://docs.github.com/github/administering-a-repository/requiring-a-linear-commit-history)\&quot; in the GitHub Help documentation. | [optional] 
**AllowForcePushes** | Pointer to **NullableBool** | Permits force pushes to the protected branch by anyone with write access to the repository. Set to &#x60;true&#x60; to allow force pushes. Set to &#x60;false&#x60; or &#x60;null&#x60; to block force pushes. Default: &#x60;false&#x60;. For more information, see \&quot;[Enabling force pushes to a protected branch](https://docs.github.com/en/github/administering-a-repository/enabling-force-pushes-to-a-protected-branch)\&quot; in the GitHub Help documentation.\&quot; | [optional] 
**AllowDeletions** | Pointer to **bool** | Allows deletion of the protected branch by anyone with write access to the repository. Set to &#x60;false&#x60; to prevent deletion of the protected branch. Default: &#x60;false&#x60;. For more information, see \&quot;[Enabling force pushes to a protected branch](https://docs.github.com/en/github/administering-a-repository/enabling-force-pushes-to-a-protected-branch)\&quot; in the GitHub Help documentation. | [optional] 
**BlockCreations** | Pointer to **bool** | If set to &#x60;true&#x60;, the &#x60;restrictions&#x60; branch protection settings which limits who can push will also block pushes which create new branches, unless the push is initiated by a user, team, or app which has the ability to push. Set to &#x60;true&#x60; to restrict new branch creation. Default: &#x60;false&#x60;. | [optional] 
**RequiredConversationResolution** | Pointer to **bool** | Requires all conversations on code to be resolved before a pull request can be merged into a branch that matches this rule. Set to &#x60;false&#x60; to disable. Default: &#x60;false&#x60;. | [optional] 

## Methods

### NewReposUpdateBranchProtectionRequest

`func NewReposUpdateBranchProtectionRequest(requiredStatusChecks NullableReposUpdateBranchProtectionRequestRequiredStatusChecks, enforceAdmins NullableBool, requiredPullRequestReviews NullableReposUpdateBranchProtectionRequestRequiredPullRequestReviews, restrictions NullableReposUpdateBranchProtectionRequestRestrictions, ) *ReposUpdateBranchProtectionRequest`

NewReposUpdateBranchProtectionRequest instantiates a new ReposUpdateBranchProtectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposUpdateBranchProtectionRequestWithDefaults

`func NewReposUpdateBranchProtectionRequestWithDefaults() *ReposUpdateBranchProtectionRequest`

NewReposUpdateBranchProtectionRequestWithDefaults instantiates a new ReposUpdateBranchProtectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequiredStatusChecks

`func (o *ReposUpdateBranchProtectionRequest) GetRequiredStatusChecks() ReposUpdateBranchProtectionRequestRequiredStatusChecks`

GetRequiredStatusChecks returns the RequiredStatusChecks field if non-nil, zero value otherwise.

### GetRequiredStatusChecksOk

`func (o *ReposUpdateBranchProtectionRequest) GetRequiredStatusChecksOk() (*ReposUpdateBranchProtectionRequestRequiredStatusChecks, bool)`

GetRequiredStatusChecksOk returns a tuple with the RequiredStatusChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredStatusChecks

`func (o *ReposUpdateBranchProtectionRequest) SetRequiredStatusChecks(v ReposUpdateBranchProtectionRequestRequiredStatusChecks)`

SetRequiredStatusChecks sets RequiredStatusChecks field to given value.


### SetRequiredStatusChecksNil

`func (o *ReposUpdateBranchProtectionRequest) SetRequiredStatusChecksNil(b bool)`

 SetRequiredStatusChecksNil sets the value for RequiredStatusChecks to be an explicit nil

### UnsetRequiredStatusChecks
`func (o *ReposUpdateBranchProtectionRequest) UnsetRequiredStatusChecks()`

UnsetRequiredStatusChecks ensures that no value is present for RequiredStatusChecks, not even an explicit nil
### GetEnforceAdmins

`func (o *ReposUpdateBranchProtectionRequest) GetEnforceAdmins() bool`

GetEnforceAdmins returns the EnforceAdmins field if non-nil, zero value otherwise.

### GetEnforceAdminsOk

`func (o *ReposUpdateBranchProtectionRequest) GetEnforceAdminsOk() (*bool, bool)`

GetEnforceAdminsOk returns a tuple with the EnforceAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceAdmins

`func (o *ReposUpdateBranchProtectionRequest) SetEnforceAdmins(v bool)`

SetEnforceAdmins sets EnforceAdmins field to given value.


### SetEnforceAdminsNil

`func (o *ReposUpdateBranchProtectionRequest) SetEnforceAdminsNil(b bool)`

 SetEnforceAdminsNil sets the value for EnforceAdmins to be an explicit nil

### UnsetEnforceAdmins
`func (o *ReposUpdateBranchProtectionRequest) UnsetEnforceAdmins()`

UnsetEnforceAdmins ensures that no value is present for EnforceAdmins, not even an explicit nil
### GetRequiredPullRequestReviews

`func (o *ReposUpdateBranchProtectionRequest) GetRequiredPullRequestReviews() ReposUpdateBranchProtectionRequestRequiredPullRequestReviews`

GetRequiredPullRequestReviews returns the RequiredPullRequestReviews field if non-nil, zero value otherwise.

### GetRequiredPullRequestReviewsOk

`func (o *ReposUpdateBranchProtectionRequest) GetRequiredPullRequestReviewsOk() (*ReposUpdateBranchProtectionRequestRequiredPullRequestReviews, bool)`

GetRequiredPullRequestReviewsOk returns a tuple with the RequiredPullRequestReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredPullRequestReviews

`func (o *ReposUpdateBranchProtectionRequest) SetRequiredPullRequestReviews(v ReposUpdateBranchProtectionRequestRequiredPullRequestReviews)`

SetRequiredPullRequestReviews sets RequiredPullRequestReviews field to given value.


### SetRequiredPullRequestReviewsNil

`func (o *ReposUpdateBranchProtectionRequest) SetRequiredPullRequestReviewsNil(b bool)`

 SetRequiredPullRequestReviewsNil sets the value for RequiredPullRequestReviews to be an explicit nil

### UnsetRequiredPullRequestReviews
`func (o *ReposUpdateBranchProtectionRequest) UnsetRequiredPullRequestReviews()`

UnsetRequiredPullRequestReviews ensures that no value is present for RequiredPullRequestReviews, not even an explicit nil
### GetRestrictions

`func (o *ReposUpdateBranchProtectionRequest) GetRestrictions() ReposUpdateBranchProtectionRequestRestrictions`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *ReposUpdateBranchProtectionRequest) GetRestrictionsOk() (*ReposUpdateBranchProtectionRequestRestrictions, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *ReposUpdateBranchProtectionRequest) SetRestrictions(v ReposUpdateBranchProtectionRequestRestrictions)`

SetRestrictions sets Restrictions field to given value.


### SetRestrictionsNil

`func (o *ReposUpdateBranchProtectionRequest) SetRestrictionsNil(b bool)`

 SetRestrictionsNil sets the value for Restrictions to be an explicit nil

### UnsetRestrictions
`func (o *ReposUpdateBranchProtectionRequest) UnsetRestrictions()`

UnsetRestrictions ensures that no value is present for Restrictions, not even an explicit nil
### GetRequiredLinearHistory

`func (o *ReposUpdateBranchProtectionRequest) GetRequiredLinearHistory() bool`

GetRequiredLinearHistory returns the RequiredLinearHistory field if non-nil, zero value otherwise.

### GetRequiredLinearHistoryOk

`func (o *ReposUpdateBranchProtectionRequest) GetRequiredLinearHistoryOk() (*bool, bool)`

GetRequiredLinearHistoryOk returns a tuple with the RequiredLinearHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredLinearHistory

`func (o *ReposUpdateBranchProtectionRequest) SetRequiredLinearHistory(v bool)`

SetRequiredLinearHistory sets RequiredLinearHistory field to given value.

### HasRequiredLinearHistory

`func (o *ReposUpdateBranchProtectionRequest) HasRequiredLinearHistory() bool`

HasRequiredLinearHistory returns a boolean if a field has been set.

### GetAllowForcePushes

`func (o *ReposUpdateBranchProtectionRequest) GetAllowForcePushes() bool`

GetAllowForcePushes returns the AllowForcePushes field if non-nil, zero value otherwise.

### GetAllowForcePushesOk

`func (o *ReposUpdateBranchProtectionRequest) GetAllowForcePushesOk() (*bool, bool)`

GetAllowForcePushesOk returns a tuple with the AllowForcePushes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowForcePushes

`func (o *ReposUpdateBranchProtectionRequest) SetAllowForcePushes(v bool)`

SetAllowForcePushes sets AllowForcePushes field to given value.

### HasAllowForcePushes

`func (o *ReposUpdateBranchProtectionRequest) HasAllowForcePushes() bool`

HasAllowForcePushes returns a boolean if a field has been set.

### SetAllowForcePushesNil

`func (o *ReposUpdateBranchProtectionRequest) SetAllowForcePushesNil(b bool)`

 SetAllowForcePushesNil sets the value for AllowForcePushes to be an explicit nil

### UnsetAllowForcePushes
`func (o *ReposUpdateBranchProtectionRequest) UnsetAllowForcePushes()`

UnsetAllowForcePushes ensures that no value is present for AllowForcePushes, not even an explicit nil
### GetAllowDeletions

`func (o *ReposUpdateBranchProtectionRequest) GetAllowDeletions() bool`

GetAllowDeletions returns the AllowDeletions field if non-nil, zero value otherwise.

### GetAllowDeletionsOk

`func (o *ReposUpdateBranchProtectionRequest) GetAllowDeletionsOk() (*bool, bool)`

GetAllowDeletionsOk returns a tuple with the AllowDeletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDeletions

`func (o *ReposUpdateBranchProtectionRequest) SetAllowDeletions(v bool)`

SetAllowDeletions sets AllowDeletions field to given value.

### HasAllowDeletions

`func (o *ReposUpdateBranchProtectionRequest) HasAllowDeletions() bool`

HasAllowDeletions returns a boolean if a field has been set.

### GetBlockCreations

`func (o *ReposUpdateBranchProtectionRequest) GetBlockCreations() bool`

GetBlockCreations returns the BlockCreations field if non-nil, zero value otherwise.

### GetBlockCreationsOk

`func (o *ReposUpdateBranchProtectionRequest) GetBlockCreationsOk() (*bool, bool)`

GetBlockCreationsOk returns a tuple with the BlockCreations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockCreations

`func (o *ReposUpdateBranchProtectionRequest) SetBlockCreations(v bool)`

SetBlockCreations sets BlockCreations field to given value.

### HasBlockCreations

`func (o *ReposUpdateBranchProtectionRequest) HasBlockCreations() bool`

HasBlockCreations returns a boolean if a field has been set.

### GetRequiredConversationResolution

`func (o *ReposUpdateBranchProtectionRequest) GetRequiredConversationResolution() bool`

GetRequiredConversationResolution returns the RequiredConversationResolution field if non-nil, zero value otherwise.

### GetRequiredConversationResolutionOk

`func (o *ReposUpdateBranchProtectionRequest) GetRequiredConversationResolutionOk() (*bool, bool)`

GetRequiredConversationResolutionOk returns a tuple with the RequiredConversationResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredConversationResolution

`func (o *ReposUpdateBranchProtectionRequest) SetRequiredConversationResolution(v bool)`

SetRequiredConversationResolution sets RequiredConversationResolution field to given value.

### HasRequiredConversationResolution

`func (o *ReposUpdateBranchProtectionRequest) HasRequiredConversationResolution() bool`

HasRequiredConversationResolution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


