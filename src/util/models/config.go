package models

type Config struct {
	Service struct {
		Path   string `json:"path"`
		Host   string `json:"host"`
		Port   string `json:"port"`
		Secret string `json:"secret"`
	} `json:"service"`
	Vault struct {
		VaultAddr string `json:"addr"`
		Token     string `json:"token"`
	} `json:"vault"`
}
