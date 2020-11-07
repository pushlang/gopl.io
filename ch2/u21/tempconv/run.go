// Cf converts its numeric argument to Celsius and Fahrenheit.
// rm u21; goimports -v -w *.go; gofmt -w *.go; go build; ./u21
package tempconv

import (
	"fmt"
	"os"
	"strconv"
)

func Run() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := Fahrenheit(t)
		c := Celsius(t)
		k := Kelvin(t)
		fmt.Printf("%s = %s, %s = %s, %s = %s, %s = %s, %s = %s, %s = %s\n",
			c, CToF(c), c, CToK(c), f, FToC(f), f, FToK(f), k, KToC(k), k, KToF(k))
	}
}
