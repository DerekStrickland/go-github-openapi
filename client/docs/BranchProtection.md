# BranchProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**RequiredStatusChecks** | Pointer to [**ProtectedBranchRequiredStatusCheck**](ProtectedBranchRequiredStatusCheck.md) |  | [optional] 
**EnforceAdmins** | Pointer to [**ProtectedBranchAdminEnforced**](ProtectedBranchAdminEnforced.md) |  | [optional] 
**RequiredPullRequestReviews** | Pointer to [**ProtectedBranchPullRequestReview**](ProtectedBranchPullRequestReview.md) |  | [optional] 
**Restrictions** | Pointer to [**BranchRestrictionPolicy**](BranchRestrictionPolicy.md) |  | [optional] 
**RequiredLinearHistory** | Pointer to [**BranchProtectionRequiredLinearHistory**](BranchProtectionRequiredLinearHistory.md) |  | [optional] 
**AllowForcePushes** | Pointer to [**BranchProtectionRequiredLinearHistory**](BranchProtectionRequiredLinearHistory.md) |  | [optional] 
**AllowDeletions** | Pointer to [**BranchProtectionRequiredLinearHistory**](BranchProtectionRequiredLinearHistory.md) |  | [optional] 
**BlockCreations** | Pointer to [**BranchProtectionRequiredLinearHistory**](BranchProtectionRequiredLinearHistory.md) |  | [optional] 
**RequiredConversationResolution** | Pointer to [**BranchProtectionRequiredLinearHistory**](BranchProtectionRequiredLinearHistory.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ProtectionUrl** | Pointer to **string** |  | [optional] 
**RequiredSignatures** | Pointer to [**BranchProtectionRequiredSignatures**](BranchProtectionRequiredSignatures.md) |  | [optional] 

## Methods

### NewBranchProtection

`func NewBranchProtection() *BranchProtection`

NewBranchProtection instantiates a new BranchProtection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchProtectionWithDefaults

`func NewBranchProtectionWithDefaults() *BranchProtection`

NewBranchProtectionWithDefaults instantiates a new BranchProtection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *BranchProtection) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BranchProtection) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BranchProtection) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *BranchProtection) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEnabled

`func (o *BranchProtection) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BranchProtection) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BranchProtection) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BranchProtection) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRequiredStatusChecks

`func (o *BranchProtection) GetRequiredStatusChecks() ProtectedBranchRequiredStatusCheck`

GetRequiredStatusChecks returns the RequiredStatusChecks field if non-nil, zero value otherwise.

### GetRequiredStatusChecksOk

`func (o *BranchProtection) GetRequiredStatusChecksOk() (*ProtectedBranchRequiredStatusCheck, bool)`

GetRequiredStatusChecksOk returns a tuple with the RequiredStatusChecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredStatusChecks

`func (o *BranchProtection) SetRequiredStatusChecks(v ProtectedBranchRequiredStatusCheck)`

SetRequiredStatusChecks sets RequiredStatusChecks field to given value.

### HasRequiredStatusChecks

`func (o *BranchProtection) HasRequiredStatusChecks() bool`

HasRequiredStatusChecks returns a boolean if a field has been set.

### GetEnforceAdmins

`func (o *BranchProtection) GetEnforceAdmins() ProtectedBranchAdminEnforced`

GetEnforceAdmins returns the EnforceAdmins field if non-nil, zero value otherwise.

### GetEnforceAdminsOk

`func (o *BranchProtection) GetEnforceAdminsOk() (*ProtectedBranchAdminEnforced, bool)`

GetEnforceAdminsOk returns a tuple with the EnforceAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceAdmins

`func (o *BranchProtection) SetEnforceAdmins(v ProtectedBranchAdminEnforced)`

SetEnforceAdmins sets EnforceAdmins field to given value.

### HasEnforceAdmins

`func (o *BranchProtection) HasEnforceAdmins() bool`

HasEnforceAdmins returns a boolean if a field has been set.

### GetRequiredPullRequestReviews

`func (o *BranchProtection) GetRequiredPullRequestReviews() ProtectedBranchPullRequestReview`

GetRequiredPullRequestReviews returns the RequiredPullRequestReviews field if non-nil, zero value otherwise.

### GetRequiredPullRequestReviewsOk

`func (o *BranchProtection) GetRequiredPullRequestReviewsOk() (*ProtectedBranchPullRequestReview, bool)`

GetRequiredPullRequestReviewsOk returns a tuple with the RequiredPullRequestReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredPullRequestReviews

`func (o *BranchProtection) SetRequiredPullRequestReviews(v ProtectedBranchPullRequestReview)`

SetRequiredPullRequestReviews sets RequiredPullRequestReviews field to given value.

### HasRequiredPullRequestReviews

`func (o *BranchProtection) HasRequiredPullRequestReviews() bool`

HasRequiredPullRequestReviews returns a boolean if a field has been set.

### GetRestrictions

`func (o *BranchProtection) GetRestrictions() BranchRestrictionPolicy`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *BranchProtection) GetRestrictionsOk() (*BranchRestrictionPolicy, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *BranchProtection) SetRestrictions(v BranchRestrictionPolicy)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *BranchProtection) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### GetRequiredLinearHistory

`func (o *BranchProtection) GetRequiredLinearHistory() BranchProtectionRequiredLinearHistory`

GetRequiredLinearHistory returns the RequiredLinearHistory field if non-nil, zero value otherwise.

### GetRequiredLinearHistoryOk

`func (o *BranchProtection) GetRequiredLinearHistoryOk() (*BranchProtectionRequiredLinearHistory, bool)`

GetRequiredLinearHistoryOk returns a tuple with the RequiredLinearHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredLinearHistory

`func (o *BranchProtection) SetRequiredLinearHistory(v BranchProtectionRequiredLinearHistory)`

SetRequiredLinearHistory sets RequiredLinearHistory field to given value.

### HasRequiredLinearHistory

`func (o *BranchProtection) HasRequiredLinearHistory() bool`

HasRequiredLinearHistory returns a boolean if a field has been set.

### GetAllowForcePushes

`func (o *BranchProtection) GetAllowForcePushes() BranchProtectionRequiredLinearHistory`

GetAllowForcePushes returns the AllowForcePushes field if non-nil, zero value otherwise.

### GetAllowForcePushesOk

`func (o *BranchProtection) GetAllowForcePushesOk() (*BranchProtectionRequiredLinearHistory, bool)`

GetAllowForcePushesOk returns a tuple with the AllowForcePushes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowForcePushes

`func (o *BranchProtection) SetAllowForcePushes(v BranchProtectionRequiredLinearHistory)`

SetAllowForcePushes sets AllowForcePushes field to given value.

### HasAllowForcePushes

`func (o *BranchProtection) HasAllowForcePushes() bool`

HasAllowForcePushes returns a boolean if a field has been set.

### GetAllowDeletions

`func (o *BranchProtection) GetAllowDeletions() BranchProtectionRequiredLinearHistory`

GetAllowDeletions returns the AllowDeletions field if non-nil, zero value otherwise.

### GetAllowDeletionsOk

`func (o *BranchProtection) GetAllowDeletionsOk() (*BranchProtectionRequiredLinearHistory, bool)`

GetAllowDeletionsOk returns a tuple with the AllowDeletions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDeletions

`func (o *BranchProtection) SetAllowDeletions(v BranchProtectionRequiredLinearHistory)`

SetAllowDeletions sets AllowDeletions field to given value.

### HasAllowDeletions

`func (o *BranchProtection) HasAllowDeletions() bool`

HasAllowDeletions returns a boolean if a field has been set.

### GetBlockCreations

`func (o *BranchProtection) GetBlockCreations() BranchProtectionRequiredLinearHistory`

GetBlockCreations returns the BlockCreations field if non-nil, zero value otherwise.

### GetBlockCreationsOk

`func (o *BranchProtection) GetBlockCreationsOk() (*BranchProtectionRequiredLinearHistory, bool)`

GetBlockCreationsOk returns a tuple with the BlockCreations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockCreations

`func (o *BranchProtection) SetBlockCreations(v BranchProtectionRequiredLinearHistory)`

SetBlockCreations sets BlockCreations field to given value.

### HasBlockCreations

`func (o *BranchProtection) HasBlockCreations() bool`

HasBlockCreations returns a boolean if a field has been set.

### GetRequiredConversationResolution

`func (o *BranchProtection) GetRequiredConversationResolution() BranchProtectionRequiredLinearHistory`

GetRequiredConversationResolution returns the RequiredConversationResolution field if non-nil, zero value otherwise.

### GetRequiredConversationResolutionOk

`func (o *BranchProtection) GetRequiredConversationResolutionOk() (*BranchProtectionRequiredLinearHistory, bool)`

GetRequiredConversationResolutionOk returns a tuple with the RequiredConversationResolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredConversationResolution

`func (o *BranchProtection) SetRequiredConversationResolution(v BranchProtectionRequiredLinearHistory)`

SetRequiredConversationResolution sets RequiredConversationResolution field to given value.

### HasRequiredConversationResolution

`func (o *BranchProtection) HasRequiredConversationResolution() bool`

HasRequiredConversationResolution returns a boolean if a field has been set.

### GetName

`func (o *BranchProtection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BranchProtection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BranchProtection) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BranchProtection) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProtectionUrl

`func (o *BranchProtection) GetProtectionUrl() string`

GetProtectionUrl returns the ProtectionUrl field if non-nil, zero value otherwise.

### GetProtectionUrlOk

`func (o *BranchProtection) GetProtectionUrlOk() (*string, bool)`

GetProtectionUrlOk returns a tuple with the ProtectionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionUrl

`func (o *BranchProtection) SetProtectionUrl(v string)`

SetProtectionUrl sets ProtectionUrl field to given value.

### HasProtectionUrl

`func (o *BranchProtection) HasProtectionUrl() bool`

HasProtectionUrl returns a boolean if a field has been set.

### GetRequiredSignatures

`func (o *BranchProtection) GetRequiredSignatures() BranchProtectionRequiredSignatures`

GetRequiredSignatures returns the RequiredSignatures field if non-nil, zero value otherwise.

### GetRequiredSignaturesOk

`func (o *BranchProtection) GetRequiredSignaturesOk() (*BranchProtectionRequiredSignatures, bool)`

GetRequiredSignaturesOk returns a tuple with the RequiredSignatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredSignatures

`func (o *BranchProtection) SetRequiredSignatures(v BranchProtectionRequiredSignatures)`

SetRequiredSignatures sets RequiredSignatures field to given value.

### HasRequiredSignatures

`func (o *BranchProtection) HasRequiredSignatures() bool`

HasRequiredSignatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


