package spec

type GenericComponent struct {
	GenericBoolean *bool        `json:"genericBoolean,omitempty"`
	GenericEnum    *GenericEnum `json:"genericEnum,omitempty"`
	GenericInteger *int         `json:"genericInteger,omitempty"`
	GenericString  *string      `json:"genericString,omitempty"`
}

func (o *GenericComponent) GetGenericBoolean() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return *o.GenericBoolean
}

func (o *GenericComponent) GetGenericBooleanOk() bool {
	if o == nil || o.GenericBoolean == nil {
		var ret bool
		return ret
	}
	return *o.GenericBoolean
}
func (o *GenericComponent) SetGenericBoolean(value bool) {
	o.GenericBoolean = &value
}

func (o *GenericComponent) GetGenericEnum() GenericEnum {
	if o == nil {
		var ret GenericEnum
		return ret
	}
	return *o.GenericEnum
}

func (o *GenericComponent) GetGenericEnumOk() GenericEnum {
	if o == nil || o.GenericEnum == nil {
		var ret GenericEnum
		return ret
	}
	return *o.GenericEnum
}
func (o *GenericComponent) SetGenericEnum(value GenericEnum) {
	o.GenericEnum = &value
}

func (o *GenericComponent) GetGenericInteger() int {
	if o == nil {
		var ret int
		return ret
	}
	return *o.GenericInteger
}

func (o *GenericComponent) GetGenericIntegerOk() int {
	if o == nil || o.GenericInteger == nil {
		var ret int
		return ret
	}
	return *o.GenericInteger
}
func (o *GenericComponent) SetGenericInteger(value int) {
	o.GenericInteger = &value
}

func (o *GenericComponent) GetGenericString() string {
	if o == nil {
		var ret string
		return ret
	}
	return *o.GenericString
}

func (o *GenericComponent) GetGenericStringOk() string {
	if o == nil || o.GenericString == nil {
		var ret string
		return ret
	}
	return *o.GenericString
}
func (o *GenericComponent) SetGenericString(value string) {
	o.GenericString = &value
}
