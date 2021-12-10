package spec

type AllOfCustomEnum struct {
	CustomEnum CustomEnum `json:"CustomEnum"`
}

func (o *AllOfCustomEnum) GetCustomEnum() CustomEnum {
	if o == nil {
		var ret CustomEnum
		return ret
	}
	return o.CustomEnum
}

func (o *AllOfCustomEnum) SetCustomEnum(value CustomEnum) {
	o.CustomEnum = value
}
