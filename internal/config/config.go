package config

const configFileName = ".gatorconfig.json"

type Config struct {
    DBUrl           string `json:"db_url"`
    CurrentUserName string
}
