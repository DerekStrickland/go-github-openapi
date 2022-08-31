# \MarkdownApi

All URIs are relative to *https://api.github.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarkdownRender**](MarkdownApi.md#MarkdownRender) | **Post** /markdown | Render a Markdown document
[**MarkdownRenderRaw**](MarkdownApi.md#MarkdownRenderRaw) | **Post** /markdown/raw | Render a Markdown document in raw mode



## MarkdownRender

> string MarkdownRender(ctx).MarkdownRenderRequest(markdownRenderRequest).Execute()

Render a Markdown document



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    markdownRenderRequest := *openapiclient.NewMarkdownRenderRequest("Text_example") // MarkdownRenderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarkdownApi.MarkdownRender(context.Background()).MarkdownRenderRequest(markdownRenderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarkdownApi.MarkdownRender``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarkdownRender`: string
    fmt.Fprintf(os.Stdout, "Response from `MarkdownApi.MarkdownRender`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkdownRenderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **markdownRenderRequest** | [**MarkdownRenderRequest**](MarkdownRenderRequest.md) |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkdownRenderRaw

> string MarkdownRenderRaw(ctx).Body(body).Execute()

Render a Markdown document in raw mode



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := "body_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarkdownApi.MarkdownRenderRaw(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarkdownApi.MarkdownRenderRaw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarkdownRenderRaw`: string
    fmt.Fprintf(os.Stdout, "Response from `MarkdownApi.MarkdownRenderRaw`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkdownRenderRawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** |  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain, text/x-markdown
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

