package main

import (
	"fmt"
	"os"

	"github.com/fatal-fruit/cosmapp/app"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/fatal-fruit/cosmapp/cmd/cosmappd/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
