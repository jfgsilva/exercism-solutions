package tree

import (
	"fmt"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	// every record is a node
	nodes := make([]Node, len(records))
	rootNode := &nodes[0]
	fmt.Println("records:", records)
	for _, record := range records {
		switch {
		// this is the root node
		// if childer is not nil the children node already populated children slice, and we would erase
		case record.ID == 0 && record.Parent == 0 && nodes[0].Children == nil:
			nodes[record.ID] = Node{ID: record.ID}
			// fmt.Println("here's the root node:", record)
			// fmt.Println(nodes[record.ID])
		case record.ID != 00:
			nodes[record.ID] = Node{ID: record.ID}
			// this seems correct, however it seems that something is missing in creating a relation between different nesting levels
			nodes[record.Parent].Children = append(nodes[record.Parent].Children, &nodes[record.ID])
			// after appending we need to reorder the children
		}
		// the children slice needs to be sorted
		children := nodes[record.Parent].Children
		sort.Slice(children, func(i, j int) bool {
			return children[i].ID < children[j].ID
		})

	}
	fmt.Println("nodes:", nodes)
	return rootNode, nil
}
