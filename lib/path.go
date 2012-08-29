package lib

import (
	"container/heap"
)

type step struct {
	c coordinate // graph location
	h float64    // heuristic value
}

type stepSlice []step

func (s stepSlice) Push(t interface{}) {
	s = append(s, t.(step))
}

func (s stepSlice) Pop() interface{} {
	l := len(s) - 1
	t := s[l]
	s = s[:l]
	return t
}

func (s stepSlice) Len() int {
	return len(s)
}

func (s stepSlice) Less(i, j int) bool {
	return s[i].h < s[j].h
}

func (s stepSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// A* algorithm for a 2D square grid, where steps are up/down/left/right.
// Requires that the lengths of the inner slices are equal.
// Returns a list of coordinates from start to end and true if a path is found,
// otherwise undefined and false.
func AStarForGrid(passable [][]bool, start, end coordinate) ([]coordinate, bool) {

	// TODO: slice of slice is bad design because the dimensions aren't strict
	// TODO: pass in a custom function to determine passable?

	xMax := len(passable)
	if xMax == 0 {
		return nil, false
	}
	yMax := len(passable[0])
	if yMax == 0 {
		return nil, false
	}

	// Explore the frontier until the end is found.
	origin := make(map[coordinate]coordinate)
	explored := map[coordinate]bool{start: true}
	frontier := stepSlice{step{start, EuclideanDistance(start.x, start.y, end.x, end.y, xMax, yMax)}}
	for _, endFound := explored[end]; !endFound; _, endFound = explored[end] {
		if len(frontier) == 0 {
			// Break early if there is no more frontier to explore.
			break
		}

		// Pick a coordinate from the frontier and add its adjacent coordinates to the frontier.
		s := heap.Pop(frontier).(step)
		for _, t := range adjacent(s.c, xMax, yMax) {
			if _, e := explored[t]; !e && passable[t.x][t.y] {
				heap.Push(frontier, step{t, EuclideanDistance(t.x, t.y, end.x, end.y, xMax, yMax)})
				origin[t] = s.c
				explored[t] = true
			}
		}
	}

	// Construct the path from start to end.
	path := make([]coordinate, 0)
	c := end
	for _, ok := origin[c]; ok; _, ok = origin[c] {
		path = append(path, c)
		c = origin[c]
	}

	// Return the path and whether it was found.
	_, endFound := explored[end]
	return path, endFound
}

func adjacent(c coordinate, xMax, yMax int) []coordinate {
	return []coordinate{
		{normalize(c.x+1, xMax), c.y},
		{normalize(c.x-1, xMax), c.y},
		{c.x, normalize(c.y+1, yMax)},
		{c.x, normalize(c.y-1, yMax)}}
}

func normalize(coord, maxCoord int) int { // TODO: bad names
	coord = coord % maxCoord
	if coord < 0 {
		coord += maxCoord
	}
	return coord
}
