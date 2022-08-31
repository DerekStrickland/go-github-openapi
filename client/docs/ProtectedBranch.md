# ProtectedBranch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**RequiredStatusChecks** | Pointer to [**StatusCheckPolicy**](StatusCheckPolicy.md) |  | [optional] 
**RequiredPullRequestReviews** | Pointer to [**ProtectedBranchRequiredPullRequestReviews**](ProtectedBranchRequiredPullRequestReviews.md) |  | [optional] 
**RequiredSignatures** | Pointer to [**BranchProtectionRequiredSignatures**](BranchProtectionRequiredSignatures.md) |  | [optional] 
**EnforceAdmins** | Pointer to [**ProtectedBranchEnforceAdmins**](ProtectedBranchEnforceAdmins.md) |  | [optional] 
**RequiredLinearHistory** | Pointer to [**ProtectedBranchRequiredLinearHistory**](ProtectedBranchRequiredLinearHistory.md) |  | [optional] 
**AllowForcePushes** | Pointer to [**ProtectedBranchRequiredLinearHistory**](ProtectedBranchRequiredLinearHistory.md) |  | [optional] 
**AllowDeletions** | Pointer to [**ProtectedBranchRequiredLinearHistory**](ProtectedBranchRequiredLinearHistory.md) |  | [optional] 
**Restrictions** | Pointer to [**BranchRestrictionPolicy**](BranchRestrictionPolicy.md) |  | [optional] 
**RequiredConversationResolution** | Pointer to [**ProtectedBranchRequiredConversationResolution**](ProtectedBranchRequiredConversationResolution.md) |  | [optional] 
**BlockCreations** | Pointer to [**ProtectedBranchRequiredLinearHistory**](ProtectedBranchRequiredLinearHistory.md) |  | [optional] 

## Methods

### NewProtectedBranch

`func NewProtectedBranch(url string, ) *ProtectedBranch`

NewProtectedBranch instantiates a new ProtectedBranch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedBranchWithDefaults

`func NewProtectedBranchWithDefaults() *ProtectedBranch`

NewProtectedBranchWithDefaults instantiates a new ProtectedBranch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ProtectedBranch) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProtectedBranch) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProtectedBranch) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetRequiredStatusChecks

`func (o *ProtectedBranch) GetRequiredStatusChecks() StatusCheckPolicy`

GetRequiredStatusChecks returns the RequiredStatusChecks field if non-nil, zero value otherwise.

### GetRequiredStatusChecksOk

`func (o *ProtectedBranch) GetRequiredStatusChecksOk() (*StatusCheckPolicy, bool)`

GetRequiredStatusChecksOk returns a tuple with the RequiredStatusChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredStatusChecks

`func (o *ProtectedBranch) SetRequiredStatusChecks(v StatusCheckPolicy)`

SetRequiredStatusChecks sets RequiredStatusChecks field to given value.

### HasRequiredStatusChecks

`func (o *ProtectedBranch) HasRequiredStatusChecks() bool`

HasRequiredStatusChecks returns a boolean if a field has been set.

### GetRequiredPullRequestReviews

`func (o *ProtectedBranch) GetRequiredPullRequestReviews() ProtectedBranchRequiredPullRequestReviews`

GetRequiredPullRequestReviews returns the RequiredPullRequestReviews field if non-nil, zero value otherwise.

### GetRequiredPullRequestReviewsOk

`func (o *ProtectedBranch) GetRequiredPullRequestReviewsOk() (*ProtectedBranchRequiredPullRequestReviews, bool)`

GetRequiredPullRequestReviewsOk returns a tuple with the RequiredPullRequestReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredPullRequestReviews

`func (o *ProtectedBranch) SetRequiredPullRequestReviews(v ProtectedBranchRequiredPullRequestReviews)`

SetRequiredPullRequestReviews sets RequiredPullRequestReviews field to given value.

### HasRequiredPullRequestReviews

`func (o *ProtectedBranch) HasRequiredPullRequestReviews() bool`

HasRequiredPullRequestReviews returns a boolean if a field has been set.

### GetRequiredSignatures

`func (o *ProtectedBranch) GetRequiredSignatures() BranchProtectionRequiredSignatures`

GetRequiredSignatures returns the RequiredSignatures field if non-nil, zero value otherwise.

### GetRequiredSignaturesOk

`func (o *ProtectedBranch) GetRequiredSignaturesOk() (*BranchProtectionRequiredSignatures, bool)`

GetRequiredSignaturesOk returns a tuple with the RequiredSignatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredSignatures

`func (o *ProtectedBranch) SetRequiredSignatures(v BranchProtectionRequiredSignatures)`

SetRequiredSignatures sets RequiredSignatures field to given value.

### HasRequiredSignatures

`func (o *ProtectedBranch) HasRequiredSignatures() bool`

HasRequiredSignatures returns a boolean if a field has been set.

### GetEnforceAdmins

`func (o *ProtectedBranch) GetEnforceAdmins() ProtectedBranchEnforceAdmins`

GetEnforceAdmins returns the EnforceAdmins field if non-nil, zero value otherwise.

### GetEnforceAdminsOk

`func (o *ProtectedBranch) GetEnforceAdminsOk() (*ProtectedBranchEnforceAdmins, bool)`

GetEnforceAdminsOk returns a tuple with the EnforceAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceAdmins

`func (o *ProtectedBranch) SetEnforceAdmins(v ProtectedBranchEnforceAdmins)`

SetEnforceAdmins sets EnforceAdmins field to given value.

### HasEnforceAdmins

`func (o *ProtectedBranch) HasEnforceAdmins() bool`

HasEnforceAdmins returns a boolean if a field has been set.

### GetRequiredLinearHistory

`func (o *ProtectedBranch) GetRequiredLinearHistory() ProtectedBranchRequiredLinearHistory`

GetRequiredLinearHistory returns the RequiredLinearHistory field if non-nil, zero value otherwise.

### GetRequiredLinearHistoryOk

`func (o *ProtectedBranch) GetRequiredLinearHistoryOk() (*ProtectedBranchRequiredLinearHistory, bool)`

GetRequiredLinearHistoryOk returns a tuple with the RequiredLinearHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredLinearHistory

`func (o *ProtectedBranch) SetRequiredLinearHistory(v ProtectedBranchRequiredLinearHistory)`

SetRequiredLinearHistory sets RequiredLinearHistory field to given value.

### HasRequiredLinearHistory

`func (o *ProtectedBranch) HasRequiredLinearHistory() bool`

HasRequiredLinearHistory returns a boolean if a field has been set.

### GetAllowForcePushes

`func (o *ProtectedBranch) GetAllowForcePushes() ProtectedBranchRequiredLinearHistory`

GetAllowForcePushes returns the AllowForcePushes field if non-nil, zero value otherwise.

### GetAllowForcePushesOk

`func (o *ProtectedBranch) GetAllowForcePushesOk() (*ProtectedBranchRequiredLinearHistory, bool)`

GetAllowForcePushesOk returns a tuple with the AllowForcePushes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowForcePushes

`func (o *ProtectedBranch) SetAllowForcePushes(v ProtectedBranchRequiredLinearHistory)`

SetAllowForcePushes sets AllowForcePushes field to given value.

### HasAllowForcePushes

`func (o *ProtectedBranch) HasAllowForcePushes() bool`

HasAllowForcePushes returns a boolean if a field has been set.

### GetAllowDeletions

`func (o *ProtectedBranch) GetAllowDeletions() ProtectedBranchRequiredLinearHistory`

GetAllowDeletions returns the AllowDeletions field if non-nil, zero value otherwise.

### GetAllowDeletionsOk

`func (o *ProtectedBranch) GetAllowDeletionsOk() (*ProtectedBranchRequiredLinearHistory, bool)`

GetAllowDeletionsOk returns a tuple with the AllowDeletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDeletions

`func (o *ProtectedBranch) SetAllowDeletions(v ProtectedBranchRequiredLinearHistory)`

SetAllowDeletions sets AllowDeletions field to given value.

### HasAllowDeletions

`func (o *ProtectedBranch) HasAllowDeletions() bool`

HasAllowDeletions returns a boolean if a field has been set.

### GetRestrictions

`func (o *ProtectedBranch) GetRestrictions() BranchRestrictionPolicy`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *ProtectedBranch) GetRestrictionsOk() (*BranchRestrictionPolicy, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *ProtectedBranch) SetRestrictions(v BranchRestrictionPolicy)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *ProtectedBranch) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### GetRequiredConversationResolution

`func (o *ProtectedBranch) GetRequiredConversationResolution() ProtectedBranchRequiredConversationResolution`

GetRequiredConversationResolution returns the RequiredConversationResolution field if non-nil, zero value otherwise.

### GetRequiredConversationResolutionOk

`func (o *ProtectedBranch) GetRequiredConversationResolutionOk() (*ProtectedBranchRequiredConversationResolution, bool)`

GetRequiredConversationResolutionOk returns a tuple with the RequiredConversationResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredConversationResolution

`func (o *ProtectedBranch) SetRequiredConversationResolution(v ProtectedBranchRequiredConversationResolution)`

SetRequiredConversationResolution sets RequiredConversationResolution field to given value.

### HasRequiredConversationResolution

`func (o *ProtectedBranch) HasRequiredConversationResolution() bool`

HasRequiredConversationResolution returns a boolean if a field has been set.

### GetBlockCreations

`func (o *ProtectedBranch) GetBlockCreations() ProtectedBranchRequiredLinearHistory`

GetBlockCreations returns the BlockCreations field if non-nil, zero value otherwise.

### GetBlockCreationsOk

`func (o *ProtectedBranch) GetBlockCreationsOk() (*ProtectedBranchRequiredLinearHistory, bool)`

GetBlockCreationsOk returns a tuple with the BlockCreations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockCreations

`func (o *ProtectedBranch) SetBlockCreations(v ProtectedBranchRequiredLinearHistory)`

SetBlockCreations sets BlockCreations field to given value.

### HasBlockCreations

`func (o *ProtectedBranch) HasBlockCreations() bool`

HasBlockCreations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


