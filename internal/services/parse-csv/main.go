package parseCsv

import (
	"fmt"

	"github.com/go-zoox/fetch"
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

type csv struct {
	files []string
}

var ParseCSV = &csv{}

func ListenDrop(win *application.WebviewWindow) {
	win.OnWindowEvent(events.Common.WindowFilesDropped, func(event *application.WindowEvent) {
		files := event.Context().DroppedFiles()
		ParseCSV.files = append(ParseCSV.files, files...)
	})
}

func parse() {
	res, err := fetch.Post("http://localhost:3000/api/utils/csv/join")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
