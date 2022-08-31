# ProjectsCreateForAuthenticatedUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the project | 
**Body** | Pointer to **NullableString** | Body of the project | [optional] 

## Methods

### NewProjectsCreateForAuthenticatedUserRequest

`func NewProjectsCreateForAuthenticatedUserRequest(name string, ) *ProjectsCreateForAuthenticatedUserRequest`

NewProjectsCreateForAuthenticatedUserRequest instantiates a new ProjectsCreateForAuthenticatedUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsCreateForAuthenticatedUserRequestWithDefaults

`func NewProjectsCreateForAuthenticatedUserRequestWithDefaults() *ProjectsCreateForAuthenticatedUserRequest`

NewProjectsCreateForAuthenticatedUserRequestWithDefaults instantiates a new ProjectsCreateForAuthenticatedUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectsCreateForAuthenticatedUserRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectsCreateForAuthenticatedUserRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectsCreateForAuthenticatedUserRequest) SetName(v string)`

SetName sets Name field to given value.


### GetBody

`func (o *ProjectsCreateForAuthenticatedUserRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ProjectsCreateForAuthenticatedUserRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ProjectsCreateForAuthenticatedUserRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *ProjectsCreateForAuthenticatedUserRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *ProjectsCreateForAuthenticatedUserRequest) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *ProjectsCreateForAuthenticatedUserRequest) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


