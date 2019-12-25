package api

type Config struct {
	Email    string
	Password string
	APIKey   string
}

type Option func(*Config)

func EmailPasswordOption(email string, password string) Option {
	return Option(func(config *Config) {
		config.Email = email
		config.Password = password
	})
}

func APIKeyOption(apiKey string) Option {
	return Option(func(config *Config) {
		config.APIKey = apiKey
	})
}
