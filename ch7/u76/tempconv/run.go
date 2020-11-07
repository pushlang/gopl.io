package tempconv

import (
	"flag"
	"fmt"
)

var temp = CelsiusFlag("temp", 20.0, "the temperature")
//var ff = FahrenheitFlag("ff", 20.0, "the temperature")
//var kf = KelvinFlag("kf", 20.0, "the temperature")

func Run() {
	flag.Parse()
	fmt.Println(*temp) 
}
