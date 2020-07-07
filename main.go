package main

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

var (
	/*
		BuildVersion, BuildDate, BuildCommitSha are filled in by the build script
	*/
	BuildVersion = "<<< filled in by build >>>"
	BuildDate    = "<<< filled in by build >>>"
	BuildComment = "<<< filled in by build >>>"
)

func main() {
	fmt.Printf("BuildVersion: %s BuildDate %s ButeComment: %s", BuildVersion, BuildDate, BuildComment)
}
