# CombinedCommitStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** |  | 
**Statuses** | [**[]SimpleCommitStatus**](SimpleCommitStatus.md) |  | 
**Sha** | **string** |  | 
**TotalCount** | **int32** |  | 
**Repository** | [**MinimalRepository**](MinimalRepository.md) |  | 
**CommitUrl** | **string** |  | 
**Url** | **string** |  | 

## Methods

### NewCombinedCommitStatus

`func NewCombinedCommitStatus(state string, statuses []SimpleCommitStatus, sha string, totalCount int32, repository MinimalRepository, commitUrl string, url string, ) *CombinedCommitStatus`

NewCombinedCommitStatus instantiates a new CombinedCommitStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCombinedCommitStatusWithDefaults

`func NewCombinedCommitStatusWithDefaults() *CombinedCommitStatus`

NewCombinedCommitStatusWithDefaults instantiates a new CombinedCommitStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *CombinedCommitStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CombinedCommitStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CombinedCommitStatus) SetState(v string)`

SetState sets State field to given value.


### GetStatuses

`func (o *CombinedCommitStatus) GetStatuses() []SimpleCommitStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *CombinedCommitStatus) GetStatusesOk() (*[]SimpleCommitStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *CombinedCommitStatus) SetStatuses(v []SimpleCommitStatus)`

SetStatuses sets Statuses field to given value.


### GetSha

`func (o *CombinedCommitStatus) GetSha() string`

GetSha returns the Sha field if non-nil, zero value otherwise.

### GetShaOk

`func (o *CombinedCommitStatus) GetShaOk() (*string, bool)`

GetShaOk returns a tuple with the Sha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha

`func (o *CombinedCommitStatus) SetSha(v string)`

SetSha sets Sha field to given value.


### GetTotalCount

`func (o *CombinedCommitStatus) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CombinedCommitStatus) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CombinedCommitStatus) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetRepository

`func (o *CombinedCommitStatus) GetRepository() MinimalRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *CombinedCommitStatus) GetRepositoryOk() (*MinimalRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *CombinedCommitStatus) SetRepository(v MinimalRepository)`

SetRepository sets Repository field to given value.


### GetCommitUrl

`func (o *CombinedCommitStatus) GetCommitUrl() string`

GetCommitUrl returns the CommitUrl field if non-nil, zero value otherwise.

### GetCommitUrlOk

`func (o *CombinedCommitStatus) GetCommitUrlOk() (*string, bool)`

GetCommitUrlOk returns a tuple with the CommitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitUrl

`func (o *CombinedCommitStatus) SetCommitUrl(v string)`

SetCommitUrl sets CommitUrl field to given value.


### GetUrl

`func (o *CombinedCommitStatus) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CombinedCommitStatus) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CombinedCommitStatus) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


