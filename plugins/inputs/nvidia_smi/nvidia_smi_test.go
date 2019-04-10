package nvidia_smi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseLineStandard(t *testing.T) {
	line := "2019/04/09 09:13:13.234, GeForce GTX 1070 Ti, P0, 0, 66 %, 324 MHz, 15 %, 324 MHz, 8114, 553, 7561, 61, GPU-xxx, Default, 0.0\n"
	tags, fields, timeStamp, err := parseLine(line)
	if err != nil {
		t.Fail()
	}
	if tags["name"] != "GeForce GTX 1070 Ti" {
		t.Fail()
	}
	if temp, ok := fields["temperature_gpu"].(int); ok && temp == 61 {
		t.Fail()
	}
	require.NotNil(t, timeStamp)
}

func TestParseLineEmptyLine(t *testing.T) {
	line := "\n"
	_, _, _, err := parseLine(line)
	if err == nil {
		t.Fail()
	}
}

func TestParseLineBad(t *testing.T) {
	line := "the quick brown fox jumped over the lazy dog"
	_, _, _, err := parseLine(line)
	if err == nil {
		t.Fail()
	}
}

func TestParseLineNotSupported(t *testing.T) {
	line := "2019/04/09 09:13:13.234, Tesla K80, P0, 0, 66 %, 324 MHz, 15 %, 324 MHz, 7606, 0, 7606, 38, GPU-xxx, Default, 0.0\n"

	_, fields, timeStamp, err := parseLine(line)
	require.NoError(t, err)
	require.NotNil(t, timeStamp)
	if temp, ok := fields["utilization_gpu"].(int); ok && temp == 66 {
		t.Fail()
	}
}
