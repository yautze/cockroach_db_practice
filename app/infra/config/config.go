package config

// C -
var C = Config{
	Cockroach: Cockroach{
		Host:     "localhost:26257",
		User:     "root",
		Password: "",
		Database: "test",
	},
}

// Config -
type Config struct {
	// Cockroach -
	Cockroach Cockroach
}

// Cockroach -
type Cockroach struct {
	// Host -
	Host string

	// User -
	User string

	// Password -
	Password string

	// Database -
	Database string
}
