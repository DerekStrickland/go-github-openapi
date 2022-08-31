# OrgsListCustomRoles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | Pointer to **int32** | The number of custom roles in this organization | [optional] 
**CustomRoles** | Pointer to [**[]OrganizationCustomRepositoryRole**](OrganizationCustomRepositoryRole.md) |  | [optional] 

## Methods

### NewOrgsListCustomRoles200Response

`func NewOrgsListCustomRoles200Response() *OrgsListCustomRoles200Response`

NewOrgsListCustomRoles200Response instantiates a new OrgsListCustomRoles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsListCustomRoles200ResponseWithDefaults

`func NewOrgsListCustomRoles200ResponseWithDefaults() *OrgsListCustomRoles200Response`

NewOrgsListCustomRoles200ResponseWithDefaults instantiates a new OrgsListCustomRoles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *OrgsListCustomRoles200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OrgsListCustomRoles200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OrgsListCustomRoles200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *OrgsListCustomRoles200Response) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetCustomRoles

`func (o *OrgsListCustomRoles200Response) GetCustomRoles() []OrganizationCustomRepositoryRole`

GetCustomRoles returns the CustomRoles field if non-nil, zero value otherwise.

### GetCustomRolesOk

`func (o *OrgsListCustomRoles200Response) GetCustomRolesOk() (*[]OrganizationCustomRepositoryRole, bool)`

GetCustomRolesOk returns a tuple with the CustomRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomRoles

`func (o *OrgsListCustomRoles200Response) SetCustomRoles(v []OrganizationCustomRepositoryRole)`

SetCustomRoles sets CustomRoles field to given value.

### HasCustomRoles

`func (o *OrgsListCustomRoles200Response) HasCustomRoles() bool`

HasCustomRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


