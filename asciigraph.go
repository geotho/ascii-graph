package asciigraph

import (
	"bufio"
	"unicode"
)

type Node struct {
	id  string
	adj []*Node
}

type Point struct {
	x, y int
}

func Parse(scanner *bufio.Scanner, g Graph) {
	var tokens []Token
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		tokens = append(tokens, Tokenise(line, y)...)
	}

	m := Map(tokens)

	for _, token := range tokens {
		if !token.IsEdge() {
			g.AddNode(token.val)
			continue
		}

		p1, p2 := token.Endpoints()
		n1, ok := m[p1]
		if !ok {
			continue
		}
		n2, ok := m[p2]
		if !ok {
			continue
		}

		g.AddEdge(n1.val, n2.val)
	}
}

func Tokenise(str string, y int) []Token {
	var tokens []Token
	tokenStart := 0
	prevSpace := true

	for x, r := range str {
		if prevSpace == unicode.IsSpace(r) {
			continue
		}

		if prevSpace {
			tokenStart = x
			prevSpace = false
			continue
		}

		tokens = append(tokens, Token{
			Point{tokenStart, y},
			str[tokenStart:x],
		})

		prevSpace = true
	}

	if tokenStart < len(str) {
		tokens = append(tokens, Token{
			Point{tokenStart, y},
			str[tokenStart:],
		})
	}

	return tokens
}

func Map(tokens []Token) map[Point]Token {
	m := make(map[Point]Token, len(tokens))
	for _, t := range tokens {
		i := 0
		for range t.val {
			m[Point{t.pos.x + i, t.pos.y}] = t
			i++
		}
	}
	return m
}
