package table_driven

import "testing"

func TestConvertFreezingCtoF(t *testing.T) {
	result, _ := ConvertTemperature(0, "C")
	if result != 32 {
		t.Errorf("ConvertTemperature is %v, expected 32", result)
	}
}

func TestConvertBoilingCtoF(t *testing.T) {
	result, _ := ConvertTemperature(100, "C")
	if result != 212 {
		t.Errorf("ConvertTemperature is %v, expected 212", result)
	}
}

func TestConvertNegativeCtoF(t *testing.T) {
	result, _ := ConvertTemperature(-40, "C")
	if result != -40 {
		t.Errorf("ConvertTemperature is %v, expected -40", result)
	}
}

func TestConvertFreezingFtoC(t *testing.T) {
	result, _ := ConvertTemperature(32, "F")
	if result != 0 {
		t.Errorf("ConvertTemperature is %v, expected 0", result)
	}
}

func TestConvertBoilingFtoC(t *testing.T) {
	result, _ := ConvertTemperature(212, "F")
	if result != 100 {
		t.Errorf("ConvertTemperature is %v, expected 100", result)
	}
}

func TestConvertUnsupportedScale(t *testing.T) {
	_, err := ConvertTemperature(100, "K")
	if err == nil {
		t.Errorf("ConvertTemperature expected error for unsupported scale K, got nil")
	}
}
