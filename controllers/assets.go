package controllers

import (
	"os"
	"time"
)

type Database struct {
	Name    string
	Pass    string
	User    string
	Host    string
	Port    string
	Storage string
}

func SetFolder() string {
	// Just gets real time
	t := time.Now()

	// Set and create our backup directory
	bkpPath := "backup-" + t.Format("2006.01.02-15-04")

	err := os.Mkdir(bkpPath, 755)
	if err != nil {
		panic(err.Error())
	}

	return bkpPath
}
