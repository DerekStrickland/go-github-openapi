# License

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Name** | **string** |  | 
**SpdxId** | **NullableString** |  | 
**Url** | **NullableString** |  | 
**NodeId** | **string** |  | 
**HtmlUrl** | **string** |  | 
**Description** | **string** |  | 
**Implementation** | **string** |  | 
**Permissions** | **[]string** |  | 
**Conditions** | **[]string** |  | 
**Limitations** | **[]string** |  | 
**Body** | **string** |  | 
**Featured** | **bool** |  | 

## Methods

### NewLicense

`func NewLicense(key string, name string, spdxId NullableString, url NullableString, nodeId string, htmlUrl string, description string, implementation string, permissions []string, conditions []string, limitations []string, body string, featured bool, ) *License`

NewLicense instantiates a new License object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseWithDefaults

`func NewLicenseWithDefaults() *License`

NewLicenseWithDefaults instantiates a new License object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *License) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *License) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *License) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *License) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *License) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *License) SetName(v string)`

SetName sets Name field to given value.


### GetSpdxId

`func (o *License) GetSpdxId() string`

GetSpdxId returns the SpdxId field if non-nil, zero value otherwise.

### GetSpdxIdOk

`func (o *License) GetSpdxIdOk() (*string, bool)`

GetSpdxIdOk returns a tuple with the SpdxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpdxId

`func (o *License) SetSpdxId(v string)`

SetSpdxId sets SpdxId field to given value.


### SetSpdxIdNil

`func (o *License) SetSpdxIdNil(b bool)`

 SetSpdxIdNil sets the value for SpdxId to be an explicit nil

### UnsetSpdxId
`func (o *License) UnsetSpdxId()`

UnsetSpdxId ensures that no value is present for SpdxId, not even an explicit nil
### GetUrl

`func (o *License) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *License) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *License) SetUrl(v string)`

SetUrl sets Url field to given value.


### SetUrlNil

`func (o *License) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *License) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetNodeId

`func (o *License) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *License) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *License) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetHtmlUrl

`func (o *License) GetHtmlUrl() string`

GetHtmlUrl returns the HtmlUrl field if non-nil, zero value otherwise.

### GetHtmlUrlOk

`func (o *License) GetHtmlUrlOk() (*string, bool)`

GetHtmlUrlOk returns a tuple with the HtmlUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlUrl

`func (o *License) SetHtmlUrl(v string)`

SetHtmlUrl sets HtmlUrl field to given value.


### GetDescription

`func (o *License) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *License) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *License) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetImplementation

`func (o *License) GetImplementation() string`

GetImplementation returns the Implementation field if non-nil, zero value otherwise.

### GetImplementationOk

`func (o *License) GetImplementationOk() (*string, bool)`

GetImplementationOk returns a tuple with the Implementation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplementation

`func (o *License) SetImplementation(v string)`

SetImplementation sets Implementation field to given value.


### GetPermissions

`func (o *License) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *License) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *License) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetConditions

`func (o *License) GetConditions() []string`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *License) GetConditionsOk() (*[]string, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *License) SetConditions(v []string)`

SetConditions sets Conditions field to given value.


### GetLimitations

`func (o *License) GetLimitations() []string`

GetLimitations returns the Limitations field if non-nil, zero value otherwise.

### GetLimitationsOk

`func (o *License) GetLimitationsOk() (*[]string, bool)`

GetLimitationsOk returns a tuple with the Limitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitations

`func (o *License) SetLimitations(v []string)`

SetLimitations sets Limitations field to given value.


### GetBody

`func (o *License) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *License) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *License) SetBody(v string)`

SetBody sets Body field to given value.


### GetFeatured

`func (o *License) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *License) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *License) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


