package spec

type GenericComponent struct {
	Enum *Enum `json:"Enum,omitempty"`
}

func (o *GenericComponent) GetEnum() Enum {
	if o == nil {
		var ret Enum
		return ret
	}
	return *o.Enum
}

func (o *GenericComponent) GetEnumOk() Enum {
	if o == nil || o.Enum == nil {
		var ret Enum
		return ret
	}
	return *o.Enum
}
func (o *GenericComponent) SetEnum(value Enum) {
	o.Enum = &value
}
