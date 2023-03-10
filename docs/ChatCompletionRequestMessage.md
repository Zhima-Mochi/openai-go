# ChatCompletionRequestMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | **string** | The role of the author of this message. | 
**Content** | **string** | The contents of the message | 
**Name** | Pointer to **string** | The name of the user in a multi-user chat | [optional] 

## Methods

### NewChatCompletionRequestMessage

`func NewChatCompletionRequestMessage(role string, content string, ) *ChatCompletionRequestMessage`

NewChatCompletionRequestMessage instantiates a new ChatCompletionRequestMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionRequestMessageWithDefaults

`func NewChatCompletionRequestMessageWithDefaults() *ChatCompletionRequestMessage`

NewChatCompletionRequestMessageWithDefaults instantiates a new ChatCompletionRequestMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *ChatCompletionRequestMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ChatCompletionRequestMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ChatCompletionRequestMessage) SetRole(v string)`

SetRole sets Role field to given value.


### GetContent

`func (o *ChatCompletionRequestMessage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ChatCompletionRequestMessage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ChatCompletionRequestMessage) SetContent(v string)`

SetContent sets Content field to given value.


### GetName

`func (o *ChatCompletionRequestMessage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChatCompletionRequestMessage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChatCompletionRequestMessage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ChatCompletionRequestMessage) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


