# SecretScanningLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The location type. Because secrets may be found in different types of resources (ie. code, comments, issues), this field identifies the type of resource where the secret was found. | 
**Details** | [**SecretScanningLocationDetails**](SecretScanningLocationDetails.md) |  | 

## Methods

### NewSecretScanningLocation

`func NewSecretScanningLocation(type_ string, details SecretScanningLocationDetails, ) *SecretScanningLocation`

NewSecretScanningLocation instantiates a new SecretScanningLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretScanningLocationWithDefaults

`func NewSecretScanningLocationWithDefaults() *SecretScanningLocation`

NewSecretScanningLocationWithDefaults instantiates a new SecretScanningLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SecretScanningLocation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecretScanningLocation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecretScanningLocation) SetType(v string)`

SetType sets Type field to given value.


### GetDetails

`func (o *SecretScanningLocation) GetDetails() SecretScanningLocationDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *SecretScanningLocation) GetDetailsOk() (*SecretScanningLocationDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *SecretScanningLocation) SetDetails(v SecretScanningLocationDetails)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


