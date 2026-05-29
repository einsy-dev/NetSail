package parseCsv

import (
	"context"

	"github.com/einsy-dev/NetSail/internal/utils"
	"github.com/wailsapp/wails/v3/pkg/application"
)

func (c *csv) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	return nil
}

func (c *csv) Clear() error {
	c.files = []string{}
	return nil
}

func (c *csv) ParseFiles() error {
	var res = utils.ReadFiles()
	c.files = append(c.files, res...)
	return nil
}
