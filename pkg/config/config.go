package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	DBType      string `json:"db_type"`
	PostgresDSN string `json:"postgres_dsn"`
	SQLitePath  string `json:"sqlite_path"`
	MongoURI    string `json:"mongodb_uri"`
	MongoDBName string `json:"mongodb_db"`
	UseMemcache bool   `json:"use_memcache"`
}

var AppConfig Config

func LoadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatal("Failed to open config:", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&AppConfig); err != nil {
		log.Fatal("Failed to parse config:", err)
	}
}
