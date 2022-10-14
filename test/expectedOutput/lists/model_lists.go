package spec

import (
	"time"
)

type Lists struct {
	BinaryList   *[][]byte    `json:"BinaryList,omitempty"`
	BooleanList  *[]bool      `json:"BooleanList,omitempty"`
	ByteList     *[][]byte    `json:"ByteList,omitempty"`
	DateList     *[]time.Time `json:"DateList,omitempty"`
	DateTimeList *[]time.Time `json:"DateTimeList,omitempty"`
	DoubleList   *[]float64   `json:"DoubleList,omitempty"`
	FloatList    *[]float32   `json:"FloatList,omitempty"`
	Int32List    *[]int32     `json:"Int32List,omitempty"`
	Int64List    *[]int64     `json:"Int64List,omitempty"`
	PasswordList *[]string    `json:"PasswordList,omitempty"`
	StringList   *[]string    `json:"StringList,omitempty"`
}

func (o *Lists) GetBinaryList() [][]byte {
	if o == nil {
		var ret [][]byte
		return ret
	}
	return *o.BinaryList
}

func (o *Lists) GetBinaryListOk() [][]byte {
	if o == nil || o.BinaryList == nil {
		var ret [][]byte
		return ret
	}
	return *o.BinaryList
}
func (o *Lists) SetBinaryList(value [][]byte) {
	o.BinaryList = &value
}

func (o *Lists) GetBooleanList() []bool {
	if o == nil {
		var ret []bool
		return ret
	}
	return *o.BooleanList
}

func (o *Lists) GetBooleanListOk() []bool {
	if o == nil || o.BooleanList == nil {
		var ret []bool
		return ret
	}
	return *o.BooleanList
}
func (o *Lists) SetBooleanList(value []bool) {
	o.BooleanList = &value
}

func (o *Lists) GetByteList() [][]byte {
	if o == nil {
		var ret [][]byte
		return ret
	}
	return *o.ByteList
}

func (o *Lists) GetByteListOk() [][]byte {
	if o == nil || o.ByteList == nil {
		var ret [][]byte
		return ret
	}
	return *o.ByteList
}
func (o *Lists) SetByteList(value [][]byte) {
	o.ByteList = &value
}

func (o *Lists) GetDateList() []time.Time {
	if o == nil {
		var ret []time.Time
		return ret
	}
	return *o.DateList
}

func (o *Lists) GetDateListOk() []time.Time {
	if o == nil || o.DateList == nil {
		var ret []time.Time
		return ret
	}
	return *o.DateList
}
func (o *Lists) SetDateList(value []time.Time) {
	o.DateList = &value
}

func (o *Lists) GetDateTimeList() []time.Time {
	if o == nil {
		var ret []time.Time
		return ret
	}
	return *o.DateTimeList
}

func (o *Lists) GetDateTimeListOk() []time.Time {
	if o == nil || o.DateTimeList == nil {
		var ret []time.Time
		return ret
	}
	return *o.DateTimeList
}
func (o *Lists) SetDateTimeList(value []time.Time) {
	o.DateTimeList = &value
}

func (o *Lists) GetDoubleList() []float64 {
	if o == nil {
		var ret []float64
		return ret
	}
	return *o.DoubleList
}

func (o *Lists) GetDoubleListOk() []float64 {
	if o == nil || o.DoubleList == nil {
		var ret []float64
		return ret
	}
	return *o.DoubleList
}
func (o *Lists) SetDoubleList(value []float64) {
	o.DoubleList = &value
}

func (o *Lists) GetFloatList() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}
	return *o.FloatList
}

func (o *Lists) GetFloatListOk() []float32 {
	if o == nil || o.FloatList == nil {
		var ret []float32
		return ret
	}
	return *o.FloatList
}
func (o *Lists) SetFloatList(value []float32) {
	o.FloatList = &value
}

func (o *Lists) GetInt32List() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return *o.Int32List
}

func (o *Lists) GetInt32ListOk() []int32 {
	if o == nil || o.Int32List == nil {
		var ret []int32
		return ret
	}
	return *o.Int32List
}
func (o *Lists) SetInt32List(value []int32) {
	o.Int32List = &value
}

func (o *Lists) GetInt64List() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return *o.Int64List
}

func (o *Lists) GetInt64ListOk() []int64 {
	if o == nil || o.Int64List == nil {
		var ret []int64
		return ret
	}
	return *o.Int64List
}
func (o *Lists) SetInt64List(value []int64) {
	o.Int64List = &value
}

func (o *Lists) GetPasswordList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return *o.PasswordList
}

func (o *Lists) GetPasswordListOk() []string {
	if o == nil || o.PasswordList == nil {
		var ret []string
		return ret
	}
	return *o.PasswordList
}
func (o *Lists) SetPasswordList(value []string) {
	o.PasswordList = &value
}

func (o *Lists) GetStringList() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return *o.StringList
}

func (o *Lists) GetStringListOk() []string {
	if o == nil || o.StringList == nil {
		var ret []string
		return ret
	}
	return *o.StringList
}
func (o *Lists) SetStringList(value []string) {
	o.StringList = &value
}
