package config 

import (
	"os"
	"strconv"
	// "fmt"
)

type Config struct {
	InCluster bool 
	APIClientSecret string
	AuthUrl string 
}


func New () Config {
	return Config {
		InCluster: getEnvAsBool("IN_CLUSTER", true),
	}
}

func getEnv ( key string, defaultValue string ) string {

	if env, exists := os.LookupEnv(key); exists {
		return env
	}
	
	return defaultValue 
	
} 

func getEnvAsBool ( key string, defaultValue bool) bool {
	envStr := getEnv(key, "")
	if boolVal, err := strconv.ParseBool(envStr); err == nil {
		return boolVal
	} 

	return defaultValue

}