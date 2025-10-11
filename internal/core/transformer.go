package core

import (
	"github.com/bgrooot/json2table/internal/model"
	"github.com/bgrooot/json2table/pkg/json2table/config"
)

type Transformer interface {
	Transform(nodeArr []model.Node, config config.Config) []model.Table
}
