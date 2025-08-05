package conf

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
	"github.com/kr/pretty"
	"github.com/spf13/viper"
)

var (
	conf *Config
	once sync.Once
)

type Config struct {
	Server Server `yaml:"server"`
	MySQL  MySQL  `yaml:"mysql"`
	Redis  Redis  `yaml:"redis"`
	JWT    JWT    `yaml:"jwt"`
}

type Server struct {
	Addr string `yaml:"addr"`
}

type MySQL struct {
	DSN string `yaml:"dsn"`
}

type Redis struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
}

type JWT struct {
	SignAlgo  string `yaml:"signAlgo"`
	SecretKey string `yaml:"secretKey"`
}

func GetConf() *Config {
	once.Do(initConf)
	return conf
}

func initConf() {
	env := getEnv()
	prefix := "conf"

	envFile := filepath.Join(prefix, filepath.Join(env, ".env"))
	err := godotenv.Load(envFile)
	if err != nil {
		panic(err)
	}

	confPath := filepath.Join(prefix, filepath.Join(env, "conf.yml"))
	viper.SetConfigFile(confPath)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	conf = new(Config)
	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}

	pretty.Printf("%#v\n", conf)
}

func getEnv() string {
	env := os.Getenv("GoEnv")
	if env == "" {
		return "test"
	}

	return env
}
