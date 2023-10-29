package er_api

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func BuildERApiManager(apiKey string) ERApiManager {
	return ERApiManager{apiKey: apiKey}
}

func (api ERApiManager) GetI10nFiles() {
	for _, language := range getLanguages() {
		req := fiber.Get(LANGUAGE_ROUTE + language)
		req.Debug()
		req.Set("accept", "application/json")
		req.Set("x-api-key", api.apiKey)

		_, data, err := req.Bytes()
		checkArr(err)

		var langResp LanguageResponse
		jsonErr := json.Unmarshal(data, &langResp)
		check(jsonErr)
		checkResp(langResp.Code, langResp.Message)

		req = fiber.Get(langResp.Data.Path)
		_, data, err = req.Bytes()
		checkArr(err)

		filename := fmt.Sprintf("%s.txt", language)
		filepath := filepath.Join("content", "l10n", filename)
		fileErr := os.WriteFile(filepath, data, 0644)
		check(fileErr)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkArr(e []error) {
	if e != nil {
		panic(e)
	}
}

func checkResp(c int, m string) {
	if c != 200 {
		panic(m)
	}
}
