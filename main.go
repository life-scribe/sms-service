package main

import (
	"os"
)

func main() {
	a := App{}
	a.Init(
		os.Getenv("LIFESCRIBE_DB_USERNAME"),
		os.Getenv("LIFESCRIBE_DB_PW"),
		os.Getenv("LIFESCRIBE_DB_NAME"))
}
