package main

import (
	"mongo"
	"process"
)

func main() {
	mongo.DbClient = mongo.Connect()

	process.SaveCertificate()
}
