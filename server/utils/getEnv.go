package utils

import (
    "log"
    "github.com/spf13/viper"
)

func ViperEnvVariable(key string) string {

  viper.SetConfigFile(".env")
  err := viper.ReadInConfig()

  if err != nil {
    log.Fatal("Error while reading config file %s", err)
  }
  value, ok := viper.Get(key).(string)

  if !ok {
    log.Fatal("Invalid type assertion")
  }

  return value
}