package internal

type Measurement struct {
	SpecimenID string  `json:"specimen_id"`
	Element    string  `json:"element"`
	Value      float64 `json:"value"`
	Unit       string  `json:"unit"`
}
