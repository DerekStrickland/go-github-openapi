# CommitComparison

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**HtmlUrl** | **string** |  | 
**PermalinkUrl** | **string** |  | 
**DiffUrl** | **string** |  | 
**PatchUrl** | **string** |  | 
**BaseCommit** | [**Commit**](Commit.md) |  | 
**MergeBaseCommit** | [**Commit**](Commit.md) |  | 
**Status** | **string** |  | 
**AheadBy** | **int32** |  | 
**BehindBy** | **int32** |  | 
**TotalCommits** | **int32** |  | 
**Commits** | [**[]Commit**](Commit.md) |  | 
**Files** | Pointer to [**[]DiffEntry**](DiffEntry.md) |  | [optional] 

## Methods

### NewCommitComparison

`func NewCommitComparison(url string, htmlUrl string, permalinkUrl string, diffUrl string, patchUrl string, baseCommit Commit, mergeBaseCommit Commit, status string, aheadBy int32, behindBy int32, totalCommits int32, commits []Commit, ) *CommitComparison`

NewCommitComparison instantiates a new CommitComparison object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommitComparisonWithDefaults

`func NewCommitComparisonWithDefaults() *CommitComparison`

NewCommitComparisonWithDefaults instantiates a new CommitComparison object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *CommitComparison) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CommitComparison) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CommitComparison) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *CommitComparison) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *CommitComparison) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *CommitComparison) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetPermalinkUrl

`func (o *CommitComparison) GetPermalinkUrl() string`

GetPermalinkUrl returns the PermalinkUrl field if non-nil, zero value otherwise.

### GetPermalinkUrlOk

`func (o *CommitComparison) GetPermalinkUrlOk() (*string, bool)`

GetPermalinkUrlOk returns a tuple with the PermalinkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermalinkUrl

`func (o *CommitComparison) SetPermalinkUrl(v string)`

SetPermalinkUrl sets PermalinkUrl field to given value.


### GetDiffUrl

`func (o *CommitComparison) GetDiffUrl() string`

GetDiffUrl returns the DiffUrl field if non-nil, zero value otherwise.

### GetDiffUrlOk

`func (o *CommitComparison) GetDiffUrlOk() (*string, bool)`

GetDiffUrlOk returns a tuple with the DiffUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiffUrl

`func (o *CommitComparison) SetDiffUrl(v string)`

SetDiffUrl sets DiffUrl field to given value.


### GetPatchUrl

`func (o *CommitComparison) GetPatchUrl() string`

GetPatchUrl returns the PatchUrl field if non-nil, zero value otherwise.

### GetPatchUrlOk

`func (o *CommitComparison) GetPatchUrlOk() (*string, bool)`

GetPatchUrlOk returns a tuple with the PatchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchUrl

`func (o *CommitComparison) SetPatchUrl(v string)`

SetPatchUrl sets PatchUrl field to given value.


### GetBaseCommit

`func (o *CommitComparison) GetBaseCommit() Commit`

GetBaseCommit returns the BaseCommit field if non-nil, zero value otherwise.

### GetBaseCommitOk

`func (o *CommitComparison) GetBaseCommitOk() (*Commit, bool)`

GetBaseCommitOk returns a tuple with the BaseCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCommit

`func (o *CommitComparison) SetBaseCommit(v Commit)`

SetBaseCommit sets BaseCommit field to given value.


### GetMergeBaseCommit

`func (o *CommitComparison) GetMergeBaseCommit() Commit`

GetMergeBaseCommit returns the MergeBaseCommit field if non-nil, zero value otherwise.

### GetMergeBaseCommitOk

`func (o *CommitComparison) GetMergeBaseCommitOk() (*Commit, bool)`

GetMergeBaseCommitOk returns a tuple with the MergeBaseCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeBaseCommit

`func (o *CommitComparison) SetMergeBaseCommit(v Commit)`

SetMergeBaseCommit sets MergeBaseCommit field to given value.


### GetStatus

`func (o *CommitComparison) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommitComparison) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommitComparison) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAheadBy

`func (o *CommitComparison) GetAheadBy() int32`

GetAheadBy returns the AheadBy field if non-nil, zero value otherwise.

### GetAheadByOk

`func (o *CommitComparison) GetAheadByOk() (*int32, bool)`

GetAheadByOk returns a tuple with the AheadBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAheadBy

`func (o *CommitComparison) SetAheadBy(v int32)`

SetAheadBy sets AheadBy field to given value.


### GetBehindBy

`func (o *CommitComparison) GetBehindBy() int32`

GetBehindBy returns the BehindBy field if non-nil, zero value otherwise.

### GetBehindByOk

`func (o *CommitComparison) GetBehindByOk() (*int32, bool)`

GetBehindByOk returns a tuple with the BehindBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehindBy

`func (o *CommitComparison) SetBehindBy(v int32)`

SetBehindBy sets BehindBy field to given value.


### GetTotalCommits

`func (o *CommitComparison) GetTotalCommits() int32`

GetTotalCommits returns the TotalCommits field if non-nil, zero value otherwise.

### GetTotalCommitsOk

`func (o *CommitComparison) GetTotalCommitsOk() (*int32, bool)`

GetTotalCommitsOk returns a tuple with the TotalCommits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCommits

`func (o *CommitComparison) SetTotalCommits(v int32)`

SetTotalCommits sets TotalCommits field to given value.


### GetCommits

`func (o *CommitComparison) GetCommits() []Commit`

GetCommits returns the Commits field if non-nil, zero value otherwise.

### GetCommitsOk

`func (o *CommitComparison) GetCommitsOk() (*[]Commit, bool)`

GetCommitsOk returns a tuple with the Commits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommits

`func (o *CommitComparison) SetCommits(v []Commit)`

SetCommits sets Commits field to given value.


### GetFiles

`func (o *CommitComparison) GetFiles() []DiffEntry`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *CommitComparison) GetFilesOk() (*[]DiffEntry, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *CommitComparison) SetFiles(v []DiffEntry)`

SetFiles sets Files field to given value.

### HasFiles

`func (o *CommitComparison) HasFiles() bool`

HasFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


