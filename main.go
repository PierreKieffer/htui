package main

import (
	// "context"
	"fmt"
	"github.com/PierreKieffer/htui/pkg/pkg/auth"
	"github.com/PierreKieffer/htui/pkg/pkg/ui"
	"os"
)

func main() {

	err := auth.Auth()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ui.App()

}
