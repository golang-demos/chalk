package chalk

import (
	"testing"
)

func Test_BasicColors(t *testing.T) {

	var colorTests = []struct {
		Expected  string
		Actual    string
		ColorName string
	}{
		{"\u001b[30m", Black().String(), "Black"},
		{"\u001b[31m", Red().String(), "Red"},
		{"\u001b[32m", Green().String(), "Green"},
		{"\u001b[33m", Yellow().String(), "Yellow"},
		{"\u001b[34m", Blue().String(), "Blue"},
		{"\u001b[35m", Magenta().String(), "Magenta"},
		{"\u001b[36m", Cyan().String(), "Cyan"},
		{"\u001b[37m", White().String(), "White"},
		{"\u001b[90m", BlackLight().String(), "BlackLight"},
		{"\u001b[91m", RedLight().String(), "RedLight"},
		{"\u001b[92m", GreenLight().String(), "GreenLight"},
		{"\u001b[93m", YellowLight().String(), "YellowLight"},
		{"\u001b[94m", BlueLight().String(), "BlueLight"},
		{"\u001b[95m", MagentaLight().String(), "MagentaLight"},
		{"\u001b[96m", CyanLight().String(), "CyanLight"},
		{"\u001b[97m", WhiteLight().String(), "WhiteLight"},
	}

	for _, test := range colorTests {
		if test.Actual != test.Expected {
			t.Error("Expected modifier is not same as actual string for Color " + test.ColorName)
		}
	}
}
