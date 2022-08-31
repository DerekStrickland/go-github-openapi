# CodespaceExportDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to **NullableString** | State of the latest export | [optional] 
**CompletedAt** | Pointer to **NullableTime** | Completion time of the last export operation | [optional] 
**Branch** | Pointer to **NullableString** | Name of the exported branch | [optional] 
**Sha** | Pointer to **NullableString** | Git commit SHA of the exported branch | [optional] 
**Id** | Pointer to **string** | Id for the export details | [optional] 
**ExportUrl** | Pointer to **string** | Url for fetching export details | [optional] 
**HtmlUrl** | Pointer to **NullableString** | Web url for the exported branch | [optional] 

## Methods

### NewCodespaceExportDetails

`func NewCodespaceExportDetails() *CodespaceExportDetails`

NewCodespaceExportDetails instantiates a new CodespaceExportDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodespaceExportDetailsWithDefaults

`func NewCodespaceExportDetailsWithDefaults() *CodespaceExportDetails`

NewCodespaceExportDetailsWithDefaults instantiates a new CodespaceExportDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *CodespaceExportDetails) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CodespaceExportDetails) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CodespaceExportDetails) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CodespaceExportDetails) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *CodespaceExportDetails) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *CodespaceExportDetails) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetCompletedAt

`func (o *CodespaceExportDetails) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *CodespaceExportDetails) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *CodespaceExportDetails) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *CodespaceExportDetails) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *CodespaceExportDetails) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *CodespaceExportDetails) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetBranch

`func (o *CodespaceExportDetails) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *CodespaceExportDetails) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *CodespaceExportDetails) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *CodespaceExportDetails) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### SetBranchNil

`func (o *CodespaceExportDetails) SetBranchNil(b bool)`

 SetBranchNil sets the value for Branch to be an explicit nil

### UnsetBranch
`func (o *CodespaceExportDetails) UnsetBranch()`

UnsetBranch ensures that no value is present for Branch, not even an explicit nil
### GetSha

`func (o *CodespaceExportDetails) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *CodespaceExportDetails) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *CodespaceExportDetails) SetSha(v string)`

SetSha sets Sha field to given value.

### HasSha

`func (o *CodespaceExportDetails) HasSha() bool`

HasSha returns a boolean if a field has been set.

### SetShaNil

`func (o *CodespaceExportDetails) SetShaNil(b bool)`

 SetShaNil sets the value for Sha to be an explicit nil

### UnsetSha
`func (o *CodespaceExportDetails) UnsetSha()`

UnsetSha ensures that no value is present for Sha, not even an explicit nil
### GetId

`func (o *CodespaceExportDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CodespaceExportDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CodespaceExportDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CodespaceExportDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExportUrl

`func (o *CodespaceExportDetails) GetExportUrl() string`

GetExportUrl returns the ExportUrl field if non-nil, zero value otherwise.

### GetExportUrlOk

`func (o *CodespaceExportDetails) GetExportUrlOk() (*string, bool)`

GetExportUrlOk returns a tuple with the ExportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportUrl

`func (o *CodespaceExportDetails) SetExportUrl(v string)`

SetExportUrl sets ExportUrl field to given value.

### HasExportUrl

`func (o *CodespaceExportDetails) HasExportUrl() bool`

HasExportUrl returns a boolean if a field has been set.

### GetHtmlUrl

`func (o *CodespaceExportDetails) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *CodespaceExportDetails) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *CodespaceExportDetails) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *CodespaceExportDetails) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.

### SetHtmlUrlNil

`func (o *CodespaceExportDetails) SetHtmlUrlNil(b bool)`

 SetHtmlUrlNil sets the value for HtmlUrl to be an explicit nil

### UnsetHtmlUrl
`func (o *CodespaceExportDetails) UnsetHtmlUrl()`

UnsetHtmlUrl ensures that no value is present for HtmlUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


