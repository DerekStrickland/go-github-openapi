# ReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | **string** | The name of the required check | 
**AppId** | Pointer to **int32** | The ID of the GitHub App that must provide this check. Omit this field to automatically select the GitHub App that has recently provided this check, or any app if it was not set by a GitHub App. Pass -1 to explicitly allow any app to set the status. | [optional] 

## Methods

### NewReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner

`func NewReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner(context string, ) *ReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner`

NewReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner instantiates a new ReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInnerWithDefaults

`func NewReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInnerWithDefaults() *ReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner`

NewReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInnerWithDefaults instantiates a new ReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *ReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner) SetContext(v string)`

SetContext sets Context field to given value.


### GetAppId

`func (o *ReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner) SetAppId(v int32)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ReposUpdateBranchProtectionRequestRequiredStatusChecksChecksInner) HasAppId() bool`

HasAppId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


