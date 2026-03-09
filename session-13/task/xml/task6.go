package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Config struct {
	Username string  `xml:"username"`
	Password string  `xml:"password"`
	Database string  `xml:"database"`
	Options  Options `xml:"options"`
}

type Options struct {
	AutoBackup     bool `xml:"auto_backup"`
	MaxConnections int  `xml:"max_connections"`
}

func main() {
	dbc := Config{
		Username: "admin",
		Password: "123456",
		Database: "Postgres",
		Options: Options{
			AutoBackup:     true,
			MaxConnections: 10,
		},
	}
	data, err := xml.Marshal(&dbc)
	if err != nil {
		fmt.Println(err)
	}
	file, err := os.OpenFile("config.xml", os.O_CREATE|os.O_RDWR, 0660)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		fmt.Println(err)
	}
}
