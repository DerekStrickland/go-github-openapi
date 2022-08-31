# PullRequestMinimal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Number** | **int32** |  | 
**Url** | **string** |  | 
**Head** | [**PullRequestMinimalHead**](PullRequestMinimalHead.md) |  | 
**Base** | [**PullRequestMinimalHead**](PullRequestMinimalHead.md) |  | 

## Methods

### NewPullRequestMinimal

`func NewPullRequestMinimal(id int32, number int32, url string, head PullRequestMinimalHead, base PullRequestMinimalHead, ) *PullRequestMinimal`

NewPullRequestMinimal instantiates a new PullRequestMinimal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestMinimalWithDefaults

`func NewPullRequestMinimalWithDefaults() *PullRequestMinimal`

NewPullRequestMinimalWithDefaults instantiates a new PullRequestMinimal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PullRequestMinimal) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PullRequestMinimal) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PullRequestMinimal) SetId(v int32)`

SetId sets Id field to given value.


### GetNumber

`func (o *PullRequestMinimal) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *PullRequestMinimal) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *PullRequestMinimal) SetNumber(v int32)`

SetNumber sets Number field to given value.


### GetUrl

`func (o *PullRequestMinimal) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PullRequestMinimal) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PullRequestMinimal) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetHead

`func (o *PullRequestMinimal) GetHead() PullRequestMinimalHead`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *PullRequestMinimal) GetHeadOk() (*PullRequestMinimalHead, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *PullRequestMinimal) SetHead(v PullRequestMinimalHead)`

SetHead sets Head field to given value.


### GetBase

`func (o *PullRequestMinimal) GetBase() PullRequestMinimalHead`

GetBase returns the Base field if non-nil, zero value otherwise.

### GetBaseOk

`func (o *PullRequestMinimal) GetBaseOk() (*PullRequestMinimalHead, bool)`

GetBaseOk returns a tuple with the Base field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase

`func (o *PullRequestMinimal) SetBase(v PullRequestMinimalHead)`

SetBase sets Base field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


