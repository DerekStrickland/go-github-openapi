# ScimUserList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schemas** | **[]string** | SCIM schema used. | 
**TotalResults** | **int32** |  | 
**ItemsPerPage** | **int32** |  | 
**StartIndex** | **int32** |  | 
**Resources** | [**[]ScimUser**](ScimUser.md) |  | 

## Methods

### NewScimUserList

`func NewScimUserList(schemas []string, totalResults int32, itemsPerPage int32, startIndex int32, resources []ScimUser, ) *ScimUserList`

NewScimUserList instantiates a new ScimUserList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScimUserListWithDefaults

`func NewScimUserListWithDefaults() *ScimUserList`

NewScimUserListWithDefaults instantiates a new ScimUserList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchemas

`func (o *ScimUserList) GetSchemas() []string`

GetSchemas returns the Schemas field if non-nil, zero value otherwise.

### GetSchemasOk

`func (o *ScimUserList) GetSchemasOk() (*[]string, bool)`

GetSchemasOk returns a tuple with the Schemas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemas

`func (o *ScimUserList) SetSchemas(v []string)`

SetSchemas sets Schemas field to given value.


### GetTotalResults

`func (o *ScimUserList) GetTotalResults() int32`

GetTotalResults returns the TotalResults field if non-nil, zero value otherwise.

### GetTotalResultsOk

`func (o *ScimUserList) GetTotalResultsOk() (*int32, bool)`

GetTotalResultsOk returns a tuple with the TotalResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResults

`func (o *ScimUserList) SetTotalResults(v int32)`

SetTotalResults sets TotalResults field to given value.


### GetItemsPerPage

`func (o *ScimUserList) GetItemsPerPage() int32`

GetItemsPerPage returns the ItemsPerPage field if non-nil, zero value otherwise.

### GetItemsPerPageOk

`func (o *ScimUserList) GetItemsPerPageOk() (*int32, bool)`

GetItemsPerPageOk returns a tuple with the ItemsPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsPerPage

`func (o *ScimUserList) SetItemsPerPage(v int32)`

SetItemsPerPage sets ItemsPerPage field to given value.


### GetStartIndex

`func (o *ScimUserList) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *ScimUserList) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *ScimUserList) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.


### GetResources

`func (o *ScimUserList) GetResources() []ScimUser`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *ScimUserList) GetResourcesOk() (*[]ScimUser, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *ScimUserList) SetResources(v []ScimUser)`

SetResources sets Resources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


