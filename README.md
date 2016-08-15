# Ascii-art to graph

This Go package turns this:

```
1   2
 \ /
  5
 / \
3   4
```

into this:

```go
g.AddEdge("1", "5")
g.AddEdge("2", "5")
g.AddEdge("3", "5")
g.AddEdge("4", "5")
```

## How?

Implement the `Graph` interface for whatever Graph backend you are using:

```go
type Graph interface {
	AddNode(nodeID string)
	AddEdge(nodeOneID, nodeTwoID string)
```

Then use the `Parse` method in this package, supplying your string and the Graph implementation.

## Todos:

  - Support straight-line edges: `|` and `-`.
  - Allow sloppy graph drawing (currently the edge has to lead directly to the node).
  - Allow longer edges.
