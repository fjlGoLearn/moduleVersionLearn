package moduleVersionLearn

import "fmt"

const Version = "V1.0.0"

func main() {
	fmt.Println(GetVersion())
}

func GetVersion() string {
	return Version
}
