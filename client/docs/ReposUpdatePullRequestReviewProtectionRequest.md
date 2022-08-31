# ReposUpdatePullRequestReviewProtectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DismissalRestrictions** | Pointer to [**ReposUpdateBranchProtectionRequestRequiredPullRequestReviewsDismissalRestrictions**](ReposUpdateBranchProtectionRequestRequiredPullRequestReviewsDismissalRestrictions.md) |  | [optional] 
**DismissStaleReviews** | Pointer to **bool** | Set to &#x60;true&#x60; if you want to automatically dismiss approving reviews when someone pushes a new commit. | [optional] 
**RequireCodeOwnerReviews** | Pointer to **bool** | Blocks merging pull requests until [code owners](https://docs.github.com/articles/about-code-owners/) have reviewed. | [optional] 
**RequiredApprovingReviewCount** | Pointer to **int32** | Specifies the number of reviewers required to approve pull requests. Use a number between 1 and 6 or 0 to not require reviewers. | [optional] 
**BypassPullRequestAllowances** | Pointer to [**ReposUpdateBranchProtectionRequestRequiredPullRequestReviewsBypassPullRequestAllowances**](ReposUpdateBranchProtectionRequestRequiredPullRequestReviewsBypassPullRequestAllowances.md) |  | [optional] 

## Methods

### NewReposUpdatePullRequestReviewProtectionRequest

`func NewReposUpdatePullRequestReviewProtectionRequest() *ReposUpdatePullRequestReviewProtectionRequest`

NewReposUpdatePullRequestReviewProtectionRequest instantiates a new ReposUpdatePullRequestReviewProtectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposUpdatePullRequestReviewProtectionRequestWithDefaults

`func NewReposUpdatePullRequestReviewProtectionRequestWithDefaults() *ReposUpdatePullRequestReviewProtectionRequest`

NewReposUpdatePullRequestReviewProtectionRequestWithDefaults instantiates a new ReposUpdatePullRequestReviewProtectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDismissalRestrictions

`func (o *ReposUpdatePullRequestReviewProtectionRequest) GetDismissalRestrictions() ReposUpdateBranchProtectionRequestRequiredPullRequestReviewsDismissalRestrictions`

GetDismissalRestrictions returns the DismissalRestrictions field if non-nil, zero value otherwise.

### GetDismissalRestrictionsOk

`func (o *ReposUpdatePullRequestReviewProtectionRequest) GetDismissalRestrictionsOk() (*ReposUpdateBranchProtectionRequestRequiredPullRequestReviewsDismissalRestrictions, bool)`

GetDismissalRestrictionsOk returns a tuple with the DismissalRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissalRestrictions

`func (o *ReposUpdatePullRequestReviewProtectionRequest) SetDismissalRestrictions(v ReposUpdateBranchProtectionRequestRequiredPullRequestReviewsDismissalRestrictions)`

SetDismissalRestrictions sets DismissalRestrictions field to given value.

### HasDismissalRestrictions

`func (o *ReposUpdatePullRequestReviewProtectionRequest) HasDismissalRestrictions() bool`

HasDismissalRestrictions returns a boolean if a field has been set.

### GetDismissStaleReviews

`func (o *ReposUpdatePullRequestReviewProtectionRequest) GetDismissStaleReviews() bool`

GetDismissStaleReviews returns the DismissStaleReviews field if non-nil, zero value otherwise.

### GetDismissStaleReviewsOk

`func (o *ReposUpdatePullRequestReviewProtectionRequest) GetDismissStaleReviewsOk() (*bool, bool)`

GetDismissStaleReviewsOk returns a tuple with the DismissStaleReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissStaleReviews

`func (o *ReposUpdatePullRequestReviewProtectionRequest) SetDismissStaleReviews(v bool)`

SetDismissStaleReviews sets DismissStaleReviews field to given value.

### HasDismissStaleReviews

`func (o *ReposUpdatePullRequestReviewProtectionRequest) HasDismissStaleReviews() bool`

HasDismissStaleReviews returns a boolean if a field has been set.

### GetRequireCodeOwnerReviews

`func (o *ReposUpdatePullRequestReviewProtectionRequest) GetRequireCodeOwnerReviews() bool`

GetRequireCodeOwnerReviews returns the RequireCodeOwnerReviews field if non-nil, zero value otherwise.

### GetRequireCodeOwnerReviewsOk

`func (o *ReposUpdatePullRequestReviewProtectionRequest) GetRequireCodeOwnerReviewsOk() (*bool, bool)`

GetRequireCodeOwnerReviewsOk returns a tuple with the RequireCodeOwnerReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireCodeOwnerReviews

`func (o *ReposUpdatePullRequestReviewProtectionRequest) SetRequireCodeOwnerReviews(v bool)`

SetRequireCodeOwnerReviews sets RequireCodeOwnerReviews field to given value.

### HasRequireCodeOwnerReviews

`func (o *ReposUpdatePullRequestReviewProtectionRequest) HasRequireCodeOwnerReviews() bool`

HasRequireCodeOwnerReviews returns a boolean if a field has been set.

### GetRequiredApprovingReviewCount

`func (o *ReposUpdatePullRequestReviewProtectionRequest) GetRequiredApprovingReviewCount() int32`

GetRequiredApprovingReviewCount returns the RequiredApprovingReviewCount field if non-nil, zero value otherwise.

### GetRequiredApprovingReviewCountOk

`func (o *ReposUpdatePullRequestReviewProtectionRequest) GetRequiredApprovingReviewCountOk() (*int32, bool)`

GetRequiredApprovingReviewCountOk returns a tuple with the RequiredApprovingReviewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredApprovingReviewCount

`func (o *ReposUpdatePullRequestReviewProtectionRequest) SetRequiredApprovingReviewCount(v int32)`

SetRequiredApprovingReviewCount sets RequiredApprovingReviewCount field to given value.

### HasRequiredApprovingReviewCount

`func (o *ReposUpdatePullRequestReviewProtectionRequest) HasRequiredApprovingReviewCount() bool`

HasRequiredApprovingReviewCount returns a boolean if a field has been set.

### GetBypassPullRequestAllowances

`func (o *ReposUpdatePullRequestReviewProtectionRequest) GetBypassPullRequestAllowances() ReposUpdateBranchProtectionRequestRequiredPullRequestReviewsBypassPullRequestAllowances`

GetBypassPullRequestAllowances returns the BypassPullRequestAllowances field if non-nil, zero value otherwise.

### GetBypassPullRequestAllowancesOk

`func (o *ReposUpdatePullRequestReviewProtectionRequest) GetBypassPullRequestAllowancesOk() (*ReposUpdateBranchProtectionRequestRequiredPullRequestReviewsBypassPullRequestAllowances, bool)`

GetBypassPullRequestAllowancesOk returns a tuple with the BypassPullRequestAllowances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassPullRequestAllowances

`func (o *ReposUpdatePullRequestReviewProtectionRequest) SetBypassPullRequestAllowances(v ReposUpdateBranchProtectionRequestRequiredPullRequestReviewsBypassPullRequestAllowances)`

SetBypassPullRequestAllowances sets BypassPullRequestAllowances field to given value.

### HasBypassPullRequestAllowances

`func (o *ReposUpdatePullRequestReviewProtectionRequest) HasBypassPullRequestAllowances() bool`

HasBypassPullRequestAllowances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


