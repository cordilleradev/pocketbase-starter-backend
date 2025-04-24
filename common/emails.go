package common

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
)

// TemplateManager manages email templates
type TemplateManager struct {
	templates map[string]*template.Template
	cssStyle  string
}

// NewTemplateManager creates and initializes a new TemplateManager
func NewTemplateManager() (*TemplateManager, error) {
	tm := &TemplateManager{
		templates: make(map[string]*template.Template),
	}

	// Load global CSS
	cssPath := filepath.Join("templates", "global.css")
	cssContent, err := os.ReadFile(cssPath)
	if err != nil {
		return nil, err
	}
	tm.cssStyle = string(cssContent)

	// Parse all templates from the templates directory
	templateFiles := []string{
		"confirm_email_change.html",
		"password_reset.html",
		"verify_email.html",
		"login_alert.html",
		"otp.html",
	}

	for _, file := range templateFiles {
		templatePath := filepath.Join("templates", file)
		tmplContent, err := os.ReadFile(templatePath)
		if err != nil {
			return nil, err
		}

		tmpl, err := template.New(file).Parse(string(tmplContent))
		if err != nil {
			return nil, err
		}
		tm.templates[file] = tmpl
	}

	return tm, nil
}

// ConfirmEmailChangeContent generates the email content for confirming email change
func (tm *TemplateManager) ConfirmEmailChangeContent(token string, appURL string, appName string) string {
	data := struct {
		Token   string
		AppURL  string
		AppName string
		CSS     string
	}{
		Token:   token,
		AppURL:  appURL,
		AppName: appName,
		CSS:     tm.cssStyle,
	}

	var buf bytes.Buffer
	if err := tm.templates["confirm_email_change.html"].Execute(&buf, data); err != nil {
		return "Error executing email template"
	}

	return buf.String()
}

// OtpContent generates the email content for OTP
func (tm *TemplateManager) OtpContent(otp string, appName string) string {
	data := struct {
		OTP     string
		AppName string
		CSS     string
	}{
		OTP:     otp,
		AppName: appName,
		CSS:     tm.cssStyle,
	}

	var buf bytes.Buffer
	if err := tm.templates["otp.html"].Execute(&buf, data); err != nil {
		return "Error executing OTP template"
	}

	return buf.String()
}

// PasswordResetContent generates the email content for password reset
func (tm *TemplateManager) PasswordResetContent(token string, appURL string, appName string) string {
	data := struct {
		Token   string
		AppURL  string
		AppName string
		CSS     string
	}{
		Token:   token,
		AppURL:  appURL,
		AppName: appName,
		CSS:     tm.cssStyle,
	}

	var buf bytes.Buffer
	if err := tm.templates["password_reset.html"].Execute(&buf, data); err != nil {
		return "Error executing password reset template"
	}

	return buf.String()
}

// VerifyEmailContent generates the email content for email verification
func (tm *TemplateManager) VerifyEmailContent(token string, appURL string, appName string) string {
	data := struct {
		Token   string
		AppURL  string
		AppName string
		CSS     string
	}{
		Token:   token,
		AppURL:  appURL,
		AppName: appName,
		CSS:     tm.cssStyle,
	}

	var buf bytes.Buffer
	if err := tm.templates["verify_email.html"].Execute(&buf, data); err != nil {
		return "Error executing email verification template"
	}

	return buf.String()
}

// LoginAlertContent generates the email content for login alerts
func (tm *TemplateManager) LoginAlertContent(appName string) string {
	data := struct {
		AppName string
		CSS     string
	}{
		AppName: appName,
		CSS:     tm.cssStyle,
	}

	var buf bytes.Buffer
	if err := tm.templates["login_alert.html"].Execute(&buf, data); err != nil {
		return "Error executing login alert template"
	}

	return buf.String()
}
