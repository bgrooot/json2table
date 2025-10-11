package json2table

import (
	"fmt"

	"github.com/bgrooot/json2table/internal/core"
	"github.com/bgrooot/json2table/internal/core/horizontal"
	"github.com/bgrooot/json2table/internal/core/vertical"
	"github.com/bgrooot/json2table/internal/input/json"
	"github.com/bgrooot/json2table/internal/model"
	"github.com/bgrooot/json2table/internal/output"
	"github.com/bgrooot/json2table/pkg/json2table/config"
)

func Json2table(jsonStr string, cfg config.Config) (string, error) {
	config.ApplyDefaultConfig(&cfg)

	omArr, err := json.Unmarshal(jsonStr)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal json: %w", err)
	}

	root := core.Parsing(omArr, cfg)

	core.SetUp(root, cfg)

	var tableArr []model.Table
	switch cfg.Orientation {
	case "horizontal":
		transformer := horizontal.Transformer{}
		tableArr = transformer.Transform(root, cfg)
	case "vertical":
		fallthrough
	default:
		transformer := vertical.Transformer{}
		tableArr = transformer.Transform(root, cfg)
	}

	html, err := output.ToHtml(tableArr, cfg)
	if err != nil {
		return "", fmt.Errorf("failed to convert to html: %w", err)
	}

	return html, nil
}
