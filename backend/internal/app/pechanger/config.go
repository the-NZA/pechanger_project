package pechanger

// Author: Roman Kozlov
// Config is a base struct for store settings
type Config struct {
	BindAddr    string `json:"bind_addr"`
	LogDebug    bool   `json:"log_debug"`
	DatabaseURL string `json:"db_url"`
	EmailFrom   string `json:"email_from"`
	EmailSMTP   string `json:"email_smtp"`
	EmailPort   string `json:"email_port"`
	EmailPass   string `json:"email_pass"`
}

// Author: Roman Kozlov
// NewConfig creates basic config
func NewConfig() *Config {
	return &Config{
		BindAddr:    ":9999",
		DatabaseURL: "NOT IMPLEMENTED YET",
		LogDebug:    false,
		EmailFrom:   "",
		EmailSMTP:   "",
		EmailPort:   "",
		EmailPass:   "",
	}
}

// Author: Roman Kozlov
func (c Config) GetAddress() string {
	return c.EmailSMTP + ":" + c.EmailPort
}
