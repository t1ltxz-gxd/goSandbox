package tasks

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) {
	<-time.After(duration)
}

func TwentyFifth() {
	fmt.Println("Start")
	Sleep(3 * time.Second)
	fmt.Println("End")
}
