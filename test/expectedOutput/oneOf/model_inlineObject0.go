package spec

type InlineObject0 struct {
	DiscriminatorTestProp *string `json:"DiscriminatorTestProp,omitempty"`
	InlineTestProp        *string `json:"InlineTestProp,omitempty"`
}

func (o *InlineObject0) GetDiscriminatorTestProp() string {
	if o == nil {
		var ret string
		return ret
	}
	return *o.DiscriminatorTestProp
}

func (o *InlineObject0) GetDiscriminatorTestPropOk() string {
	if o == nil || o.DiscriminatorTestProp == nil {
		var ret string
		return ret
	}
	return *o.DiscriminatorTestProp
}
func (o *InlineObject0) SetDiscriminatorTestProp(value string) {
	o.DiscriminatorTestProp = &value
}

func (o *InlineObject0) GetInlineTestProp() string {
	if o == nil {
		var ret string
		return ret
	}
	return *o.InlineTestProp
}

func (o *InlineObject0) GetInlineTestPropOk() string {
	if o == nil || o.InlineTestProp == nil {
		var ret string
		return ret
	}
	return *o.InlineTestProp
}
func (o *InlineObject0) SetInlineTestProp(value string) {
	o.InlineTestProp = &value
}
