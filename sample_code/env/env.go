package env

import "os"

func ProcessEnvVars() Config {
	env, _ := os.LookupEnv("OUTPUT_FORMAT")
	return Config{
		OutputFormat: env,
	}
}

type Config struct {
	OutputFormat string
}
