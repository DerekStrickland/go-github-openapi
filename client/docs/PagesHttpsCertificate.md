# PagesHttpsCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** |  | 
**Description** | **string** |  | 
**Domains** | **[]string** | Array of the domain set and its alternate name (if it is configured) | 
**ExpiresAt** | Pointer to **string** |  | [optional] 

## Methods

### NewPagesHttpsCertificate

`func NewPagesHttpsCertificate(state string, description string, domains []string, ) *PagesHttpsCertificate`

NewPagesHttpsCertificate instantiates a new PagesHttpsCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagesHttpsCertificateWithDefaults

`func NewPagesHttpsCertificateWithDefaults() *PagesHttpsCertificate`

NewPagesHttpsCertificateWithDefaults instantiates a new PagesHttpsCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *PagesHttpsCertificate) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PagesHttpsCertificate) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PagesHttpsCertificate) SetState(v string)`

SetState sets State field to given value.


### GetDescription

`func (o *PagesHttpsCertificate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PagesHttpsCertificate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PagesHttpsCertificate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDomains

`func (o *PagesHttpsCertificate) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *PagesHttpsCertificate) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *PagesHttpsCertificate) SetDomains(v []string)`

SetDomains sets Domains field to given value.


### GetExpiresAt

`func (o *PagesHttpsCertificate) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PagesHttpsCertificate) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PagesHttpsCertificate) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PagesHttpsCertificate) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


