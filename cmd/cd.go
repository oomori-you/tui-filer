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

		list := tview.NewList()
		list.ShowSecondaryText(false).SetBorder(true).SetTitle("Middle Proportion")
		files, _ := os.ReadDir(wdPath)
		for _, v := range files {
			list.AddItem(v.Name(), "", 0, nil)
		}

		flex := tview.NewFlex().
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(textview, 3, 0, false).
				AddItem(list, 0, 1, false).
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
