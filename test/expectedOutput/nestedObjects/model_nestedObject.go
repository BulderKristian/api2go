package spec

type NestedObject struct {
	Test *string `json:"Test,omitempty"`
}

func (o *NestedObject) GetTest() string {
	if o == nil {
		var ret string
		return ret
	}
	return *o.Test
}

func (o *NestedObject) GetTestOk() string {
	if o == nil || o.Test == nil {
		var ret string
		return ret
	}
	return *o.Test
}
func (o *NestedObject) SetTest(value string) {
	o.Test = &value
}
