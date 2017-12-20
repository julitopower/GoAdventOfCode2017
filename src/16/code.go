package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	in := make([]byte, 16)
	for i := 0; i < 16; i++ {
		in[i] = 'a' + byte(i)
	}
	fmt.Println(in)
	content, _ := ioutil.ReadFile("input.txt")
	str := string(content)
	tokens := strings.Split(str, ",")

	for _, token := range tokens {
		switch token[0] {
		case 's':
			count, _ := strconv.Atoi(string(token[1:]))
			prefix := in[16-count:]
			suffix := in[:16-count]
			in = append(prefix, suffix...)
		case 'x':
			token = token[1:]
			positions := strings.Split(token, "/")
			a, _ := strconv.Atoi(positions[0])
			b, _ := strconv.Atoi(positions[1])
			aux := in[a]
			in[a] = in[b]
			in[b] = aux
		case 'p':
			token = token[1:]
			chars := strings.Split(token, "/")
			a, b := 0, 0
			for i, c := range in {
				switch string(c) {
				case chars[0]:
					a = i
				case chars[1]:
					b = i
				}
			}
			aux := in[a]
			in[a] = in[b]
			in[b] = aux
		}
	}

	fmt.Println(string(in))
}
