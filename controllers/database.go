package controllers

import (
	"os"
	"os/exec"
)

func DbDump(dir string, data *Database) {
	// Create new path inside backup directory
	err := os.Mkdir(dir+"/Banco de dados", 755)

	// Create .sql file inside our database directory
	sqlFile, err := os.Create(dir + "/Banco de dados/database.sql")
	if err != nil {
		panic(err.Error())
	}

	// Running MySQL dump
	dump := exec.Command("mysqldump", "-u", data.User, "-p"+data.Pass, "-h", data.Host, "-P", data.Port, data.Name)

	output, err := dump.Output()
	if err != nil {
		panic(err.Error())
		os.Exit(1)
	}

	// Set our .sql file as output for mysql dump command
	dump.Stdout = sqlFile

	// Write data onto .sql file
	_, err = sqlFile.Write(output)
	if err != nil {
		panic(err.Error())
		os.Exit(1)
	}

	// Closing .sql file in system
	defer sqlFile.Close()
}
