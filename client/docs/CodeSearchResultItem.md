# CodeSearchResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Path** | **string** |  | 
**Sha** | **string** |  | 
**Url** | **string** |  | 
**GitUrl** | **string** |  | 
**HtmlUrl** | **string** |  | 
**Repository** | [**MinimalRepository**](MinimalRepository.md) |  | 
**Score** | **float32** |  | 
**FileSize** | Pointer to **int32** |  | [optional] 
**Language** | Pointer to **NullableString** |  | [optional] 
**LastModifiedAt** | Pointer to **time.Time** |  | [optional] 
**LineNumbers** | Pointer to **[]string** |  | [optional] 
**TextMatches** | Pointer to [**[]SearchResultTextMatchesInner**](SearchResultTextMatchesInner.md) |  | [optional] 

## Methods

### NewCodeSearchResultItem

`func NewCodeSearchResultItem(name string, path string, sha string, url string, gitUrl string, htmlUrl string, repository MinimalRepository, score float32, ) *CodeSearchResultItem`

NewCodeSearchResultItem instantiates a new CodeSearchResultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeSearchResultItemWithDefaults

`func NewCodeSearchResultItemWithDefaults() *CodeSearchResultItem`

NewCodeSearchResultItemWithDefaults instantiates a new CodeSearchResultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CodeSearchResultItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CodeSearchResultItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CodeSearchResultItem) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *CodeSearchResultItem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CodeSearchResultItem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CodeSearchResultItem) SetPath(v string)`

SetPath sets Path field to given value.


### GetSha

`func (o *CodeSearchResultItem) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *CodeSearchResultItem) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *CodeSearchResultItem) SetSha(v string)`

SetSha sets Sha field to given value.


### GetUrl

`func (o *CodeSearchResultItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CodeSearchResultItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CodeSearchResultItem) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetGitUrl

`func (o *CodeSearchResultItem) GetGitUrl() string`

GetGitUrl returns the GitUrl field if non-nil, zero value otherwise.

### GetGitUrlOk

`func (o *CodeSearchResultItem) GetGitUrlOk() (*string, bool)`

GetGitUrlOk returns a tuple with the GitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitUrl

`func (o *CodeSearchResultItem) SetGitUrl(v string)`

SetGitUrl sets GitUrl field to given value.


### GetHtmlUrl

`func (o *CodeSearchResultItem) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *CodeSearchResultItem) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *CodeSearchResultItem) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetRepository

`func (o *CodeSearchResultItem) GetRepository() MinimalRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *CodeSearchResultItem) GetRepositoryOk() (*MinimalRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *CodeSearchResultItem) SetRepository(v MinimalRepository)`

SetRepository sets Repository field to given value.


### GetScore

`func (o *CodeSearchResultItem) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *CodeSearchResultItem) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *CodeSearchResultItem) SetScore(v float32)`

SetScore sets Score field to given value.


### GetFileSize

`func (o *CodeSearchResultItem) GetFileSize() int32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *CodeSearchResultItem) GetFileSizeOk() (*int32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *CodeSearchResultItem) SetFileSize(v int32)`

SetFileSize sets FileSize field to given value.

### HasFileSize

`func (o *CodeSearchResultItem) HasFileSize() bool`

HasFileSize returns a boolean if a field has been set.

### GetLanguage

`func (o *CodeSearchResultItem) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *CodeSearchResultItem) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *CodeSearchResultItem) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *CodeSearchResultItem) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *CodeSearchResultItem) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *CodeSearchResultItem) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil
### GetLastModifiedAt

`func (o *CodeSearchResultItem) GetLastModifiedAt() time.Time`

GetLastModifiedAt returns the LastModifiedAt field if non-nil, zero value otherwise.

### GetLastModifiedAtOk

`func (o *CodeSearchResultItem) GetLastModifiedAtOk() (*time.Time, bool)`

GetLastModifiedAtOk returns a tuple with the LastModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedAt

`func (o *CodeSearchResultItem) SetLastModifiedAt(v time.Time)`

SetLastModifiedAt sets LastModifiedAt field to given value.

### HasLastModifiedAt

`func (o *CodeSearchResultItem) HasLastModifiedAt() bool`

HasLastModifiedAt returns a boolean if a field has been set.

### GetLineNumbers

`func (o *CodeSearchResultItem) GetLineNumbers() []string`

GetLineNumbers returns the LineNumbers field if non-nil, zero value otherwise.

### GetLineNumbersOk

`func (o *CodeSearchResultItem) GetLineNumbersOk() (*[]string, bool)`

GetLineNumbersOk returns a tuple with the LineNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineNumbers

`func (o *CodeSearchResultItem) SetLineNumbers(v []string)`

SetLineNumbers sets LineNumbers field to given value.

### HasLineNumbers

`func (o *CodeSearchResultItem) HasLineNumbers() bool`

HasLineNumbers returns a boolean if a field has been set.

### GetTextMatches

`func (o *CodeSearchResultItem) GetTextMatches() []SearchResultTextMatchesInner`

GetTextMatches returns the TextMatches field if non-nil, zero value otherwise.

### GetTextMatchesOk

`func (o *CodeSearchResultItem) GetTextMatchesOk() (*[]SearchResultTextMatchesInner, bool)`

GetTextMatchesOk returns a tuple with the TextMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextMatches

`func (o *CodeSearchResultItem) SetTextMatches(v []SearchResultTextMatchesInner)`

SetTextMatches sets TextMatches field to given value.

### HasTextMatches

`func (o *CodeSearchResultItem) HasTextMatches() bool`

HasTextMatches returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


