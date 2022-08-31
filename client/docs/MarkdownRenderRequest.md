# MarkdownRenderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | The Markdown text to render in HTML. | 
**Mode** | Pointer to **string** | The rendering mode. Can be either &#x60;markdown&#x60; or &#x60;gfm&#x60;. | [optional] [default to "markdown"]
**Context** | Pointer to **string** | The repository context to use when creating references in &#x60;gfm&#x60; mode.  For example, setting &#x60;context&#x60; to &#x60;octo-org/octo-repo&#x60; will change the text &#x60;#42&#x60; into an HTML link to issue 42 in the &#x60;octo-org/octo-repo&#x60; repository. | [optional] 

## Methods

### NewMarkdownRenderRequest

`func NewMarkdownRenderRequest(text string, ) *MarkdownRenderRequest`

NewMarkdownRenderRequest instantiates a new MarkdownRenderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarkdownRenderRequestWithDefaults

`func NewMarkdownRenderRequestWithDefaults() *MarkdownRenderRequest`

NewMarkdownRenderRequestWithDefaults instantiates a new MarkdownRenderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *MarkdownRenderRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MarkdownRenderRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MarkdownRenderRequest) SetText(v string)`

SetText sets Text field to given value.


### GetMode

`func (o *MarkdownRenderRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *MarkdownRenderRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *MarkdownRenderRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *MarkdownRenderRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetContext

`func (o *MarkdownRenderRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *MarkdownRenderRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *MarkdownRenderRequest) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *MarkdownRenderRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


