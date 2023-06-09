package generator

type Generator interface {
	OutFile() string
	ToMarkdown() ([]byte, error)
}
