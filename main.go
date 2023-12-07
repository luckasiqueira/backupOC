package main

import "backupOC/controllers"

var dir = controllers.SetFolder()
var data = controllers.GetData()

func main() {
	controllers.DbDump(dir, data)
}
