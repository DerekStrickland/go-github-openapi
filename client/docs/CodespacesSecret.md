# CodespacesSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the secret | 
**CreatedAt** | **time.Time** | The date and time at which the secret was created, in ISO 8601 format&#39;:&#39; YYYY-MM-DDTHH:MM:SSZ. | 
**UpdatedAt** | **time.Time** | The date and time at which the secret was last updated, in ISO 8601 format&#39;:&#39; YYYY-MM-DDTHH:MM:SSZ. | 
**Visibility** | **string** | The type of repositories in the organization that the secret is visible to | 
**SelectedRepositoriesUrl** | **string** | The API URL at which the list of repositories this secret is visible to can be retrieved | 

## Methods

### NewCodespacesSecret

`func NewCodespacesSecret(name string, createdAt time.Time, updatedAt time.Time, visibility string, selectedRepositoriesUrl string, ) *CodespacesSecret`

NewCodespacesSecret instantiates a new CodespacesSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodespacesSecretWithDefaults

`func NewCodespacesSecretWithDefaults() *CodespacesSecret`

NewCodespacesSecretWithDefaults instantiates a new CodespacesSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CodespacesSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CodespacesSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CodespacesSecret) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *CodespacesSecret) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CodespacesSecret) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CodespacesSecret) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CodespacesSecret) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CodespacesSecret) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CodespacesSecret) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVisibility

`func (o *CodespacesSecret) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CodespacesSecret) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CodespacesSecret) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetSelectedRepositoriesUrl

`func (o *CodespacesSecret) GetSelectedRepositoriesUrl() string`

GetSelectedRepositoriesUrl returns the SelectedRepositoriesUrl field if non-nil, zero value otherwise.

### GetSelectedRepositoriesUrlOk

`func (o *CodespacesSecret) GetSelectedRepositoriesUrlOk() (*string, bool)`

GetSelectedRepositoriesUrlOk returns a tuple with the SelectedRepositoriesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedRepositoriesUrl

`func (o *CodespacesSecret) SetSelectedRepositoriesUrl(v string)`

SetSelectedRepositoriesUrl sets SelectedRepositoriesUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


