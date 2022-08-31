# PagesHealthCheck

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to [**PagesHealthCheckDomain**](PagesHealthCheckDomain.md) |  | [optional] 
**AltDomain** | Pointer to [**NullablePagesHealthCheckAltDomain**](PagesHealthCheckAltDomain.md) |  | [optional] 

## Methods

### NewPagesHealthCheck

`func NewPagesHealthCheck() *PagesHealthCheck`

NewPagesHealthCheck instantiates a new PagesHealthCheck object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagesHealthCheckWithDefaults

`func NewPagesHealthCheckWithDefaults() *PagesHealthCheck`

NewPagesHealthCheckWithDefaults instantiates a new PagesHealthCheck object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *PagesHealthCheck) GetDomain() PagesHealthCheckDomain`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PagesHealthCheck) GetDomainOk() (*PagesHealthCheckDomain, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PagesHealthCheck) SetDomain(v PagesHealthCheckDomain)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *PagesHealthCheck) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetAltDomain

`func (o *PagesHealthCheck) GetAltDomain() PagesHealthCheckAltDomain`

GetAltDomain returns the AltDomain field if non-nil, zero value otherwise.

### GetAltDomainOk

`func (o *PagesHealthCheck) GetAltDomainOk() (*PagesHealthCheckAltDomain, bool)`

GetAltDomainOk returns a tuple with the AltDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltDomain

`func (o *PagesHealthCheck) SetAltDomain(v PagesHealthCheckAltDomain)`

SetAltDomain sets AltDomain field to given value.

### HasAltDomain

`func (o *PagesHealthCheck) HasAltDomain() bool`

HasAltDomain returns a boolean if a field has been set.

### SetAltDomainNil

`func (o *PagesHealthCheck) SetAltDomainNil(b bool)`

 SetAltDomainNil sets the value for AltDomain to be an explicit nil

### UnsetAltDomain
`func (o *PagesHealthCheck) UnsetAltDomain()`

UnsetAltDomain ensures that no value is present for AltDomain, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


