# ReposUpdateInformationAboutPagesSiteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cname** | Pointer to **NullableString** | Specify a custom domain for the repository. Sending a &#x60;null&#x60; value will remove the custom domain. For more about custom domains, see \&quot;[Using a custom domain with GitHub Pages](https://docs.github.com/articles/using-a-custom-domain-with-github-pages/).\&quot; | [optional] 
**HttpsEnforced** | Pointer to **bool** | Specify whether HTTPS should be enforced for the repository. | [optional] 
**Public** | Pointer to **bool** | Configures access controls for the GitHub Pages site. If public is set to &#x60;true&#x60;, the site is accessible to anyone on the internet. If set to &#x60;false&#x60;, the site will only be accessible to users who have at least &#x60;read&#x60; access to the repository that published the site. This includes anyone in your Enterprise if the repository is set to &#x60;internal&#x60; visibility. This feature is only available to repositories in an organization on an Enterprise plan. | [optional] 
**BuildType** | Pointer to **string** | The process by which the GitHub Pages site will be built. &#x60;workflow&#x60; means that the site is built by a custom GitHub Actions workflow. &#x60;legacy&#x60; means that the site is built by GitHub when changes are pushed to a specific branch. | [optional] 
**Source** | Pointer to [**ReposUpdateInformationAboutPagesSiteRequestSource**](ReposUpdateInformationAboutPagesSiteRequestSource.md) |  | [optional] 

## Methods

### NewReposUpdateInformationAboutPagesSiteRequest

`func NewReposUpdateInformationAboutPagesSiteRequest() *ReposUpdateInformationAboutPagesSiteRequest`

NewReposUpdateInformationAboutPagesSiteRequest instantiates a new ReposUpdateInformationAboutPagesSiteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposUpdateInformationAboutPagesSiteRequestWithDefaults

`func NewReposUpdateInformationAboutPagesSiteRequestWithDefaults() *ReposUpdateInformationAboutPagesSiteRequest`

NewReposUpdateInformationAboutPagesSiteRequestWithDefaults instantiates a new ReposUpdateInformationAboutPagesSiteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCname

`func (o *ReposUpdateInformationAboutPagesSiteRequest) GetCname() string`

GetCname returns the Cname field if non-nil, zero value otherwise.

### GetCnameOk

`func (o *ReposUpdateInformationAboutPagesSiteRequest) GetCnameOk() (*string, bool)`

GetCnameOk returns a tuple with the Cname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCname

`func (o *ReposUpdateInformationAboutPagesSiteRequest) SetCname(v string)`

SetCname sets Cname field to given value.

### HasCname

`func (o *ReposUpdateInformationAboutPagesSiteRequest) HasCname() bool`

HasCname returns a boolean if a field has been set.

### SetCnameNil

`func (o *ReposUpdateInformationAboutPagesSiteRequest) SetCnameNil(b bool)`

 SetCnameNil sets the value for Cname to be an explicit nil

### UnsetCname
`func (o *ReposUpdateInformationAboutPagesSiteRequest) UnsetCname()`

UnsetCname ensures that no value is present for Cname, not even an explicit nil
### GetHttpsEnforced

`func (o *ReposUpdateInformationAboutPagesSiteRequest) GetHttpsEnforced() bool`

GetHttpsEnforced returns the HttpsEnforced field if non-nil, zero value otherwise.

### GetHttpsEnforcedOk

`func (o *ReposUpdateInformationAboutPagesSiteRequest) GetHttpsEnforcedOk() (*bool, bool)`

GetHttpsEnforcedOk returns a tuple with the HttpsEnforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsEnforced

`func (o *ReposUpdateInformationAboutPagesSiteRequest) SetHttpsEnforced(v bool)`

SetHttpsEnforced sets HttpsEnforced field to given value.

### HasHttpsEnforced

`func (o *ReposUpdateInformationAboutPagesSiteRequest) HasHttpsEnforced() bool`

HasHttpsEnforced returns a boolean if a field has been set.

### GetPublic

`func (o *ReposUpdateInformationAboutPagesSiteRequest) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *ReposUpdateInformationAboutPagesSiteRequest) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *ReposUpdateInformationAboutPagesSiteRequest) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *ReposUpdateInformationAboutPagesSiteRequest) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetBuildType

`func (o *ReposUpdateInformationAboutPagesSiteRequest) GetBuildType() string`

GetBuildType returns the BuildType field if non-nil, zero value otherwise.

### GetBuildTypeOk

`func (o *ReposUpdateInformationAboutPagesSiteRequest) GetBuildTypeOk() (*string, bool)`

GetBuildTypeOk returns a tuple with the BuildType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildType

`func (o *ReposUpdateInformationAboutPagesSiteRequest) SetBuildType(v string)`

SetBuildType sets BuildType field to given value.

### HasBuildType

`func (o *ReposUpdateInformationAboutPagesSiteRequest) HasBuildType() bool`

HasBuildType returns a boolean if a field has been set.

### GetSource

`func (o *ReposUpdateInformationAboutPagesSiteRequest) GetSource() ReposUpdateInformationAboutPagesSiteRequestSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ReposUpdateInformationAboutPagesSiteRequest) GetSourceOk() (*ReposUpdateInformationAboutPagesSiteRequestSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ReposUpdateInformationAboutPagesSiteRequest) SetSource(v ReposUpdateInformationAboutPagesSiteRequestSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *ReposUpdateInformationAboutPagesSiteRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


