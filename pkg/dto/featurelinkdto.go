package dto

type FeatureLinkDto struct {
	ID            uint
	RoadName      string
	FromFeatureId uint `yaml:"fromfeatureid"`
	ToFeatureId   uint `yaml:"tofeatureid"`
	Highways      []string
}
