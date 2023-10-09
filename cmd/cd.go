/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/spf13/cobra"
)

// cdCmd represents the cd command
var cdCmd = &cobra.Command{
	Use:   "cd",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		app := tview.NewApplication()

		wdPath, _ := os.Getwd()
		textview := tview.NewTextView()
		textview.SetText(wdPath).SetTitle("Top").SetBorder(true)

		inputField := tview.NewInputField()
		inputField.SetDoneFunc(func(key tcell.Key) {
			app.Stop()
		})

		flex := tview.NewFlex().
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				//AddItem(textview.SetBorder(true).SetTitle("Top - Fixed - 1"), 30, 0, false).
				AddItem(textview, 3, 0, false).
				AddItem(tview.NewBox().SetBorder(true).SetTitle("Middle Proportion"), 0, 1, false).
				AddItem(inputField, 1, 0, true),
				0, 1, true)

		if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(cdCmd)
}
