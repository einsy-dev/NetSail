package services

import (
	"github.com/einsy-dev/NetSail/internal/services/clipboard"
	"github.com/einsy-dev/NetSail/internal/services/extension"
	parseCsv "github.com/einsy-dev/NetSail/internal/services/parse-csv"
	"github.com/einsy-dev/NetSail/internal/services/server"
)

type Clip = clipboard.Clipboard
type Knot = extension.Extension
type Player = extension.Player
type Api = server.Api

var ParseCSV = parseCsv.ParseCSV
