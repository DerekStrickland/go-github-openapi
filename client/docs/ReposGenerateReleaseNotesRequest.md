# ReposGenerateReleaseNotesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TagName** | **string** | The tag name for the release. This can be an existing tag or a new one. | 
**TargetCommitish** | Pointer to **string** | Specifies the commitish value that will be the target for the release&#39;s tag. Required if the supplied tag_name does not reference an existing tag. Ignored if the tag_name already exists. | [optional] 
**PreviousTagName** | Pointer to **string** | The name of the previous tag to use as the starting point for the release notes. Use to manually specify the range for the set of changes considered as part this release. | [optional] 
**ConfigurationFilePath** | Pointer to **string** | Specifies a path to a file in the repository containing configuration settings used for generating the release notes. If unspecified, the configuration file located in the repository at &#39;.github/release.yml&#39; or &#39;.github/release.yaml&#39; will be used. If that is not present, the default configuration will be used. | [optional] 

## Methods

### NewReposGenerateReleaseNotesRequest

`func NewReposGenerateReleaseNotesRequest(tagName string, ) *ReposGenerateReleaseNotesRequest`

NewReposGenerateReleaseNotesRequest instantiates a new ReposGenerateReleaseNotesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposGenerateReleaseNotesRequestWithDefaults

`func NewReposGenerateReleaseNotesRequestWithDefaults() *ReposGenerateReleaseNotesRequest`

NewReposGenerateReleaseNotesRequestWithDefaults instantiates a new ReposGenerateReleaseNotesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTagName

`func (o *ReposGenerateReleaseNotesRequest) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *ReposGenerateReleaseNotesRequest) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *ReposGenerateReleaseNotesRequest) SetTagName(v string)`

SetTagName sets TagName field to given value.


### GetTargetCommitish

`func (o *ReposGenerateReleaseNotesRequest) GetTargetCommitish() string`

GetTargetCommitish returns the TargetCommitish field if non-nil, zero value otherwise.

### GetTargetCommitishOk

`func (o *ReposGenerateReleaseNotesRequest) GetTargetCommitishOk() (*string, bool)`

GetTargetCommitishOk returns a tuple with the TargetCommitish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCommitish

`func (o *ReposGenerateReleaseNotesRequest) SetTargetCommitish(v string)`

SetTargetCommitish sets TargetCommitish field to given value.

### HasTargetCommitish

`func (o *ReposGenerateReleaseNotesRequest) HasTargetCommitish() bool`

HasTargetCommitish returns a boolean if a field has been set.

### GetPreviousTagName

`func (o *ReposGenerateReleaseNotesRequest) GetPreviousTagName() string`

GetPreviousTagName returns the PreviousTagName field if non-nil, zero value otherwise.

### GetPreviousTagNameOk

`func (o *ReposGenerateReleaseNotesRequest) GetPreviousTagNameOk() (*string, bool)`

GetPreviousTagNameOk returns a tuple with the PreviousTagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousTagName

`func (o *ReposGenerateReleaseNotesRequest) SetPreviousTagName(v string)`

SetPreviousTagName sets PreviousTagName field to given value.

### HasPreviousTagName

`func (o *ReposGenerateReleaseNotesRequest) HasPreviousTagName() bool`

HasPreviousTagName returns a boolean if a field has been set.

### GetConfigurationFilePath

`func (o *ReposGenerateReleaseNotesRequest) GetConfigurationFilePath() string`

GetConfigurationFilePath returns the ConfigurationFilePath field if non-nil, zero value otherwise.

### GetConfigurationFilePathOk

`func (o *ReposGenerateReleaseNotesRequest) GetConfigurationFilePathOk() (*string, bool)`

GetConfigurationFilePathOk returns a tuple with the ConfigurationFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationFilePath

`func (o *ReposGenerateReleaseNotesRequest) SetConfigurationFilePath(v string)`

SetConfigurationFilePath sets ConfigurationFilePath field to given value.

### HasConfigurationFilePath

`func (o *ReposGenerateReleaseNotesRequest) HasConfigurationFilePath() bool`

HasConfigurationFilePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


