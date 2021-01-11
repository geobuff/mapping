package mapping

// Mapping represents a mapping entry.
type Mapping struct {
	Name             string   `json:"name"`
	Code             string   `json:"code"`
	SVGName          string   `json:"svgName"`
	AlternativeNames []string `json:"alternativeNames"`
	Prefixes         []string `json:"prefixes"`
}
