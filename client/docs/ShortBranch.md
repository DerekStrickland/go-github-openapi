# ShortBranch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Commit** | [**ShortBranchCommit**](ShortBranchCommit.md) |  | 
**Protected** | **bool** |  | 
**Protection** | Pointer to [**BranchProtection**](BranchProtection.md) |  | [optional] 
**ProtectionUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewShortBranch

`func NewShortBranch(name string, commit ShortBranchCommit, protected bool, ) *ShortBranch`

NewShortBranch instantiates a new ShortBranch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShortBranchWithDefaults

`func NewShortBranchWithDefaults() *ShortBranch`

NewShortBranchWithDefaults instantiates a new ShortBranch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ShortBranch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ShortBranch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ShortBranch) SetName(v string)`

SetName sets Name field to given value.


### GetCommit

`func (o *ShortBranch) GetCommit() ShortBranchCommit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *ShortBranch) GetCommitOk() (*ShortBranchCommit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *ShortBranch) SetCommit(v ShortBranchCommit)`

SetCommit sets Commit field to given value.


### GetProtected

`func (o *ShortBranch) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *ShortBranch) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *ShortBranch) SetProtected(v bool)`

SetProtected sets Protected field to given value.


### GetProtection

`func (o *ShortBranch) GetProtection() BranchProtection`

GetProtection returns the Protection field if non-nil, zero value otherwise.

### GetProtectionOk

`func (o *ShortBranch) GetProtectionOk() (*BranchProtection, bool)`

GetProtectionOk returns a tuple with the Protection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtection

`func (o *ShortBranch) SetProtection(v BranchProtection)`

SetProtection sets Protection field to given value.

### HasProtection

`func (o *ShortBranch) HasProtection() bool`

HasProtection returns a boolean if a field has been set.

### GetProtectionUrl

`func (o *ShortBranch) GetProtectionUrl() string`

GetProtectionUrl returns the ProtectionUrl field if non-nil, zero value otherwise.

### GetProtectionUrlOk

`func (o *ShortBranch) GetProtectionUrlOk() (*string, bool)`

GetProtectionUrlOk returns a tuple with the ProtectionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionUrl

`func (o *ShortBranch) SetProtectionUrl(v string)`

SetProtectionUrl sets ProtectionUrl field to given value.

### HasProtectionUrl

`func (o *ShortBranch) HasProtectionUrl() bool`

HasProtectionUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


