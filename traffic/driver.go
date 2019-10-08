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
  setLight(0, 1)
  setLight(1, 0)
  setLight(2, 0)
}

func Yellow(){
  fmt.Println("Setting color to yellow")
  setLight(0, 0)
  setLight(1, 1)
  setLight(2, 0)
}

func Green(){
  fmt.Println("Setting color to green")
  setLight(0, 0)
  setLight(1, 0)
  setLight(2, 1)
}

func setLight(num, state int) {
  cmd := exec.Command("clewarecontrol", "-c", "1", "-as", strconv.Itoa(num), strconv.Itoa(state))
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("output: %q\n", out.String())
}
