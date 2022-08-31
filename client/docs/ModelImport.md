# ModelImport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vcs** | **NullableString** |  | 
**UseLfs** | Pointer to **bool** |  | [optional] 
**VcsUrl** | **string** | The URL of the originating repository. | 
**SvcRoot** | Pointer to **string** |  | [optional] 
**TfvcProject** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**StatusText** | Pointer to **NullableString** |  | [optional] 
**FailedStep** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**ImportPercent** | Pointer to **NullableInt32** |  | [optional] 
**CommitCount** | Pointer to **NullableInt32** |  | [optional] 
**PushPercent** | Pointer to **NullableInt32** |  | [optional] 
**HasLargeFiles** | Pointer to **bool** |  | [optional] 
**LargeFilesSize** | Pointer to **int32** |  | [optional] 
**LargeFilesCount** | Pointer to **int32** |  | [optional] 
**ProjectChoices** | Pointer to [**[]ImportProjectChoicesInner**](ImportProjectChoicesInner.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**AuthorsCount** | Pointer to **NullableInt32** |  | [optional] 
**Url** | **string** |  | 
**HtmlUrl** | **string** |  | 
**AuthorsUrl** | **string** |  | 
**RepositoryUrl** | **string** |  | 
**SvnRoot** | Pointer to **string** |  | [optional] 

## Methods

### NewModelImport

`func NewModelImport(vcs NullableString, vcsUrl string, status string, url string, htmlUrl string, authorsUrl string, repositoryUrl string, ) *ModelImport`

NewModelImport instantiates a new ModelImport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelImportWithDefaults

`func NewModelImportWithDefaults() *ModelImport`

NewModelImportWithDefaults instantiates a new ModelImport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVcs

`func (o *ModelImport) GetVcs() string`

GetVcs returns the Vcs field if non-nil, zero value otherwise.

### GetVcsOk

`func (o *ModelImport) GetVcsOk() (*string, bool)`

GetVcsOk returns a tuple with the Vcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcs

`func (o *ModelImport) SetVcs(v string)`

SetVcs sets Vcs field to given value.


### SetVcsNil

`func (o *ModelImport) SetVcsNil(b bool)`

 SetVcsNil sets the value for Vcs to be an explicit nil

### UnsetVcs
`func (o *ModelImport) UnsetVcs()`

UnsetVcs ensures that no value is present for Vcs, not even an explicit nil
### GetUseLfs

`func (o *ModelImport) GetUseLfs() bool`

GetUseLfs returns the UseLfs field if non-nil, zero value otherwise.

### GetUseLfsOk

`func (o *ModelImport) GetUseLfsOk() (*bool, bool)`

GetUseLfsOk returns a tuple with the UseLfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLfs

`func (o *ModelImport) SetUseLfs(v bool)`

SetUseLfs sets UseLfs field to given value.

### HasUseLfs

`func (o *ModelImport) HasUseLfs() bool`

HasUseLfs returns a boolean if a field has been set.

### GetVcsUrl

`func (o *ModelImport) GetVcsUrl() string`

GetVcsUrl returns the VcsUrl field if non-nil, zero value otherwise.

### GetVcsUrlOk

`func (o *ModelImport) GetVcsUrlOk() (*string, bool)`

GetVcsUrlOk returns a tuple with the VcsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcsUrl

`func (o *ModelImport) SetVcsUrl(v string)`

SetVcsUrl sets VcsUrl field to given value.


### GetSvcRoot

`func (o *ModelImport) GetSvcRoot() string`

GetSvcRoot returns the SvcRoot field if non-nil, zero value otherwise.

### GetSvcRootOk

`func (o *ModelImport) GetSvcRootOk() (*string, bool)`

GetSvcRootOk returns a tuple with the SvcRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvcRoot

`func (o *ModelImport) SetSvcRoot(v string)`

SetSvcRoot sets SvcRoot field to given value.

### HasSvcRoot

`func (o *ModelImport) HasSvcRoot() bool`

HasSvcRoot returns a boolean if a field has been set.

### GetTfvcProject

`func (o *ModelImport) GetTfvcProject() string`

GetTfvcProject returns the TfvcProject field if non-nil, zero value otherwise.

### GetTfvcProjectOk

`func (o *ModelImport) GetTfvcProjectOk() (*string, bool)`

GetTfvcProjectOk returns a tuple with the TfvcProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfvcProject

`func (o *ModelImport) SetTfvcProject(v string)`

SetTfvcProject sets TfvcProject field to given value.

### HasTfvcProject

`func (o *ModelImport) HasTfvcProject() bool`

HasTfvcProject returns a boolean if a field has been set.

### GetStatus

`func (o *ModelImport) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelImport) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelImport) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusText

`func (o *ModelImport) GetStatusText() string`

GetStatusText returns the StatusText field if non-nil, zero value otherwise.

### GetStatusTextOk

`func (o *ModelImport) GetStatusTextOk() (*string, bool)`

GetStatusTextOk returns a tuple with the StatusText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusText

`func (o *ModelImport) SetStatusText(v string)`

SetStatusText sets StatusText field to given value.

### HasStatusText

`func (o *ModelImport) HasStatusText() bool`

HasStatusText returns a boolean if a field has been set.

### SetStatusTextNil

`func (o *ModelImport) SetStatusTextNil(b bool)`

 SetStatusTextNil sets the value for StatusText to be an explicit nil

### UnsetStatusText
`func (o *ModelImport) UnsetStatusText()`

UnsetStatusText ensures that no value is present for StatusText, not even an explicit nil
### GetFailedStep

`func (o *ModelImport) GetFailedStep() string`

GetFailedStep returns the FailedStep field if non-nil, zero value otherwise.

### GetFailedStepOk

`func (o *ModelImport) GetFailedStepOk() (*string, bool)`

GetFailedStepOk returns a tuple with the FailedStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedStep

`func (o *ModelImport) SetFailedStep(v string)`

SetFailedStep sets FailedStep field to given value.

### HasFailedStep

`func (o *ModelImport) HasFailedStep() bool`

HasFailedStep returns a boolean if a field has been set.

### SetFailedStepNil

`func (o *ModelImport) SetFailedStepNil(b bool)`

 SetFailedStepNil sets the value for FailedStep to be an explicit nil

### UnsetFailedStep
`func (o *ModelImport) UnsetFailedStep()`

UnsetFailedStep ensures that no value is present for FailedStep, not even an explicit nil
### GetErrorMessage

`func (o *ModelImport) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ModelImport) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ModelImport) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ModelImport) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ModelImport) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ModelImport) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetImportPercent

`func (o *ModelImport) GetImportPercent() int32`

GetImportPercent returns the ImportPercent field if non-nil, zero value otherwise.

### GetImportPercentOk

`func (o *ModelImport) GetImportPercentOk() (*int32, bool)`

GetImportPercentOk returns a tuple with the ImportPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportPercent

`func (o *ModelImport) SetImportPercent(v int32)`

SetImportPercent sets ImportPercent field to given value.

### HasImportPercent

`func (o *ModelImport) HasImportPercent() bool`

HasImportPercent returns a boolean if a field has been set.

### SetImportPercentNil

`func (o *ModelImport) SetImportPercentNil(b bool)`

 SetImportPercentNil sets the value for ImportPercent to be an explicit nil

### UnsetImportPercent
`func (o *ModelImport) UnsetImportPercent()`

UnsetImportPercent ensures that no value is present for ImportPercent, not even an explicit nil
### GetCommitCount

`func (o *ModelImport) GetCommitCount() int32`

GetCommitCount returns the CommitCount field if non-nil, zero value otherwise.

### GetCommitCountOk

`func (o *ModelImport) GetCommitCountOk() (*int32, bool)`

GetCommitCountOk returns a tuple with the CommitCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitCount

`func (o *ModelImport) SetCommitCount(v int32)`

SetCommitCount sets CommitCount field to given value.

### HasCommitCount

`func (o *ModelImport) HasCommitCount() bool`

HasCommitCount returns a boolean if a field has been set.

### SetCommitCountNil

`func (o *ModelImport) SetCommitCountNil(b bool)`

 SetCommitCountNil sets the value for CommitCount to be an explicit nil

### UnsetCommitCount
`func (o *ModelImport) UnsetCommitCount()`

UnsetCommitCount ensures that no value is present for CommitCount, not even an explicit nil
### GetPushPercent

`func (o *ModelImport) GetPushPercent() int32`

GetPushPercent returns the PushPercent field if non-nil, zero value otherwise.

### GetPushPercentOk

`func (o *ModelImport) GetPushPercentOk() (*int32, bool)`

GetPushPercentOk returns a tuple with the PushPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushPercent

`func (o *ModelImport) SetPushPercent(v int32)`

SetPushPercent sets PushPercent field to given value.

### HasPushPercent

`func (o *ModelImport) HasPushPercent() bool`

HasPushPercent returns a boolean if a field has been set.

### SetPushPercentNil

`func (o *ModelImport) SetPushPercentNil(b bool)`

 SetPushPercentNil sets the value for PushPercent to be an explicit nil

### UnsetPushPercent
`func (o *ModelImport) UnsetPushPercent()`

UnsetPushPercent ensures that no value is present for PushPercent, not even an explicit nil
### GetHasLargeFiles

`func (o *ModelImport) GetHasLargeFiles() bool`

GetHasLargeFiles returns the HasLargeFiles field if non-nil, zero value otherwise.

### GetHasLargeFilesOk

`func (o *ModelImport) GetHasLargeFilesOk() (*bool, bool)`

GetHasLargeFilesOk returns a tuple with the HasLargeFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLargeFiles

`func (o *ModelImport) SetHasLargeFiles(v bool)`

SetHasLargeFiles sets HasLargeFiles field to given value.

### HasHasLargeFiles

`func (o *ModelImport) HasHasLargeFiles() bool`

HasHasLargeFiles returns a boolean if a field has been set.

### GetLargeFilesSize

`func (o *ModelImport) GetLargeFilesSize() int32`

GetLargeFilesSize returns the LargeFilesSize field if non-nil, zero value otherwise.

### GetLargeFilesSizeOk

`func (o *ModelImport) GetLargeFilesSizeOk() (*int32, bool)`

GetLargeFilesSizeOk returns a tuple with the LargeFilesSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeFilesSize

`func (o *ModelImport) SetLargeFilesSize(v int32)`

SetLargeFilesSize sets LargeFilesSize field to given value.

### HasLargeFilesSize

`func (o *ModelImport) HasLargeFilesSize() bool`

HasLargeFilesSize returns a boolean if a field has been set.

### GetLargeFilesCount

`func (o *ModelImport) GetLargeFilesCount() int32`

GetLargeFilesCount returns the LargeFilesCount field if non-nil, zero value otherwise.

### GetLargeFilesCountOk

`func (o *ModelImport) GetLargeFilesCountOk() (*int32, bool)`

GetLargeFilesCountOk returns a tuple with the LargeFilesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeFilesCount

`func (o *ModelImport) SetLargeFilesCount(v int32)`

SetLargeFilesCount sets LargeFilesCount field to given value.

### HasLargeFilesCount

`func (o *ModelImport) HasLargeFilesCount() bool`

HasLargeFilesCount returns a boolean if a field has been set.

### GetProjectChoices

`func (o *ModelImport) GetProjectChoices() []ImportProjectChoicesInner`

GetProjectChoices returns the ProjectChoices field if non-nil, zero value otherwise.

### GetProjectChoicesOk

`func (o *ModelImport) GetProjectChoicesOk() (*[]ImportProjectChoicesInner, bool)`

GetProjectChoicesOk returns a tuple with the ProjectChoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectChoices

`func (o *ModelImport) SetProjectChoices(v []ImportProjectChoicesInner)`

SetProjectChoices sets ProjectChoices field to given value.

### HasProjectChoices

`func (o *ModelImport) HasProjectChoices() bool`

HasProjectChoices returns a boolean if a field has been set.

### GetMessage

`func (o *ModelImport) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ModelImport) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ModelImport) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ModelImport) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetAuthorsCount

`func (o *ModelImport) GetAuthorsCount() int32`

GetAuthorsCount returns the AuthorsCount field if non-nil, zero value otherwise.

### GetAuthorsCountOk

`func (o *ModelImport) GetAuthorsCountOk() (*int32, bool)`

GetAuthorsCountOk returns a tuple with the AuthorsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorsCount

`func (o *ModelImport) SetAuthorsCount(v int32)`

SetAuthorsCount sets AuthorsCount field to given value.

### HasAuthorsCount

`func (o *ModelImport) HasAuthorsCount() bool`

HasAuthorsCount returns a boolean if a field has been set.

### SetAuthorsCountNil

`func (o *ModelImport) SetAuthorsCountNil(b bool)`

 SetAuthorsCountNil sets the value for AuthorsCount to be an explicit nil

### UnsetAuthorsCount
`func (o *ModelImport) UnsetAuthorsCount()`

UnsetAuthorsCount ensures that no value is present for AuthorsCount, not even an explicit nil
### GetUrl

`func (o *ModelImport) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ModelImport) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ModelImport) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *ModelImport) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *ModelImport) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *ModelImport) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetAuthorsUrl

`func (o *ModelImport) GetAuthorsUrl() string`

GetAuthorsUrl returns the AuthorsUrl field if non-nil, zero value otherwise.

### GetAuthorsUrlOk

`func (o *ModelImport) GetAuthorsUrlOk() (*string, bool)`

GetAuthorsUrlOk returns a tuple with the AuthorsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorsUrl

`func (o *ModelImport) SetAuthorsUrl(v string)`

SetAuthorsUrl sets AuthorsUrl field to given value.


### GetRepositoryUrl

`func (o *ModelImport) GetRepositoryUrl() string`

GetRepositoryUrl returns the RepositoryUrl field if non-nil, zero value otherwise.

### GetRepositoryUrlOk

`func (o *ModelImport) GetRepositoryUrlOk() (*string, bool)`

GetRepositoryUrlOk returns a tuple with the RepositoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositoryUrl

`func (o *ModelImport) SetRepositoryUrl(v string)`

SetRepositoryUrl sets RepositoryUrl field to given value.


### GetSvnRoot

`func (o *ModelImport) GetSvnRoot() string`

GetSvnRoot returns the SvnRoot field if non-nil, zero value otherwise.

### GetSvnRootOk

`func (o *ModelImport) GetSvnRootOk() (*string, bool)`

GetSvnRootOk returns a tuple with the SvnRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvnRoot

`func (o *ModelImport) SetSvnRoot(v string)`

SetSvnRoot sets SvnRoot field to given value.

### HasSvnRoot

`func (o *ModelImport) HasSvnRoot() bool`

HasSvnRoot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


