package controllers

import (
	"fmt"
	"os"
	"os/exec"
)

func DbDump(dir string) {
	// Create new path inside backup directory
	err := os.Mkdir(dir+"/Banco de dados", 755)

	// Create .sql file inside our database directory
	sqlFile, err := os.Create(dir + "/Banco de dados/database.sql")
	if err != nil {
		panic(err.Error())
	}

	// Running MySQL dump
	dump := exec.Command("mysqldump", "-u", "root", "-p"+"", "-h", "", "")

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
		fmt.Printf("Error when writing to .sql file: %v\n", err)
		os.Exit(1)
	}

	// Closing .sql file in system
	defer sqlFile.Close()
}
