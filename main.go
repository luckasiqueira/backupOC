package main

import "backupOC/controllers"

var dir = controllers.SetFolder()

func main() {
	controllers.GetData()
}
