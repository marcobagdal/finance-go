package util

import (
	"testing"
	"time"
)

func TestStringToDateTime(t *testing.T) {
	//Caso de test 1: entrada válida
	input := time.Date(2019, time.September, 12, 10, 10, 10, 0, time.UTC)
	expectedOutput := time.Date(2019, time.September, 12, 10, 10, 10, 0, time.UTC)

	output, err := StringToDateTime(input)
	if err != nil {
		t.Errorf("Esperado %v, mas recebido foi %v", expectedOutput, output)
	}

	if !output.Equal(expectedOutput) {
		t.Errorf("Esperado %v, mas recebido foi %v", expectedOutput, output)
	}
}
