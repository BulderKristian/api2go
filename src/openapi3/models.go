package openapi3

type Openapi3 struct {
	Openapi    string     `json:"openapi" yaml:"openapi"`
	Components Components `json:"components" yaml:"components"`
}

type Components struct {
	Schemas   map[string]Schema    `json:"schemas,omitempty" yaml:"schemas,omitempty"`
	Responses map[string]Responses `json:"responses,omitempty" yaml:"Responses,omitempty"`
}

type Schema struct {
	Ref          string            `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Title        string            `json:"title,omitempty" yaml:"title,omitempty"`
	Required     []string          `json:"required,omitempty" yaml:"required,omitempty"`
	Type         string            `json:"type,omitempty" yaml:"type,omitempty"`
	Format       string            `json:"format,omitempty" yaml:"format,omitempty"`
	Items        map[string]string `json:"items,omitempty" yaml:"items,omitempty"`
	Enums        []string          `json:"enum,omitempty" yaml:"enum,omitempty"`
	Properties   map[string]Schema `json:"properties,omitempty" yaml:"properties,omitempty"` //RawProperties to get x-extensions?
	AllOf        []Schema          `json:"allOf,omitempty" yaml:"allOf,omitempty"`
	OneOf        []Schema          `json:"oneOf,omitempty" yaml:"oneOf,omitempty"`
	AnyOf        []Schema          `json:"anyOf,omitempty" yaml:"anyOf,omitempty"`
	XGoCustomTag string            `json:"x-go-custom-tag,omitempty" yaml:"x-go-custom-tag,omitempty"`
}

type Responses struct {
	Ref string `json:"$ref,omitempty" yaml:"$ref,omitempty"`
}
