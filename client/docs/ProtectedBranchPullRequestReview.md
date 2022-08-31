# ProtectedBranchPullRequestReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] 
**DismissalRestrictions** | Pointer to [**ProtectedBranchPullRequestReviewDismissalRestrictions**](ProtectedBranchPullRequestReviewDismissalRestrictions.md) |  | [optional] 
**BypassPullRequestAllowances** | Pointer to [**ProtectedBranchPullRequestReviewBypassPullRequestAllowances**](ProtectedBranchPullRequestReviewBypassPullRequestAllowances.md) |  | [optional] 
**DismissStaleReviews** | **bool** |  | 
**RequireCodeOwnerReviews** | **bool** |  | 
**RequiredApprovingReviewCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewProtectedBranchPullRequestReview

`func NewProtectedBranchPullRequestReview(dismissStaleReviews bool, requireCodeOwnerReviews bool, ) *ProtectedBranchPullRequestReview`

NewProtectedBranchPullRequestReview instantiates a new ProtectedBranchPullRequestReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedBranchPullRequestReviewWithDefaults

`func NewProtectedBranchPullRequestReviewWithDefaults() *ProtectedBranchPullRequestReview`

NewProtectedBranchPullRequestReviewWithDefaults instantiates a new ProtectedBranchPullRequestReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ProtectedBranchPullRequestReview) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProtectedBranchPullRequestReview) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProtectedBranchPullRequestReview) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ProtectedBranchPullRequestReview) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDismissalRestrictions

`func (o *ProtectedBranchPullRequestReview) GetDismissalRestrictions() ProtectedBranchPullRequestReviewDismissalRestrictions`

GetDismissalRestrictions returns the DismissalRestrictions field if non-nil, zero value otherwise.

### GetDismissalRestrictionsOk

`func (o *ProtectedBranchPullRequestReview) GetDismissalRestrictionsOk() (*ProtectedBranchPullRequestReviewDismissalRestrictions, bool)`

GetDismissalRestrictionsOk returns a tuple with the DismissalRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissalRestrictions

`func (o *ProtectedBranchPullRequestReview) SetDismissalRestrictions(v ProtectedBranchPullRequestReviewDismissalRestrictions)`

SetDismissalRestrictions sets DismissalRestrictions field to given value.

### HasDismissalRestrictions

`func (o *ProtectedBranchPullRequestReview) HasDismissalRestrictions() bool`

HasDismissalRestrictions returns a boolean if a field has been set.

### GetBypassPullRequestAllowances

`func (o *ProtectedBranchPullRequestReview) GetBypassPullRequestAllowances() ProtectedBranchPullRequestReviewBypassPullRequestAllowances`

GetBypassPullRequestAllowances returns the BypassPullRequestAllowances field if non-nil, zero value otherwise.

### GetBypassPullRequestAllowancesOk

`func (o *ProtectedBranchPullRequestReview) GetBypassPullRequestAllowancesOk() (*ProtectedBranchPullRequestReviewBypassPullRequestAllowances, bool)`

GetBypassPullRequestAllowancesOk returns a tuple with the BypassPullRequestAllowances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassPullRequestAllowances

`func (o *ProtectedBranchPullRequestReview) SetBypassPullRequestAllowances(v ProtectedBranchPullRequestReviewBypassPullRequestAllowances)`

SetBypassPullRequestAllowances sets BypassPullRequestAllowances field to given value.

### HasBypassPullRequestAllowances

`func (o *ProtectedBranchPullRequestReview) HasBypassPullRequestAllowances() bool`

HasBypassPullRequestAllowances returns a boolean if a field has been set.

### GetDismissStaleReviews

`func (o *ProtectedBranchPullRequestReview) GetDismissStaleReviews() bool`

GetDismissStaleReviews returns the DismissStaleReviews field if non-nil, zero value otherwise.

### GetDismissStaleReviewsOk

`func (o *ProtectedBranchPullRequestReview) GetDismissStaleReviewsOk() (*bool, bool)`

GetDismissStaleReviewsOk returns a tuple with the DismissStaleReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissStaleReviews

`func (o *ProtectedBranchPullRequestReview) SetDismissStaleReviews(v bool)`

SetDismissStaleReviews sets DismissStaleReviews field to given value.


### GetRequireCodeOwnerReviews

`func (o *ProtectedBranchPullRequestReview) GetRequireCodeOwnerReviews() bool`

GetRequireCodeOwnerReviews returns the RequireCodeOwnerReviews field if non-nil, zero value otherwise.

### GetRequireCodeOwnerReviewsOk

`func (o *ProtectedBranchPullRequestReview) GetRequireCodeOwnerReviewsOk() (*bool, bool)`

GetRequireCodeOwnerReviewsOk returns a tuple with the RequireCodeOwnerReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireCodeOwnerReviews

`func (o *ProtectedBranchPullRequestReview) SetRequireCodeOwnerReviews(v bool)`

SetRequireCodeOwnerReviews sets RequireCodeOwnerReviews field to given value.


### GetRequiredApprovingReviewCount

`func (o *ProtectedBranchPullRequestReview) GetRequiredApprovingReviewCount() int32`

GetRequiredApprovingReviewCount returns the RequiredApprovingReviewCount field if non-nil, zero value otherwise.

### GetRequiredApprovingReviewCountOk

`func (o *ProtectedBranchPullRequestReview) GetRequiredApprovingReviewCountOk() (*int32, bool)`

GetRequiredApprovingReviewCountOk returns a tuple with the RequiredApprovingReviewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredApprovingReviewCount

`func (o *ProtectedBranchPullRequestReview) SetRequiredApprovingReviewCount(v int32)`

SetRequiredApprovingReviewCount sets RequiredApprovingReviewCount field to given value.

### HasRequiredApprovingReviewCount

`func (o *ProtectedBranchPullRequestReview) HasRequiredApprovingReviewCount() bool`

HasRequiredApprovingReviewCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


