# ReposCreatePagesSiteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildType** | Pointer to **string** | The process in which the Page will be built. Possible values are &#x60;\&quot;legacy\&quot;&#x60; and &#x60;\&quot;workflow\&quot;&#x60;. | [optional] 
**Source** | Pointer to [**ReposCreatePagesSiteRequestSource**](ReposCreatePagesSiteRequestSource.md) |  | [optional] 

## Methods

### NewReposCreatePagesSiteRequest

`func NewReposCreatePagesSiteRequest() *ReposCreatePagesSiteRequest`

NewReposCreatePagesSiteRequest instantiates a new ReposCreatePagesSiteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposCreatePagesSiteRequestWithDefaults

`func NewReposCreatePagesSiteRequestWithDefaults() *ReposCreatePagesSiteRequest`

NewReposCreatePagesSiteRequestWithDefaults instantiates a new ReposCreatePagesSiteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildType

`func (o *ReposCreatePagesSiteRequest) GetBuildType() string`

GetBuildType returns the BuildType field if non-nil, zero value otherwise.

### GetBuildTypeOk

`func (o *ReposCreatePagesSiteRequest) GetBuildTypeOk() (*string, bool)`

GetBuildTypeOk returns a tuple with the BuildType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildType

`func (o *ReposCreatePagesSiteRequest) SetBuildType(v string)`

SetBuildType sets BuildType field to given value.

### HasBuildType

`func (o *ReposCreatePagesSiteRequest) HasBuildType() bool`

HasBuildType returns a boolean if a field has been set.

### GetSource

`func (o *ReposCreatePagesSiteRequest) GetSource() ReposCreatePagesSiteRequestSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ReposCreatePagesSiteRequest) GetSourceOk() (*ReposCreatePagesSiteRequestSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ReposCreatePagesSiteRequest) SetSource(v ReposCreatePagesSiteRequestSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *ReposCreatePagesSiteRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


