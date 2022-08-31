# GitCreateRefRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | **string** | The name of the fully qualified reference (ie: &#x60;refs/heads/master&#x60;). If it doesn&#39;t start with &#39;refs&#39; and have at least two slashes, it will be rejected. | 
**Sha** | **string** | The SHA1 value for this reference. | 
**Key** | Pointer to **string** |  | [optional] 

## Methods

### NewGitCreateRefRequest

`func NewGitCreateRefRequest(ref string, sha string, ) *GitCreateRefRequest`

NewGitCreateRefRequest instantiates a new GitCreateRefRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitCreateRefRequestWithDefaults

`func NewGitCreateRefRequestWithDefaults() *GitCreateRefRequest`

NewGitCreateRefRequestWithDefaults instantiates a new GitCreateRefRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GitCreateRefRequest) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GitCreateRefRequest) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GitCreateRefRequest) SetRef(v string)`

SetRef sets Ref field to given value.


### GetSha

`func (o *GitCreateRefRequest) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *GitCreateRefRequest) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *GitCreateRefRequest) SetSha(v string)`

SetSha sets Sha field to given value.


### GetKey

`func (o *GitCreateRefRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GitCreateRefRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GitCreateRefRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GitCreateRefRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


