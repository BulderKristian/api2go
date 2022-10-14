package spec

type MultiObject struct {
	WithDiscriminator    *OneOfWithDiscriminator    `json:"WithDiscriminator,omitempty"`
	WithoutDiscriminator *OneOfWithoutDiscriminator `json:"WithoutDiscriminator,omitempty"`
}

func (o *MultiObject) GetWithDiscriminator() OneOfWithDiscriminator {
	if o == nil {
		var ret OneOfWithDiscriminator
		return ret
	}
	return *o.WithDiscriminator
}

func (o *MultiObject) GetWithDiscriminatorOk() OneOfWithDiscriminator {
	if o == nil || o.WithDiscriminator == nil {
		var ret OneOfWithDiscriminator
		return ret
	}
	return *o.WithDiscriminator
}
func (o *MultiObject) SetWithDiscriminator(value OneOfWithDiscriminator) {
	o.WithDiscriminator = &value
}

func (o *MultiObject) GetWithoutDiscriminator() OneOfWithoutDiscriminator {
	if o == nil {
		var ret OneOfWithoutDiscriminator
		return ret
	}
	return *o.WithoutDiscriminator
}

func (o *MultiObject) GetWithoutDiscriminatorOk() OneOfWithoutDiscriminator {
	if o == nil || o.WithoutDiscriminator == nil {
		var ret OneOfWithoutDiscriminator
		return ret
	}
	return *o.WithoutDiscriminator
}
func (o *MultiObject) SetWithoutDiscriminator(value OneOfWithoutDiscriminator) {
	o.WithoutDiscriminator = &value
}
