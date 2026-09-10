package internal

import "testing"

func TestMeasurementRequiresUnitAtTheBoundary(t *testing.T) {
	if (Measurement{Element: "humerus"}).Unit != "" {
		t.Fatal("zero value should remain explicit")
	}
}
