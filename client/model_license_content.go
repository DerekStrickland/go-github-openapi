/*
GitHub v3 REST API

GitHub's v3 REST API.

API version: 1.1.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package github

import (
	"encoding/json"
)

// LicenseContent License Content
type LicenseContent struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Sha string `json:"sha"`
	Size int32 `json:"size"`
	Url string `json:"url"`
	HtmlUrl NullableString `json:"html_url"`
	GitUrl NullableString `json:"git_url"`
	DownloadUrl NullableString `json:"download_url"`
	Type string `json:"type"`
	Content string `json:"content"`
	Encoding string `json:"encoding"`
	Links ContentTreeEntriesInnerLinks `json:"_links"`
	License NullableNullableLicenseSimple `json:"license"`
}

// NewLicenseContent instantiates a new LicenseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseContent(name string, path string, sha string, size int32, url string, htmlUrl NullableString, gitUrl NullableString, downloadUrl NullableString, type_ string, content string, encoding string, links ContentTreeEntriesInnerLinks, license NullableNullableLicenseSimple) *LicenseContent {
	this := LicenseContent{}
	this.Name = name
	this.Path = path
	this.Sha = sha
	this.Size = size
	this.Url = url
	this.HtmlUrl = htmlUrl
	this.GitUrl = gitUrl
	this.DownloadUrl = downloadUrl
	this.Type = type_
	this.Content = content
	this.Encoding = encoding
	this.Links = links
	this.License = license
	return &this
}

// NewLicenseContentWithDefaults instantiates a new LicenseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseContentWithDefaults() *LicenseContent {
	this := LicenseContent{}
	return &this
}

// GetName returns the Name field value
func (o *LicenseContent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LicenseContent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LicenseContent) SetName(v string) {
	o.Name = v
}

// GetPath returns the Path field value
func (o *LicenseContent) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *LicenseContent) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *LicenseContent) SetPath(v string) {
	o.Path = v
}

// GetSha returns the Sha field value
func (o *LicenseContent) GetSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha
}

// GetShaOk returns a tuple with the Sha field value
// and a boolean to check if the value has been set.
func (o *LicenseContent) GetShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sha, true
}

// SetSha sets field value
func (o *LicenseContent) SetSha(v string) {
	o.Sha = v
}

// GetSize returns the Size field value
func (o *LicenseContent) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *LicenseContent) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *LicenseContent) SetSize(v int32) {
	o.Size = v
}

// GetUrl returns the Url field value
func (o *LicenseContent) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *LicenseContent) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *LicenseContent) SetUrl(v string) {
	o.Url = v
}

// GetHtmlUrl returns the HtmlUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LicenseContent) GetHtmlUrl() string {
	if o == nil || o.HtmlUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.HtmlUrl.Get()
}

// GetHtmlUrlOk returns a tuple with the HtmlUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicenseContent) GetHtmlUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HtmlUrl.Get(), o.HtmlUrl.IsSet()
}

// SetHtmlUrl sets field value
func (o *LicenseContent) SetHtmlUrl(v string) {
	o.HtmlUrl.Set(&v)
}

// GetGitUrl returns the GitUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LicenseContent) GetGitUrl() string {
	if o == nil || o.GitUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.GitUrl.Get()
}

// GetGitUrlOk returns a tuple with the GitUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicenseContent) GetGitUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GitUrl.Get(), o.GitUrl.IsSet()
}

// SetGitUrl sets field value
func (o *LicenseContent) SetGitUrl(v string) {
	o.GitUrl.Set(&v)
}

// GetDownloadUrl returns the DownloadUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *LicenseContent) GetDownloadUrl() string {
	if o == nil || o.DownloadUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.DownloadUrl.Get()
}

// GetDownloadUrlOk returns a tuple with the DownloadUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicenseContent) GetDownloadUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DownloadUrl.Get(), o.DownloadUrl.IsSet()
}

// SetDownloadUrl sets field value
func (o *LicenseContent) SetDownloadUrl(v string) {
	o.DownloadUrl.Set(&v)
}

// GetType returns the Type field value
func (o *LicenseContent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LicenseContent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LicenseContent) SetType(v string) {
	o.Type = v
}

// GetContent returns the Content field value
func (o *LicenseContent) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *LicenseContent) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *LicenseContent) SetContent(v string) {
	o.Content = v
}

// GetEncoding returns the Encoding field value
func (o *LicenseContent) GetEncoding() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value
// and a boolean to check if the value has been set.
func (o *LicenseContent) GetEncodingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Encoding, true
}

// SetEncoding sets field value
func (o *LicenseContent) SetEncoding(v string) {
	o.Encoding = v
}

// GetLinks returns the Links field value
func (o *LicenseContent) GetLinks() ContentTreeEntriesInnerLinks {
	if o == nil {
		var ret ContentTreeEntriesInnerLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *LicenseContent) GetLinksOk() (*ContentTreeEntriesInnerLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *LicenseContent) SetLinks(v ContentTreeEntriesInnerLinks) {
	o.Links = v
}

// GetLicense returns the License field value
// If the value is explicit nil, the zero value for NullableLicenseSimple will be returned
func (o *LicenseContent) GetLicense() NullableLicenseSimple {
	if o == nil || o.License.Get() == nil {
		var ret NullableLicenseSimple
		return ret
	}

	return *o.License.Get()
}

// GetLicenseOk returns a tuple with the License field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LicenseContent) GetLicenseOk() (*NullableLicenseSimple, bool) {
	if o == nil {
		return nil, false
	}
	return o.License.Get(), o.License.IsSet()
}

// SetLicense sets field value
func (o *LicenseContent) SetLicense(v NullableLicenseSimple) {
	o.License.Set(&v)
}

func (o LicenseContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["sha"] = o.Sha
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["url"] = o.Url
	}
	if true {
		toSerialize["html_url"] = o.HtmlUrl.Get()
	}
	if true {
		toSerialize["git_url"] = o.GitUrl.Get()
	}
	if true {
		toSerialize["download_url"] = o.DownloadUrl.Get()
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["content"] = o.Content
	}
	if true {
		toSerialize["encoding"] = o.Encoding
	}
	if true {
		toSerialize["_links"] = o.Links
	}
	if true {
		toSerialize["license"] = o.License.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableLicenseContent struct {
	value *LicenseContent
	isSet bool
}

func (v NullableLicenseContent) Get() *LicenseContent {
	return v.value
}

func (v *NullableLicenseContent) Set(val *LicenseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseContent(val *LicenseContent) *NullableLicenseContent {
	return &NullableLicenseContent{value: val, isSet: true}
}

func (v NullableLicenseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


