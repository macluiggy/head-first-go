package main
import (
	// "fmt"
	"strings"
)


func main()  {
	broken := "G# r#cks"
	// fmt.Println(strings.Replace(broken, "#", "o", -1))
	replacer := strings.NewReplacer("#", "o")
	print(replacer)
	fixed := replacer.Replace(broken)
	print(fixed)
	// fmt.Println(fixed)
}