# \DocumentsAPI

All URIs are relative to *https://api-sandbox.synctera.com/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDocument**](DocumentsAPI.md#CreateDocument) | **Post** /documents | Create a document
[**CreateDocumentVersion**](DocumentsAPI.md#CreateDocumentVersion) | **Post** /documents/{document_id}/versions | Create a new document version
[**GetDocument**](DocumentsAPI.md#GetDocument) | **Get** /documents/{document_id} | Get a document
[**GetDocumentContents**](DocumentsAPI.md#GetDocumentContents) | **Get** /documents/{document_id}/contents | Get contents of latest document version
[**GetDocumentVersion**](DocumentsAPI.md#GetDocumentVersion) | **Get** /documents/{document_id}/versions/{document_version} | Get a document by version
[**GetDocumentVersionContents**](DocumentsAPI.md#GetDocumentVersionContents) | **Get** /documents/{document_id}/versions/{document_version}/contents | Get document contents by version
[**ListDocuments**](DocumentsAPI.md#ListDocuments) | **Get** /documents | List documents
[**UpdateDocument**](DocumentsAPI.md#UpdateDocument) | **Patch** /documents/{document_id} | Update a document



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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    file := os.NewFile(1234, "some_file") // *os.File | The file contents
    description := "description_example" // string | A description of the attached document (optional)
    encryption := openapiclient.encryption("NOT_REQUIRED") // Encryption |  (optional) (default to "NOT_REQUIRED")
    metadata := map[string]interface{}{ ... } // map[string]interface{} | Optional field to store additional information about the resource. Intended to be used by the integrator to store non-sensitive data.  (optional)
    name := "name_example" // string | A user-friendly name for the document (optional)
    relatedResourceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the resource related to the document (optional)
    relatedResourceType := openapiclient.related_resource_type("ACCOUNT") // RelatedResourceType |  (optional)
    type_ := openapiclient.document_type("APPLICATION_DOCUMENTATION") // DocumentType |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsAPI.CreateDocument(context.Background()).File(file).Description(description).Encryption(encryption).Metadata(metadata).Name(name).RelatedResourceId(relatedResourceId).RelatedResourceType(relatedResourceType).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.CreateDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDocument`: Document
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.CreateDocument`: %v\n", resp)
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


## CreateDocumentVersion

> Document CreateDocumentVersion(ctx, documentId).File(file).Description(description).Encryption(encryption).Name(name).Type_(type_).Execute()

Create a new document version



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    documentId := "b01db9c7-78f2-4a99-8aca-1231d32f9b96" // string | The unique identifier of the document.
    file := os.NewFile(1234, "some_file") // *os.File | The file contents
    description := "description_example" // string | A description of the attached document (optional)
    encryption := openapiclient.encryption("NOT_REQUIRED") // Encryption |  (optional) (default to "NOT_REQUIRED")
    name := "name_example" // string | A user-friendly name for the document (optional)
    type_ := openapiclient.document_type("APPLICATION_DOCUMENTATION") // DocumentType |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsAPI.CreateDocumentVersion(context.Background(), documentId).File(file).Description(description).Encryption(encryption).Name(name).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.CreateDocumentVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDocumentVersion`: Document
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.CreateDocumentVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** | The unique identifier of the document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDocumentVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | The file contents | 
 **description** | **string** | A description of the attached document | 
 **encryption** | [**Encryption**](Encryption.md) |  | [default to &quot;NOT_REQUIRED&quot;]
 **name** | **string** | A user-friendly name for the document | 
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

> Document GetDocument(ctx, documentId).Execute()

Get a document



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    documentId := "b01db9c7-78f2-4a99-8aca-1231d32f9b96" // string | The unique identifier of the document.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsAPI.GetDocument(context.Background(), documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.GetDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocument`: Document
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.GetDocument`: %v\n", resp)
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

[**Document**](Document.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentContents

> *os.File GetDocumentContents(ctx, documentId).Execute()

Get contents of latest document version



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    documentId := "b01db9c7-78f2-4a99-8aca-1231d32f9b96" // string | The unique identifier of the document.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsAPI.GetDocumentContents(context.Background(), documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.GetDocumentContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentContents`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.GetDocumentContents`: %v\n", resp)
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


## GetDocumentVersion

> Document GetDocumentVersion(ctx, documentVersion, documentId).Execute()

Get a document by version



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    documentVersion := int32(1) // int32 | The document version.
    documentId := "b01db9c7-78f2-4a99-8aca-1231d32f9b96" // string | The unique identifier of the document.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsAPI.GetDocumentVersion(context.Background(), documentVersion, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.GetDocumentVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentVersion`: Document
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.GetDocumentVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentVersion** | **int32** | The document version. | 
**documentId** | **string** | The unique identifier of the document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Document**](Document.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDocumentVersionContents

> *os.File GetDocumentVersionContents(ctx, documentVersion, documentId).Execute()

Get document contents by version



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    documentVersion := int32(1) // int32 | The document version.
    documentId := "b01db9c7-78f2-4a99-8aca-1231d32f9b96" // string | The unique identifier of the document.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsAPI.GetDocumentVersionContents(context.Background(), documentVersion, documentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.GetDocumentVersionContents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDocumentVersionContents`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.GetDocumentVersionContents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentVersion** | **int32** | The document version. | 
**documentId** | **string** | The unique identifier of the document. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDocumentVersionContentsRequest struct via the builder pattern


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

> DocumentList ListDocuments(ctx).RelatedResourceId(relatedResourceId).Type_(type_).Encryption(encryption).PageToken(pageToken).Id(id).RelatedResourceType(relatedResourceType).Limit(limit).Execute()

List documents



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    relatedResourceId := "b01db9c7-78f2-4a99-8aca-1231d32f9b96" // string | Return documents that are related to resources with the specified ID (optional)
    type_ := openapiclient.document_type("APPLICATION_DOCUMENTATION") // DocumentType | The type of documents (optional)
    encryption := "encryption_example" // string | Whether the file should be encrypted and access restricted, e.g. if the file contains PII (optional)
    pageToken := "a8937a0d" // string |  (optional)
    id := []string{"7d943c51-e4ff-4e57-9558-08cab6b963c7"} // []string | Unique resource identifier (optional)
    relatedResourceType := openapiclient.related_resource_type("ACCOUNT") // RelatedResourceType | Return documents that are related to resources of the specified type (optional)
    limit := int32(100) // int32 |  (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsAPI.ListDocuments(context.Background()).RelatedResourceId(relatedResourceId).Type_(type_).Encryption(encryption).PageToken(pageToken).Id(id).RelatedResourceType(relatedResourceType).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.ListDocuments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDocuments`: DocumentList
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.ListDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **relatedResourceId** | **string** | Return documents that are related to resources with the specified ID | 
 **type_** | [**DocumentType**](DocumentType.md) | The type of documents | 
 **encryption** | **string** | Whether the file should be encrypted and access restricted, e.g. if the file contains PII | 
 **pageToken** | **string** |  | 
 **id** | **[]string** | Unique resource identifier | 
 **relatedResourceType** | [**RelatedResourceType**](RelatedResourceType.md) | Return documents that are related to resources of the specified type | 
 **limit** | **int32** |  | [default to 100]

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
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    documentId := "b01db9c7-78f2-4a99-8aca-1231d32f9b96" // string | The unique identifier of the document.
    patchDocument := *openapiclient.NewPatchDocument() // PatchDocument | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentsAPI.UpdateDocument(context.Background(), documentId).PatchDocument(patchDocument).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.UpdateDocument``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDocument`: Document
    fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.UpdateDocument`: %v\n", resp)
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

