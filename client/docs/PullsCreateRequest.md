# PullsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | The title of the new pull request. Required unless &#x60;issue&#x60; is specified. | [optional] 
**Head** | **string** | The name of the branch where your changes are implemented. For cross-repository pull requests in the same network, namespace &#x60;head&#x60; with a user like this: &#x60;username:branch&#x60;. | 
**Base** | **string** | The name of the branch you want the changes pulled into. This should be an existing branch on the current repository. You cannot submit a pull request to one repository that requests a merge to a base of another repository. | 
**Body** | Pointer to **string** | The contents of the pull request. | [optional] 
**MaintainerCanModify** | Pointer to **bool** | Indicates whether [maintainers can modify](https://docs.github.com/articles/allowing-changes-to-a-pull-request-branch-created-from-a-fork/) the pull request. | [optional] 
**Draft** | Pointer to **bool** | Indicates whether the pull request is a draft. See \&quot;[Draft Pull Requests](https://docs.github.com/en/articles/about-pull-requests#draft-pull-requests)\&quot; in the GitHub Help documentation to learn more. | [optional] 
**Issue** | Pointer to **int32** | An issue in the repository to convert to a pull request. The issue title, body, and comments will become the title, body, and comments on the new pull request. Required unless &#x60;title&#x60; is specified. | [optional] 

## Methods

### NewPullsCreateRequest

`func NewPullsCreateRequest(head string, base string, ) *PullsCreateRequest`

NewPullsCreateRequest instantiates a new PullsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullsCreateRequestWithDefaults

`func NewPullsCreateRequestWithDefaults() *PullsCreateRequest`

NewPullsCreateRequestWithDefaults instantiates a new PullsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *PullsCreateRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PullsCreateRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PullsCreateRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PullsCreateRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetHead

`func (o *PullsCreateRequest) GetHead() string`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *PullsCreateRequest) GetHeadOk() (*string, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *PullsCreateRequest) SetHead(v string)`

SetHead sets Head field to given value.


### GetBase

`func (o *PullsCreateRequest) GetBase() string`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *PullsCreateRequest) GetBaseOk() (*string, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *PullsCreateRequest) SetBase(v string)`

SetBase sets Base field to given value.


### GetBody

`func (o *PullsCreateRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PullsCreateRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PullsCreateRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *PullsCreateRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetMaintainerCanModify

`func (o *PullsCreateRequest) GetMaintainerCanModify() bool`

GetMaintainerCanModify returns the MaintainerCanModify field if non-nil, zero value otherwise.

### GetMaintainerCanModifyOk

`func (o *PullsCreateRequest) GetMaintainerCanModifyOk() (*bool, bool)`

GetMaintainerCanModifyOk returns a tuple with the MaintainerCanModify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintainerCanModify

`func (o *PullsCreateRequest) SetMaintainerCanModify(v bool)`

SetMaintainerCanModify sets MaintainerCanModify field to given value.

### HasMaintainerCanModify

`func (o *PullsCreateRequest) HasMaintainerCanModify() bool`

HasMaintainerCanModify returns a boolean if a field has been set.

### GetDraft

`func (o *PullsCreateRequest) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *PullsCreateRequest) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *PullsCreateRequest) SetDraft(v bool)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *PullsCreateRequest) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetIssue

`func (o *PullsCreateRequest) GetIssue() int32`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *PullsCreateRequest) GetIssueOk() (*int32, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *PullsCreateRequest) SetIssue(v int32)`

SetIssue sets Issue field to given value.

### HasIssue

`func (o *PullsCreateRequest) HasIssue() bool`

HasIssue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


