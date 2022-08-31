# AutoMerge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnabledBy** | [**SimpleUser**](SimpleUser.md) |  | 
**MergeMethod** | **string** | The merge method to use. | 
**CommitTitle** | **string** | Title for the merge commit message. | 
**CommitMessage** | **string** | Commit message for the merge commit. | 

## Methods

### NewAutoMerge

`func NewAutoMerge(enabledBy SimpleUser, mergeMethod string, commitTitle string, commitMessage string, ) *AutoMerge`

NewAutoMerge instantiates a new AutoMerge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoMergeWithDefaults

`func NewAutoMergeWithDefaults() *AutoMerge`

NewAutoMergeWithDefaults instantiates a new AutoMerge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabledBy

`func (o *AutoMerge) GetEnabledBy() SimpleUser`

GetEnabledBy returns the EnabledBy field if non-nil, zero value otherwise.

### GetEnabledByOk

`func (o *AutoMerge) GetEnabledByOk() (*SimpleUser, bool)`

GetEnabledByOk returns a tuple with the EnabledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledBy

`func (o *AutoMerge) SetEnabledBy(v SimpleUser)`

SetEnabledBy sets EnabledBy field to given value.


### GetMergeMethod

`func (o *AutoMerge) GetMergeMethod() string`

GetMergeMethod returns the MergeMethod field if non-nil, zero value otherwise.

### GetMergeMethodOk

`func (o *AutoMerge) GetMergeMethodOk() (*string, bool)`

GetMergeMethodOk returns a tuple with the MergeMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeMethod

`func (o *AutoMerge) SetMergeMethod(v string)`

SetMergeMethod sets MergeMethod field to given value.


### GetCommitTitle

`func (o *AutoMerge) GetCommitTitle() string`

GetCommitTitle returns the CommitTitle field if non-nil, zero value otherwise.

### GetCommitTitleOk

`func (o *AutoMerge) GetCommitTitleOk() (*string, bool)`

GetCommitTitleOk returns a tuple with the CommitTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitTitle

`func (o *AutoMerge) SetCommitTitle(v string)`

SetCommitTitle sets CommitTitle field to given value.


### GetCommitMessage

`func (o *AutoMerge) GetCommitMessage() string`

GetCommitMessage returns the CommitMessage field if non-nil, zero value otherwise.

### GetCommitMessageOk

`func (o *AutoMerge) GetCommitMessageOk() (*string, bool)`

GetCommitMessageOk returns a tuple with the CommitMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitMessage

`func (o *AutoMerge) SetCommitMessage(v string)`

SetCommitMessage sets CommitMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


