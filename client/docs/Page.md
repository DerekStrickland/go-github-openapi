# Page

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The API address for accessing this Page resource. | 
**Status** | **NullableString** | The status of the most recent build of the Page. | 
**Cname** | **NullableString** | The Pages site&#39;s custom domain | 
**ProtectedDomainState** | Pointer to **NullableString** | The state if the domain is verified | [optional] 
**PendingDomainUnverifiedAt** | Pointer to **NullableTime** | The timestamp when a pending domain becomes unverified. | [optional] 
**Custom404** | **bool** | Whether the Page has a custom 404 page. | [default to false]
**HtmlUrl** | Pointer to **string** | The web address the Page can be accessed from. | [optional] 
**BuildType** | Pointer to **NullableString** | The process in which the Page will be built. | [optional] 
**Source** | Pointer to [**PagesSourceHash**](PagesSourceHash.md) |  | [optional] 
**Public** | **bool** | Whether the GitHub Pages site is publicly visible. If set to &#x60;true&#x60;, the site is accessible to anyone on the internet. If set to &#x60;false&#x60;, the site will only be accessible to users who have at least &#x60;read&#x60; access to the repository that published the site. | 
**HttpsCertificate** | Pointer to [**PagesHttpsCertificate**](PagesHttpsCertificate.md) |  | [optional] 
**HttpsEnforced** | Pointer to **bool** | Whether https is enabled on the domain | [optional] 

## Methods

### NewPage

`func NewPage(url string, status NullableString, cname NullableString, custom404 bool, public bool, ) *Page`

NewPage instantiates a new Page object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageWithDefaults

`func NewPageWithDefaults() *Page`

NewPageWithDefaults instantiates a new Page object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Page) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Page) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Page) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetStatus

`func (o *Page) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Page) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Page) SetStatus(v string)`

SetStatus sets Status field to given value.


### SetStatusNil

`func (o *Page) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *Page) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetCname

`func (o *Page) GetCname() string`

GetCname returns the Cname field if non-nil, zero value otherwise.

### GetCnameOk

`func (o *Page) GetCnameOk() (*string, bool)`

GetCnameOk returns a tuple with the Cname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCname

`func (o *Page) SetCname(v string)`

SetCname sets Cname field to given value.


### SetCnameNil

`func (o *Page) SetCnameNil(b bool)`

 SetCnameNil sets the value for Cname to be an explicit nil

### UnsetCname
`func (o *Page) UnsetCname()`

UnsetCname ensures that no value is present for Cname, not even an explicit nil
### GetProtectedDomainState

`func (o *Page) GetProtectedDomainState() string`

GetProtectedDomainState returns the ProtectedDomainState field if non-nil, zero value otherwise.

### GetProtectedDomainStateOk

`func (o *Page) GetProtectedDomainStateOk() (*string, bool)`

GetProtectedDomainStateOk returns a tuple with the ProtectedDomainState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedDomainState

`func (o *Page) SetProtectedDomainState(v string)`

SetProtectedDomainState sets ProtectedDomainState field to given value.

### HasProtectedDomainState

`func (o *Page) HasProtectedDomainState() bool`

HasProtectedDomainState returns a boolean if a field has been set.

### SetProtectedDomainStateNil

`func (o *Page) SetProtectedDomainStateNil(b bool)`

 SetProtectedDomainStateNil sets the value for ProtectedDomainState to be an explicit nil

### UnsetProtectedDomainState
`func (o *Page) UnsetProtectedDomainState()`

UnsetProtectedDomainState ensures that no value is present for ProtectedDomainState, not even an explicit nil
### GetPendingDomainUnverifiedAt

`func (o *Page) GetPendingDomainUnverifiedAt() time.Time`

GetPendingDomainUnverifiedAt returns the PendingDomainUnverifiedAt field if non-nil, zero value otherwise.

### GetPendingDomainUnverifiedAtOk

`func (o *Page) GetPendingDomainUnverifiedAtOk() (*time.Time, bool)`

GetPendingDomainUnverifiedAtOk returns a tuple with the PendingDomainUnverifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDomainUnverifiedAt

`func (o *Page) SetPendingDomainUnverifiedAt(v time.Time)`

SetPendingDomainUnverifiedAt sets PendingDomainUnverifiedAt field to given value.

### HasPendingDomainUnverifiedAt

`func (o *Page) HasPendingDomainUnverifiedAt() bool`

HasPendingDomainUnverifiedAt returns a boolean if a field has been set.

### SetPendingDomainUnverifiedAtNil

`func (o *Page) SetPendingDomainUnverifiedAtNil(b bool)`

 SetPendingDomainUnverifiedAtNil sets the value for PendingDomainUnverifiedAt to be an explicit nil

### UnsetPendingDomainUnverifiedAt
`func (o *Page) UnsetPendingDomainUnverifiedAt()`

UnsetPendingDomainUnverifiedAt ensures that no value is present for PendingDomainUnverifiedAt, not even an explicit nil
### GetCustom404

`func (o *Page) GetCustom404() bool`

GetCustom404 returns the Custom404 field if non-nil, zero value otherwise.

### GetCustom404Ok

`func (o *Page) GetCustom404Ok() (*bool, bool)`

GetCustom404Ok returns a tuple with the Custom404 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom404

`func (o *Page) SetCustom404(v bool)`

SetCustom404 sets Custom404 field to given value.


### GetHtmlUrl

`func (o *Page) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *Page) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *Page) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.

### HasHtmlUrl

`func (o *Page) HasHtmlUrl() bool`

HasHtmlUrl returns a boolean if a field has been set.

### GetBuildType

`func (o *Page) GetBuildType() string`

GetBuildType returns the BuildType field if non-nil, zero value otherwise.

### GetBuildTypeOk

`func (o *Page) GetBuildTypeOk() (*string, bool)`

GetBuildTypeOk returns a tuple with the BuildType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildType

`func (o *Page) SetBuildType(v string)`

SetBuildType sets BuildType field to given value.

### HasBuildType

`func (o *Page) HasBuildType() bool`

HasBuildType returns a boolean if a field has been set.

### SetBuildTypeNil

`func (o *Page) SetBuildTypeNil(b bool)`

 SetBuildTypeNil sets the value for BuildType to be an explicit nil

### UnsetBuildType
`func (o *Page) UnsetBuildType()`

UnsetBuildType ensures that no value is present for BuildType, not even an explicit nil
### GetSource

`func (o *Page) GetSource() PagesSourceHash`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Page) GetSourceOk() (*PagesSourceHash, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Page) SetSource(v PagesSourceHash)`

SetSource sets Source field to given value.

### HasSource

`func (o *Page) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPublic

`func (o *Page) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *Page) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *Page) SetPublic(v bool)`

SetPublic sets Public field to given value.


### GetHttpsCertificate

`func (o *Page) GetHttpsCertificate() PagesHttpsCertificate`

GetHttpsCertificate returns the HttpsCertificate field if non-nil, zero value otherwise.

### GetHttpsCertificateOk

`func (o *Page) GetHttpsCertificateOk() (*PagesHttpsCertificate, bool)`

GetHttpsCertificateOk returns a tuple with the HttpsCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsCertificate

`func (o *Page) SetHttpsCertificate(v PagesHttpsCertificate)`

SetHttpsCertificate sets HttpsCertificate field to given value.

### HasHttpsCertificate

`func (o *Page) HasHttpsCertificate() bool`

HasHttpsCertificate returns a boolean if a field has been set.

### GetHttpsEnforced

`func (o *Page) GetHttpsEnforced() bool`

GetHttpsEnforced returns the HttpsEnforced field if non-nil, zero value otherwise.

### GetHttpsEnforcedOk

`func (o *Page) GetHttpsEnforcedOk() (*bool, bool)`

GetHttpsEnforcedOk returns a tuple with the HttpsEnforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpsEnforced

`func (o *Page) SetHttpsEnforced(v bool)`

SetHttpsEnforced sets HttpsEnforced field to given value.

### HasHttpsEnforced

`func (o *Page) HasHttpsEnforced() bool`

HasHttpsEnforced returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


