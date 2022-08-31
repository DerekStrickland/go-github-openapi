# ReposCreateUsingTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Owner** | Pointer to **string** | The organization or person who will own the new repository. To create a new repository in an organization, the authenticated user must be a member of the specified organization. | [optional] 
**Name** | **string** | The name of the new repository. | 
**Description** | Pointer to **string** | A short description of the new repository. | [optional] 
**IncludeAllBranches** | Pointer to **bool** | Set to &#x60;true&#x60; to include the directory structure and files from all branches in the template repository, and not just the default branch. Default: &#x60;false&#x60;. | [optional] [default to false]
**Private** | Pointer to **bool** | Either &#x60;true&#x60; to create a new private repository or &#x60;false&#x60; to create a new public one. | [optional] [default to false]

## Methods

### NewReposCreateUsingTemplateRequest

`func NewReposCreateUsingTemplateRequest(name string, ) *ReposCreateUsingTemplateRequest`

NewReposCreateUsingTemplateRequest instantiates a new ReposCreateUsingTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposCreateUsingTemplateRequestWithDefaults

`func NewReposCreateUsingTemplateRequestWithDefaults() *ReposCreateUsingTemplateRequest`

NewReposCreateUsingTemplateRequestWithDefaults instantiates a new ReposCreateUsingTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOwner

`func (o *ReposCreateUsingTemplateRequest) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ReposCreateUsingTemplateRequest) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ReposCreateUsingTemplateRequest) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ReposCreateUsingTemplateRequest) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetName

`func (o *ReposCreateUsingTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReposCreateUsingTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReposCreateUsingTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ReposCreateUsingTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReposCreateUsingTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReposCreateUsingTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReposCreateUsingTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIncludeAllBranches

`func (o *ReposCreateUsingTemplateRequest) GetIncludeAllBranches() bool`

GetIncludeAllBranches returns the IncludeAllBranches field if non-nil, zero value otherwise.

### GetIncludeAllBranchesOk

`func (o *ReposCreateUsingTemplateRequest) GetIncludeAllBranchesOk() (*bool, bool)`

GetIncludeAllBranchesOk returns a tuple with the IncludeAllBranches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllBranches

`func (o *ReposCreateUsingTemplateRequest) SetIncludeAllBranches(v bool)`

SetIncludeAllBranches sets IncludeAllBranches field to given value.

### HasIncludeAllBranches

`func (o *ReposCreateUsingTemplateRequest) HasIncludeAllBranches() bool`

HasIncludeAllBranches returns a boolean if a field has been set.

### GetPrivate

`func (o *ReposCreateUsingTemplateRequest) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *ReposCreateUsingTemplateRequest) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *ReposCreateUsingTemplateRequest) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *ReposCreateUsingTemplateRequest) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


