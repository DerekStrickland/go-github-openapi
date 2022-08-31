# ProjectsAddCollaboratorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permission** | Pointer to **string** | The permission to grant the collaborator. | [optional] [default to "write"]

## Methods

### NewProjectsAddCollaboratorRequest

`func NewProjectsAddCollaboratorRequest() *ProjectsAddCollaboratorRequest`

NewProjectsAddCollaboratorRequest instantiates a new ProjectsAddCollaboratorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsAddCollaboratorRequestWithDefaults

`func NewProjectsAddCollaboratorRequestWithDefaults() *ProjectsAddCollaboratorRequest`

NewProjectsAddCollaboratorRequestWithDefaults instantiates a new ProjectsAddCollaboratorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermission

`func (o *ProjectsAddCollaboratorRequest) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *ProjectsAddCollaboratorRequest) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *ProjectsAddCollaboratorRequest) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *ProjectsAddCollaboratorRequest) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


