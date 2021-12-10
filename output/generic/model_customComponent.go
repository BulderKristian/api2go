package spec

type CustomComponent struct {
	CustomBoolean *bool       `json:"customBoolean,omitempty"`
	CustomEnum    *CustomEnum `json:"customEnum,omitempty"`
	CustomInteger *int        `json:"customInteger,omitempty"`
	CustomString  *string     `json:"customString,omitempty"`
}

func (o *CustomComponent) GetCustomBoolean() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return *o.CustomBoolean
}

func (o *CustomComponent) GetCustomBooleanOk() bool {
	if o == nil || o.CustomBoolean == nil {
		var ret bool
		return ret
	}
	return *o.CustomBoolean
}
func (o *CustomComponent) SetCustomBoolean(value bool) {
	o.CustomBoolean = &value
}

func (o *CustomComponent) GetCustomEnum() CustomEnum {
	if o == nil {
		var ret CustomEnum
		return ret
	}
	return *o.CustomEnum
}

func (o *CustomComponent) GetCustomEnumOk() CustomEnum {
	if o == nil || o.CustomEnum == nil {
		var ret CustomEnum
		return ret
	}
	return *o.CustomEnum
}
func (o *CustomComponent) SetCustomEnum(value CustomEnum) {
	o.CustomEnum = &value
}

func (o *CustomComponent) GetCustomInteger() int {
	if o == nil {
		var ret int
		return ret
	}
	return *o.CustomInteger
}

func (o *CustomComponent) GetCustomIntegerOk() int {
	if o == nil || o.CustomInteger == nil {
		var ret int
		return ret
	}
	return *o.CustomInteger
}
func (o *CustomComponent) SetCustomInteger(value int) {
	o.CustomInteger = &value
}

func (o *CustomComponent) GetCustomString() string {
	if o == nil {
		var ret string
		return ret
	}
	return *o.CustomString
}

func (o *CustomComponent) GetCustomStringOk() string {
	if o == nil || o.CustomString == nil {
		var ret string
		return ret
	}
	return *o.CustomString
}
func (o *CustomComponent) SetCustomString(value string) {
	o.CustomString = &value
}
