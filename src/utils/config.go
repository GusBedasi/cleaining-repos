package configs

import "os"

func GetKey(key string) string {
	return os.Getenv(key)
}
