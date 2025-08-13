package config

import "os"

type Config struct {
	AppPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPass     string
	DBName     string
	JWTSecret  string
	AdminName  string
	AdminEmail string
	AdminPass  string
}

func Load() *Config {
	return &Config{
		AppPort:    getenv("APP_PORT", ":8080"),
		DBHost:     getenv("DB_HOST", "127.0.0.1"),
		DBPort:     getenv("DB_PORT", "3306"),
		DBUser:     getenv("DB_USER", "root"),
		DBPass:     getenv("DB_PASS", "example"),
		DBName:     getenv("DB_NAME", "portfolio"),
		JWTSecret:  getenv("JWT_SECRET", "change_me"),
		AdminName:  getenv("ADMIN_NAME", "Owner"),
		AdminEmail: getenv("ADMIN_EMAIL", "admin@example.com"),
		AdminPass:  getenv("ADMIN_PASSWORD", "changeme"),
	}
}

func getenv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
