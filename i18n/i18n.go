package i18n

import (
	"encoding/json"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var (
	//Bundle holds all translations
	Bundle *i18n.Bundle
	//Localizer is used to translate messages to the desired language
	Localizer *i18n.Localizer
)

func init() {
	//Initialize the bundle with default language
	Bundle = i18n.NewBundle(language.English)

	//Register JSON as the format for translation files
	Bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	//Load the translation files
	Bundle.MustLoadMessageFile("i18n/en.json")

	//Create a localizer for the default language
	Localizer = i18n.NewLocalizer(Bundle, "en")
}

func T(messageID string, templateData map[string]interface{}) string {
	message, err := Localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: templateData})
	if err != nil {
		return messageID
	}
	return message
}

// SetLanguage sets the language for the localizer
func SetLanguage(lang string) {
	Localizer = i18n.NewLocalizer(Bundle, lang)
}
