package spec

type OneOfWithDiscriminator struct {
	InlineObject0   *InlineObject0   `json:"InlineObject0,omitempty"`
	ReferencedOneOf *ReferencedOneOf `json:"ReferencedOneOf,omitempty"`
}

func (o *OneOfWithDiscriminator) GetInlineObject0() InlineObject0 {
	if o == nil {
		var ret InlineObject0
		return ret
	}
	return *o.InlineObject0
}

func (o *OneOfWithDiscriminator) GetInlineObject0Ok() InlineObject0 {
	if o == nil || o.InlineObject0 == nil {
		var ret InlineObject0
		return ret
	}
	return *o.InlineObject0
}
func (o *OneOfWithDiscriminator) SetInlineObject0(value InlineObject0) {
	o.InlineObject0 = &value
}

func (o *OneOfWithDiscriminator) GetReferencedOneOf() ReferencedOneOf {
	if o == nil {
		var ret ReferencedOneOf
		return ret
	}
	return *o.ReferencedOneOf
}

func (o *OneOfWithDiscriminator) GetReferencedOneOfOk() ReferencedOneOf {
	if o == nil || o.ReferencedOneOf == nil {
		var ret ReferencedOneOf
		return ret
	}
	return *o.ReferencedOneOf
}
func (o *OneOfWithDiscriminator) SetReferencedOneOf(value ReferencedOneOf) {
	o.ReferencedOneOf = &value
}
