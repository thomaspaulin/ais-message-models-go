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

// ExtendedClassBPositionReport struct for ExtendedClassBPositionReport
type ExtendedClassBPositionReport struct {
	MessageID int32 `json:"MessageID"`
	RepeatIndicator int32 `json:"RepeatIndicator"`
	UserID int32 `json:"UserID"`
	Valid bool `json:"Valid"`
	Spare1 int32 `json:"Spare1"`
	Sog float64 `json:"Sog"`
	PositionAccuracy bool `json:"PositionAccuracy"`
	Longitude float64 `json:"Longitude"`
	Latitude float64 `json:"Latitude"`
	Cog float64 `json:"Cog"`
	TrueHeading int32 `json:"TrueHeading"`
	Timestamp int32 `json:"Timestamp"`
	Spare2 int32 `json:"Spare2"`
	Name string `json:"Name"`
	Type int32 `json:"Type"`
	Dimension ShipStaticDataDimension `json:"Dimension"`
	FixType int32 `json:"FixType"`
	Raim bool `json:"Raim"`
	Dte bool `json:"Dte"`
	AssignedMode bool `json:"AssignedMode"`
	Spare3 int32 `json:"Spare3"`
}

// NewExtendedClassBPositionReport instantiates a new ExtendedClassBPositionReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtendedClassBPositionReport(messageID int32, repeatIndicator int32, userID int32, valid bool, spare1 int32, sog float64, positionAccuracy bool, longitude float64, latitude float64, cog float64, trueHeading int32, timestamp int32, spare2 int32, name string, type_ int32, dimension ShipStaticDataDimension, fixType int32, raim bool, dte bool, assignedMode bool, spare3 int32) *ExtendedClassBPositionReport {
	this := ExtendedClassBPositionReport{}
	this.MessageID = messageID
	this.RepeatIndicator = repeatIndicator
	this.UserID = userID
	this.Valid = valid
	this.Spare1 = spare1
	this.Sog = sog
	this.PositionAccuracy = positionAccuracy
	this.Longitude = longitude
	this.Latitude = latitude
	this.Cog = cog
	this.TrueHeading = trueHeading
	this.Timestamp = timestamp
	this.Spare2 = spare2
	this.Name = name
	this.Type = type_
	this.Dimension = dimension
	this.FixType = fixType
	this.Raim = raim
	this.Dte = dte
	this.AssignedMode = assignedMode
	this.Spare3 = spare3
	return &this
}

// NewExtendedClassBPositionReportWithDefaults instantiates a new ExtendedClassBPositionReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtendedClassBPositionReportWithDefaults() *ExtendedClassBPositionReport {
	this := ExtendedClassBPositionReport{}
	return &this
}

// GetMessageID returns the MessageID field value
func (o *ExtendedClassBPositionReport) GetMessageID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MessageID
}

// GetMessageIDOk returns a tuple with the MessageID field value
// and a boolean to check if the value has been set.
func (o *ExtendedClassBPositionReport) GetMessageIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageID, true
}

// SetMessageID sets field value
func (o *ExtendedClassBPositionReport) SetMessageID(v int32) {
	o.MessageID = v
}

// GetRepeatIndicator returns the RepeatIndicator field value
func (o *ExtendedClassBPositionReport) GetRepeatIndicator() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RepeatIndicator
}

// GetRepeatIndicatorOk returns a tuple with the RepeatIndicator field value
// and a boolean to check if the value has been set.
func (o *ExtendedClassBPositionReport) GetRepeatIndicatorOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RepeatIndicator, true
}

// SetRepeatIndicator sets field value
func (o *ExtendedClassBPositionReport) SetRepeatIndicator(v int32) {
	o.RepeatIndicator = v
}

// GetUserID returns the UserID field value
func (o *ExtendedClassBPositionReport) GetUserID() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value
// and a boolean to check if the value has been set.
func (o *ExtendedClassBPositionReport) GetUserIDOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserID, true
}

// SetUserID sets field value
func (o *ExtendedClassBPositionReport) SetUserID(v int32) {
	o.UserID = v
}

// GetValid returns the Valid field value
func (o *ExtendedClassBPositionReport) GetValid() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Valid
}

// GetValidOk returns a tuple with the Valid field value
// and a boolean to check if the value has been set.
func (o *ExtendedClassBPositionReport) GetValidOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Valid, true
}

// SetValid sets field value
func (o *ExtendedClassBPositionReport) SetValid(v bool) {
	o.Valid = v
}

// GetSpare1 returns the Spare1 field value
func (o *ExtendedClassBPositionReport) GetSpare1() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Spare1
}

// GetSpare1Ok returns a tuple with the Spare1 field value
// and a boolean to check if the value has been set.
func (o *ExtendedClassBPositionReport) GetSpare1Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spare1, true
}

// SetSpare1 sets field value
func (o *ExtendedClassBPositionReport) SetSpare1(v int32) {
	o.Spare1 = v
}

// GetSog returns the Sog field value
func (o *ExtendedClassBPositionReport) GetSog() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Sog
}

// GetSogOk returns a tuple with the Sog field value
// and a boolean to check if the value has been set.
func (o *ExtendedClassBPositionReport) GetSogOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sog, true
}

// SetSog sets field value
func (o *ExtendedClassBPositionReport) SetSog(v float64) {
	o.Sog = v
}

// GetPositionAccuracy returns the PositionAccuracy field value
func (o *ExtendedClassBPositionReport) GetPositionAccuracy() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PositionAccuracy
}

// GetPositionAccuracyOk returns a tuple with the PositionAccuracy field value
// and a boolean to check if the value has been set.
func (o *ExtendedClassBPositionReport) GetPositionAccuracyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PositionAccuracy, true
}

// SetPositionAccuracy sets field value
func (o *ExtendedClassBPositionReport) SetPositionAccuracy(v bool) {
	o.PositionAccuracy = v
}

// GetLongitude returns the Longitude field value
func (o *ExtendedClassBPositionReport) GetLongitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *ExtendedClassBPositionReport) GetLongitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *ExtendedClassBPositionReport) SetLongitude(v float64) {
	o.Longitude = v
}

// GetLatitude returns the Latitude field value
func (o *ExtendedClassBPositionReport) GetLatitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *ExtendedClassBPositionReport) GetLatitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *ExtendedClassBPositionReport) SetLatitude(v float64) {
	o.Latitude = v
}

// GetCog returns the Cog field value
func (o *ExtendedClassBPositionReport) GetCog() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Cog
}

// GetCogOk returns a tuple with the Cog field value
// and a boolean to check if the value has been set.
func (o *ExtendedClassBPositionReport) GetCogOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cog, true
}

// SetCog sets field value
func (o *ExtendedClassBPositionReport) SetCog(v float64) {
	o.Cog = v
}

// GetTrueHeading returns the TrueHeading field value
func (o *ExtendedClassBPositionReport) GetTrueHeading() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TrueHeading
}

// GetTrueHeadingOk returns a tuple with the TrueHeading field value
// and a boolean to check if the value has been set.
func (o *ExtendedClassBPositionReport) GetTrueHeadingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrueHeading, true
}

// SetTrueHeading sets field value
func (o *ExtendedClassBPositionReport) SetTrueHeading(v int32) {
	o.TrueHeading = v
}

// GetTimestamp returns the Timestamp field value
func (o *ExtendedClassBPositionReport) GetTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ExtendedClassBPositionReport) GetTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ExtendedClassBPositionReport) SetTimestamp(v int32) {
	o.Timestamp = v
}

// GetSpare2 returns the Spare2 field value
func (o *ExtendedClassBPositionReport) GetSpare2() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Spare2
}

// GetSpare2Ok returns a tuple with the Spare2 field value
// and a boolean to check if the value has been set.
func (o *ExtendedClassBPositionReport) GetSpare2Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spare2, true
}

// SetSpare2 sets field value
func (o *ExtendedClassBPositionReport) SetSpare2(v int32) {
	o.Spare2 = v
}

// GetName returns the Name field value
func (o *ExtendedClassBPositionReport) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExtendedClassBPositionReport) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExtendedClassBPositionReport) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *ExtendedClassBPositionReport) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ExtendedClassBPositionReport) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ExtendedClassBPositionReport) SetType(v int32) {
	o.Type = v
}

// GetDimension returns the Dimension field value
func (o *ExtendedClassBPositionReport) GetDimension() ShipStaticDataDimension {
	if o == nil {
		var ret ShipStaticDataDimension
		return ret
	}

	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value
// and a boolean to check if the value has been set.
func (o *ExtendedClassBPositionReport) GetDimensionOk() (*ShipStaticDataDimension, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimension, true
}

// SetDimension sets field value
func (o *ExtendedClassBPositionReport) SetDimension(v ShipStaticDataDimension) {
	o.Dimension = v
}

// GetFixType returns the FixType field value
func (o *ExtendedClassBPositionReport) GetFixType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FixType
}

// GetFixTypeOk returns a tuple with the FixType field value
// and a boolean to check if the value has been set.
func (o *ExtendedClassBPositionReport) GetFixTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FixType, true
}

// SetFixType sets field value
func (o *ExtendedClassBPositionReport) SetFixType(v int32) {
	o.FixType = v
}

// GetRaim returns the Raim field value
func (o *ExtendedClassBPositionReport) GetRaim() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Raim
}

// GetRaimOk returns a tuple with the Raim field value
// and a boolean to check if the value has been set.
func (o *ExtendedClassBPositionReport) GetRaimOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Raim, true
}

// SetRaim sets field value
func (o *ExtendedClassBPositionReport) SetRaim(v bool) {
	o.Raim = v
}

// GetDte returns the Dte field value
func (o *ExtendedClassBPositionReport) GetDte() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Dte
}

// GetDteOk returns a tuple with the Dte field value
// and a boolean to check if the value has been set.
func (o *ExtendedClassBPositionReport) GetDteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dte, true
}

// SetDte sets field value
func (o *ExtendedClassBPositionReport) SetDte(v bool) {
	o.Dte = v
}

// GetAssignedMode returns the AssignedMode field value
func (o *ExtendedClassBPositionReport) GetAssignedMode() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AssignedMode
}

// GetAssignedModeOk returns a tuple with the AssignedMode field value
// and a boolean to check if the value has been set.
func (o *ExtendedClassBPositionReport) GetAssignedModeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssignedMode, true
}

// SetAssignedMode sets field value
func (o *ExtendedClassBPositionReport) SetAssignedMode(v bool) {
	o.AssignedMode = v
}

// GetSpare3 returns the Spare3 field value
func (o *ExtendedClassBPositionReport) GetSpare3() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Spare3
}

// GetSpare3Ok returns a tuple with the Spare3 field value
// and a boolean to check if the value has been set.
func (o *ExtendedClassBPositionReport) GetSpare3Ok() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spare3, true
}

// SetSpare3 sets field value
func (o *ExtendedClassBPositionReport) SetSpare3(v int32) {
	o.Spare3 = v
}

func (o ExtendedClassBPositionReport) MarshalJSON() ([]byte, error) {
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
		toSerialize["Spare1"] = o.Spare1
	}
	if true {
		toSerialize["Sog"] = o.Sog
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
		toSerialize["Cog"] = o.Cog
	}
	if true {
		toSerialize["TrueHeading"] = o.TrueHeading
	}
	if true {
		toSerialize["Timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["Spare2"] = o.Spare2
	}
	if true {
		toSerialize["Name"] = o.Name
	}
	if true {
		toSerialize["Type"] = o.Type
	}
	if true {
		toSerialize["Dimension"] = o.Dimension
	}
	if true {
		toSerialize["FixType"] = o.FixType
	}
	if true {
		toSerialize["Raim"] = o.Raim
	}
	if true {
		toSerialize["Dte"] = o.Dte
	}
	if true {
		toSerialize["AssignedMode"] = o.AssignedMode
	}
	if true {
		toSerialize["Spare3"] = o.Spare3
	}
	return json.Marshal(toSerialize)
}

type NullableExtendedClassBPositionReport struct {
	value *ExtendedClassBPositionReport
	isSet bool
}

func (v NullableExtendedClassBPositionReport) Get() *ExtendedClassBPositionReport {
	return v.value
}

func (v *NullableExtendedClassBPositionReport) Set(val *ExtendedClassBPositionReport) {
	v.value = val
	v.isSet = true
}

func (v NullableExtendedClassBPositionReport) IsSet() bool {
	return v.isSet
}

func (v *NullableExtendedClassBPositionReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtendedClassBPositionReport(val *ExtendedClassBPositionReport) *NullableExtendedClassBPositionReport {
	return &NullableExtendedClassBPositionReport{value: val, isSet: true}
}

func (v NullableExtendedClassBPositionReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtendedClassBPositionReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


