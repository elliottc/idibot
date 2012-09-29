package grid

import (
	"container/heap"
)

type step struct {
	c Coordinate // graph location
	h float64    // heuristic value
}

type stepSlice []step

func (s *stepSlice) Push(t interface{}) {
	*s = append(*s, t.(step))
}

func (s *stepSlice) Pop() interface{} {
	l := len(*s) - 1
	t := (*s)[l]
	*s = (*s)[:l]
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

func checkCoordinate(g Interface, c Coordinate) bool {
	if c.Row < 0 || c.Row >= g.NumRows() || c.Col < 0 || c.Col >= g.NumCols() {
		return false
	}
	return true
}

// A* algorithm for a 2D square grid, where steps are up/down/left/right.
// Returns a list of coordinates from start to end and true if a path is found,
// otherwise undefined and false.
func AStarForGrid(g Interface, start, end Coordinate) ([]Coordinate, bool) {
	if !checkCoordinate(g, start) {
		panic("Invalid start coordinate")
	}
	if !checkCoordinate(g, end) {
		panic("Invalid end coordinate")
	}

	// Explore the frontier until the end is found.
	origin := make(map[Coordinate]Coordinate)
	explored := map[Coordinate]bool{start: true}
	frontier := stepSlice{step{start, EuclideanDistance(g, start, end)}}

	for _, endFound := explored[end]; !endFound && len(frontier) > 0; _, endFound = explored[end] {
		// Pick a coordinate from the frontier and add its adjacent coordinates to the frontier.
		s := heap.Pop(&frontier).(step)
		for _, t := range adjacent(g, s.c) {
			if _, e := explored[t]; !e && g.IsPassable(t) {
				heap.Push(&frontier, step{t, EuclideanDistance(g, t, end)})
				origin[t] = s.c
				explored[t] = true
			}
		}
	}

	// Construct the path from start to end.
	path := make([]Coordinate, 0)
	c := end
	for o, ok := origin[c]; ok; o, ok = origin[c] {
		path = append(path, c)
		c = o
	}
	// Reverse the path.
	for i, j := 0, len(path)-1; i < j; {
		path[i], path[j] = path[j], path[i]
		i++
		j--
	}

	// Return the path and whether it was found.
	_, endFound := explored[end]
	return path, endFound
}

func adjacent(g Interface, c Coordinate) []Coordinate {
	return []Coordinate{
		{normalize(c.Row+1, g.NumRows()), c.Col},
		{normalize(c.Row-1, g.NumRows()), c.Col},
		{c.Row, normalize(c.Col+1, g.NumCols())},
		{c.Row, normalize(c.Col-1, g.NumCols())}}
}

func normalize(coord, maxCoord int) int {
	coord = coord % maxCoord
	if coord < 0 {
		coord += maxCoord
	}
	return coord
}
