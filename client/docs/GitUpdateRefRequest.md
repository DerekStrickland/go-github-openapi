# GitUpdateRefRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sha** | **string** | The SHA1 value to set this reference to | 
**Force** | Pointer to **bool** | Indicates whether to force the update or to make sure the update is a fast-forward update. Leaving this out or setting it to &#x60;false&#x60; will make sure you&#39;re not overwriting work. | [optional] [default to false]

## Methods

### NewGitUpdateRefRequest

`func NewGitUpdateRefRequest(sha string, ) *GitUpdateRefRequest`

NewGitUpdateRefRequest instantiates a new GitUpdateRefRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitUpdateRefRequestWithDefaults

`func NewGitUpdateRefRequestWithDefaults() *GitUpdateRefRequest`

NewGitUpdateRefRequestWithDefaults instantiates a new GitUpdateRefRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSha

`func (o *GitUpdateRefRequest) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *GitUpdateRefRequest) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *GitUpdateRefRequest) SetSha(v string)`

SetSha sets Sha field to given value.


### GetForce

`func (o *GitUpdateRefRequest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *GitUpdateRefRequest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *GitUpdateRefRequest) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *GitUpdateRefRequest) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


