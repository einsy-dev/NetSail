package utils

import (
	"github.com/einsy-dev/NetSail/internal/app"
)

func ReadFiles() []string {
	if app.App == nil {
		return nil
	}

	filePath, err := app.App.Dialog.
		OpenFile().
		SetTitle("Select a Text File").
		AddFilter("Text Files", "*.csv;"). // Limit file extensions
		PromptForMultipleSelection()

	if err != nil {
		return nil
	}

	return filePath
}
