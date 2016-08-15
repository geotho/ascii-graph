package asciigraph

import (
	"bufio"
	"strings"
	"testing"
)

const g1 = `
1   2
 \ /
  5
 / \
3   4
`

const g2 = `
      17
     /
1   2
 \ / \
  5   15
 / \
3   4
`

const g3 = `
1 - 2 - 3
|   |   |
4 - 5 - 6
`

func AssertCorrectGraph(t *testing.T, g string, expNodes []string, expEdges []TestEdge) {
	graph := NewTestGraph()
	Parse(bufio.NewScanner(strings.NewReader(g)), graph)

	if len(expNodes) != len(graph.Nodes) {
		t.Errorf("Nodes is the wrong length")
	}

	for _, s := range expNodes {
		_, ok := graph.Nodes[s]
		if !ok {
			t.Errorf("Node %s missing", s)
		}
	}

	if len(expEdges) != len(graph.Edges) {
		t.Errorf("Edges is the wrong length")
	}

	for _, e := range expEdges {
		_, ok := graph.Edges[e]
		if !ok {
			t.Errorf("Edge %+v missing", e)
		}
	}
}

func TestParse(t *testing.T) {
	type testcase struct {
		G        string
		ExpNodes []string
		ExpEdges []TestEdge
	}

	testcases := []testcase{
		{
			G: `
1   2
 \ /
  5
 / \
3   4
`,
			ExpNodes: []string{"1", "2", "3", "4", "5"},
			ExpEdges: []TestEdge{
				{"1", "5"},
				{"2", "5"},
				{"3", "5"},
				{"4", "5"},
				{"5", "1"},
				{"5", "2"},
				{"5", "3"},
				{"5", "4"},
			},
		},
		{
			G: `
      17
     /
1   2
 \ / \
  5   15
 / \
3   4
`,
			ExpNodes: []string{"1", "2", "3", "4", "5", "15", "17"},
			ExpEdges: []TestEdge{
				{"1", "5"},
				{"2", "5"},
				{"3", "5"},
				{"4", "5"},
				{"2", "15"},
				{"2", "17"},
				{"5", "1"},
				{"5", "2"},
				{"5", "3"},
				{"5", "4"},
				{"17", "2"},
				{"15", "2"},
			},
		},
	}

	for _, test := range testcases {
		AssertCorrectGraph(t, test.G, test.ExpNodes, test.ExpEdges)
	}
}

func TestTokenise(t *testing.T) {
	type testcase struct {
		input  string
		output []Token
	}

	tests := []testcase{
		{
			input: "abcd",
			output: []Token{
				{
					pos: Point{0, 1},
					val: "abcd",
				},
			},
		},
		{
			input: "     17 a b cd",
			output: []Token{
				{
					pos: Point{5, 1},
					val: "17",
				},
				{
					pos: Point{8, 1},
					val: "a",
				},
				{
					pos: Point{10, 1},
					val: "b",
				},
				{
					pos: Point{12, 1},
					val: "cd",
				},
			},
		},
	}

	for _, test := range tests {
		x := Tokenise(test.input, 1)
		if l1, l2 := len(x), len(test.output); l1 != l2 {
			t.Errorf("\nactualLen=%d, expectedLen=%d, \nactual=%+v, \nexpected=%+v", l1, l2, x, test.output)
			return
		}
		for i, token := range x {
			if token != test.output[i] {
				t.Errorf("\ninput=%q,\nactual=  %+v,\nexpected=%+v", test.input, x, test.output)
			}
		}
	}
}
