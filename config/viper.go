package config

import (
	"errors"

	"github.com/spf13/viper"
)

type EnvVars struct {
  Port string `mapstructure:"PORT"`
}

func LoadEnv() (EnvVars, error){
  viper.SetConfigType("env")
  viper.AddConfigPath(".")
  viper.SetConfigFile(".env")
  err := viper.ReadInConfig()

  config := EnvVars {}
  if err != nil {
    return config, err
  } 

  err = viper.Unmarshal(&config)
  if err != nil {
    return config, err
  }
  
  if config.Port == "" {
    return config, errors.New("PORT not defined.")
  }

  return config, nil
}
