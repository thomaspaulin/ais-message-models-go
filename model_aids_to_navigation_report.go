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

// AidsToNavigationReport struct for AidsToNavigationReport
type AidsToNavigationReport struct {
	MessageID int32 `json:"MessageID"`
	RepeatIndicator int32 `json:"RepeatIndicator"`
	UserID int32 `json:"UserID"`
	Valid bool `json:"Valid"`
	Type int32 `json:"Type"`
	Name string `json:"Name"`
	PositionAccuracy bool `json:"PositionAccuracy"`
	Longitude float64 `json:"Longitude"`
	Latitude float64 `json:"Latitude"`
	Dimension ShipStaticDataDimension `json:"Dimension"`
	Fixtype int32 `json:"Fixtype"`
	Timestamp int32 `json:"Timestamp"`
	OffPosition bool `json:"OffPosition"`
	AtoN int32 `json:"AtoN"`
	Raim bool `json:"Raim"`
	VirtualAtoN bool `json:"VirtualAtoN"`
	AssignedMode bool `json:"AssignedMode"`
	Spare bool `json:"Spare"`
	NameExtension string `json:"NameExtension"`
}

// NewAidsToNavigationReport instantiates a new AidsToNavigationReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAidsToNavigationReport(messageID int32, repeatIndicator int32, userID int32, valid bool, type_ int32, name string, positionAccuracy bool, longitude float64, latitude float64, dimension ShipStaticDataDimension, fixtype int32, timestamp int32, offPosition bool, atoN int32, raim bool, virtualAtoN bool, assignedMode bool, spare bool, nameExtension string) *AidsToNavigationReport {
	this := AidsToNavigationReport{}
	this.MessageID = messageID
	this.RepeatIndicator = repeatIndicator
	this.UserID = userID
	this.Valid = valid
	this.Type = type_
	this.Name = name
	this.PositionAccuracy = positionAccuracy
	this.Longitude = longitude
	this.Latitude = latitude
	this.Dimension = dimension
	this.Fixtype = fixtype
	this.Timestamp = timestamp
	this.OffPosition = offPosition
	this.AtoN = atoN
	this.Raim = raim
	this.VirtualAtoN = virtualAtoN
	this.AssignedMode = assignedMode
	this.Spare = spare
	this.NameExtension = nameExtension
	return &this
}

// NewAidsToNavigationReportWithDefaults instantiates a new AidsToNavigationReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAidsToNavigationReportWithDefaults() *AidsToNavigationReport {
	this := AidsToNavigationReport{}
	return &this
}

// GetMessageID returns the MessageID field value
func (o *AidsToNavigationReport) GetMessageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MessageID
}

// GetMessageIDOk returns a tuple with the MessageID field value
// and a boolean to check if the value has been set.
func (o *AidsToNavigationReport) GetMessageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageID, true
}

// SetMessageID sets field value
func (o *AidsToNavigationReport) SetMessageID(v int32) {
	o.MessageID = v
}

// GetRepeatIndicator returns the RepeatIndicator field value
func (o *AidsToNavigationReport) GetRepeatIndicator() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RepeatIndicator
}

// GetRepeatIndicatorOk returns a tuple with the RepeatIndicator field value
// and a boolean to check if the value has been set.
func (o *AidsToNavigationReport) GetRepeatIndicatorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepeatIndicator, true
}

// SetRepeatIndicator sets field value
func (o *AidsToNavigationReport) SetRepeatIndicator(v int32) {
	o.RepeatIndicator = v
}

// GetUserID returns the UserID field value
func (o *AidsToNavigationReport) GetUserID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value
// and a boolean to check if the value has been set.
func (o *AidsToNavigationReport) GetUserIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserID, true
}

// SetUserID sets field value
func (o *AidsToNavigationReport) SetUserID(v int32) {
	o.UserID = v
}

// GetValid returns the Valid field value
func (o *AidsToNavigationReport) GetValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Valid
}

// GetValidOk returns a tuple with the Valid field value
// and a boolean to check if the value has been set.
func (o *AidsToNavigationReport) GetValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Valid, true
}

// SetValid sets field value
func (o *AidsToNavigationReport) SetValid(v bool) {
	o.Valid = v
}

// GetType returns the Type field value
func (o *AidsToNavigationReport) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AidsToNavigationReport) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AidsToNavigationReport) SetType(v int32) {
	o.Type = v
}

// GetName returns the Name field value
func (o *AidsToNavigationReport) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AidsToNavigationReport) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AidsToNavigationReport) SetName(v string) {
	o.Name = v
}

// GetPositionAccuracy returns the PositionAccuracy field value
func (o *AidsToNavigationReport) GetPositionAccuracy() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PositionAccuracy
}

// GetPositionAccuracyOk returns a tuple with the PositionAccuracy field value
// and a boolean to check if the value has been set.
func (o *AidsToNavigationReport) GetPositionAccuracyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PositionAccuracy, true
}

// SetPositionAccuracy sets field value
func (o *AidsToNavigationReport) SetPositionAccuracy(v bool) {
	o.PositionAccuracy = v
}

// GetLongitude returns the Longitude field value
func (o *AidsToNavigationReport) GetLongitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *AidsToNavigationReport) GetLongitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *AidsToNavigationReport) SetLongitude(v float64) {
	o.Longitude = v
}

// GetLatitude returns the Latitude field value
func (o *AidsToNavigationReport) GetLatitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *AidsToNavigationReport) GetLatitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *AidsToNavigationReport) SetLatitude(v float64) {
	o.Latitude = v
}

// GetDimension returns the Dimension field value
func (o *AidsToNavigationReport) GetDimension() ShipStaticDataDimension {
	if o == nil {
		var ret ShipStaticDataDimension
		return ret
	}

	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value
// and a boolean to check if the value has been set.
func (o *AidsToNavigationReport) GetDimensionOk() (*ShipStaticDataDimension, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimension, true
}

// SetDimension sets field value
func (o *AidsToNavigationReport) SetDimension(v ShipStaticDataDimension) {
	o.Dimension = v
}

// GetFixtype returns the Fixtype field value
func (o *AidsToNavigationReport) GetFixtype() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Fixtype
}

// GetFixtypeOk returns a tuple with the Fixtype field value
// and a boolean to check if the value has been set.
func (o *AidsToNavigationReport) GetFixtypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fixtype, true
}

// SetFixtype sets field value
func (o *AidsToNavigationReport) SetFixtype(v int32) {
	o.Fixtype = v
}

// GetTimestamp returns the Timestamp field value
func (o *AidsToNavigationReport) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *AidsToNavigationReport) GetTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *AidsToNavigationReport) SetTimestamp(v int32) {
	o.Timestamp = v
}

// GetOffPosition returns the OffPosition field value
func (o *AidsToNavigationReport) GetOffPosition() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.OffPosition
}

// GetOffPositionOk returns a tuple with the OffPosition field value
// and a boolean to check if the value has been set.
func (o *AidsToNavigationReport) GetOffPositionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OffPosition, true
}

// SetOffPosition sets field value
func (o *AidsToNavigationReport) SetOffPosition(v bool) {
	o.OffPosition = v
}

// GetAtoN returns the AtoN field value
func (o *AidsToNavigationReport) GetAtoN() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AtoN
}

// GetAtoNOk returns a tuple with the AtoN field value
// and a boolean to check if the value has been set.
func (o *AidsToNavigationReport) GetAtoNOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AtoN, true
}

// SetAtoN sets field value
func (o *AidsToNavigationReport) SetAtoN(v int32) {
	o.AtoN = v
}

// GetRaim returns the Raim field value
func (o *AidsToNavigationReport) GetRaim() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Raim
}

// GetRaimOk returns a tuple with the Raim field value
// and a boolean to check if the value has been set.
func (o *AidsToNavigationReport) GetRaimOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Raim, true
}

// SetRaim sets field value
func (o *AidsToNavigationReport) SetRaim(v bool) {
	o.Raim = v
}

// GetVirtualAtoN returns the VirtualAtoN field value
func (o *AidsToNavigationReport) GetVirtualAtoN() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.VirtualAtoN
}

// GetVirtualAtoNOk returns a tuple with the VirtualAtoN field value
// and a boolean to check if the value has been set.
func (o *AidsToNavigationReport) GetVirtualAtoNOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VirtualAtoN, true
}

// SetVirtualAtoN sets field value
func (o *AidsToNavigationReport) SetVirtualAtoN(v bool) {
	o.VirtualAtoN = v
}

// GetAssignedMode returns the AssignedMode field value
func (o *AidsToNavigationReport) GetAssignedMode() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AssignedMode
}

// GetAssignedModeOk returns a tuple with the AssignedMode field value
// and a boolean to check if the value has been set.
func (o *AidsToNavigationReport) GetAssignedModeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedMode, true
}

// SetAssignedMode sets field value
func (o *AidsToNavigationReport) SetAssignedMode(v bool) {
	o.AssignedMode = v
}

// GetSpare returns the Spare field value
func (o *AidsToNavigationReport) GetSpare() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Spare
}

// GetSpareOk returns a tuple with the Spare field value
// and a boolean to check if the value has been set.
func (o *AidsToNavigationReport) GetSpareOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spare, true
}

// SetSpare sets field value
func (o *AidsToNavigationReport) SetSpare(v bool) {
	o.Spare = v
}

// GetNameExtension returns the NameExtension field value
func (o *AidsToNavigationReport) GetNameExtension() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NameExtension
}

// GetNameExtensionOk returns a tuple with the NameExtension field value
// and a boolean to check if the value has been set.
func (o *AidsToNavigationReport) GetNameExtensionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameExtension, true
}

// SetNameExtension sets field value
func (o *AidsToNavigationReport) SetNameExtension(v string) {
	o.NameExtension = v
}

func (o AidsToNavigationReport) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["MessageID"] = o.MessageID
	}
	if true {
		toSerialize["RepeatIndicator"] = o.RepeatIndicator
	}
	if true {
		toSerialize["UserID"] = o.UserID
	}
	if true {
		toSerialize["Valid"] = o.Valid
	}
	if true {
		toSerialize["Type"] = o.Type
	}
	if true {
		toSerialize["Name"] = o.Name
	}
	if true {
		toSerialize["PositionAccuracy"] = o.PositionAccuracy
	}
	if true {
		toSerialize["Longitude"] = o.Longitude
	}
	if true {
		toSerialize["Latitude"] = o.Latitude
	}
	if true {
		toSerialize["Dimension"] = o.Dimension
	}
	if true {
		toSerialize["Fixtype"] = o.Fixtype
	}
	if true {
		toSerialize["Timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["OffPosition"] = o.OffPosition
	}
	if true {
		toSerialize["AtoN"] = o.AtoN
	}
	if true {
		toSerialize["Raim"] = o.Raim
	}
	if true {
		toSerialize["VirtualAtoN"] = o.VirtualAtoN
	}
	if true {
		toSerialize["AssignedMode"] = o.AssignedMode
	}
	if true {
		toSerialize["Spare"] = o.Spare
	}
	if true {
		toSerialize["NameExtension"] = o.NameExtension
	}
	return json.Marshal(toSerialize)
}

type NullableAidsToNavigationReport struct {
	value *AidsToNavigationReport
	isSet bool
}

func (v NullableAidsToNavigationReport) Get() *AidsToNavigationReport {
	return v.value
}

func (v *NullableAidsToNavigationReport) Set(val *AidsToNavigationReport) {
	v.value = val
	v.isSet = true
}

func (v NullableAidsToNavigationReport) IsSet() bool {
	return v.isSet
}

func (v *NullableAidsToNavigationReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAidsToNavigationReport(val *AidsToNavigationReport) *NullableAidsToNavigationReport {
	return &NullableAidsToNavigationReport{value: val, isSet: true}
}

func (v NullableAidsToNavigationReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAidsToNavigationReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


