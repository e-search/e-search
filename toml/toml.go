package toml

import (
	"github.com/BurntSushi/toml"
)

// Config コンフィグ
type Config struct {
	Postgres DBConfig
}

// DBConfig dbコンフィグ
type DBConfig struct {
	DBName  string
	User    string
	Host    string
	Sslmode string
}

// LoadDBToml tomlからDB情報を取得
func LoadDBToml() (string, error) {
	var config Config
	_, err := toml.DecodeFile("dbconf.toml", &config)
	if err != nil {
		return "", err
	}

	psql := config.Postgres
	str := "dbname=" + psql.DBName + " user=" + psql.User + " host=" + psql.Host + " sslmode=" + psql.Sslmode

	return str, nil
}
