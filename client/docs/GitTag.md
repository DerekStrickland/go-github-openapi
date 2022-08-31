# GitTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | **string** |  | 
**Tag** | **string** | Name of the tag | 
**Sha** | **string** |  | 
**Url** | **string** | URL for the tag | 
**Message** | **string** | Message describing the purpose of the tag | 
**Tagger** | [**GitTagTagger**](GitTagTagger.md) |  | 
**Object** | [**GitTagObject**](GitTagObject.md) |  | 
**Verification** | Pointer to [**Verification**](Verification.md) |  | [optional] 

## Methods

### NewGitTag

`func NewGitTag(nodeId string, tag string, sha string, url string, message string, tagger GitTagTagger, object GitTagObject, ) *GitTag`

NewGitTag instantiates a new GitTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitTagWithDefaults

`func NewGitTagWithDefaults() *GitTag`

NewGitTagWithDefaults instantiates a new GitTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *GitTag) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *GitTag) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *GitTag) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetTag

`func (o *GitTag) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GitTag) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GitTag) SetTag(v string)`

SetTag sets Tag field to given value.


### GetSha

`func (o *GitTag) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *GitTag) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *GitTag) SetSha(v string)`

SetSha sets Sha field to given value.


### GetUrl

`func (o *GitTag) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GitTag) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GitTag) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMessage

`func (o *GitTag) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GitTag) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GitTag) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTagger

`func (o *GitTag) GetTagger() GitTagTagger`

GetTagger returns the Tagger field if non-nil, zero value otherwise.

### GetTaggerOk

`func (o *GitTag) GetTaggerOk() (*GitTagTagger, bool)`

GetTaggerOk returns a tuple with the Tagger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagger

`func (o *GitTag) SetTagger(v GitTagTagger)`

SetTagger sets Tagger field to given value.


### GetObject

`func (o *GitTag) GetObject() GitTagObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *GitTag) GetObjectOk() (*GitTagObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *GitTag) SetObject(v GitTagObject)`

SetObject sets Object field to given value.


### GetVerification

`func (o *GitTag) GetVerification() Verification`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *GitTag) GetVerificationOk() (*Verification, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *GitTag) SetVerification(v Verification)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *GitTag) HasVerification() bool`

HasVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


