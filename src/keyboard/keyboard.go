// get the keyboard input from the user
package keyboard

import (
	"bufio"
	// "github.com/headfirstgo/greeting"
	"os"
	"strconv"
	"strings"
)

// get the keyboard input from the user
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
