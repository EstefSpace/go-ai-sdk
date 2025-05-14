package goaisdk

import (
	"fmt"

	"github.com/estefspace/go-ai-sdk/deepseek"
	"github.com/estefspace/go-ai-sdk/gemini"
)

var (
	Version   = "1.0"
	Developer = "EstefDev"
)

type ClientAI struct {
	apiKey string
	model  string
}

func NewClient(apiKey string, model string) *ClientAI {
	return &ClientAI{
		apiKey: apiKey,
		model:  model,
	}
}

func (c *ClientAI) GenerateContent(prompt string, instructions string) (string, error) {
	switch c.model {

	// MODELOS DE GEMINI
	case "gemini-2.0-flash", "gemini-2.5-pro-exp-03-25", "gemini-2.5-flash-preview-04-17":
		content, err := gemini.Ask(prompt, instructions, c.apiKey, c.model)

		if err != nil {
			return "", fmt.Errorf("revisa que hayas elegido un modelo valido de gemini, visita goaisdk.info/docs/models/gemini para más información: %w", err)
		}

		return content, nil

		// MODELOS DE DEEPSEEK
	case "deepseek-chat", "deepseek-reasoner":
		content, err := deepseek.Ask(prompt, instructions, c.apiKey, c.model)

		if err != nil {
			return "", fmt.Errorf("revisa que hayas elegido un modelo valido de gemini, visita goaisdk.info/docs/models/deepseek para más información: %w", err)
		}

		return content, nil
	default:
		return "", fmt.Errorf("revisa que hayas elegido un modelo valido, visita goaisdk.info/docs/models para más información")
	}

}
