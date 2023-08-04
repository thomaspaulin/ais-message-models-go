/*
Ais-Stream WebsocketObjects

OpenAPI 3.0 definitions for the data models used by aisstream.io.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package aisStream

import (
	"encoding/json"
)

// AisStreamMessage struct for AisStreamMessage
type AisStreamMessage struct {
	MetaData map[string]interface{} `json:"MetaData"`
	MessageType AisMessageTypes `json:"MessageType"`
	Message AisStreamMessageMessage `json:"Message"`
}

// NewAisStreamMessage instantiates a new AisStreamMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAisStreamMessage(metaData map[string]interface{}, messageType AisMessageTypes, message AisStreamMessageMessage) *AisStreamMessage {
	this := AisStreamMessage{}
	this.MetaData = metaData
	this.MessageType = messageType
	this.Message = message
	return &this
}

// NewAisStreamMessageWithDefaults instantiates a new AisStreamMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAisStreamMessageWithDefaults() *AisStreamMessage {
	this := AisStreamMessage{}
	return &this
}

// GetMetaData returns the MetaData field value
func (o *AisStreamMessage) GetMetaData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.MetaData
}

// GetMetaDataOk returns a tuple with the MetaData field value
// and a boolean to check if the value has been set.
func (o *AisStreamMessage) GetMetaDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetaData, true
}

// SetMetaData sets field value
func (o *AisStreamMessage) SetMetaData(v map[string]interface{}) {
	o.MetaData = v
}

// GetMessageType returns the MessageType field value
func (o *AisStreamMessage) GetMessageType() AisMessageTypes {
	if o == nil {
		var ret AisMessageTypes
		return ret
	}

	return o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value
// and a boolean to check if the value has been set.
func (o *AisStreamMessage) GetMessageTypeOk() (*AisMessageTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageType, true
}

// SetMessageType sets field value
func (o *AisStreamMessage) SetMessageType(v AisMessageTypes) {
	o.MessageType = v
}

// GetMessage returns the Message field value
func (o *AisStreamMessage) GetMessage() AisStreamMessageMessage {
	if o == nil {
		var ret AisStreamMessageMessage
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *AisStreamMessage) GetMessageOk() (*AisStreamMessageMessage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *AisStreamMessage) SetMessage(v AisStreamMessageMessage) {
	o.Message = v
}

func (o AisStreamMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["MetaData"] = o.MetaData
	}
	if true {
		toSerialize["MessageType"] = o.MessageType
	}
	if true {
		toSerialize["Message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableAisStreamMessage struct {
	value *AisStreamMessage
	isSet bool
}

func (v NullableAisStreamMessage) Get() *AisStreamMessage {
	return v.value
}

func (v *NullableAisStreamMessage) Set(val *AisStreamMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableAisStreamMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableAisStreamMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAisStreamMessage(val *AisStreamMessage) *NullableAisStreamMessage {
	return &NullableAisStreamMessage{value: val, isSet: true}
}

func (v NullableAisStreamMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAisStreamMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

