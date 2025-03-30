package translate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	applicationError "github.com/changchanghwang/wdwb_back/pkg/application-error"
)

type Translator struct {
	translations map[string]map[string]map[string]string
}

func New() *Translator {
	translator := &Translator{
		translations: make(map[string]map[string]map[string]string),
	}
	translator.LoadTranslations()
	return translator
}

func (t *Translator) LoadTranslations() error {
	baseDir := filepath.Join("internal", "libs", "translate", "languages")

	entries, err := os.ReadDir(baseDir)
	if err != nil {
		return applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to read translation directory. %s", err.Error()), "AAA000")
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		category := entry.Name()
		langDir := filepath.Join(baseDir, category)
		langFiles, err := os.ReadDir(langDir)
		if err != nil {
			return fmt.Errorf("언어 디렉토리 읽기 실패: %w", err)
		}

		for _, langFile := range langFiles {
			if !langFile.IsDir() && strings.HasSuffix(langFile.Name(), ".json") {
				lang := strings.Split(langFile.Name(), ".")[0]
				filePath := filepath.Join(langDir, langFile.Name())

				data, err := os.ReadFile(filePath)
				if err != nil {
					return applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to read translation file. %s", err.Error()), "AAA000")
				}

				var trans map[string]string
				if err := json.Unmarshal(data, &trans); err != nil {
					return applicationError.New(http.StatusInternalServerError, fmt.Sprintf("Failed to parse translation file. %s", err.Error()), "AAA000")
				}

				if t.translations[category] == nil {
					t.translations[category] = make(map[string]map[string]string)
				}

				t.translations[category][lang] = trans
			}
		}
	}

	return nil
}

func (t *Translator) Translate(category, lang, key string) string {
	if trans, ok := t.translations[category]; ok {
		if langTrans, ok := trans[lang]; ok {
			if value, ok := langTrans[key]; ok {
				return value
			}
		}
	}

	return key
}
