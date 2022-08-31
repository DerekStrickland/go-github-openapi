# GitCreateTagRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tag** | **string** | The tag&#39;s name. This is typically a version (e.g., \&quot;v0.0.1\&quot;). | 
**Message** | **string** | The tag message. | 
**Object** | **string** | The SHA of the git object this is tagging. | 
**Type** | **string** | The type of the object we&#39;re tagging. Normally this is a &#x60;commit&#x60; but it can also be a &#x60;tree&#x60; or a &#x60;blob&#x60;. | 
**Tagger** | Pointer to [**GitCreateTagRequestTagger**](GitCreateTagRequestTagger.md) |  | [optional] 

## Methods

### NewGitCreateTagRequest

`func NewGitCreateTagRequest(tag string, message string, object string, type_ string, ) *GitCreateTagRequest`

NewGitCreateTagRequest instantiates a new GitCreateTagRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitCreateTagRequestWithDefaults

`func NewGitCreateTagRequestWithDefaults() *GitCreateTagRequest`

NewGitCreateTagRequestWithDefaults instantiates a new GitCreateTagRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTag

`func (o *GitCreateTagRequest) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *GitCreateTagRequest) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *GitCreateTagRequest) SetTag(v string)`

SetTag sets Tag field to given value.


### GetMessage

`func (o *GitCreateTagRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GitCreateTagRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GitCreateTagRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetObject

`func (o *GitCreateTagRequest) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *GitCreateTagRequest) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *GitCreateTagRequest) SetObject(v string)`

SetObject sets Object field to given value.


### GetType

`func (o *GitCreateTagRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GitCreateTagRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GitCreateTagRequest) SetType(v string)`

SetType sets Type field to given value.


### GetTagger

`func (o *GitCreateTagRequest) GetTagger() GitCreateTagRequestTagger`

GetTagger returns the Tagger field if non-nil, zero value otherwise.

### GetTaggerOk

`func (o *GitCreateTagRequest) GetTaggerOk() (*GitCreateTagRequestTagger, bool)`

GetTaggerOk returns a tuple with the Tagger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagger

`func (o *GitCreateTagRequest) SetTagger(v GitCreateTagRequestTagger)`

SetTagger sets Tagger field to given value.

### HasTagger

`func (o *GitCreateTagRequest) HasTagger() bool`

HasTagger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


