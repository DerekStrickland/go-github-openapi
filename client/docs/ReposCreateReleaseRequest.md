# ReposCreateReleaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagName** | **string** | The name of the tag. | 
**TargetCommitish** | Pointer to **string** | Specifies the commitish value that determines where the Git tag is created from. Can be any branch or commit SHA. Unused if the Git tag already exists. Default: the repository&#39;s default branch (usually &#x60;master&#x60;). | [optional] 
**Name** | Pointer to **string** | The name of the release. | [optional] 
**Body** | Pointer to **string** | Text describing the contents of the tag. | [optional] 
**Draft** | Pointer to **bool** | &#x60;true&#x60; to create a draft (unpublished) release, &#x60;false&#x60; to create a published one. | [optional] [default to false]
**Prerelease** | Pointer to **bool** | &#x60;true&#x60; to identify the release as a prerelease. &#x60;false&#x60; to identify the release as a full release. | [optional] [default to false]
**DiscussionCategoryName** | Pointer to **string** | If specified, a discussion of the specified category is created and linked to the release. The value must be a category that already exists in the repository. For more information, see \&quot;[Managing categories for discussions in your repository](https://docs.github.com/discussions/managing-discussions-for-your-community/managing-categories-for-discussions-in-your-repository).\&quot; | [optional] 
**GenerateReleaseNotes** | Pointer to **bool** | Whether to automatically generate the name and body for this release. If &#x60;name&#x60; is specified, the specified name will be used; otherwise, a name will be automatically generated. If &#x60;body&#x60; is specified, the body will be pre-pended to the automatically generated notes. | [optional] [default to false]

## Methods

### NewReposCreateReleaseRequest

`func NewReposCreateReleaseRequest(tagName string, ) *ReposCreateReleaseRequest`

NewReposCreateReleaseRequest instantiates a new ReposCreateReleaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposCreateReleaseRequestWithDefaults

`func NewReposCreateReleaseRequestWithDefaults() *ReposCreateReleaseRequest`

NewReposCreateReleaseRequestWithDefaults instantiates a new ReposCreateReleaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagName

`func (o *ReposCreateReleaseRequest) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *ReposCreateReleaseRequest) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *ReposCreateReleaseRequest) SetTagName(v string)`

SetTagName sets TagName field to given value.


### GetTargetCommitish

`func (o *ReposCreateReleaseRequest) GetTargetCommitish() string`

GetTargetCommitish returns the TargetCommitish field if non-nil, zero value otherwise.

### GetTargetCommitishOk

`func (o *ReposCreateReleaseRequest) GetTargetCommitishOk() (*string, bool)`

GetTargetCommitishOk returns a tuple with the TargetCommitish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCommitish

`func (o *ReposCreateReleaseRequest) SetTargetCommitish(v string)`

SetTargetCommitish sets TargetCommitish field to given value.

### HasTargetCommitish

`func (o *ReposCreateReleaseRequest) HasTargetCommitish() bool`

HasTargetCommitish returns a boolean if a field has been set.

### GetName

`func (o *ReposCreateReleaseRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReposCreateReleaseRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReposCreateReleaseRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReposCreateReleaseRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBody

`func (o *ReposCreateReleaseRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ReposCreateReleaseRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ReposCreateReleaseRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ReposCreateReleaseRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetDraft

`func (o *ReposCreateReleaseRequest) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *ReposCreateReleaseRequest) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *ReposCreateReleaseRequest) SetDraft(v bool)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *ReposCreateReleaseRequest) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetPrerelease

`func (o *ReposCreateReleaseRequest) GetPrerelease() bool`

GetPrerelease returns the Prerelease field if non-nil, zero value otherwise.

### GetPrereleaseOk

`func (o *ReposCreateReleaseRequest) GetPrereleaseOk() (*bool, bool)`

GetPrereleaseOk returns a tuple with the Prerelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerelease

`func (o *ReposCreateReleaseRequest) SetPrerelease(v bool)`

SetPrerelease sets Prerelease field to given value.

### HasPrerelease

`func (o *ReposCreateReleaseRequest) HasPrerelease() bool`

HasPrerelease returns a boolean if a field has been set.

### GetDiscussionCategoryName

`func (o *ReposCreateReleaseRequest) GetDiscussionCategoryName() string`

GetDiscussionCategoryName returns the DiscussionCategoryName field if non-nil, zero value otherwise.

### GetDiscussionCategoryNameOk

`func (o *ReposCreateReleaseRequest) GetDiscussionCategoryNameOk() (*string, bool)`

GetDiscussionCategoryNameOk returns a tuple with the DiscussionCategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscussionCategoryName

`func (o *ReposCreateReleaseRequest) SetDiscussionCategoryName(v string)`

SetDiscussionCategoryName sets DiscussionCategoryName field to given value.

### HasDiscussionCategoryName

`func (o *ReposCreateReleaseRequest) HasDiscussionCategoryName() bool`

HasDiscussionCategoryName returns a boolean if a field has been set.

### GetGenerateReleaseNotes

`func (o *ReposCreateReleaseRequest) GetGenerateReleaseNotes() bool`

GetGenerateReleaseNotes returns the GenerateReleaseNotes field if non-nil, zero value otherwise.

### GetGenerateReleaseNotesOk

`func (o *ReposCreateReleaseRequest) GetGenerateReleaseNotesOk() (*bool, bool)`

GetGenerateReleaseNotesOk returns a tuple with the GenerateReleaseNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateReleaseNotes

`func (o *ReposCreateReleaseRequest) SetGenerateReleaseNotes(v bool)`

SetGenerateReleaseNotes sets GenerateReleaseNotes field to given value.

### HasGenerateReleaseNotes

`func (o *ReposCreateReleaseRequest) HasGenerateReleaseNotes() bool`

HasGenerateReleaseNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


