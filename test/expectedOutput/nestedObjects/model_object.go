package spec

type Object struct {
	NestedObject *NestedObject `json:"NestedObject,omitempty"`
}

func (o *Object) GetNestedObject() NestedObject {
	if o == nil {
		var ret NestedObject
		return ret
	}
	return *o.NestedObject
}

func (o *Object) GetNestedObjectOk() NestedObject {
	if o == nil || o.NestedObject == nil {
		var ret NestedObject
		return ret
	}
	return *o.NestedObject
}
func (o *Object) SetNestedObject(value NestedObject) {
	o.NestedObject = &value
}
