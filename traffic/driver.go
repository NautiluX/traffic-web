package traffic

import (
  "fmt"
  "os/exec"
  "strconv"
  "bytes"
  "log"
)

func Red(){
  fmt.Println("Setting color to red")
  setColors(1,0,0)
}

func Yellow(){
  fmt.Println("Setting color to yellow")
  setColors(0,1,0)
}

func Green(){
  fmt.Println("Setting color to green")
  setColors(0,0,1)
}

func Off(){
  fmt.Println("Turning off traffic light")
  setColors(0,0,0)
}

func setColors(red, yellow, green int) {
  setLight(0, red)
  setLight(1, yellow)
  setLight(2, green)
}

func setLight(num, state int) {
  cmd := exec.Command("clewarecontrol", "-c", "1", "-as", strconv.Itoa(num), strconv.Itoa(state))
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("output: %q\n", out.String())
}
