package tree

import (
	"errors"
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
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	for index, record := range records {
		switch {
		case record.ID != index:
			return nil, errors.New("record ID doesn't match index: duplicate")
		case record.ID != 0 && record.ID == record.Parent:
			return nil, errors.New("direct cycle: only root can be it's own parent")
		case record.ID < record.Parent:
			return nil, errors.New("record id cannot be higher than parent id")
		// this is the root node
		// if childer is not nil the children node already populated children slice, and we would erase
		case record.ID == 0 && record.Parent != 0:
			return nil, errors.New("root node can't have parent")
		case record.ID == 0 && record.Parent == 0 && nodes[0].Children == nil:
			nodes[record.ID].ID = record.ID
		case record.ID != 00 && index == 0:
			return nil, errors.New("single record which isn't root")
		case record.ID != 00 && index != 0:
			nodes[record.ID].ID = record.ID
			nodes[record.Parent].Children = append(nodes[record.Parent].Children, &nodes[record.ID])
		}
	}
	return rootNode, nil
}
