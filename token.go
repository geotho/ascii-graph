package asciigraph

type Token struct {
	pos Point
	val string
}

func (t Token) IsEdge() bool {
	switch t.val {
	case "/", `\`:
		return true
	}

	return false
}

func (t Token) Endpoints() (left, right Point) {
	switch t.val {
	case `\`:
		left = Point{t.pos.x - 1, t.pos.y - 1}
		right = Point{t.pos.x + 1, t.pos.y + 1}
		return
	default: // /
		left = Point{t.pos.x - 1, t.pos.y + 1}
		right = Point{t.pos.x + 1, t.pos.y - 1}
		return
	}
}
