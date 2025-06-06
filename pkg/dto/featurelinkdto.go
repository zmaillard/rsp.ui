package dto

type FeatureLinkDto struct {
	ID            uint
	RoadName      string
	FromFeatureId uint `yaml:"fromfeatureid"`
	ToFeatureId   uint `yaml:"tofeatureid"`
	Highways      []string
	Direction     float64
	Bearing       string
	HighwayName   *string `yaml:"highwayName,omitempty"`
}
