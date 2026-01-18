package models

type ConfigFile struct {
	AutoStart map[string]ServiceConfig `json:"autostart"`
}

type ServiceConfig struct {
	Path    string `json:"path"`
	Tracked bool   `json:"tracked"`
}
