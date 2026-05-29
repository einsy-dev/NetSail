package main

import (
	"embed"
	"fmt"

	a "github.com/einsy-dev/NetSail/internal/app"
	"github.com/einsy-dev/NetSail/internal/services"
	parseCsv "github.com/einsy-dev/NetSail/internal/services/parse-csv"
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	app := application.New(application.Options{
		Name: "NetSail",
		Services: []application.Service{
			application.NewService(&services.Clip{}),
			application.NewService(&services.Knot{}),
			application.NewService(&services.Player{}),
			application.NewService(&services.Api{}),
			application.NewService(services.ParseCSV),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		FileAssociations: []string{".csv"},
	})
	a.App = app // global app pointer

	app.Event.OnApplicationEvent(events.Common.ApplicationOpenedWithFile, func(event *application.ApplicationEvent) {
		associatedFile := event.Context().Filename()
		app.Dialog.Info().SetMessage("Application opened with file: " + associatedFile).Show()
	})

	// gh, _ := github.New(github.Config{Repository: "github.com/einsy-dev/NetSail"})
	// if err := app.Updater.Init(updater.Config{
	// 	CurrentVersion: "1.0.0",
	// 	Providers:      []updater.Provider{gh},
	// }); err != nil {
	// 	log.Fatal(err)
	// }

	// if err := app.Updater.CheckAndInstall(context.Background()); err != nil {
	// 	log.Printf("update: %v", err)
	// }

	win := app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:          "NetSail",
		X:              1000,
		Y:              500,
		Width:          350,
		MinWidth:       350,
		MaxWidth:       700,
		Height:         500,
		MinHeight:      500,
		MaxHeight:      800,
		AlwaysOnTop:    true,
		EnableFileDrop: true,
	})

	parseCsv.ListenDrop(win)

	err := app.Run()
	if err != nil {
		panic(err)
	}
}
