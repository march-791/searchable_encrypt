package db

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gotest/tool"
	"io/ioutil"
	"log"
)

var Database *gorm.DB

type Data struct {
	Mysql struct {
		Ip       string `json:"ip"`
		Port     string `json:"port"`
		Pass     string `json:"pass"`
		Database string `json:"database"`
	} `json:"mysql"`
}

func Config_get() Data {
	//first read the `config.json` file
	abspath := tool.GetCurrentAbPath()
	content, err := ioutil.ReadFile(abspath + "/config/config.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	//unmarshall the data into `payload`
	var payload Data
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// print the unmarshalled data
	fmt.Printf("mysql config:")
	fmt.Println(payload)
	return payload
}

func DB_init() {
	config_data := Config_get()
	var err error
	dsn := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config_data.Mysql.Pass, config_data.Mysql.Ip, config_data.Mysql.Port, config_data.Mysql.Database)
	Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, error=" + err.Error())
	}
	fmt.Println(Database)
}
