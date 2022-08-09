package main

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"os"
	"strings"

	"golang.org/x/tour/pic"
)

func main() {

	fmt.Println("started")
	fmt.Println("")
	fmt.Println("")
	typSwi()
	fmt.Println("")
	fmt.Println("stringers")
	exer()
	fmt.Println("")
	fmt.Println("rot13Reader")
	Rot13Read()
	fmt.Println("image")
	imgTest()

}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func typSwi() {
	do(21)
	do("hello")
	do(true)
}

// Stringers exercise

type IPAddr [4]byte

func (ip IPAddr) string() string {
	IPAddrString := []string{}
	for _, num := range ip {
		IPAddrString = append(IPAddrString, fmt.Sprint(int(num)))
	}
	return strings.Join(IPAddrString, ".")
}

func exer() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

// Exercise rot13Reader

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'M':
		b = (b - 'A') + 'N'
	case 'N' <= b && b <= 'Z':
		b = (b - 'N') + 'A'
	case 'a' <= b && b <= 'm':
		b = (b - 'a') + 'n'
	case 'n' <= b && b <= 'z':
		b = (b - 'n') + 'a'
	}
	return b
}

func (b *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = b.r.Read(p)
	for i := range p[:n] {
		p[i] = rot13(p[i])
	}
	return
}

func Rot13Read() {
	s := strings.NewReader(
		"Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

//Exercise Image

type Image struct {
	width, height int
	colr          uint8
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.width, i.height)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{i.colr + uint8(x), i.colr + uint8(y), 255, 255}
}

func imgTest() {
	m := Image{100, 100, 128}
	pic.ShowImage(m)
}
