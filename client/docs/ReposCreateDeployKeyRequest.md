# ReposCreateDeployKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | A name for the key. | [optional] 
**Key** | **string** | The contents of the key. | 
**ReadOnly** | Pointer to **bool** | If &#x60;true&#x60;, the key will only be able to read repository contents. Otherwise, the key will be able to read and write.      Deploy keys with write access can perform the same actions as an organization member with admin access, or a collaborator on a personal repository. For more information, see \&quot;[Repository permission levels for an organization](https://docs.github.com/articles/repository-permission-levels-for-an-organization/)\&quot; and \&quot;[Permission levels for a user account repository](https://docs.github.com/articles/permission-levels-for-a-user-account-repository/).\&quot; | [optional] 

## Methods

### NewReposCreateDeployKeyRequest

`func NewReposCreateDeployKeyRequest(key string, ) *ReposCreateDeployKeyRequest`

NewReposCreateDeployKeyRequest instantiates a new ReposCreateDeployKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposCreateDeployKeyRequestWithDefaults

`func NewReposCreateDeployKeyRequestWithDefaults() *ReposCreateDeployKeyRequest`

NewReposCreateDeployKeyRequestWithDefaults instantiates a new ReposCreateDeployKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ReposCreateDeployKeyRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ReposCreateDeployKeyRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ReposCreateDeployKeyRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ReposCreateDeployKeyRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetKey

`func (o *ReposCreateDeployKeyRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ReposCreateDeployKeyRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ReposCreateDeployKeyRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetReadOnly

`func (o *ReposCreateDeployKeyRequest) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ReposCreateDeployKeyRequest) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ReposCreateDeployKeyRequest) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ReposCreateDeployKeyRequest) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


