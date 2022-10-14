package spec

type ReferencedOneOf struct {
	DiscriminatorTestProp *string `json:"DiscriminatorTestProp,omitempty"`
	ReferencedTestProp    *string `json:"ReferencedTestProp,omitempty"`
}

func (o *ReferencedOneOf) GetDiscriminatorTestProp() string {
	if o == nil {
		var ret string
		return ret
	}
	return *o.DiscriminatorTestProp
}

func (o *ReferencedOneOf) GetDiscriminatorTestPropOk() string {
	if o == nil || o.DiscriminatorTestProp == nil {
		var ret string
		return ret
	}
	return *o.DiscriminatorTestProp
}
func (o *ReferencedOneOf) SetDiscriminatorTestProp(value string) {
	o.DiscriminatorTestProp = &value
}

func (o *ReferencedOneOf) GetReferencedTestProp() string {
	if o == nil {
		var ret string
		return ret
	}
	return *o.ReferencedTestProp
}

func (o *ReferencedOneOf) GetReferencedTestPropOk() string {
	if o == nil || o.ReferencedTestProp == nil {
		var ret string
		return ret
	}
	return *o.ReferencedTestProp
}
func (o *ReferencedOneOf) SetReferencedTestProp(value string) {
	o.ReferencedTestProp = &value
}
