# Verification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verified** | **bool** |  | 
**Reason** | **string** |  | 
**Payload** | **NullableString** |  | 
**Signature** | **NullableString** |  | 

## Methods

### NewVerification

`func NewVerification(verified bool, reason string, payload NullableString, signature NullableString, ) *Verification`

NewVerification instantiates a new Verification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationWithDefaults

`func NewVerificationWithDefaults() *Verification`

NewVerificationWithDefaults instantiates a new Verification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerified

`func (o *Verification) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *Verification) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *Verification) SetVerified(v bool)`

SetVerified sets Verified field to given value.


### GetReason

`func (o *Verification) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Verification) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Verification) SetReason(v string)`

SetReason sets Reason field to given value.


### GetPayload

`func (o *Verification) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *Verification) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *Verification) SetPayload(v string)`

SetPayload sets Payload field to given value.


### SetPayloadNil

`func (o *Verification) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *Verification) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetSignature

`func (o *Verification) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *Verification) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *Verification) SetSignature(v string)`

SetSignature sets Signature field to given value.


### SetSignatureNil

`func (o *Verification) SetSignatureNil(b bool)`

 SetSignatureNil sets the value for Signature to be an explicit nil

### UnsetSignature
`func (o *Verification) UnsetSignature()`

UnsetSignature ensures that no value is present for Signature, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


