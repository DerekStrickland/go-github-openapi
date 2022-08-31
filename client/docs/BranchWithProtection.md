# BranchWithProtection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Commit** | [**Commit**](Commit.md) |  | 
**Links** | [**BranchWithProtectionLinks**](BranchWithProtectionLinks.md) |  | 
**Protected** | **bool** |  | 
**Protection** | [**BranchProtection**](BranchProtection.md) |  | 
**ProtectionUrl** | **string** |  | 
**Pattern** | Pointer to **string** |  | [optional] 
**RequiredApprovingReviewCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewBranchWithProtection

`func NewBranchWithProtection(name string, commit Commit, links BranchWithProtectionLinks, protected bool, protection BranchProtection, protectionUrl string, ) *BranchWithProtection`

NewBranchWithProtection instantiates a new BranchWithProtection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchWithProtectionWithDefaults

`func NewBranchWithProtectionWithDefaults() *BranchWithProtection`

NewBranchWithProtectionWithDefaults instantiates a new BranchWithProtection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BranchWithProtection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BranchWithProtection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BranchWithProtection) SetName(v string)`

SetName sets Name field to given value.


### GetCommit

`func (o *BranchWithProtection) GetCommit() Commit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *BranchWithProtection) GetCommitOk() (*Commit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *BranchWithProtection) SetCommit(v Commit)`

SetCommit sets Commit field to given value.


### GetLinks

`func (o *BranchWithProtection) GetLinks() BranchWithProtectionLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BranchWithProtection) GetLinksOk() (*BranchWithProtectionLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BranchWithProtection) SetLinks(v BranchWithProtectionLinks)`

SetLinks sets Links field to given value.


### GetProtected

`func (o *BranchWithProtection) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *BranchWithProtection) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *BranchWithProtection) SetProtected(v bool)`

SetProtected sets Protected field to given value.


### GetProtection

`func (o *BranchWithProtection) GetProtection() BranchProtection`

GetProtection returns the Protection field if non-nil, zero value otherwise.

### GetProtectionOk

`func (o *BranchWithProtection) GetProtectionOk() (*BranchProtection, bool)`

GetProtectionOk returns a tuple with the Protection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtection

`func (o *BranchWithProtection) SetProtection(v BranchProtection)`

SetProtection sets Protection field to given value.


### GetProtectionUrl

`func (o *BranchWithProtection) GetProtectionUrl() string`

GetProtectionUrl returns the ProtectionUrl field if non-nil, zero value otherwise.

### GetProtectionUrlOk

`func (o *BranchWithProtection) GetProtectionUrlOk() (*string, bool)`

GetProtectionUrlOk returns a tuple with the ProtectionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionUrl

`func (o *BranchWithProtection) SetProtectionUrl(v string)`

SetProtectionUrl sets ProtectionUrl field to given value.


### GetPattern

`func (o *BranchWithProtection) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *BranchWithProtection) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *BranchWithProtection) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *BranchWithProtection) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetRequiredApprovingReviewCount

`func (o *BranchWithProtection) GetRequiredApprovingReviewCount() int32`

GetRequiredApprovingReviewCount returns the RequiredApprovingReviewCount field if non-nil, zero value otherwise.

### GetRequiredApprovingReviewCountOk

`func (o *BranchWithProtection) GetRequiredApprovingReviewCountOk() (*int32, bool)`

GetRequiredApprovingReviewCountOk returns a tuple with the RequiredApprovingReviewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredApprovingReviewCount

`func (o *BranchWithProtection) SetRequiredApprovingReviewCount(v int32)`

SetRequiredApprovingReviewCount sets RequiredApprovingReviewCount field to given value.

### HasRequiredApprovingReviewCount

`func (o *BranchWithProtection) HasRequiredApprovingReviewCount() bool`

HasRequiredApprovingReviewCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


