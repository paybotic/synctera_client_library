# \NotesAPI

All URIs are relative to *https://api-sandbox.synctera.com/v0*

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
	openapiclient "./openapi"
)

func main() {
	noteCreate := *openapiclient.NewNoteCreate("Customer was frozen to investigate fraud.", "7d943c51-e4ff-4e57-9558-08cab6b963c7", openapiclient.related_resource_type2("ACCOUNT")) // NoteCreate | note to create

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

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotes

> NoteList ListNotes(ctx).AuthorUserId(authorUserId).RelatedResourceId(relatedResourceId).Id(id).RelatedResourceType(relatedResourceType).PageToken(pageToken).Limit(limit).Tenant(tenant).Execute()

List notes



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
	authorUserId := "64438afd-fa20-4010-a573-2bbdca77cdb6" // string |  (optional)
	relatedResourceId := []string{"Inner_example"} // []string | Only return notes that have a related resource with the specified ID. (optional)
	id := []string{"Inner_example"} // []string | Return the note with the specified id. Multiple IDs can be provided as a comma-separated list.  (optional)
	relatedResourceType := openapiclient.related_resource_type2("ACCOUNT") // RelatedResourceType2 | Only return notes that have a related resource with the specified type. (optional)
	pageToken := "a8937a0d" // string |  (optional)
	limit := int32(100) // int32 |  (optional) (default to 100)
	tenant := "abcdef_ghijkl" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotesAPI.ListNotes(context.Background()).AuthorUserId(authorUserId).RelatedResourceId(relatedResourceId).Id(id).RelatedResourceType(relatedResourceType).PageToken(pageToken).Limit(limit).Tenant(tenant).Execute()
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
 **authorUserId** | **string** |  | 
 **relatedResourceId** | **[]string** | Only return notes that have a related resource with the specified ID. | 
 **id** | **[]string** | Return the note with the specified id. Multiple IDs can be provided as a comma-separated list.  | 
 **relatedResourceType** | [**RelatedResourceType2**](RelatedResourceType2.md) | Only return notes that have a related resource with the specified type. | 
 **pageToken** | **string** |  | 
 **limit** | **int32** |  | [default to 100]
 **tenant** | **string** |  | 

### Return type

[**NoteList**](NoteList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchNote

> NoteResponse PatchNote(ctx, noteId).PatchNote(patchNote).IdempotencyKey(idempotencyKey).RelatedResourceType(relatedResourceType).Execute()

Patch Note



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
	noteId := "64438afd-fa20-4010-a573-2bbdca77cdb6" // string | The unique identifier of a note
	patchNote := *openapiclient.NewPatchNote() // PatchNote | 
	idempotencyKey := "7d943c51-e4ff-4e57-9558-08cab6b963c7" // string | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key. A different key must be used for each request, unless it is a retry. (optional)
	relatedResourceType := openapiclient.related_resource_type2("ACCOUNT") // RelatedResourceType2 | Only return notes that have a related resource with the specified type. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotesAPI.PatchNote(context.Background(), noteId).PatchNote(patchNote).IdempotencyKey(idempotencyKey).RelatedResourceType(relatedResourceType).Execute()
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

 **patchNote** | [**PatchNote**](PatchNote.md) |  | 
 **idempotencyKey** | **string** | An idempotency key is an arbitrary unique value generated by client to detect subsequent retries of the same request. It is recommended that a UUID or a similar random identifier be used as an idempotency key. A different key must be used for each request, unless it is a retry. | 
 **relatedResourceType** | [**RelatedResourceType2**](RelatedResourceType2.md) | Only return notes that have a related resource with the specified type. | 

### Return type

[**NoteResponse**](NoteResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, application/problem+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

