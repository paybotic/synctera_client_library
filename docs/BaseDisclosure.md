# BaseDisclosure

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessId** | **string** |  | [optional] [default to null]
**CreationTime** | [**time.Time**](time.Time.md) | The date and time the resource was created. | [optional] [default to null]
**DisclosureDate** | [**time.Time**](time.Time.md) | Date and time the disclosure was made. | [optional] [default to null]
**EventType** | **string** | Describes how the disclosure was shown and what the user did as a result. One of the following: * &#x60;DISPLAYED&#x60; —     The document was made visible to the user,     but they did not interact with it. * &#x60;VIEWED&#x60; —     The document was made visible to the user,     and they interacted enough to see the whole document (e.g. scrolled to the bottom). * &#x60;ACKNOWLEDGED&#x60; —     The document was made visible to the user,     and they took positive action to confirm that they have read and accepted the document.  | [optional] [default to null]
**Id** | **string** |  | [optional] [default to null]
**LastUpdatedTime** | [**time.Time**](time.Time.md) | The date and time the resource was last updated. | [optional] [default to null]
**Metadata** | [***Metadata**](metadata.md) |  | [optional] [default to null]
**PersonId** | **string** |  | [optional] [default to null]
**Type_** | [***DisclosureType**](disclosure_type.md) |  | [optional] [default to null]
**Version** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

