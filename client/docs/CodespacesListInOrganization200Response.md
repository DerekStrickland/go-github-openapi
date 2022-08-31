# CodespacesListInOrganization200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**Codespaces** | [**[]Codespace**](Codespace.md) |  | 

## Methods

### NewCodespacesListInOrganization200Response

`func NewCodespacesListInOrganization200Response(totalCount int32, codespaces []Codespace, ) *CodespacesListInOrganization200Response`

NewCodespacesListInOrganization200Response instantiates a new CodespacesListInOrganization200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodespacesListInOrganization200ResponseWithDefaults

`func NewCodespacesListInOrganization200ResponseWithDefaults() *CodespacesListInOrganization200Response`

NewCodespacesListInOrganization200ResponseWithDefaults instantiates a new CodespacesListInOrganization200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *CodespacesListInOrganization200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CodespacesListInOrganization200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CodespacesListInOrganization200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetCodespaces

`func (o *CodespacesListInOrganization200Response) GetCodespaces() []Codespace`

GetCodespaces returns the Codespaces field if non-nil, zero value otherwise.

### GetCodespacesOk

`func (o *CodespacesListInOrganization200Response) GetCodespacesOk() (*[]Codespace, bool)`

GetCodespacesOk returns a tuple with the Codespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodespaces

`func (o *CodespacesListInOrganization200Response) SetCodespaces(v []Codespace)`

SetCodespaces sets Codespaces field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


