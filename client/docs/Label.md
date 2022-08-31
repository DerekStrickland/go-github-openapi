# Label

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**NodeId** | **string** |  | 
**Url** | **string** | URL for the label | 
**Name** | **string** | The name of the label. | 
**Description** | **NullableString** |  | 
**Color** | **string** | 6-character hex code, without the leading #, identifying the color | 
**Default** | **bool** |  | 

## Methods

### NewLabel

`func NewLabel(id int64, nodeId string, url string, name string, description NullableString, color string, default_ bool, ) *Label`

NewLabel instantiates a new Label object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelWithDefaults

`func NewLabelWithDefaults() *Label`

NewLabelWithDefaults instantiates a new Label object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Label) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Label) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Label) SetId(v int64)`

SetId sets Id field to given value.


### GetNodeId

`func (o *Label) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *Label) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *Label) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *Label) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Label) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Label) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *Label) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Label) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Label) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Label) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Label) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Label) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Label) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Label) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetColor

`func (o *Label) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Label) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Label) SetColor(v string)`

SetColor sets Color field to given value.


### GetDefault

`func (o *Label) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Label) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Label) SetDefault(v bool)`

SetDefault sets Default field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


