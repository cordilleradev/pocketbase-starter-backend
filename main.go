package main

import (
	"log"
	"os"
	"slices"

	"github.com/cordilleradev/stream2/common"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	isDev := slices.Contains(os.Args, "--dev")

	config, err := common.NewConfig(isDev)
	if err != nil {
		log.Fatal(err)
	}
	app := pocketbase.New()
	app.OnBootstrap().BindFunc(func(e *core.BootstrapEvent) error {
		if err := e.Next(); err != nil {
			return err
		}

		e.App.Settings().Meta.AppName = config.AppName
		e.App.Settings().Meta.AppURL = config.FrontendURL
		e.App.Settings().Meta.SenderName = config.SendingName
		e.App.Settings().Meta.SenderAddress = config.SendingAddress
		e.App.Settings().SMTP.Host = config.SMTPHost
		e.App.Settings().SMTP.Port = config.SMTPPort
		e.App.Settings().SMTP.Enabled = config.SMTPEnabled
		e.App.Settings().SMTP.Username = config.SMTPUsername
		e.App.Settings().SMTP.Password = config.SMTPPassword

		return e.App.Save(e.App.Settings())
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
