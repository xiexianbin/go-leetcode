package t5303

import (
	"fmt"
	"strings"
	"testing"
)

func freqAlphabets(s string) string {
	result := []string{}
	dict := map[string]string{}
	dict["1"] = "a"
	dict["2"] = "b"
	dict["3"] = "c"
	dict["4"] = "d"
	dict["5"] = "e"
	dict["6"] = "f"
	dict["7"] = "g"
	dict["8"] = "h"
	dict["9"] = "i"
	dict["10#"] = "j"
	dict["11#"] = "k"
	dict["12#"] = "l"
	dict["13#"] = "m"
	dict["14#"] = "n"
	dict["15#"] = "o"
	dict["16#"] = "p"
	dict["17#"] = "q"
	dict["18#"] = "r"
	dict["19#"] = "s"
	dict["20#"] = "t"
	dict["21#"] = "u"
	dict["22#"] = "v"
	dict["23#"] = "w"
	dict["24#"] = "x"
	dict["25#"] = "y"
	dict["26#"] = "z"

	sL := strings.Split(s, "")
	for {
		if len(sL) == 0 {
			break
		}
		if len(sL) >= 3 && sL[2] == "#" {
			k := strings.Join(sL[:3], "")
			result = append(result, dict[k])
			sL = sL[3:]
		} else {
			k := sL[0]
			result = append(result, dict[k])
			sL = sL[1:]
		}
	}

	return strings.Join(result, "")
}

func TestTest(t *testing.T) {
	for _, s := range []string{"1326#", "25#", "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#", "10#11#12"}{
		fmt.Println(freqAlphabets(s))
	}
}
