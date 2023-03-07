/*
OpenAI API

APIs for sampling from and fine-tuning language models

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CreateAnswerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAnswerRequest{}

// CreateAnswerRequest struct for CreateAnswerRequest
type CreateAnswerRequest struct {
	// ID of the model to use for completion. You can select one of `ada`, `babbage`, `curie`, or `davinci`.
	Model string `json:"model"`
	// Question to get answered.
	Question string `json:"question"`
	// List of (question, answer) pairs that will help steer the model towards the tone and answer format you'd like. We recommend adding 2 to 3 examples.
	Examples [][]string `json:"examples"`
	// A text snippet containing the contextual information used to generate the answers for the `examples` you provide.
	ExamplesContext string `json:"examples_context"`
	// List of documents from which the answer for the input `question` should be derived. If this is an empty list, the question will be answered based on the question-answer examples.  You should specify either `documents` or a `file`, but not both. 
	Documents []string `json:"documents,omitempty"`
	// The ID of an uploaded file that contains documents to search over. See [upload file](/docs/api-reference/files/upload) for how to upload a file of the desired format and purpose.  You should specify either `documents` or a `file`, but not both. 
	File NullableString `json:"file,omitempty"`
	// ID of the model to use for [Search](/docs/api-reference/searches/create). You can select one of `ada`, `babbage`, `curie`, or `davinci`.
	SearchModel NullableString `json:"search_model,omitempty"`
	// The maximum number of documents to be ranked by [Search](/docs/api-reference/searches/create) when using `file`. Setting it to a higher value leads to improved accuracy but with increased latency and cost.
	MaxRerank NullableInt32 `json:"max_rerank,omitempty"`
	// What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.
	Temperature NullableFloat32 `json:"temperature,omitempty"`
	// Include the log probabilities on the `logprobs` most likely tokens, as well the chosen tokens. For example, if `logprobs` is 5, the API will return a list of the 5 most likely tokens. The API will always return the `logprob` of the sampled token, so there may be up to `logprobs+1` elements in the response.  The maximum value for `logprobs` is 5. If you need more than this, please contact us through our [Help center](https://help.openai.com) and describe your use case.  When `logprobs` is set, `completion` will be automatically added into `expand` to get the logprobs. 
	Logprobs NullableInt32 `json:"logprobs,omitempty"`
	// The maximum number of tokens allowed for the generated answer
	MaxTokens NullableInt32 `json:"max_tokens,omitempty"`
	Stop NullableCreateAnswerRequestStop `json:"stop,omitempty"`
	// How many answers to generate for each question.
	N NullableInt32 `json:"n,omitempty"`
	// Modify the likelihood of specified tokens appearing in the completion.  Accepts a json object that maps tokens (specified by their token ID in the GPT tokenizer) to an associated bias value from -100 to 100. You can use this [tokenizer tool](/tokenizer?view=bpe) (which works for both GPT-2 and GPT-3) to convert text to token IDs. Mathematically, the bias is added to the logits generated by the model prior to sampling. The exact effect will vary per model, but values between -1 and 1 should decrease or increase likelihood of selection; values like -100 or 100 should result in a ban or exclusive selection of the relevant token.  As an example, you can pass `{\"50256\": -100}` to prevent the <|endoftext|> token from being generated. 
	LogitBias map[string]interface{} `json:"logit_bias,omitempty"`
	// A special boolean flag for showing metadata. If set to `true`, each document entry in the returned JSON will contain a \"metadata\" field.  This flag only takes effect when `file` is set. 
	ReturnMetadata NullableBool `json:"return_metadata,omitempty"`
	// If set to `true`, the returned JSON will include a \"prompt\" field containing the final prompt that was used to request a completion. This is mainly useful for debugging purposes.
	ReturnPrompt NullableBool `json:"return_prompt,omitempty"`
	// If an object name is in the list, we provide the full information of the object; otherwise, we only provide the object ID. Currently we support `completion` and `file` objects for expansion.
	Expand []interface{} `json:"expand,omitempty"`
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. [Learn more](/docs/guides/safety-best-practices/end-user-ids). 
	User *string `json:"user,omitempty"`
}

// NewCreateAnswerRequest instantiates a new CreateAnswerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAnswerRequest(model string, question string, examples [][]string, examplesContext string) *CreateAnswerRequest {
	this := CreateAnswerRequest{}
	this.Model = model
	this.Question = question
	this.Examples = examples
	this.ExamplesContext = examplesContext
	var searchModel string = "ada"
	this.SearchModel = *NewNullableString(&searchModel)
	var maxRerank int32 = 200
	this.MaxRerank = *NewNullableInt32(&maxRerank)
	var temperature float32 = 0
	this.Temperature = *NewNullableFloat32(&temperature)
	var maxTokens int32 = 16
	this.MaxTokens = *NewNullableInt32(&maxTokens)
	var stop CreateAnswerRequestStop
	this.Stop = *NewNullableCreateAnswerRequestStop(&stop)
	var n int32 = 1
	this.N = *NewNullableInt32(&n)
	var returnMetadata bool = false
	this.ReturnMetadata = *NewNullableBool(&returnMetadata)
	var returnPrompt bool = false
	this.ReturnPrompt = *NewNullableBool(&returnPrompt)
	return &this
}

// NewCreateAnswerRequestWithDefaults instantiates a new CreateAnswerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAnswerRequestWithDefaults() *CreateAnswerRequest {
	this := CreateAnswerRequest{}
	var searchModel string = "ada"
	this.SearchModel = *NewNullableString(&searchModel)
	var maxRerank int32 = 200
	this.MaxRerank = *NewNullableInt32(&maxRerank)
	var temperature float32 = 0
	this.Temperature = *NewNullableFloat32(&temperature)
	var maxTokens int32 = 16
	this.MaxTokens = *NewNullableInt32(&maxTokens)
	var stop CreateAnswerRequestStop
	this.Stop = *NewNullableCreateAnswerRequestStop(&stop)
	var n int32 = 1
	this.N = *NewNullableInt32(&n)
	var returnMetadata bool = false
	this.ReturnMetadata = *NewNullableBool(&returnMetadata)
	var returnPrompt bool = false
	this.ReturnPrompt = *NewNullableBool(&returnPrompt)
	return &this
}

// GetModel returns the Model field value
func (o *CreateAnswerRequest) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *CreateAnswerRequest) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *CreateAnswerRequest) SetModel(v string) {
	o.Model = v
}

// GetQuestion returns the Question field value
func (o *CreateAnswerRequest) GetQuestion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Question
}

// GetQuestionOk returns a tuple with the Question field value
// and a boolean to check if the value has been set.
func (o *CreateAnswerRequest) GetQuestionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Question, true
}

// SetQuestion sets field value
func (o *CreateAnswerRequest) SetQuestion(v string) {
	o.Question = v
}

// GetExamples returns the Examples field value
func (o *CreateAnswerRequest) GetExamples() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}

	return o.Examples
}

// GetExamplesOk returns a tuple with the Examples field value
// and a boolean to check if the value has been set.
func (o *CreateAnswerRequest) GetExamplesOk() ([][]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Examples, true
}

// SetExamples sets field value
func (o *CreateAnswerRequest) SetExamples(v [][]string) {
	o.Examples = v
}

// GetExamplesContext returns the ExamplesContext field value
func (o *CreateAnswerRequest) GetExamplesContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExamplesContext
}

// GetExamplesContextOk returns a tuple with the ExamplesContext field value
// and a boolean to check if the value has been set.
func (o *CreateAnswerRequest) GetExamplesContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExamplesContext, true
}

// SetExamplesContext sets field value
func (o *CreateAnswerRequest) SetExamplesContext(v string) {
	o.ExamplesContext = v
}

// GetDocuments returns the Documents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAnswerRequest) GetDocuments() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAnswerRequest) GetDocumentsOk() ([]string, bool) {
	if o == nil || IsNil(o.Documents) {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *CreateAnswerRequest) HasDocuments() bool {
	if o != nil && IsNil(o.Documents) {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []string and assigns it to the Documents field.
func (o *CreateAnswerRequest) SetDocuments(v []string) {
	o.Documents = v
}

// GetFile returns the File field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAnswerRequest) GetFile() string {
	if o == nil || IsNil(o.File.Get()) {
		var ret string
		return ret
	}
	return *o.File.Get()
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAnswerRequest) GetFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.File.Get(), o.File.IsSet()
}

// HasFile returns a boolean if a field has been set.
func (o *CreateAnswerRequest) HasFile() bool {
	if o != nil && o.File.IsSet() {
		return true
	}

	return false
}

// SetFile gets a reference to the given NullableString and assigns it to the File field.
func (o *CreateAnswerRequest) SetFile(v string) {
	o.File.Set(&v)
}
// SetFileNil sets the value for File to be an explicit nil
func (o *CreateAnswerRequest) SetFileNil() {
	o.File.Set(nil)
}

// UnsetFile ensures that no value is present for File, not even an explicit nil
func (o *CreateAnswerRequest) UnsetFile() {
	o.File.Unset()
}

// GetSearchModel returns the SearchModel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAnswerRequest) GetSearchModel() string {
	if o == nil || IsNil(o.SearchModel.Get()) {
		var ret string
		return ret
	}
	return *o.SearchModel.Get()
}

// GetSearchModelOk returns a tuple with the SearchModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAnswerRequest) GetSearchModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SearchModel.Get(), o.SearchModel.IsSet()
}

// HasSearchModel returns a boolean if a field has been set.
func (o *CreateAnswerRequest) HasSearchModel() bool {
	if o != nil && o.SearchModel.IsSet() {
		return true
	}

	return false
}

// SetSearchModel gets a reference to the given NullableString and assigns it to the SearchModel field.
func (o *CreateAnswerRequest) SetSearchModel(v string) {
	o.SearchModel.Set(&v)
}
// SetSearchModelNil sets the value for SearchModel to be an explicit nil
func (o *CreateAnswerRequest) SetSearchModelNil() {
	o.SearchModel.Set(nil)
}

// UnsetSearchModel ensures that no value is present for SearchModel, not even an explicit nil
func (o *CreateAnswerRequest) UnsetSearchModel() {
	o.SearchModel.Unset()
}

// GetMaxRerank returns the MaxRerank field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAnswerRequest) GetMaxRerank() int32 {
	if o == nil || IsNil(o.MaxRerank.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxRerank.Get()
}

// GetMaxRerankOk returns a tuple with the MaxRerank field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAnswerRequest) GetMaxRerankOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxRerank.Get(), o.MaxRerank.IsSet()
}

// HasMaxRerank returns a boolean if a field has been set.
func (o *CreateAnswerRequest) HasMaxRerank() bool {
	if o != nil && o.MaxRerank.IsSet() {
		return true
	}

	return false
}

// SetMaxRerank gets a reference to the given NullableInt32 and assigns it to the MaxRerank field.
func (o *CreateAnswerRequest) SetMaxRerank(v int32) {
	o.MaxRerank.Set(&v)
}
// SetMaxRerankNil sets the value for MaxRerank to be an explicit nil
func (o *CreateAnswerRequest) SetMaxRerankNil() {
	o.MaxRerank.Set(nil)
}

// UnsetMaxRerank ensures that no value is present for MaxRerank, not even an explicit nil
func (o *CreateAnswerRequest) UnsetMaxRerank() {
	o.MaxRerank.Unset()
}

// GetTemperature returns the Temperature field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAnswerRequest) GetTemperature() float32 {
	if o == nil || IsNil(o.Temperature.Get()) {
		var ret float32
		return ret
	}
	return *o.Temperature.Get()
}

// GetTemperatureOk returns a tuple with the Temperature field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAnswerRequest) GetTemperatureOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Temperature.Get(), o.Temperature.IsSet()
}

// HasTemperature returns a boolean if a field has been set.
func (o *CreateAnswerRequest) HasTemperature() bool {
	if o != nil && o.Temperature.IsSet() {
		return true
	}

	return false
}

// SetTemperature gets a reference to the given NullableFloat32 and assigns it to the Temperature field.
func (o *CreateAnswerRequest) SetTemperature(v float32) {
	o.Temperature.Set(&v)
}
// SetTemperatureNil sets the value for Temperature to be an explicit nil
func (o *CreateAnswerRequest) SetTemperatureNil() {
	o.Temperature.Set(nil)
}

// UnsetTemperature ensures that no value is present for Temperature, not even an explicit nil
func (o *CreateAnswerRequest) UnsetTemperature() {
	o.Temperature.Unset()
}

// GetLogprobs returns the Logprobs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAnswerRequest) GetLogprobs() int32 {
	if o == nil || IsNil(o.Logprobs.Get()) {
		var ret int32
		return ret
	}
	return *o.Logprobs.Get()
}

// GetLogprobsOk returns a tuple with the Logprobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAnswerRequest) GetLogprobsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Logprobs.Get(), o.Logprobs.IsSet()
}

// HasLogprobs returns a boolean if a field has been set.
func (o *CreateAnswerRequest) HasLogprobs() bool {
	if o != nil && o.Logprobs.IsSet() {
		return true
	}

	return false
}

// SetLogprobs gets a reference to the given NullableInt32 and assigns it to the Logprobs field.
func (o *CreateAnswerRequest) SetLogprobs(v int32) {
	o.Logprobs.Set(&v)
}
// SetLogprobsNil sets the value for Logprobs to be an explicit nil
func (o *CreateAnswerRequest) SetLogprobsNil() {
	o.Logprobs.Set(nil)
}

// UnsetLogprobs ensures that no value is present for Logprobs, not even an explicit nil
func (o *CreateAnswerRequest) UnsetLogprobs() {
	o.Logprobs.Unset()
}

// GetMaxTokens returns the MaxTokens field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAnswerRequest) GetMaxTokens() int32 {
	if o == nil || IsNil(o.MaxTokens.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxTokens.Get()
}

// GetMaxTokensOk returns a tuple with the MaxTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAnswerRequest) GetMaxTokensOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxTokens.Get(), o.MaxTokens.IsSet()
}

// HasMaxTokens returns a boolean if a field has been set.
func (o *CreateAnswerRequest) HasMaxTokens() bool {
	if o != nil && o.MaxTokens.IsSet() {
		return true
	}

	return false
}

// SetMaxTokens gets a reference to the given NullableInt32 and assigns it to the MaxTokens field.
func (o *CreateAnswerRequest) SetMaxTokens(v int32) {
	o.MaxTokens.Set(&v)
}
// SetMaxTokensNil sets the value for MaxTokens to be an explicit nil
func (o *CreateAnswerRequest) SetMaxTokensNil() {
	o.MaxTokens.Set(nil)
}

// UnsetMaxTokens ensures that no value is present for MaxTokens, not even an explicit nil
func (o *CreateAnswerRequest) UnsetMaxTokens() {
	o.MaxTokens.Unset()
}

// GetStop returns the Stop field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAnswerRequest) GetStop() CreateAnswerRequestStop {
	if o == nil || IsNil(o.Stop.Get()) {
		var ret CreateAnswerRequestStop
		return ret
	}
	return *o.Stop.Get()
}

// GetStopOk returns a tuple with the Stop field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAnswerRequest) GetStopOk() (*CreateAnswerRequestStop, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stop.Get(), o.Stop.IsSet()
}

// HasStop returns a boolean if a field has been set.
func (o *CreateAnswerRequest) HasStop() bool {
	if o != nil && o.Stop.IsSet() {
		return true
	}

	return false
}

// SetStop gets a reference to the given NullableCreateAnswerRequestStop and assigns it to the Stop field.
func (o *CreateAnswerRequest) SetStop(v CreateAnswerRequestStop) {
	o.Stop.Set(&v)
}
// SetStopNil sets the value for Stop to be an explicit nil
func (o *CreateAnswerRequest) SetStopNil() {
	o.Stop.Set(nil)
}

// UnsetStop ensures that no value is present for Stop, not even an explicit nil
func (o *CreateAnswerRequest) UnsetStop() {
	o.Stop.Unset()
}

// GetN returns the N field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAnswerRequest) GetN() int32 {
	if o == nil || IsNil(o.N.Get()) {
		var ret int32
		return ret
	}
	return *o.N.Get()
}

// GetNOk returns a tuple with the N field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAnswerRequest) GetNOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.N.Get(), o.N.IsSet()
}

// HasN returns a boolean if a field has been set.
func (o *CreateAnswerRequest) HasN() bool {
	if o != nil && o.N.IsSet() {
		return true
	}

	return false
}

// SetN gets a reference to the given NullableInt32 and assigns it to the N field.
func (o *CreateAnswerRequest) SetN(v int32) {
	o.N.Set(&v)
}
// SetNNil sets the value for N to be an explicit nil
func (o *CreateAnswerRequest) SetNNil() {
	o.N.Set(nil)
}

// UnsetN ensures that no value is present for N, not even an explicit nil
func (o *CreateAnswerRequest) UnsetN() {
	o.N.Unset()
}

// GetLogitBias returns the LogitBias field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAnswerRequest) GetLogitBias() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.LogitBias
}

// GetLogitBiasOk returns a tuple with the LogitBias field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAnswerRequest) GetLogitBiasOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.LogitBias) {
		return map[string]interface{}{}, false
	}
	return o.LogitBias, true
}

// HasLogitBias returns a boolean if a field has been set.
func (o *CreateAnswerRequest) HasLogitBias() bool {
	if o != nil && IsNil(o.LogitBias) {
		return true
	}

	return false
}

// SetLogitBias gets a reference to the given map[string]interface{} and assigns it to the LogitBias field.
func (o *CreateAnswerRequest) SetLogitBias(v map[string]interface{}) {
	o.LogitBias = v
}

// GetReturnMetadata returns the ReturnMetadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAnswerRequest) GetReturnMetadata() bool {
	if o == nil || IsNil(o.ReturnMetadata.Get()) {
		var ret bool
		return ret
	}
	return *o.ReturnMetadata.Get()
}

// GetReturnMetadataOk returns a tuple with the ReturnMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAnswerRequest) GetReturnMetadataOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReturnMetadata.Get(), o.ReturnMetadata.IsSet()
}

// HasReturnMetadata returns a boolean if a field has been set.
func (o *CreateAnswerRequest) HasReturnMetadata() bool {
	if o != nil && o.ReturnMetadata.IsSet() {
		return true
	}

	return false
}

// SetReturnMetadata gets a reference to the given NullableBool and assigns it to the ReturnMetadata field.
func (o *CreateAnswerRequest) SetReturnMetadata(v bool) {
	o.ReturnMetadata.Set(&v)
}
// SetReturnMetadataNil sets the value for ReturnMetadata to be an explicit nil
func (o *CreateAnswerRequest) SetReturnMetadataNil() {
	o.ReturnMetadata.Set(nil)
}

// UnsetReturnMetadata ensures that no value is present for ReturnMetadata, not even an explicit nil
func (o *CreateAnswerRequest) UnsetReturnMetadata() {
	o.ReturnMetadata.Unset()
}

// GetReturnPrompt returns the ReturnPrompt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAnswerRequest) GetReturnPrompt() bool {
	if o == nil || IsNil(o.ReturnPrompt.Get()) {
		var ret bool
		return ret
	}
	return *o.ReturnPrompt.Get()
}

// GetReturnPromptOk returns a tuple with the ReturnPrompt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAnswerRequest) GetReturnPromptOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReturnPrompt.Get(), o.ReturnPrompt.IsSet()
}

// HasReturnPrompt returns a boolean if a field has been set.
func (o *CreateAnswerRequest) HasReturnPrompt() bool {
	if o != nil && o.ReturnPrompt.IsSet() {
		return true
	}

	return false
}

// SetReturnPrompt gets a reference to the given NullableBool and assigns it to the ReturnPrompt field.
func (o *CreateAnswerRequest) SetReturnPrompt(v bool) {
	o.ReturnPrompt.Set(&v)
}
// SetReturnPromptNil sets the value for ReturnPrompt to be an explicit nil
func (o *CreateAnswerRequest) SetReturnPromptNil() {
	o.ReturnPrompt.Set(nil)
}

// UnsetReturnPrompt ensures that no value is present for ReturnPrompt, not even an explicit nil
func (o *CreateAnswerRequest) UnsetReturnPrompt() {
	o.ReturnPrompt.Unset()
}

// GetExpand returns the Expand field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateAnswerRequest) GetExpand() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}
	return o.Expand
}

// GetExpandOk returns a tuple with the Expand field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateAnswerRequest) GetExpandOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Expand) {
		return nil, false
	}
	return o.Expand, true
}

// HasExpand returns a boolean if a field has been set.
func (o *CreateAnswerRequest) HasExpand() bool {
	if o != nil && IsNil(o.Expand) {
		return true
	}

	return false
}

// SetExpand gets a reference to the given []interface{} and assigns it to the Expand field.
func (o *CreateAnswerRequest) SetExpand(v []interface{}) {
	o.Expand = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *CreateAnswerRequest) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAnswerRequest) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *CreateAnswerRequest) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *CreateAnswerRequest) SetUser(v string) {
	o.User = &v
}

func (o CreateAnswerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAnswerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["model"] = o.Model
	toSerialize["question"] = o.Question
	toSerialize["examples"] = o.Examples
	toSerialize["examples_context"] = o.ExamplesContext
	if o.Documents != nil {
		toSerialize["documents"] = o.Documents
	}
	if o.File.IsSet() {
		toSerialize["file"] = o.File.Get()
	}
	if o.SearchModel.IsSet() {
		toSerialize["search_model"] = o.SearchModel.Get()
	}
	if o.MaxRerank.IsSet() {
		toSerialize["max_rerank"] = o.MaxRerank.Get()
	}
	if o.Temperature.IsSet() {
		toSerialize["temperature"] = o.Temperature.Get()
	}
	if o.Logprobs.IsSet() {
		toSerialize["logprobs"] = o.Logprobs.Get()
	}
	if o.MaxTokens.IsSet() {
		toSerialize["max_tokens"] = o.MaxTokens.Get()
	}
	if o.Stop.IsSet() {
		toSerialize["stop"] = o.Stop.Get()
	}
	if o.N.IsSet() {
		toSerialize["n"] = o.N.Get()
	}
	if o.LogitBias != nil {
		toSerialize["logit_bias"] = o.LogitBias
	}
	if o.ReturnMetadata.IsSet() {
		toSerialize["return_metadata"] = o.ReturnMetadata.Get()
	}
	if o.ReturnPrompt.IsSet() {
		toSerialize["return_prompt"] = o.ReturnPrompt.Get()
	}
	if o.Expand != nil {
		toSerialize["expand"] = o.Expand
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableCreateAnswerRequest struct {
	value *CreateAnswerRequest
	isSet bool
}

func (v NullableCreateAnswerRequest) Get() *CreateAnswerRequest {
	return v.value
}

func (v *NullableCreateAnswerRequest) Set(val *CreateAnswerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAnswerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAnswerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAnswerRequest(val *CreateAnswerRequest) *NullableCreateAnswerRequest {
	return &NullableCreateAnswerRequest{value: val, isSet: true}
}

func (v NullableCreateAnswerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAnswerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


