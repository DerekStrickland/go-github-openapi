# ReposUpdateReleaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagName** | Pointer to **string** | The name of the tag. | [optional] 
**TargetCommitish** | Pointer to **string** | Specifies the commitish value that determines where the Git tag is created from. Can be any branch or commit SHA. Unused if the Git tag already exists. Default: the repository&#39;s default branch (usually &#x60;master&#x60;). | [optional] 
**Name** | Pointer to **string** | The name of the release. | [optional] 
**Body** | Pointer to **string** | Text describing the contents of the tag. | [optional] 
**Draft** | Pointer to **bool** | &#x60;true&#x60; makes the release a draft, and &#x60;false&#x60; publishes the release. | [optional] 
**Prerelease** | Pointer to **bool** | &#x60;true&#x60; to identify the release as a prerelease, &#x60;false&#x60; to identify the release as a full release. | [optional] 
**DiscussionCategoryName** | Pointer to **string** | If specified, a discussion of the specified category is created and linked to the release. The value must be a category that already exists in the repository. If there is already a discussion linked to the release, this parameter is ignored. For more information, see \&quot;[Managing categories for discussions in your repository](https://docs.github.com/discussions/managing-discussions-for-your-community/managing-categories-for-discussions-in-your-repository).\&quot; | [optional] 

## Methods

### NewReposUpdateReleaseRequest

`func NewReposUpdateReleaseRequest() *ReposUpdateReleaseRequest`

NewReposUpdateReleaseRequest instantiates a new ReposUpdateReleaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposUpdateReleaseRequestWithDefaults

`func NewReposUpdateReleaseRequestWithDefaults() *ReposUpdateReleaseRequest`

NewReposUpdateReleaseRequestWithDefaults instantiates a new ReposUpdateReleaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagName

`func (o *ReposUpdateReleaseRequest) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *ReposUpdateReleaseRequest) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *ReposUpdateReleaseRequest) SetTagName(v string)`

SetTagName sets TagName field to given value.

### HasTagName

`func (o *ReposUpdateReleaseRequest) HasTagName() bool`

HasTagName returns a boolean if a field has been set.

### GetTargetCommitish

`func (o *ReposUpdateReleaseRequest) GetTargetCommitish() string`

GetTargetCommitish returns the TargetCommitish field if non-nil, zero value otherwise.

### GetTargetCommitishOk

`func (o *ReposUpdateReleaseRequest) GetTargetCommitishOk() (*string, bool)`

GetTargetCommitishOk returns a tuple with the TargetCommitish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCommitish

`func (o *ReposUpdateReleaseRequest) SetTargetCommitish(v string)`

SetTargetCommitish sets TargetCommitish field to given value.

### HasTargetCommitish

`func (o *ReposUpdateReleaseRequest) HasTargetCommitish() bool`

HasTargetCommitish returns a boolean if a field has been set.

### GetName

`func (o *ReposUpdateReleaseRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReposUpdateReleaseRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReposUpdateReleaseRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReposUpdateReleaseRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBody

`func (o *ReposUpdateReleaseRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ReposUpdateReleaseRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ReposUpdateReleaseRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ReposUpdateReleaseRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetDraft

`func (o *ReposUpdateReleaseRequest) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *ReposUpdateReleaseRequest) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *ReposUpdateReleaseRequest) SetDraft(v bool)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *ReposUpdateReleaseRequest) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetPrerelease

`func (o *ReposUpdateReleaseRequest) GetPrerelease() bool`

GetPrerelease returns the Prerelease field if non-nil, zero value otherwise.

### GetPrereleaseOk

`func (o *ReposUpdateReleaseRequest) GetPrereleaseOk() (*bool, bool)`

GetPrereleaseOk returns a tuple with the Prerelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerelease

`func (o *ReposUpdateReleaseRequest) SetPrerelease(v bool)`

SetPrerelease sets Prerelease field to given value.

### HasPrerelease

`func (o *ReposUpdateReleaseRequest) HasPrerelease() bool`

HasPrerelease returns a boolean if a field has been set.

### GetDiscussionCategoryName

`func (o *ReposUpdateReleaseRequest) GetDiscussionCategoryName() string`

GetDiscussionCategoryName returns the DiscussionCategoryName field if non-nil, zero value otherwise.

### GetDiscussionCategoryNameOk

`func (o *ReposUpdateReleaseRequest) GetDiscussionCategoryNameOk() (*string, bool)`

GetDiscussionCategoryNameOk returns a tuple with the DiscussionCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscussionCategoryName

`func (o *ReposUpdateReleaseRequest) SetDiscussionCategoryName(v string)`

SetDiscussionCategoryName sets DiscussionCategoryName field to given value.

### HasDiscussionCategoryName

`func (o *ReposUpdateReleaseRequest) HasDiscussionCategoryName() bool`

HasDiscussionCategoryName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


