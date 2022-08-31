# BranchShort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Commit** | [**BranchShortCommit**](BranchShortCommit.md) |  | 
**Protected** | **bool** |  | 

## Methods

### NewBranchShort

`func NewBranchShort(name string, commit BranchShortCommit, protected bool, ) *BranchShort`

NewBranchShort instantiates a new BranchShort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBranchShortWithDefaults

`func NewBranchShortWithDefaults() *BranchShort`

NewBranchShortWithDefaults instantiates a new BranchShort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BranchShort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BranchShort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BranchShort) SetName(v string)`

SetName sets Name field to given value.


### GetCommit

`func (o *BranchShort) GetCommit() BranchShortCommit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *BranchShort) GetCommitOk() (*BranchShortCommit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *BranchShort) SetCommit(v BranchShortCommit)`

SetCommit sets Commit field to given value.


### GetProtected

`func (o *BranchShort) GetProtected() bool`

GetProtected returns the Protected field if non-nil, zero value otherwise.

### GetProtectedOk

`func (o *BranchShort) GetProtectedOk() (*bool, bool)`

GetProtectedOk returns a tuple with the Protected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtected

`func (o *BranchShort) SetProtected(v bool)`

SetProtected sets Protected field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


