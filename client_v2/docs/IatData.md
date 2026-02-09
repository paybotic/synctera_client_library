# IatData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ForeignPaymentAmount** | **int32** | For inbound IAT payments this field should contain the USD amount or may be blank. Amount is in cents ($100 would be 10000). | 
**ForeignTraceNumber** | **string** |  | 
**Name** | **string** | Receiving Company Name/Individual Name | 
**OdfiBranchCountryCode** | **string** | For Inbound IATs: This 3 position field contains a 2-character code as approved by the International Organization for Standardization (ISO) used to identify the country in which the branch of the bank that originated the entry is located. Values for other countries can be found on the International Organization for Standardization website: www.iso.org.  | 
**OdfiIdNumberQualifier** | **string** | For Inbound IATs: The 2-digit code that identifies the numbering scheme used in the Foreign DFI Identification Number field: 01 &#x3D; National Clearing System 02 &#x3D; BIC Code 03 &#x3D; IBAN Code  | 
**OdfiIdentification** | **string** | For Inbound IATs: This field contains the bank ID number of the Foreign Bank providing funding for the payment transaction.  | 
**OdfiName** | **string** | For Inbound IATs: The name of the foreign bank providing funding for the payment transaction  | 
**OriginatorAddress** | [**IatAddress**](IatAddress.md) |  | 
**OriginatorName** | **string** | The originators name | 
**RdfiBranchCountryCode** | **string** | This 3 position field contains a 2-character code as approved by the International Organization for Standardization (ISO) used to identify the country in which the branch of the bank that receives the entry is located. Values for other countries can be found on the International Organization for Standardization website: www.iso.org  | 
**RdfiIdNumberQualifier** | **string** | The 2-digit code that identifies the numbering scheme used in the Receiving DFI Identification Number field: 01 &#x3D; National Clearing System 02 &#x3D; BIC Code 03 &#x3D; IBAN Code  | 
**RdfiIdentification** | **string** | The bank identification number of the DFI at which the Receiver maintains his account. | 
**RdfiName** | **string** | Name of the Receiver&#39;s bank | 
**ReceiverAddress** | [**IatAddress**](IatAddress.md) |  | 
**ReceiverIdNumber** | **string** | The accounting number by which the Originator is known to the Receiver for descriptive purposes. | 
**TransactionTypeCode** | **string** | Transaction Type Code Describes the type of payment ANN &#x3D; Annuity, BUS &#x3D; Business/Commercial, DEP &#x3D; Deposit, LOA &#x3D; Loan, MIS &#x3D; Miscellaneous, MOR &#x3D; Mortgage PEN &#x3D; Pension, RLS &#x3D; Rent/Lease, REM &#x3D; Remittance2, SAL &#x3D; Salary/Payroll, TAX &#x3D; Tax, TEL &#x3D; Telephone-Initiated Transaction WEB &#x3D; Internet-Initiated Transaction, ARC &#x3D; Accounts Receivable Entry, BOC &#x3D; Back Office Conversion Entry, POP &#x3D; Point of Purchase Entry, RCK &#x3D; Re-presented Check Entry  | 

## Methods

### NewIatData

`func NewIatData(foreignPaymentAmount int32, foreignTraceNumber string, name string, odfiBranchCountryCode string, odfiIdNumberQualifier string, odfiIdentification string, odfiName string, originatorAddress IatAddress, originatorName string, rdfiBranchCountryCode string, rdfiIdNumberQualifier string, rdfiIdentification string, rdfiName string, receiverAddress IatAddress, receiverIdNumber string, transactionTypeCode string, ) *IatData`

NewIatData instantiates a new IatData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIatDataWithDefaults

`func NewIatDataWithDefaults() *IatData`

NewIatDataWithDefaults instantiates a new IatData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetForeignPaymentAmount

`func (o *IatData) GetForeignPaymentAmount() int32`

GetForeignPaymentAmount returns the ForeignPaymentAmount field if non-nil, zero value otherwise.

### GetForeignPaymentAmountOk

`func (o *IatData) GetForeignPaymentAmountOk() (*int32, bool)`

GetForeignPaymentAmountOk returns a tuple with the ForeignPaymentAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignPaymentAmount

`func (o *IatData) SetForeignPaymentAmount(v int32)`

SetForeignPaymentAmount sets ForeignPaymentAmount field to given value.


### GetForeignTraceNumber

`func (o *IatData) GetForeignTraceNumber() string`

GetForeignTraceNumber returns the ForeignTraceNumber field if non-nil, zero value otherwise.

### GetForeignTraceNumberOk

`func (o *IatData) GetForeignTraceNumberOk() (*string, bool)`

GetForeignTraceNumberOk returns a tuple with the ForeignTraceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeignTraceNumber

`func (o *IatData) SetForeignTraceNumber(v string)`

SetForeignTraceNumber sets ForeignTraceNumber field to given value.


### GetName

`func (o *IatData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IatData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IatData) SetName(v string)`

SetName sets Name field to given value.


### GetOdfiBranchCountryCode

`func (o *IatData) GetOdfiBranchCountryCode() string`

GetOdfiBranchCountryCode returns the OdfiBranchCountryCode field if non-nil, zero value otherwise.

### GetOdfiBranchCountryCodeOk

`func (o *IatData) GetOdfiBranchCountryCodeOk() (*string, bool)`

GetOdfiBranchCountryCodeOk returns a tuple with the OdfiBranchCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdfiBranchCountryCode

`func (o *IatData) SetOdfiBranchCountryCode(v string)`

SetOdfiBranchCountryCode sets OdfiBranchCountryCode field to given value.


### GetOdfiIdNumberQualifier

`func (o *IatData) GetOdfiIdNumberQualifier() string`

GetOdfiIdNumberQualifier returns the OdfiIdNumberQualifier field if non-nil, zero value otherwise.

### GetOdfiIdNumberQualifierOk

`func (o *IatData) GetOdfiIdNumberQualifierOk() (*string, bool)`

GetOdfiIdNumberQualifierOk returns a tuple with the OdfiIdNumberQualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdfiIdNumberQualifier

`func (o *IatData) SetOdfiIdNumberQualifier(v string)`

SetOdfiIdNumberQualifier sets OdfiIdNumberQualifier field to given value.


### GetOdfiIdentification

`func (o *IatData) GetOdfiIdentification() string`

GetOdfiIdentification returns the OdfiIdentification field if non-nil, zero value otherwise.

### GetOdfiIdentificationOk

`func (o *IatData) GetOdfiIdentificationOk() (*string, bool)`

GetOdfiIdentificationOk returns a tuple with the OdfiIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdfiIdentification

`func (o *IatData) SetOdfiIdentification(v string)`

SetOdfiIdentification sets OdfiIdentification field to given value.


### GetOdfiName

`func (o *IatData) GetOdfiName() string`

GetOdfiName returns the OdfiName field if non-nil, zero value otherwise.

### GetOdfiNameOk

`func (o *IatData) GetOdfiNameOk() (*string, bool)`

GetOdfiNameOk returns a tuple with the OdfiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOdfiName

`func (o *IatData) SetOdfiName(v string)`

SetOdfiName sets OdfiName field to given value.


### GetOriginatorAddress

`func (o *IatData) GetOriginatorAddress() IatAddress`

GetOriginatorAddress returns the OriginatorAddress field if non-nil, zero value otherwise.

### GetOriginatorAddressOk

`func (o *IatData) GetOriginatorAddressOk() (*IatAddress, bool)`

GetOriginatorAddressOk returns a tuple with the OriginatorAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatorAddress

`func (o *IatData) SetOriginatorAddress(v IatAddress)`

SetOriginatorAddress sets OriginatorAddress field to given value.


### GetOriginatorName

`func (o *IatData) GetOriginatorName() string`

GetOriginatorName returns the OriginatorName field if non-nil, zero value otherwise.

### GetOriginatorNameOk

`func (o *IatData) GetOriginatorNameOk() (*string, bool)`

GetOriginatorNameOk returns a tuple with the OriginatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatorName

`func (o *IatData) SetOriginatorName(v string)`

SetOriginatorName sets OriginatorName field to given value.


### GetRdfiBranchCountryCode

`func (o *IatData) GetRdfiBranchCountryCode() string`

GetRdfiBranchCountryCode returns the RdfiBranchCountryCode field if non-nil, zero value otherwise.

### GetRdfiBranchCountryCodeOk

`func (o *IatData) GetRdfiBranchCountryCodeOk() (*string, bool)`

GetRdfiBranchCountryCodeOk returns a tuple with the RdfiBranchCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdfiBranchCountryCode

`func (o *IatData) SetRdfiBranchCountryCode(v string)`

SetRdfiBranchCountryCode sets RdfiBranchCountryCode field to given value.


### GetRdfiIdNumberQualifier

`func (o *IatData) GetRdfiIdNumberQualifier() string`

GetRdfiIdNumberQualifier returns the RdfiIdNumberQualifier field if non-nil, zero value otherwise.

### GetRdfiIdNumberQualifierOk

`func (o *IatData) GetRdfiIdNumberQualifierOk() (*string, bool)`

GetRdfiIdNumberQualifierOk returns a tuple with the RdfiIdNumberQualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdfiIdNumberQualifier

`func (o *IatData) SetRdfiIdNumberQualifier(v string)`

SetRdfiIdNumberQualifier sets RdfiIdNumberQualifier field to given value.


### GetRdfiIdentification

`func (o *IatData) GetRdfiIdentification() string`

GetRdfiIdentification returns the RdfiIdentification field if non-nil, zero value otherwise.

### GetRdfiIdentificationOk

`func (o *IatData) GetRdfiIdentificationOk() (*string, bool)`

GetRdfiIdentificationOk returns a tuple with the RdfiIdentification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdfiIdentification

`func (o *IatData) SetRdfiIdentification(v string)`

SetRdfiIdentification sets RdfiIdentification field to given value.


### GetRdfiName

`func (o *IatData) GetRdfiName() string`

GetRdfiName returns the RdfiName field if non-nil, zero value otherwise.

### GetRdfiNameOk

`func (o *IatData) GetRdfiNameOk() (*string, bool)`

GetRdfiNameOk returns a tuple with the RdfiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdfiName

`func (o *IatData) SetRdfiName(v string)`

SetRdfiName sets RdfiName field to given value.


### GetReceiverAddress

`func (o *IatData) GetReceiverAddress() IatAddress`

GetReceiverAddress returns the ReceiverAddress field if non-nil, zero value otherwise.

### GetReceiverAddressOk

`func (o *IatData) GetReceiverAddressOk() (*IatAddress, bool)`

GetReceiverAddressOk returns a tuple with the ReceiverAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverAddress

`func (o *IatData) SetReceiverAddress(v IatAddress)`

SetReceiverAddress sets ReceiverAddress field to given value.


### GetReceiverIdNumber

`func (o *IatData) GetReceiverIdNumber() string`

GetReceiverIdNumber returns the ReceiverIdNumber field if non-nil, zero value otherwise.

### GetReceiverIdNumberOk

`func (o *IatData) GetReceiverIdNumberOk() (*string, bool)`

GetReceiverIdNumberOk returns a tuple with the ReceiverIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverIdNumber

`func (o *IatData) SetReceiverIdNumber(v string)`

SetReceiverIdNumber sets ReceiverIdNumber field to given value.


### GetTransactionTypeCode

`func (o *IatData) GetTransactionTypeCode() string`

GetTransactionTypeCode returns the TransactionTypeCode field if non-nil, zero value otherwise.

### GetTransactionTypeCodeOk

`func (o *IatData) GetTransactionTypeCodeOk() (*string, bool)`

GetTransactionTypeCodeOk returns a tuple with the TransactionTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionTypeCode

`func (o *IatData) SetTransactionTypeCode(v string)`

SetTransactionTypeCode sets TransactionTypeCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


