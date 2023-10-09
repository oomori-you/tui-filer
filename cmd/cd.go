/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

// cdCmd represents the cd command
var cdCmd = &cobra.Command{
	Use:   "cd",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		app := tview.NewApplication()
		flex := tview.NewFlex().
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(tview.NewBox().SetBorder(true).SetTitle("Top - Fixed - 1"), 3, 0, false).
				AddItem(tview.NewBox().SetBorder(true).SetTitle("Middle Proportion"), 0, 1, false).
				AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom Fixed - 1"), 3, 0, false),
				0, 1, false)

		if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(cdCmd)
}
