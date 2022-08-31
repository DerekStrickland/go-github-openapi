# IssuePullRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MergedAt** | Pointer to **NullableTime** |  | [optional] 
**DiffUrl** | **NullableString** |  | 
**HtmlUrl** | **NullableString** |  | 
**PatchUrl** | **NullableString** |  | 
**Url** | **NullableString** |  | 

## Methods

### NewIssuePullRequest

`func NewIssuePullRequest(diffUrl NullableString, htmlUrl NullableString, patchUrl NullableString, url NullableString, ) *IssuePullRequest`

NewIssuePullRequest instantiates a new IssuePullRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuePullRequestWithDefaults

`func NewIssuePullRequestWithDefaults() *IssuePullRequest`

NewIssuePullRequestWithDefaults instantiates a new IssuePullRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMergedAt

`func (o *IssuePullRequest) GetMergedAt() time.Time`

GetMergedAt returns the MergedAt field if non-nil, zero value otherwise.

### GetMergedAtOk

`func (o *IssuePullRequest) GetMergedAtOk() (*time.Time, bool)`

GetMergedAtOk returns a tuple with the MergedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedAt

`func (o *IssuePullRequest) SetMergedAt(v time.Time)`

SetMergedAt sets MergedAt field to given value.

### HasMergedAt

`func (o *IssuePullRequest) HasMergedAt() bool`

HasMergedAt returns a boolean if a field has been set.

### SetMergedAtNil

`func (o *IssuePullRequest) SetMergedAtNil(b bool)`

 SetMergedAtNil sets the value for MergedAt to be an explicit nil

### UnsetMergedAt
`func (o *IssuePullRequest) UnsetMergedAt()`

UnsetMergedAt ensures that no value is present for MergedAt, not even an explicit nil
### GetDiffUrl

`func (o *IssuePullRequest) GetDiffUrl() string`

GetDiffUrl returns the DiffUrl field if non-nil, zero value otherwise.

### GetDiffUrlOk

`func (o *IssuePullRequest) GetDiffUrlOk() (*string, bool)`

GetDiffUrlOk returns a tuple with the DiffUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffUrl

`func (o *IssuePullRequest) SetDiffUrl(v string)`

SetDiffUrl sets DiffUrl field to given value.


### SetDiffUrlNil

`func (o *IssuePullRequest) SetDiffUrlNil(b bool)`

 SetDiffUrlNil sets the value for DiffUrl to be an explicit nil

### UnsetDiffUrl
`func (o *IssuePullRequest) UnsetDiffUrl()`

UnsetDiffUrl ensures that no value is present for DiffUrl, not even an explicit nil
### GetHtmlUrl

`func (o *IssuePullRequest) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *IssuePullRequest) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *IssuePullRequest) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### SetHtmlUrlNil

`func (o *IssuePullRequest) SetHtmlUrlNil(b bool)`

 SetHtmlUrlNil sets the value for HtmlUrl to be an explicit nil

### UnsetHtmlUrl
`func (o *IssuePullRequest) UnsetHtmlUrl()`

UnsetHtmlUrl ensures that no value is present for HtmlUrl, not even an explicit nil
### GetPatchUrl

`func (o *IssuePullRequest) GetPatchUrl() string`

GetPatchUrl returns the PatchUrl field if non-nil, zero value otherwise.

### GetPatchUrlOk

`func (o *IssuePullRequest) GetPatchUrlOk() (*string, bool)`

GetPatchUrlOk returns a tuple with the PatchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchUrl

`func (o *IssuePullRequest) SetPatchUrl(v string)`

SetPatchUrl sets PatchUrl field to given value.


### SetPatchUrlNil

`func (o *IssuePullRequest) SetPatchUrlNil(b bool)`

 SetPatchUrlNil sets the value for PatchUrl to be an explicit nil

### UnsetPatchUrl
`func (o *IssuePullRequest) UnsetPatchUrl()`

UnsetPatchUrl ensures that no value is present for PatchUrl, not even an explicit nil
### GetUrl

`func (o *IssuePullRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *IssuePullRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *IssuePullRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### SetUrlNil

`func (o *IssuePullRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *IssuePullRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


