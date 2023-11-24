package moduleVersionLearn

import "fmt"

const Version = "V1.0.0"

func main() {
	fmt.Println(GetVersion())
}

func GetVersion() string {
	return Version
}

func GetVersionWithInput(input string) string {
	return input + Version
}
