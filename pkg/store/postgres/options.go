package postgres

type Option func(*Settings)

func WithHost(host string) Option {
	return func(settings *Settings) {
		settings.host = host
	}
}

func WithPort(port string) Option {
	return func(settings *Settings) {
		settings.port = port
	}
}

func WithUsername(username string) Option {
	return func(settings *Settings) {
		settings.username = username
	}
}

func WithPassword(password string) Option {
	return func(settings *Settings) {
		settings.password = password
	}
}

func WithDBName(db string) Option {
	return func(settings *Settings) {
		settings.db = db
	}
}
