package main

import (
	"fmt"
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

	templateManager, err := common.NewTemplateManager()
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

		users, err := e.App.FindCollectionByNameOrId("users")
		if err != nil {
			return err
		}
		users.OTP.Enabled = true
		users.OTP.Length = 6
		err = e.App.Save(users)
		if err != nil {
			return err
		}
		return e.App.Save(e.App.Settings())
	})

	app.OnMailerRecordAuthAlertSend().BindFunc(func(e *core.MailerRecordEvent) error {
		e.Message.HTML = templateManager.LoginAlertContent(config.AppName)
		return e.Next()
	})

	app.OnMailerRecordEmailChangeSend().BindFunc(func(e *core.MailerRecordEvent) error {
		e.Message.HTML = templateManager.ConfirmEmailChangeContent(
			e.Meta["token"].(string),
			config.FrontendURL,
			config.AppName,
		)
		return e.Next()
	})

	app.OnMailerRecordOTPSend().BindFunc(func(e *core.MailerRecordEvent) error {
		fmt.Println(e.Meta)
		e.Message.HTML = templateManager.OtpContent(
			e.Meta["password"].(string),
			config.AppName,
		)
		return e.Next()
	})

	app.OnMailerRecordPasswordResetSend().BindFunc(func(e *core.MailerRecordEvent) error {
		e.Message.HTML = templateManager.PasswordResetContent(
			e.Meta["token"].(string),
			config.FrontendURL,
			config.AppName,
		)
		return e.Next()
	})

	app.OnMailerRecordVerificationSend().BindFunc(func(e *core.MailerRecordEvent) error {
		e.Message.HTML = templateManager.VerifyEmailContent(
			e.Meta["token"].(string),
			config.FrontendURL,
			config.AppName,
		)
		return e.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
