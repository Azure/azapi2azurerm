package coverage

import (
	_ "embed"
	"encoding/json"
	"strings"
)

//go:embed tf.json
var coverageJson string

var cov []Coverage

func init() {
	_ = json.Unmarshal([]byte(coverageJson), &cov)
	for i := range cov {
		cov[i].IdPattern = strings.ReplaceAll(cov[i].IdPattern, "/{}", "")
	}
}

func GetPutCoverage(props []string, idPattern string) ([]string, []string) {
	return getCoverage(props, "PUT", idPattern)
}

func GetGetCoverage(props []string, idPattern string) ([]string, []string) {
	return getCoverage(props, "GET", idPattern)
}

func getCoverage(props []string, operation, idPattern string) ([]string, []string) {
	for _, r := range cov {
		if r.Operation != operation {
			continue
		}
		if !strings.EqualFold(idPattern, r.IdPattern) {
			continue
		}
		propsSet := make(map[string]bool)
		propsSet["name"] = true
		for _, prop := range r.Properties {
			propsSet[strings.ReplaceAll(prop, "/", ".")] = true
		}
		covered := make([]string, 0)
		uncovered := make([]string, 0)
		for _, prop := range props {
			if propsSet[prop] {
				covered = append(covered, prop)
			} else {
				uncovered = append(uncovered, prop)
			}
		}
		return covered, uncovered
	}
	return []string{}, props
}
