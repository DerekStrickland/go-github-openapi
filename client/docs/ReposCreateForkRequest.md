# ReposCreateForkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organization** | Pointer to **string** | Optional parameter to specify the organization name if forking into an organization. | [optional] 
**Name** | Pointer to **string** | When forking from an existing repository, a new name for the fork. | [optional] 
**DefaultBranchOnly** | Pointer to **bool** | When forking from an existing repository, fork with only the default branch. | [optional] 

## Methods

### NewReposCreateForkRequest

`func NewReposCreateForkRequest() *ReposCreateForkRequest`

NewReposCreateForkRequest instantiates a new ReposCreateForkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReposCreateForkRequestWithDefaults

`func NewReposCreateForkRequestWithDefaults() *ReposCreateForkRequest`

NewReposCreateForkRequestWithDefaults instantiates a new ReposCreateForkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganization

`func (o *ReposCreateForkRequest) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *ReposCreateForkRequest) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *ReposCreateForkRequest) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *ReposCreateForkRequest) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetName

`func (o *ReposCreateForkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReposCreateForkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReposCreateForkRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReposCreateForkRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultBranchOnly

`func (o *ReposCreateForkRequest) GetDefaultBranchOnly() bool`

GetDefaultBranchOnly returns the DefaultBranchOnly field if non-nil, zero value otherwise.

### GetDefaultBranchOnlyOk

`func (o *ReposCreateForkRequest) GetDefaultBranchOnlyOk() (*bool, bool)`

GetDefaultBranchOnlyOk returns a tuple with the DefaultBranchOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBranchOnly

`func (o *ReposCreateForkRequest) SetDefaultBranchOnly(v bool)`

SetDefaultBranchOnly sets DefaultBranchOnly field to given value.

### HasDefaultBranchOnly

`func (o *ReposCreateForkRequest) HasDefaultBranchOnly() bool`

HasDefaultBranchOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


