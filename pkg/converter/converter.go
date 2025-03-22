package converter

import (
	"highway-sign-portal-builder/pkg/generator"
	"iter"
)

type Converter interface {
	Convert() iter.Seq[generator.Generator]
}
