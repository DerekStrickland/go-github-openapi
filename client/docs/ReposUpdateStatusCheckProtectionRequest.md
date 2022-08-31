# ReposUpdateStatusCheckProtectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Strict** | Pointer to **bool** | Require branches to be up to date before merging. | [optional] 
**Contexts** | Pointer to **[]string** | **Deprecated**: The list of status checks to require in order to merge into this branch. If any of these checks have recently been set by a particular GitHub App, they will be required to come from that app in future for the branch to merge. Use &#x60;checks&#x60; instead of &#x60;contexts&#x60; for more fine-grained control.  | [optional] 
**Checks** | Pointer to [**[]ReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner**](ReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner.md) | The list of status checks to require in order to merge into this branch. | [optional] 

## Methods

### NewReposUpdateStatusCheckProtectionRequest

`func NewReposUpdateStatusCheckProtectionRequest() *ReposUpdateStatusCheckProtectionRequest`

NewReposUpdateStatusCheckProtectionRequest instantiates a new ReposUpdateStatusCheckProtectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposUpdateStatusCheckProtectionRequestWithDefaults

`func NewReposUpdateStatusCheckProtectionRequestWithDefaults() *ReposUpdateStatusCheckProtectionRequest`

NewReposUpdateStatusCheckProtectionRequestWithDefaults instantiates a new ReposUpdateStatusCheckProtectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrict

`func (o *ReposUpdateStatusCheckProtectionRequest) GetStrict() bool`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *ReposUpdateStatusCheckProtectionRequest) GetStrictOk() (*bool, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *ReposUpdateStatusCheckProtectionRequest) SetStrict(v bool)`

SetStrict sets Strict field to given value.

### HasStrict

`func (o *ReposUpdateStatusCheckProtectionRequest) HasStrict() bool`

HasStrict returns a boolean if a field has been set.

### GetContexts

`func (o *ReposUpdateStatusCheckProtectionRequest) GetContexts() []string`

GetContexts returns the Contexts field if non-nil, zero value otherwise.

### GetContextsOk

`func (o *ReposUpdateStatusCheckProtectionRequest) GetContextsOk() (*[]string, bool)`

GetContextsOk returns a tuple with the Contexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContexts

`func (o *ReposUpdateStatusCheckProtectionRequest) SetContexts(v []string)`

SetContexts sets Contexts field to given value.

### HasContexts

`func (o *ReposUpdateStatusCheckProtectionRequest) HasContexts() bool`

HasContexts returns a boolean if a field has been set.

### GetChecks

`func (o *ReposUpdateStatusCheckProtectionRequest) GetChecks() []ReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *ReposUpdateStatusCheckProtectionRequest) GetChecksOk() (*[]ReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *ReposUpdateStatusCheckProtectionRequest) SetChecks(v []ReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *ReposUpdateStatusCheckProtectionRequest) HasChecks() bool`

HasChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


