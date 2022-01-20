package xiaobuzzer

import (
	"strconv"
	"time"

	utils "github.com/therebelrobot/ghostfacer/utils/shared"
	"tinygo.org/x/drivers/buzzer"
)

func PlayTone(buzzer buzzer.Device, tone int, duration int) {
	utils.Log("PlayTone " + strconv.FormatInt(int64(tone), 10) + " " + strconv.FormatInt(int64(duration), 10))
	for i := 0; i < duration*1000; i += tone * 2 {
		buzzer.On()
		DelayMicroseconds(tone)
		buzzer.Off()
		DelayMicroseconds(tone)
	}
	utils.Log("after PlayTone")

}

func DelayMicroseconds(duration int) {
	utils.Log("DelayMicroseconds")
	time.Sleep(time.Duration(duration) * time.Microsecond)
}

func PlayNote(buzzer buzzer.Device, note string, duration int) {
	utils.Log("PlayNote")
	Names, Tones := GetNotesAndTones()
	SPEE := 5

	// play the tone corresponding to the note name

	for i := 0; i < len(Names); i++ {
		utils.Log("i " + strconv.FormatInt(int64(i), 10))
		if Names[i] == note {
			utils.Log("Names[i] " + Names[i])

			newduration := duration / SPEE
			PlayTone(buzzer, Tones[i], newduration)
		}
	}
}

func PlaySong(buzzer buzzer.Device, notesStr string, beats []int, tempo int) {
	utils.Log("PlaySong")
	notes := []rune(notesStr)
	for i := 0; i < len(beats); i++ {
		if notes[i] == ' ' {
			time.Sleep(time.Duration(beats[i]*tempo) * time.Millisecond) // rest
		} else {
			PlayNote(buzzer, string(notes[i]), beats[i]*tempo)
		}
		// pause between notes
		time.Sleep(time.Duration(tempo) * time.Millisecond)
	}
}
func GetNotesAndTones() (notes []string, tones []int) {
	utils.Log("GetNotesAndTones")
	Names := make([]string, 16)
	Tones := make([]int, 16)
	Names[0] = "C"
	Tones[0] = 1915
	Names[1] = "D"
	Tones[1] = 1700
	Names[2] = "E"
	Tones[2] = 1519
	Names[3] = "F"
	Tones[3] = 1432
	Names[4] = "G"
	Tones[4] = 1275
	Names[5] = "A"
	Tones[5] = 1136
	Names[6] = "B"
	Tones[6] = 1014
	Names[7] = "c"
	Tones[7] = 956
	Names[8] = "d"
	Tones[8] = 834
	Names[9] = "e"
	Tones[9] = 765
	Names[10] = "f"
	Tones[10] = 593
	Names[11] = "g"
	Tones[11] = 468
	Names[12] = "a"
	Tones[12] = 346
	Names[13] = "b"
	Tones[13] = 224
	Names[14] = "x"
	Tones[14] = 655
	Names[15] = "y"
	Tones[15] = 715
	return Names, Tones
}
