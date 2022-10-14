package spec

import (
	"time"
)

type Simple struct {
	Boolean  *bool      `json:"Boolean,omitempty"`
	Byte     *[]byte    `json:"Byte,omitempty"`
	Date     *time.Time `json:"Date,omitempty"`
	DateTime *time.Time `json:"DateTime,omitempty"`
	Double   *float64   `json:"Double,omitempty"`
	Float    *float32   `json:"Float,omitempty"`
	Int32    *int32     `json:"Int32,omitempty"`
	Int64    *int64     `json:"Int64,omitempty"`
	Password *string    `json:"Password,omitempty"`
	String   *string    `json:"String,omitempty"`
}

func (o *Simple) GetBoolean() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return *o.Boolean
}

func (o *Simple) GetBooleanOk() bool {
	if o == nil || o.Boolean == nil {
		var ret bool
		return ret
	}
	return *o.Boolean
}
func (o *Simple) SetBoolean(value bool) {
	o.Boolean = &value
}

func (o *Simple) GetByte() []byte {
	if o == nil {
		var ret []byte
		return ret
	}
	return *o.Byte
}

func (o *Simple) GetByteOk() []byte {
	if o == nil || o.Byte == nil {
		var ret []byte
		return ret
	}
	return *o.Byte
}
func (o *Simple) SetByte(value []byte) {
	o.Byte = &value
}

func (o *Simple) GetDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return *o.Date
}

func (o *Simple) GetDateOk() time.Time {
	if o == nil || o.Date == nil {
		var ret time.Time
		return ret
	}
	return *o.Date
}
func (o *Simple) SetDate(value time.Time) {
	o.Date = &value
}

func (o *Simple) GetDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}
	return *o.DateTime
}

func (o *Simple) GetDateTimeOk() time.Time {
	if o == nil || o.DateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.DateTime
}
func (o *Simple) SetDateTime(value time.Time) {
	o.DateTime = &value
}

func (o *Simple) GetDouble() float64 {
	if o == nil {
		var ret float64
		return ret
	}
	return *o.Double
}

func (o *Simple) GetDoubleOk() float64 {
	if o == nil || o.Double == nil {
		var ret float64
		return ret
	}
	return *o.Double
}
func (o *Simple) SetDouble(value float64) {
	o.Double = &value
}

func (o *Simple) GetFloat() float32 {
	if o == nil {
		var ret float32
		return ret
	}
	return *o.Float
}

func (o *Simple) GetFloatOk() float32 {
	if o == nil || o.Float == nil {
		var ret float32
		return ret
	}
	return *o.Float
}
func (o *Simple) SetFloat(value float32) {
	o.Float = &value
}

func (o *Simple) GetInt32() int32 {
	if o == nil {
		var ret int32
		return ret
	}
	return *o.Int32
}

func (o *Simple) GetInt32Ok() int32 {
	if o == nil || o.Int32 == nil {
		var ret int32
		return ret
	}
	return *o.Int32
}
func (o *Simple) SetInt32(value int32) {
	o.Int32 = &value
}

func (o *Simple) GetInt64() int64 {
	if o == nil {
		var ret int64
		return ret
	}
	return *o.Int64
}

func (o *Simple) GetInt64Ok() int64 {
	if o == nil || o.Int64 == nil {
		var ret int64
		return ret
	}
	return *o.Int64
}
func (o *Simple) SetInt64(value int64) {
	o.Int64 = &value
}

func (o *Simple) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}
	return *o.Password
}

func (o *Simple) GetPasswordOk() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}
func (o *Simple) SetPassword(value string) {
	o.Password = &value
}

func (o *Simple) GetString() string {
	if o == nil {
		var ret string
		return ret
	}
	return *o.String
}

func (o *Simple) GetStringOk() string {
	if o == nil || o.String == nil {
		var ret string
		return ret
	}
	return *o.String
}
func (o *Simple) SetString(value string) {
	o.String = &value
}
