package controllers

import "os"

func DbDump(dir string) {
	// Create new path inside backup directory
	err := os.Mkdir(dir+"/Banco de dados", 755)

	// Create .sql file inside our database directory
	sqlFile, err := os.Create(dir + "/Banco de dados/database.sql")
	if err != nil {
		panic(err.Error())
	}

	defer sqlFile.Close()
}
