package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetUserInput() []string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")
	if scanner.Scan() {
		return strings.Split(scanner.Text(), " ")
	}

	return nil
}
