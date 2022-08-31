# ProtectedBranchRequiredPullRequestReviews

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**DismissStaleReviews** | Pointer to **bool** |  | [optional] 
**RequireCodeOwnerReviews** | Pointer to **bool** |  | [optional] 
**RequiredApprovingReviewCount** | Pointer to **int32** |  | [optional] 
**DismissalRestrictions** | Pointer to [**ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions**](ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions.md) |  | [optional] 
**BypassPullRequestAllowances** | Pointer to [**ProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowances**](ProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowances.md) |  | [optional] 

## Methods

### NewProtectedBranchRequiredPullRequestReviews

`func NewProtectedBranchRequiredPullRequestReviews(url string, ) *ProtectedBranchRequiredPullRequestReviews`

NewProtectedBranchRequiredPullRequestReviews instantiates a new ProtectedBranchRequiredPullRequestReviews object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtectedBranchRequiredPullRequestReviewsWithDefaults

`func NewProtectedBranchRequiredPullRequestReviewsWithDefaults() *ProtectedBranchRequiredPullRequestReviews`

NewProtectedBranchRequiredPullRequestReviewsWithDefaults instantiates a new ProtectedBranchRequiredPullRequestReviews object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ProtectedBranchRequiredPullRequestReviews) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ProtectedBranchRequiredPullRequestReviews) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ProtectedBranchRequiredPullRequestReviews) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDismissStaleReviews

`func (o *ProtectedBranchRequiredPullRequestReviews) GetDismissStaleReviews() bool`

GetDismissStaleReviews returns the DismissStaleReviews field if non-nil, zero value otherwise.

### GetDismissStaleReviewsOk

`func (o *ProtectedBranchRequiredPullRequestReviews) GetDismissStaleReviewsOk() (*bool, bool)`

GetDismissStaleReviewsOk returns a tuple with the DismissStaleReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissStaleReviews

`func (o *ProtectedBranchRequiredPullRequestReviews) SetDismissStaleReviews(v bool)`

SetDismissStaleReviews sets DismissStaleReviews field to given value.

### HasDismissStaleReviews

`func (o *ProtectedBranchRequiredPullRequestReviews) HasDismissStaleReviews() bool`

HasDismissStaleReviews returns a boolean if a field has been set.

### GetRequireCodeOwnerReviews

`func (o *ProtectedBranchRequiredPullRequestReviews) GetRequireCodeOwnerReviews() bool`

GetRequireCodeOwnerReviews returns the RequireCodeOwnerReviews field if non-nil, zero value otherwise.

### GetRequireCodeOwnerReviewsOk

`func (o *ProtectedBranchRequiredPullRequestReviews) GetRequireCodeOwnerReviewsOk() (*bool, bool)`

GetRequireCodeOwnerReviewsOk returns a tuple with the RequireCodeOwnerReviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireCodeOwnerReviews

`func (o *ProtectedBranchRequiredPullRequestReviews) SetRequireCodeOwnerReviews(v bool)`

SetRequireCodeOwnerReviews sets RequireCodeOwnerReviews field to given value.

### HasRequireCodeOwnerReviews

`func (o *ProtectedBranchRequiredPullRequestReviews) HasRequireCodeOwnerReviews() bool`

HasRequireCodeOwnerReviews returns a boolean if a field has been set.

### GetRequiredApprovingReviewCount

`func (o *ProtectedBranchRequiredPullRequestReviews) GetRequiredApprovingReviewCount() int32`

GetRequiredApprovingReviewCount returns the RequiredApprovingReviewCount field if non-nil, zero value otherwise.

### GetRequiredApprovingReviewCountOk

`func (o *ProtectedBranchRequiredPullRequestReviews) GetRequiredApprovingReviewCountOk() (*int32, bool)`

GetRequiredApprovingReviewCountOk returns a tuple with the RequiredApprovingReviewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredApprovingReviewCount

`func (o *ProtectedBranchRequiredPullRequestReviews) SetRequiredApprovingReviewCount(v int32)`

SetRequiredApprovingReviewCount sets RequiredApprovingReviewCount field to given value.

### HasRequiredApprovingReviewCount

`func (o *ProtectedBranchRequiredPullRequestReviews) HasRequiredApprovingReviewCount() bool`

HasRequiredApprovingReviewCount returns a boolean if a field has been set.

### GetDismissalRestrictions

`func (o *ProtectedBranchRequiredPullRequestReviews) GetDismissalRestrictions() ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions`

GetDismissalRestrictions returns the DismissalRestrictions field if non-nil, zero value otherwise.

### GetDismissalRestrictionsOk

`func (o *ProtectedBranchRequiredPullRequestReviews) GetDismissalRestrictionsOk() (*ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions, bool)`

GetDismissalRestrictionsOk returns a tuple with the DismissalRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissalRestrictions

`func (o *ProtectedBranchRequiredPullRequestReviews) SetDismissalRestrictions(v ProtectedBranchRequiredPullRequestReviewsDismissalRestrictions)`

SetDismissalRestrictions sets DismissalRestrictions field to given value.

### HasDismissalRestrictions

`func (o *ProtectedBranchRequiredPullRequestReviews) HasDismissalRestrictions() bool`

HasDismissalRestrictions returns a boolean if a field has been set.

### GetBypassPullRequestAllowances

`func (o *ProtectedBranchRequiredPullRequestReviews) GetBypassPullRequestAllowances() ProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowances`

GetBypassPullRequestAllowances returns the BypassPullRequestAllowances field if non-nil, zero value otherwise.

### GetBypassPullRequestAllowancesOk

`func (o *ProtectedBranchRequiredPullRequestReviews) GetBypassPullRequestAllowancesOk() (*ProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowances, bool)`

GetBypassPullRequestAllowancesOk returns a tuple with the BypassPullRequestAllowances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassPullRequestAllowances

`func (o *ProtectedBranchRequiredPullRequestReviews) SetBypassPullRequestAllowances(v ProtectedBranchRequiredPullRequestReviewsBypassPullRequestAllowances)`

SetBypassPullRequestAllowances sets BypassPullRequestAllowances field to given value.

### HasBypassPullRequestAllowances

`func (o *ProtectedBranchRequiredPullRequestReviews) HasBypassPullRequestAllowances() bool`

HasBypassPullRequestAllowances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


