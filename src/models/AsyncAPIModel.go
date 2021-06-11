package models

type AsyncApi struct {
	Asyncapi   string     `yaml:"asyncapi"`
	Components *Component `yaml:"components"`
}

type Component struct {
	Messages *map[string]Message `yaml:"messages,omitempty"`
	Schemas  *map[string]Schema  `json:"schemas,omitempty"`
}

type Message struct {
	Payload Payload `yaml:"payload"`
}

type Payload struct {
	Type string `yaml:"type"`
	//Properties map[string]interface{} `json:"properties"`
	Properties map[string]Property `yaml:"properties"`
}

type Property struct {
	Type        string            `yaml:"type,omitempty"`
	Format      string            `yaml:"format,omitempty"`
	Example     interface{}       `yaml:"example,omitempty"`
	Pattern     string            `yaml:"pattern,omitempty"`
	Description string            `yaml:"description,omitempty"`
	Nullable    bool              `yaml:"nullable,omitempty"`
	Enum        []string          `yaml:"enum,omitempty"`
	Items       map[string]string `yaml:"items,omitempty"`
	Ref         string            `yaml:"$ref,omitempty"`
}

type Schema struct {
	Type       string              `yaml:"type"`
	Enums      []string            `yaml:"enum,omitempty"`
	Example    string              `yaml:"example,omitempty"`
	Pattern    string              `yaml:"pattern,omitempty"`
	Properties map[string]Property `yaml:"properties"`
}
