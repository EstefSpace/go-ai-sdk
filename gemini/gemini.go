package gemini

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Ask(prompt string, instructions string, apiKey string, model string) (string, error) {

	url := "https://generativelanguage.googleapis.com/v1beta/models/" + model + ":generateContent?key=" + apiKey

	payload := map[string]interface{}{
		"system_instruction": map[string]interface{}{
			"parts": []map[string]interface{}{
				{
					"text": instructions,
				},
			},
		},
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]interface{}{
					{
						"text": prompt,
					},
				},
			},
		},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("error al convertir el payload a JSON: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", fmt.Errorf("error al crear la petición: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error al realizar la petición: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error al leer la respuesta: %w", err)
	}

	// Decodificar la respuesta JSON y extraer el texto
	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("error al decodificar la respuesta JSON: %w", err)
	}

	candidates, ok := response["candidates"].([]interface{})
	if !ok || len(candidates) == 0 {
		return "", fmt.Errorf("respuesta de API inesperada: %s", string(body))
	}

	candidate, ok := candidates[0].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("respuesta de API inesperada: %s", string(body))
	}

	content, ok := candidate["content"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("respuesta de API inesperada: %s", string(body))
	}

	parts, ok := content["parts"].([]interface{})
	if !ok || len(parts) == 0 {
		return "", fmt.Errorf("respuesta de API inesperada: %s", string(body))
	}

	part, ok := parts[0].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("respuesta de API inesperada: %s", string(body))
	}

	text, ok := part["text"].(string)
	if !ok {
		return "", fmt.Errorf("respuesta de API inesperada: %s", string(body))
	}

	return text, nil
}
