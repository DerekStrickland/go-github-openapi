# Release

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**HtmlUrl** | **string** |  | 
**AssetsUrl** | **string** |  | 
**UploadUrl** | **string** |  | 
**TarballUrl** | **NullableString** |  | 
**ZipballUrl** | **NullableString** |  | 
**Id** | **int32** |  | 
**NodeId** | **string** |  | 
**TagName** | **string** | The name of the tag. | 
**TargetCommitish** | **string** | Specifies the commitish value that determines where the Git tag is created from. | 
**Name** | **NullableString** |  | 
**Body** | Pointer to **NullableString** |  | [optional] 
**Draft** | **bool** | true to create a draft (unpublished) release, false to create a published one. | 
**Prerelease** | **bool** | Whether to identify the release as a prerelease or a full release. | 
**CreatedAt** | **time.Time** |  | 
**PublishedAt** | **NullableTime** |  | 
**Author** | [**SimpleUser**](SimpleUser.md) |  | 
**Assets** | [**[]ReleaseAsset**](ReleaseAsset.md) |  | 
**BodyHtml** | Pointer to **string** |  | [optional] 
**BodyText** | Pointer to **string** |  | [optional] 
**MentionsCount** | Pointer to **int32** |  | [optional] 
**DiscussionUrl** | Pointer to **string** | The URL of the release discussion. | [optional] 
**Reactions** | Pointer to [**ReactionRollup**](ReactionRollup.md) |  | [optional] 

## Methods

### NewRelease

`func NewRelease(url string, htmlUrl string, assetsUrl string, uploadUrl string, tarballUrl NullableString, zipballUrl NullableString, id int32, nodeId string, tagName string, targetCommitish string, name NullableString, draft bool, prerelease bool, createdAt time.Time, publishedAt NullableTime, author SimpleUser, assets []ReleaseAsset, ) *Release`

NewRelease instantiates a new Release object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReleaseWithDefaults

`func NewReleaseWithDefaults() *Release`

NewReleaseWithDefaults instantiates a new Release object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *Release) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Release) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Release) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHtmlUrl

`func (o *Release) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *Release) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *Release) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetAssetsUrl

`func (o *Release) GetAssetsUrl() string`

GetAssetsUrl returns the AssetsUrl field if non-nil, zero value otherwise.

### GetAssetsUrlOk

`func (o *Release) GetAssetsUrlOk() (*string, bool)`

GetAssetsUrlOk returns a tuple with the AssetsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetsUrl

`func (o *Release) SetAssetsUrl(v string)`

SetAssetsUrl sets AssetsUrl field to given value.


### GetUploadUrl

`func (o *Release) GetUploadUrl() string`

GetUploadUrl returns the UploadUrl field if non-nil, zero value otherwise.

### GetUploadUrlOk

`func (o *Release) GetUploadUrlOk() (*string, bool)`

GetUploadUrlOk returns a tuple with the UploadUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadUrl

`func (o *Release) SetUploadUrl(v string)`

SetUploadUrl sets UploadUrl field to given value.


### GetTarballUrl

`func (o *Release) GetTarballUrl() string`

GetTarballUrl returns the TarballUrl field if non-nil, zero value otherwise.

### GetTarballUrlOk

`func (o *Release) GetTarballUrlOk() (*string, bool)`

GetTarballUrlOk returns a tuple with the TarballUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarballUrl

`func (o *Release) SetTarballUrl(v string)`

SetTarballUrl sets TarballUrl field to given value.


### SetTarballUrlNil

`func (o *Release) SetTarballUrlNil(b bool)`

 SetTarballUrlNil sets the value for TarballUrl to be an explicit nil

### UnsetTarballUrl
`func (o *Release) UnsetTarballUrl()`

UnsetTarballUrl ensures that no value is present for TarballUrl, not even an explicit nil
### GetZipballUrl

`func (o *Release) GetZipballUrl() string`

GetZipballUrl returns the ZipballUrl field if non-nil, zero value otherwise.

### GetZipballUrlOk

`func (o *Release) GetZipballUrlOk() (*string, bool)`

GetZipballUrlOk returns a tuple with the ZipballUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZipballUrl

`func (o *Release) SetZipballUrl(v string)`

SetZipballUrl sets ZipballUrl field to given value.


### SetZipballUrlNil

`func (o *Release) SetZipballUrlNil(b bool)`

 SetZipballUrlNil sets the value for ZipballUrl to be an explicit nil

### UnsetZipballUrl
`func (o *Release) UnsetZipballUrl()`

UnsetZipballUrl ensures that no value is present for ZipballUrl, not even an explicit nil
### GetId

`func (o *Release) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Release) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Release) SetId(v int32)`

SetId sets Id field to given value.


### GetNodeId

`func (o *Release) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *Release) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *Release) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetTagName

`func (o *Release) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *Release) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *Release) SetTagName(v string)`

SetTagName sets TagName field to given value.


### GetTargetCommitish

`func (o *Release) GetTargetCommitish() string`

GetTargetCommitish returns the TargetCommitish field if non-nil, zero value otherwise.

### GetTargetCommitishOk

`func (o *Release) GetTargetCommitishOk() (*string, bool)`

GetTargetCommitishOk returns a tuple with the TargetCommitish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCommitish

`func (o *Release) SetTargetCommitish(v string)`

SetTargetCommitish sets TargetCommitish field to given value.


### GetName

`func (o *Release) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Release) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Release) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Release) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Release) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetBody

`func (o *Release) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Release) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Release) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *Release) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *Release) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *Release) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetDraft

`func (o *Release) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *Release) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *Release) SetDraft(v bool)`

SetDraft sets Draft field to given value.


### GetPrerelease

`func (o *Release) GetPrerelease() bool`

GetPrerelease returns the Prerelease field if non-nil, zero value otherwise.

### GetPrereleaseOk

`func (o *Release) GetPrereleaseOk() (*bool, bool)`

GetPrereleaseOk returns a tuple with the Prerelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrerelease

`func (o *Release) SetPrerelease(v bool)`

SetPrerelease sets Prerelease field to given value.


### GetCreatedAt

`func (o *Release) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Release) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Release) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetPublishedAt

`func (o *Release) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *Release) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *Release) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.


### SetPublishedAtNil

`func (o *Release) SetPublishedAtNil(b bool)`

 SetPublishedAtNil sets the value for PublishedAt to be an explicit nil

### UnsetPublishedAt
`func (o *Release) UnsetPublishedAt()`

UnsetPublishedAt ensures that no value is present for PublishedAt, not even an explicit nil
### GetAuthor

`func (o *Release) GetAuthor() SimpleUser`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Release) GetAuthorOk() (*SimpleUser, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Release) SetAuthor(v SimpleUser)`

SetAuthor sets Author field to given value.


### GetAssets

`func (o *Release) GetAssets() []ReleaseAsset`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *Release) GetAssetsOk() (*[]ReleaseAsset, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *Release) SetAssets(v []ReleaseAsset)`

SetAssets sets Assets field to given value.


### GetBodyHtml

`func (o *Release) GetBodyHtml() string`

GetBodyHtml returns the BodyHtml field if non-nil, zero value otherwise.

### GetBodyHtmlOk

`func (o *Release) GetBodyHtmlOk() (*string, bool)`

GetBodyHtmlOk returns a tuple with the BodyHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyHtml

`func (o *Release) SetBodyHtml(v string)`

SetBodyHtml sets BodyHtml field to given value.

### HasBodyHtml

`func (o *Release) HasBodyHtml() bool`

HasBodyHtml returns a boolean if a field has been set.

### GetBodyText

`func (o *Release) GetBodyText() string`

GetBodyText returns the BodyText field if non-nil, zero value otherwise.

### GetBodyTextOk

`func (o *Release) GetBodyTextOk() (*string, bool)`

GetBodyTextOk returns a tuple with the BodyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyText

`func (o *Release) SetBodyText(v string)`

SetBodyText sets BodyText field to given value.

### HasBodyText

`func (o *Release) HasBodyText() bool`

HasBodyText returns a boolean if a field has been set.

### GetMentionsCount

`func (o *Release) GetMentionsCount() int32`

GetMentionsCount returns the MentionsCount field if non-nil, zero value otherwise.

### GetMentionsCountOk

`func (o *Release) GetMentionsCountOk() (*int32, bool)`

GetMentionsCountOk returns a tuple with the MentionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionsCount

`func (o *Release) SetMentionsCount(v int32)`

SetMentionsCount sets MentionsCount field to given value.

### HasMentionsCount

`func (o *Release) HasMentionsCount() bool`

HasMentionsCount returns a boolean if a field has been set.

### GetDiscussionUrl

`func (o *Release) GetDiscussionUrl() string`

GetDiscussionUrl returns the DiscussionUrl field if non-nil, zero value otherwise.

### GetDiscussionUrlOk

`func (o *Release) GetDiscussionUrlOk() (*string, bool)`

GetDiscussionUrlOk returns a tuple with the DiscussionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscussionUrl

`func (o *Release) SetDiscussionUrl(v string)`

SetDiscussionUrl sets DiscussionUrl field to given value.

### HasDiscussionUrl

`func (o *Release) HasDiscussionUrl() bool`

HasDiscussionUrl returns a boolean if a field has been set.

### GetReactions

`func (o *Release) GetReactions() ReactionRollup`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *Release) GetReactionsOk() (*ReactionRollup, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *Release) SetReactions(v ReactionRollup)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *Release) HasReactions() bool`

HasReactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


