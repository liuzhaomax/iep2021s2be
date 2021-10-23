/****************************************************************************
 * @copyright   LIU Zhao
 * @authors     LIU Zhao (liuzhaomax@163.com)
 * @date        2021/8/1 18:19
 * @version     v1.0
 * @filename    config.go
 * @description
 ***************************************************************************/
package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

var cfg *Config
var once sync.Once

func init() {
	once.Do(func() {
		cfg = &Config{}
	})
}

func GetInstanceOfConfig() *Config {
	return cfg
}

type Config struct {
	mux     sync.Mutex
	RunMode string  `json:"run_mode"`
	App     App     `json:"app"`
	Http    Http    `json:"http"`
	DB      DB      `json:"db"`
	Mysql   Mysql   `json:"mysql"`
	Admin   Admin   `json:"admin"`
	FEAdmin FEAdmin `json:"fe_admin"`
	FEMain  FEMain  `json:"fe_main"`
}

type App struct {
	AppName string `json:"app_name"`
	Version string `json:"version"`
}

type Http struct {
	Protocol         string `json:"protocol"`
	ProtocolVersion  string `json:"protocol_version"`
	Domain           string `json:"domain"`
	Host             string `json:"host"`
	Port             string `json:"port"`
	CertDir          string `json:"certi_dir"`
	KeyDir           string `json:"key_dir"`
	ShutdownTimeout  int    `json:"shutdown_timeout"`
	MaxContentLength int64  `json:"max_content_length"`
	MaxLoggerLength  int    `json:"max_logger_length"`
}

type DB struct {
	Type         string `json:"type"`
	Debug        bool   `json:"debug"`
	DSN          string `json:"dsn"`
	MaxLifetime  int    `json:"max_life_time"`
	MaxOpenConns int    `json:"max_open_conns"`
	MaxIdleConns int    `json:"max_idle_conns"`
}

type Mysql struct {
	Host      string `json:"host"`
	LocalHost string `json:"local_host"`
	Port      string `json:"port"`
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
	DBName    string `json:"db_name"`
	Params    string `json:"params"`
}

type Admin struct {
	UserEmail    string `json:"user_email"`
	UserPassword string `json:"user_password"`
	UserNickName string `json:"user_nick_name"`
}

type FEAdmin struct {
	FEAdminHttp FEAdminHttp `json:"http"`
}

type FEAdminHttp struct {
	LocalProtocol   string `json:"local_protocol"`
	RemoteProtocol  string `json:"remote_protocol"`
	ProtocolVersion string `json:"protocol_version"`
	Domain          string `json:"domain"`
	Host            string `json:"host"`
	Port            string `json:"port"`
}

type FEMain struct {
	FEMainHttp FEMainHttp `json:"http"`
}

type FEMainHttp struct {
	LocalProtocol   string `json:"local_protocol"`
	RemoteProtocol  string `json:"remote_protocol"`
	ProtocolVersion string `json:"protocol_version"`
	Domain          string `json:"domain"`
	Host            string `json:"host"`
	Port            string `json:"port"`
}

func (mysql *Mysql) DSN() string {
	if mysql.Password == "" {
		mysql.Password = "123456"
	}
	if cfg.RunMode == "release" {
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
			mysql.UserName, mysql.Password, mysql.Host, mysql.Port, mysql.DBName, mysql.Params)
	} else {
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
			mysql.UserName, mysql.Password, mysql.LocalHost, mysql.Port, mysql.DBName, mysql.Params)
	}
}

func (cfg *Config) Load(configDir string, configName string) error {
	dataJson, err := ioutil.ReadFile(configDir + "/" + configName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(dataJson, cfg)
	if err != nil {
		return err
	}
	return nil
}
