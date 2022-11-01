package template

import (
	"html/template"
	"strings"

	"github.com/matsuev/klsh-email-sender/internal/config"
)

// EmailTemplate struct
type EmailTemplate struct {
	tpl  *template.Template
	keys map[string]int
}

// Create function
func Create(cfg *config.AppConfig) (*EmailTemplate, error) {
	t, err := template.ParseFiles(cfg.TemplatePath)
	if err != nil {
		return nil, err
	}

	keys := make(map[string]int)
	for _, v := range t.Tree.Root.Nodes {
		if v.Type() == 1 {
			key := strings.TrimSuffix(strings.TrimPrefix(v.String(), "{{."), "}}")
			keys[key]++
		}
	}

	return &EmailTemplate{
		tpl:  t,
		keys: keys,
	}, nil
}

// GetKeys function
func (t *EmailTemplate) GetKeys() map[string]int {
	return t.keys
}

// CheckKeys function
func (t *EmailTemplate) CheckKeys(keys []string) bool {
	for _, key := range keys {
		if _, ok := t.keys[key]; !ok {
			return false
		}
	}

	return true
}
