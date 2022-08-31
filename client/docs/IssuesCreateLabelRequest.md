# IssuesCreateLabelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the label. Emoji can be added to label names, using either native emoji or colon-style markup. For example, typing &#x60;:strawberry:&#x60; will render the emoji ![:strawberry:](https://github.githubassets.com/images/icons/emoji/unicode/1f353.png \&quot;:strawberry:\&quot;). For a full list of available emoji and codes, see \&quot;[Emoji cheat sheet](https://github.com/ikatyang/emoji-cheat-sheet).\&quot; | 
**Color** | Pointer to **string** | The [hexadecimal color code](http://www.color-hex.com/) for the label, without the leading &#x60;#&#x60;. | [optional] 
**Description** | Pointer to **string** | A short description of the label. Must be 100 characters or fewer. | [optional] 

## Methods

### NewIssuesCreateLabelRequest

`func NewIssuesCreateLabelRequest(name string, ) *IssuesCreateLabelRequest`

NewIssuesCreateLabelRequest instantiates a new IssuesCreateLabelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssuesCreateLabelRequestWithDefaults

`func NewIssuesCreateLabelRequestWithDefaults() *IssuesCreateLabelRequest`

NewIssuesCreateLabelRequestWithDefaults instantiates a new IssuesCreateLabelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IssuesCreateLabelRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssuesCreateLabelRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssuesCreateLabelRequest) SetName(v string)`

SetName sets Name field to given value.


### GetColor

`func (o *IssuesCreateLabelRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *IssuesCreateLabelRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *IssuesCreateLabelRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *IssuesCreateLabelRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDescription

`func (o *IssuesCreateLabelRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssuesCreateLabelRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssuesCreateLabelRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssuesCreateLabelRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


