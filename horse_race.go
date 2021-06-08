// You can directly run this program here: https://play.golang.org/p/OIQCnBzJhgd

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func horseman(name string) {

	for i := 0; i <= 10; i++ {
		speed := rand.Intn(6)

		if speed == 1 {
			fmt.Printf("- %s is on fire!\n", name)
		} else if speed == 3 {
			fmt.Printf("- %s is keeping a constant speed.\n", name)
		} else if speed == 5 {
			fmt.Printf("- %s is slowing down!\n", name)
		}

		time.Sleep(time.Duration(speed) * time.Second)

		switch i {
		case 5:
			fmt.Printf("- %s is halfway through!\n", name)
		case 8:
			fmt.Printf("- %s is heading towards the end!\n", name)
		case 10:
			fmt.Printf("- %s is almost here!\n", name)
		}
	}

	fmt.Printf("\n-- %s CROSSED THE FINISHING LINE! --\n\n", name)
}

func main() {
	fmt.Println("WELCOME TO THE HORSE RACE!\nLet's see what positive things these gentlemen bring to us!")
	time.Sleep(2 * time.Second)
	fmt.Println("\n- Ready... GO! (No pun intended)")
	time.Sleep(4 * time.Second)

	go horseman("DEATH")
	go horseman("FAMINE")
	go horseman("PESTILENCE")
	go horseman("WAR")

	time.Sleep(60 * time.Second)
}

