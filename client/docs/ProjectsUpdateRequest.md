# ProjectsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the project | [optional] 
**Body** | Pointer to **NullableString** | Body of the project | [optional] 
**State** | Pointer to **string** | State of the project; either &#39;open&#39; or &#39;closed&#39; | [optional] 
**OrganizationPermission** | Pointer to **string** | The baseline permission that all organization members have on this project | [optional] 
**Private** | Pointer to **bool** | Whether or not this project can be seen by everyone. | [optional] 

## Methods

### NewProjectsUpdateRequest

`func NewProjectsUpdateRequest() *ProjectsUpdateRequest`

NewProjectsUpdateRequest instantiates a new ProjectsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsUpdateRequestWithDefaults

`func NewProjectsUpdateRequestWithDefaults() *ProjectsUpdateRequest`

NewProjectsUpdateRequestWithDefaults instantiates a new ProjectsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectsUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectsUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectsUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectsUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBody

`func (o *ProjectsUpdateRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ProjectsUpdateRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ProjectsUpdateRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ProjectsUpdateRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *ProjectsUpdateRequest) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *ProjectsUpdateRequest) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetState

`func (o *ProjectsUpdateRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ProjectsUpdateRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ProjectsUpdateRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ProjectsUpdateRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetOrganizationPermission

`func (o *ProjectsUpdateRequest) GetOrganizationPermission() string`

GetOrganizationPermission returns the OrganizationPermission field if non-nil, zero value otherwise.

### GetOrganizationPermissionOk

`func (o *ProjectsUpdateRequest) GetOrganizationPermissionOk() (*string, bool)`

GetOrganizationPermissionOk returns a tuple with the OrganizationPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationPermission

`func (o *ProjectsUpdateRequest) SetOrganizationPermission(v string)`

SetOrganizationPermission sets OrganizationPermission field to given value.

### HasOrganizationPermission

`func (o *ProjectsUpdateRequest) HasOrganizationPermission() bool`

HasOrganizationPermission returns a boolean if a field has been set.

### GetPrivate

`func (o *ProjectsUpdateRequest) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *ProjectsUpdateRequest) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *ProjectsUpdateRequest) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *ProjectsUpdateRequest) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


