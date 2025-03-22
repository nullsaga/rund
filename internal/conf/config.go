package conf

type Webhook struct {
	Provider string `json:"provider"`
	Secret   string `json:"secret"`
}

type Deploy struct {
	Exec string `json:"exec"`
}

type Project struct {
	Secret  string   `json:"secret"`
	Webhook *Webhook `json:"webhook"`
	Deploy  *Deploy  `json:"deploy"`
}

type RootConf struct {
	Projects map[string]*Project `yaml:"projects"`
}
