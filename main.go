package main

import (
	"fmt"
	// "encoding/binary"
	"time"
	// "strings"
	"crypto/rand"
	"math/big"
)


type Color struct {
	R, G, B float64
}

func NewColor(r, g, b float64) Color {
	return Color{R: r, G: g, B: b}
}

func softClearScreen() {
	fmt.Print("\033[2J\033[H")
}

func hardClearScreen() {
	fmt.Print("\033c")
}

func do(s any) {
	str, ok := s.(string)
	if ok {
		fmt.Printf(str)
	} else {
		fmt.Printf(string(s.(byte)))
	}
}

func getForeGroundColor(rgb ...float64) string {
	r := int(rgb[0])
	g := int(rgb[1])
	b := int(rgb[2])
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func getBackGroundColor(rgb ...float64) string {
	r := int(rgb[0])
	g := int(rgb[1])
	b := int(rgb[2])
	return fmt.Sprintf("\033[48;2;%d;%d;%dm", r, g, b)
}

func getReset() string {
	return fmt.Sprintf("\033[0m")
}

func getNLinesAbove(n int) string {
	return fmt.Sprintf("\033[%dA", n)
}

// func colorPrint(r, g, b int, text string) {
// 	foreground := getForeGroundColor(r, g, b)
// 	background := getBackGroundColor(r, g, b)
// 	fmt.Printf(foreground + background + text + getReset())
// }

func getRand(mod float64) float64 {
	if mod <= 0 {
		return 0.
	}

	ng, err := rand.Int(rand.Reader, big.NewInt(int64(mod)))
	if err != nil {
		fmt.Println("error")
	}
	return float64(ng.Int64())
	// var random int = 0
	// binary.Read(rand.Reader, binary.LittleEndian, &random)
	// return random % mod
}

func RandAB(a, b float64) float64 {
	// assume a < b
	rng := b - a
	return getRand(rng) + a
}

func randC(color int) string {
	a := float64(color)// + int(getRand(256)) / 30 //+ 120
	b := float64(color) //+ int(getRand(256)) / 20 //+ 220
	c := float64(color) //+ int(getRand(256)) / 30 //+ 220
	// color = getRand(50)
	// a := color / 20 + 55
	// b := color / 2  + 134
	// c := color / 50 + 10
	// a := int(getRand(256)) + 35
	// b := int(getRand(256)) - 33
	// c := int(getRand(256)) + 32
	// fmt.Println(a, b, c)
	return getForeGroundColor(a, b, c)
}

func secret(flame, i, j, n, m int) []float64 {
	// saltr := 55 saltg := 55 saltb := 25 delta := flame / n


	// var jitter float64 = 6.
	var k float64  = 0.95
	var M float64 = float64(flame) / float64(n)
	r := min(255., M * float64( i ) * k)
	g := min(255., M * float64( i ) * k)
	b := min(255., M * float64( i ) * k)

	// var delta float64 = 20 * float64(n) / float64(i + 1) 
	// fmt.Println(delta)
	rands := []float64{r, g, b}
	rands[0] = min(245., r + RandAB(240, 250))
	rands[1] = min(245., g + RandAB(50, 140))
	rands[2] = min(245., b + RandAB(60, 50))
	return rands
}

func eventLoop(script byte, n, m, flame int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			color := secret(flame, i, j, n, m)
			// do(randC(color))
			do(getForeGroundColor(color...))
			do(script)
			do(getReset())
		}
		do("\n")
		// color -= (i + 1) * (5 * delta)
		// color = color - 2 * delta

	}
		// for j := 0; j < 10; j++ { s.WriteByte(' ') }
}

func wrapUp(s string) string {
	return getForeGroundColor(3, 3, 3)
}

func Run() {
	refreshRate := 25 * time.Millisecond
	script := byte('S')
	flame := 256
	n := 10
	m := 20

	for {
		softClearScreen()
		eventLoop(script, n, m, flame)
		 // break
		do(getNLinesAbove(n - 1))
		time.Sleep(refreshRate)
	}
}

func main() {
	Run()
}
