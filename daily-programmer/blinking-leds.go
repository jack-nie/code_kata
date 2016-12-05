//	Description
//
//	Mark saw someone doing experiments with blinking LEDs (imagine something like this ) and became fascinated by it. He wants to know more about it. He knows you are good with computers, so he comes to you asking if you can teach him how it works. You agree, but as you don't have any LEDs with you at the moment, you suggest: "Let's build an emulator with which we can see what's happening inside". And that's today's challenge.
//	1st Part
//	The 1st part should be easy, even though the description is rather verbose. If you want more challenge try the 2nd part afterwards.
//	Our system has 8 LEDs, we represent their state with a text output. When all LEDs are off, it is printed as string of eight dots "........". When a led is on, it is printed as "*". LED-0 is on the right side (least significant bit), LED-7 is on the left side. Having LEDs 0 and 1 on and all others off is written as "......**"
//	On input you get a sequence of lines forming a program. Read all lines of the input (detect EOF, or make the first line contain number of lines that follow, whichever is more convenient for you). Afterwards, print LED states as they are whenever the program performs an out instruction.
//	Each line is in the following format:
//	<line>: <whitespace> <instruction> |
//	        <empty>
//
//	<instruction> : ld a,<num> |
//	                out (0),a
//	<whitespace> is one or more of characters " " or "\t". <num> is a number between 0 and 255.
//	Instruction ld a,<num> sets internal 8-bit register A to the given number. Instruction out (0),a updates the LEDs according to the current number in A. The LED-0's state corresponds to bit 0 of number in A, when that number is represented in binary. For example, when A = 5, the LED state after out instruction is ".....*.*".
//	You should output the LED states after each out instruction.
//	Challenge input 1:
//	  ld a,14
//	  out (0),a
//	  ld a,12
//	  out (0),a
//	  ld a,8
//	  out (0),a
//
//	  out (0),a
//	  ld a,12
//	  out (0),a
//	  ld a,14
//	  out (0),a
//	Expected output:
//	....***.
//	....**..
//	....*...
//	....*...
//	....**..
//	....***.

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func processingLeds(filePath string) {
	dat, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(dat)
	var formerText string
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}

		matched, err := regexp.MatchString("^ld", text)
		if err != nil {
			log.Fatal(err)
		}

		if matched {
			re := regexp.MustCompile("[0-9]+")
			str := re.FindString(text)
			num, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}

			bnum := strconv.FormatInt(int64(num), 2)
			reForOne := regexp.MustCompile("1")
			reForZero := regexp.MustCompile("0")

			result := reForOne.ReplaceAllString(bnum, "*")
			result = reForZero.ReplaceAllString(result, ".")
			if len(result) < 8 {
				var buffer bytes.Buffer
				length := len(result)
				for i := 0; i < 8-length; i++ {
					buffer.WriteString(".")
				}
				result = buffer.String() + result
			}

			formerText = result
		} else {
			fmt.Println(formerText)
		}
	}
}

func main() {
	processingLeds("tmp/leds.txt")
}
