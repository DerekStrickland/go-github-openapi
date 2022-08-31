# OrgsListAppInstallations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCount** | **int32** |  | 
**Installations** | [**[]Installation**](Installation.md) |  | 

## Methods

### NewOrgsListAppInstallations200Response

`func NewOrgsListAppInstallations200Response(totalCount int32, installations []Installation, ) *OrgsListAppInstallations200Response`

NewOrgsListAppInstallations200Response instantiates a new OrgsListAppInstallations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgsListAppInstallations200ResponseWithDefaults

`func NewOrgsListAppInstallations200ResponseWithDefaults() *OrgsListAppInstallations200Response`

NewOrgsListAppInstallations200ResponseWithDefaults instantiates a new OrgsListAppInstallations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCount

`func (o *OrgsListAppInstallations200Response) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *OrgsListAppInstallations200Response) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *OrgsListAppInstallations200Response) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.


### GetInstallations

`func (o *OrgsListAppInstallations200Response) GetInstallations() []Installation`

GetInstallations returns the Installations field if non-nil, zero value otherwise.

### GetInstallationsOk

`func (o *OrgsListAppInstallations200Response) GetInstallationsOk() (*[]Installation, bool)`

GetInstallationsOk returns a tuple with the Installations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallations

`func (o *OrgsListAppInstallations200Response) SetInstallations(v []Installation)`

SetInstallations sets Installations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


