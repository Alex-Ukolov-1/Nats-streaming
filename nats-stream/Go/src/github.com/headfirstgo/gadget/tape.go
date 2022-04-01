package gadget

import (
	"fmt"
)

type Point struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Playing")
}

type TapeRecorder struct {
	Microphones int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Recording", song)
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}
