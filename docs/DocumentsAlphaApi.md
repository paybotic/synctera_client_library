# \DocumentsAlphaApi

All URIs are relative to *https://api.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDocument**](DocumentsAlphaApi.md#CreateDocument) | **Post** /documents | Create a document
[**GetDocument**](DocumentsAlphaApi.md#GetDocument) | **Get** /documents/{document_id} | Get a document
[**GetDocumentContents**](DocumentsAlphaApi.md#GetDocumentContents) | **Get** /documents/{document_id}/contents | Get document contents
[**ListDocuments**](DocumentsAlphaApi.md#ListDocuments) | **Get** /documents | List documents
[**UpdateDocument**](DocumentsAlphaApi.md#UpdateDocument) | **Patch** /documents/{document_id} | Update a document



## CreateDocument

> Document CreateDocument(ctx).File(file).Description(description).Encryption(encryption).Metadata(metadata).Name(name).RelatedResourceId(relatedResourceId).RelatedResourceType(relatedResourceType).Type_(type_).Execute()

Create a document



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
    file := os.NewFile(1234, "some_file") // *os.File | The file contents
    description := "description_example" // string | A description of the attached document (optional)
    encryption := openapiclient.encryption("REQUIRED") // Encryption |  (optional) (default to "NOT_REQUIRED")
    metadata := map[string]interface{}{ ... } // map[string]interface{} | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  (optional)
    name := "name_example" // string | A user-friendly name for the document (optional)
    relatedResourceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the resource related to the document (optional)
    relatedResourceType := openapiclient.related_resource_type("CUSTOMER") // RelatedResourceType |  (optional)
    type_ := openapiclient.document_type("APPLICATION_DOCUMENTATION") // DocumentType |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsAlphaApi.CreateDocument(context.Background()).File(file).Description(description).Encryption(encryption).Metadata(metadata).Name(name).RelatedResourceId(relatedResourceId).RelatedResourceType(relatedResourceType).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAlphaApi.CreateDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDocument`: Document
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAlphaApi.CreateDocument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The file contents | 
 **description** | **string** | A description of the attached document | 
 **encryption** | [**Encryption**](Encryption.md) |  | [default to &quot;NOT_REQUIRED&quot;]
 **metadata** | [**map[string]interface{}**](map[string]interface{}.md) | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  | 
 **name** | **string** | A user-friendly name for the document | 
 **relatedResourceId** | **string** | The ID of the resource related to the document | 
 **relatedResourceType** | [**RelatedResourceType**](RelatedResourceType.md) |  | 
 **type_** | [**DocumentType**](DocumentType.md) |  | 

### Return type

[**Document**](Document.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocument

> *os.File GetDocument(ctx, documentId).Execute()

Get a document



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
    documentId := "2e3304dc-a1c2-427e-ac5a-a2586a95ce1f" // string | The unique identifier of the document.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsAlphaApi.GetDocument(context.Background(), documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAlphaApi.GetDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocument`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAlphaApi.GetDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** | The unique identifier of the document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentContents

> *os.File GetDocumentContents(ctx, documentId).Execute()

Get document contents



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
    documentId := "2e3304dc-a1c2-427e-ac5a-a2586a95ce1f" // string | The unique identifier of the document.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsAlphaApi.GetDocumentContents(context.Background(), documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAlphaApi.GetDocumentContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentContents`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAlphaApi.GetDocumentContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** | The unique identifier of the document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentContentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDocuments

> DocumentList ListDocuments(ctx).Id(id).Limit(limit).PageToken(pageToken).RelatedResourceType(relatedResourceType).RelatedResourceId(relatedResourceId).Encryption(encryption).Type_(type_).Execute()

List documents



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
    id := []string{"7d943c51-e4ff-4e57-9558-08cab6b963c7"} // []string | Unique resource identifier (optional)
    limit := int32(100) // int32 |  (optional) (default to 100)
    pageToken := "a8937a0d" // string |  (optional)
    relatedResourceType := openapiclient.related_resource_type("CUSTOMER") // RelatedResourceType | Return documents that are related to resources of the specified type (optional)
    relatedResourceId := "1985f769-dd31-4128-95db-f765355e6631" // string | Return documents that are related to resources with the specified ID (optional)
    encryption := "encryption_example" // string | Whether the file should be encrypted and access restricted, e.g. if the file contains PII (optional)
    type_ := openapiclient.document_type("APPLICATION_DOCUMENTATION") // DocumentType | The type of documents (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsAlphaApi.ListDocuments(context.Background()).Id(id).Limit(limit).PageToken(pageToken).RelatedResourceType(relatedResourceType).RelatedResourceId(relatedResourceId).Encryption(encryption).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAlphaApi.ListDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDocuments`: DocumentList
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAlphaApi.ListDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** | Unique resource identifier | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 
 **relatedResourceType** | [**RelatedResourceType**](RelatedResourceType.md) | Return documents that are related to resources of the specified type | 
 **relatedResourceId** | **string** | Return documents that are related to resources with the specified ID | 
 **encryption** | **string** | Whether the file should be encrypted and access restricted, e.g. if the file contains PII | 
 **type_** | [**DocumentType**](DocumentType.md) | The type of documents | 

### Return type

[**DocumentList**](DocumentList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDocument

> Document UpdateDocument(ctx, documentId).PatchDocument(patchDocument).Execute()

Update a document



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
    documentId := "2e3304dc-a1c2-427e-ac5a-a2586a95ce1f" // string | The unique identifier of the document.
    patchDocument := *openapiclient.NewPatchDocument() // PatchDocument | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsAlphaApi.UpdateDocument(context.Background(), documentId).PatchDocument(patchDocument).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAlphaApi.UpdateDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDocument`: Document
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAlphaApi.UpdateDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** | The unique identifier of the document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchDocument** | [**PatchDocument**](PatchDocument.md) |  | 

### Return type

[**Document**](Document.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

