package tkn2

type Secret struct {
	APIVersion string      `yaml:"apiVersion,omitempty"`
	Kind       string      `yaml:"kind,omitempty"`
	Metadata   *Metadata   `yaml:"metadata,omitempty"`
	Type       string      `yaml:"type,omitempty"`
	StringData *StringData `yaml:"stringData,omitempty"`
	SecretName string      `yaml:"secretName,omitempty"`
}

type StringData struct {
	Pwd   string `yaml:"pwd,omitempty"`
	User  string `yaml:"user,omitempty"`
	Token string `yaml:"token,omitempty"`
}