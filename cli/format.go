package cli

import "fmt"

// \033[38;2;<r>;<g>;<b>m -- forground
const CHECK = "\xE2\x9C\x94"
const WARN = "\xE2\x9A\xA0"
const CROSS = "\xF0\x9D\x97\xAB"
const CLEAR = "\033[0m"
const RED = "\033[38;2;255;0;0m" + CROSS + CLEAR
const YELLOW = "\033[38;2;217;158;1m" + WARN + CLEAR
const GREEN = "\033[38;2;0;255;0m" + CHECK + CLEAR

const HEADER = "---------- Home Lab Stats ----------"

func defaultFormat(input map[string]StatusResponse) {
	var color string
	fmt.Println(HEADER)
	for key, val := range input {
		switch val.status {
		case "OK":
			color = GREEN
		case "WARN":
			color = YELLOW
		case "FAIL":
			color = RED
		}
		fmt.Printf("   %s %s\n", color, key)
	}
}
