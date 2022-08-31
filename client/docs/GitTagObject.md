# GitTagObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sha** | **string** |  | 
**Type** | **string** |  | 
**Url** | **string** |  | 

## Methods

### NewGitTagObject

`func NewGitTagObject(sha string, type_ string, url string, ) *GitTagObject`

NewGitTagObject instantiates a new GitTagObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitTagObjectWithDefaults

`func NewGitTagObjectWithDefaults() *GitTagObject`

NewGitTagObjectWithDefaults instantiates a new GitTagObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSha

`func (o *GitTagObject) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *GitTagObject) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *GitTagObject) SetSha(v string)`

SetSha sets Sha field to given value.


### GetType

`func (o *GitTagObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GitTagObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GitTagObject) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *GitTagObject) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GitTagObject) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GitTagObject) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


