# ReposCreatePagesSiteRequestSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | **string** | The repository branch used to publish your site&#39;s source files. | 
**Path** | Pointer to **string** | The repository directory that includes the source files for the Pages site. Allowed paths are &#x60;/&#x60; or &#x60;/docs&#x60;. Default: &#x60;/&#x60; | [optional] [default to "/"]

## Methods

### NewReposCreatePagesSiteRequestSource

`func NewReposCreatePagesSiteRequestSource(branch string, ) *ReposCreatePagesSiteRequestSource`

NewReposCreatePagesSiteRequestSource instantiates a new ReposCreatePagesSiteRequestSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposCreatePagesSiteRequestSourceWithDefaults

`func NewReposCreatePagesSiteRequestSourceWithDefaults() *ReposCreatePagesSiteRequestSource`

NewReposCreatePagesSiteRequestSourceWithDefaults instantiates a new ReposCreatePagesSiteRequestSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *ReposCreatePagesSiteRequestSource) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ReposCreatePagesSiteRequestSource) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ReposCreatePagesSiteRequestSource) SetBranch(v string)`

SetBranch sets Branch field to given value.


### GetPath

`func (o *ReposCreatePagesSiteRequestSource) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ReposCreatePagesSiteRequestSource) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ReposCreatePagesSiteRequestSource) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ReposCreatePagesSiteRequestSource) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


