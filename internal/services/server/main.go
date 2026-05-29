package server

import (
	"context"

	"github.com/einsy-dev/NetSail/internal/utils"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type Api struct{}

func (s *Api) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {

	return nil
}

func (s *Api) ParseCsv() string {
	utils.ReadFiles()

	// res, err := fetch.Post("http://localhost:3000/api/utils/csv/join", &fetch.Config{Body: nBody, Headers: fetch.Headers{
	// 	"Authorization": "Bearer 147852",
	// }})

	// if err != nil {
	// 	fmt.Println("err", err.Error())
	// }
	// fmt.Println(res)
	// result, err := res.JSON()

	// return result
	// nbody file content in string format

	return ""
}
