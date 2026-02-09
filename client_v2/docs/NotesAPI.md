# \NotesAPI

All URIs are relative to *https://api.synctera.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNote**](NotesAPI.md#CreateNote) | **Post** /notes | Create a note
[**ListNotes**](NotesAPI.md#ListNotes) | **Get** /notes | List notes
[**PatchNote**](NotesAPI.md#PatchNote) | **Patch** /notes/{note_id} | Patch Note



## CreateNote

> NoteResponse CreateNote(ctx).NoteCreate(noteCreate).Execute()

Create a note



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/paybotic/synctera_client_library/client_v2"
)

func main() {
	noteCreate := *openapiclient.NewNoteCreate("Customer was frozen to investigate fraud.", "7d943c51-e4ff-4e57-9558-08cab6b963c7", openapiclient.related_resource_type("CUSTOMER")) // NoteCreate | note to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotesAPI.CreateNote(context.Background()).NoteCreate(noteCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.CreateNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNote`: NoteResponse
	fmt.Fprintf(os.Stdout, "Response from `NotesAPI.CreateNote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **noteCreate** | [**NoteCreate**](NoteCreate.md) | note to create | 

### Return type

[**NoteResponse**](NoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotes

> NoteList ListNotes(ctx).RelatedResourceType(relatedResourceType).Id(id).RelatedResourceId(relatedResourceId).AuthorUserId(authorUserId).Tenant(tenant).Limit(limit).PageToken(pageToken).Execute()

List notes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/paybotic/synctera_client_library/client_v2"
)

func main() {
	relatedResourceType := []openapiclient.RelatedResourceType{openapiclient.related_resource_type("CUSTOMER")} // []RelatedResourceType | Type of related resource. Multiple values can be provided as a comma-separated list. 
	id := []string{"Inner_example"} // []string | Return the note with the specified id. Multiple IDs can be provided as a comma-separated list.  (optional)
	relatedResourceId := []string{"Inner_example"} // []string | Unique identifier for the related resource. Multiple IDs can be provided as a comma-separated list.  (optional)
	authorUserId := "64438afd-fa20-4010-a573-2bbdca77cdb6" // string |  (optional)
	tenant := "abcdef_ghijkl" // string |  (optional)
	limit := int32(100) // int32 |  (optional) (default to 100)
	pageToken := "a8937a0d" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotesAPI.ListNotes(context.Background()).RelatedResourceType(relatedResourceType).Id(id).RelatedResourceId(relatedResourceId).AuthorUserId(authorUserId).Tenant(tenant).Limit(limit).PageToken(pageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.ListNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNotes`: NoteList
	fmt.Fprintf(os.Stdout, "Response from `NotesAPI.ListNotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **relatedResourceType** | [**[]RelatedResourceType**](RelatedResourceType.md) | Type of related resource. Multiple values can be provided as a comma-separated list.  | 
 **id** | **[]string** | Return the note with the specified id. Multiple IDs can be provided as a comma-separated list.  | 
 **relatedResourceId** | **[]string** | Unique identifier for the related resource. Multiple IDs can be provided as a comma-separated list.  | 
 **authorUserId** | **string** |  | 
 **tenant** | **string** |  | 
 **limit** | **int32** |  | [default to 100]
 **pageToken** | **string** |  | 

### Return type

[**NoteList**](NoteList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchNote

> NoteResponse PatchNote(ctx, noteId).RelatedResourceType(relatedResourceType).PatchNote(patchNote).IdempotencyKey(idempotencyKey).Execute()

Patch Note



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/paybotic/synctera_client_library/client_v2"
)

func main() {
	relatedResourceType := []openapiclient.RelatedResourceType{openapiclient.related_resource_type("CUSTOMER")} // []RelatedResourceType | Type of related resource. Multiple values can be provided as a comma-separated list. 
	noteId := "64438afd-fa20-4010-a573-2bbdca77cdb6" // string | The unique identifier of a note
	patchNote := *openapiclient.NewPatchNote() // PatchNote | 
	idempotencyKey := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key. A different key must be used for each request, unless it is a retry. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotesAPI.PatchNote(context.Background(), noteId).RelatedResourceType(relatedResourceType).PatchNote(patchNote).IdempotencyKey(idempotencyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.PatchNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchNote`: NoteResponse
	fmt.Fprintf(os.Stdout, "Response from `NotesAPI.PatchNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**noteId** | **string** | The unique identifier of a note | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **relatedResourceType** | [**[]RelatedResourceType**](RelatedResourceType.md) | Type of related resource. Multiple values can be provided as a comma-separated list.  | 

 **patchNote** | [**PatchNote**](PatchNote.md) |  | 
 **idempotencyKey** | **string** | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key. A different key must be used for each request, unless it is a retry. | 

### Return type

[**NoteResponse**](NoteResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

