# PullRequestLabelsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**NodeId** | **string** |  | 
**Url** | **string** |  | 
**Name** | **string** |  | 
**Description** | **NullableString** |  | 
**Color** | **string** |  | 
**Default** | **bool** |  | 

## Methods

### NewPullRequestLabelsInner

`func NewPullRequestLabelsInner(id int64, nodeId string, url string, name string, description NullableString, color string, default_ bool, ) *PullRequestLabelsInner`

NewPullRequestLabelsInner instantiates a new PullRequestLabelsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPullRequestLabelsInnerWithDefaults

`func NewPullRequestLabelsInnerWithDefaults() *PullRequestLabelsInner`

NewPullRequestLabelsInnerWithDefaults instantiates a new PullRequestLabelsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PullRequestLabelsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PullRequestLabelsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PullRequestLabelsInner) SetId(v int64)`

SetId sets Id field to given value.


### GetNodeId

`func (o *PullRequestLabelsInner) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *PullRequestLabelsInner) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *PullRequestLabelsInner) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetUrl

`func (o *PullRequestLabelsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PullRequestLabelsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PullRequestLabelsInner) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *PullRequestLabelsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PullRequestLabelsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PullRequestLabelsInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PullRequestLabelsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PullRequestLabelsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PullRequestLabelsInner) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *PullRequestLabelsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PullRequestLabelsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetColor

`func (o *PullRequestLabelsInner) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PullRequestLabelsInner) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PullRequestLabelsInner) SetColor(v string)`

SetColor sets Color field to given value.


### GetDefault

`func (o *PullRequestLabelsInner) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *PullRequestLabelsInner) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *PullRequestLabelsInner) SetDefault(v bool)`

SetDefault sets Default field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


