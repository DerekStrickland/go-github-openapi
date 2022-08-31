# OrganizationDependabotSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the secret. | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Visibility** | **string** | Visibility of a secret | 
**SelectedRepositoriesUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewOrganizationDependabotSecret

`func NewOrganizationDependabotSecret(name string, createdAt time.Time, updatedAt time.Time, visibility string, ) *OrganizationDependabotSecret`

NewOrganizationDependabotSecret instantiates a new OrganizationDependabotSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationDependabotSecretWithDefaults

`func NewOrganizationDependabotSecretWithDefaults() *OrganizationDependabotSecret`

NewOrganizationDependabotSecretWithDefaults instantiates a new OrganizationDependabotSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OrganizationDependabotSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationDependabotSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationDependabotSecret) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *OrganizationDependabotSecret) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationDependabotSecret) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationDependabotSecret) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *OrganizationDependabotSecret) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationDependabotSecret) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationDependabotSecret) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVisibility

`func (o *OrganizationDependabotSecret) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *OrganizationDependabotSecret) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *OrganizationDependabotSecret) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetSelectedRepositoriesUrl

`func (o *OrganizationDependabotSecret) GetSelectedRepositoriesUrl() string`

GetSelectedRepositoriesUrl returns the SelectedRepositoriesUrl field if non-nil, zero value otherwise.

### GetSelectedRepositoriesUrlOk

`func (o *OrganizationDependabotSecret) GetSelectedRepositoriesUrlOk() (*string, bool)`

GetSelectedRepositoriesUrlOk returns a tuple with the SelectedRepositoriesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRepositoriesUrl

`func (o *OrganizationDependabotSecret) SetSelectedRepositoriesUrl(v string)`

SetSelectedRepositoriesUrl sets SelectedRepositoriesUrl field to given value.

### HasSelectedRepositoriesUrl

`func (o *OrganizationDependabotSecret) HasSelectedRepositoriesUrl() bool`

HasSelectedRepositoriesUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


