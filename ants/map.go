package ants

import (
	"idibot/lib/grid"
	"log"
)

//Item represents all the various items that may be on the map
type Item int8

const (
	UNKNOWN Item = iota - 5
	WATER
	FOOD
	LAND
	DEAD
	MY_ANT //= 0
	ANT_1
	ANT_2
	ANT_3
	ANT_4
	ANT_5
	ANT_6
	ANT_7
	ANT_8
	ANT_9

	MY_HILL //= 10
	HILL_1
	HILL_2
	HILL_3
	HILL_4
	HILL_5
	HILL_6
	HILL_7
	HILL_8
	HILL_9

	MY_OCCUPIED_HILL //= 20
	OCCUPIED_HILL_1
	OCCUPIED_HILL_2
	OCCUPIED_HILL_3
	OCCUPIED_HILL_4
	OCCUPIED_HILL_5
	OCCUPIED_HILL_6
	OCCUPIED_HILL_7
	OCCUPIED_HILL_8
	OCCUPIED_HILL_9
)

//IsOccupied returns true if hillOrAnt is an occupied ant hill.
func (hillOrAnt Item) IsOccupied() bool {
	if hillOrAnt >= MY_OCCUPIED_HILL && hillOrAnt <= OCCUPIED_HILL_9 {
		return true
	}
	return false
}

//IsAnt returns true if o is an ant or occupied hill
func (o Item) IsAnt() bool {
	if o >= MY_OCCUPIED_HILL && o <= OCCUPIED_HILL_9 {
		return true
	}
	if o >= MY_ANT && o <= ANT_9 {
		return true
	}
	return false
}

//IsHill returns true if o is an un/occupied hill
func (o Item) IsHill() bool {
	if o >= MY_OCCUPIED_HILL && o <= OCCUPIED_HILL_9 {
		return true
	}
	if o >= MY_HILL && o <= HILL_9 {
		return true
	}
	return false
}

//Player returns the player number of the given ant/hill (0 - 9)
func (o Item) Player() int {
	if o < 0 || o > OCCUPIED_HILL_9 {
		log.Panicf("Expected an ant or a hill!")
	}
	return int(o) % 10
}

//ToUnoccupied returns the HILL_X version of the given hill or ant.
func (hillOrAnt Item) ToUnoccupied() Item {
	return Item(hillOrAnt.Player()) + MY_HILL
}

//ToOccupied returns the OCCUPIED_HILL_X version of the given hill or ant.
func (hillOrAnt Item) ToOccupied() Item {
	return Item(hillOrAnt.Player()) + MY_OCCUPIED_HILL
}

//ToAnt returns the ANT_X version of the given hill or ant.
func (hillOrAnt Item) ToAnt() Item {
	return Item(hillOrAnt.Player())
}

//Symbol returns the symbol for the ascii diagram
func (o Item) Symbol() byte {
	switch o {
	case UNKNOWN:
		return '.'
	case WATER:
		return '%'
	case FOOD:
		return '*'
	case LAND:
		return ' '
	case DEAD:
		return '!'
	}
	if o >= MY_HILL && o <= HILL_9 {
		return byte(o-MY_HILL) + '0'
	}
	if o >= MY_OCCUPIED_HILL && o <= OCCUPIED_HILL_9 {
		return byte(o-MY_OCCUPIED_HILL) + 'A'
	}
	if o < MY_ANT || o > ANT_9 {
		log.Panicf("invalid item: %v", o)
	}
	return byte(o) + 'a'
}

//FromSymbol reverses Symbol
func FromSymbol(ch byte) Item {
	switch ch {
	case '.':
		return UNKNOWN
	case '%':
		return WATER
	case '*':
		return FOOD
	case ' ':
		return LAND
	case '!':
		return DEAD
	}
	if ch >= '0' && ch <= '9' {
		return MY_HILL + Item(ch-'0')
	}
	if ch >= 'A' && ch <= 'J' {
		return MY_OCCUPIED_HILL + Item(ch-'A')
	}
	if ch < 'a' || ch > 'j' {
		log.Panicf("invalid item symbol: %v", ch)
	}
	return Item(ch) + 'a'
}

type Map struct {
	Rows int
	Cols int

	itemGrid []Item

	Ants         map[grid.Location]Item
	Hills        map[grid.Location]Item
	Dead         map[grid.Location]Item
	Water        map[grid.Location]bool
	Food         map[grid.Location]bool
	Destinations map[grid.Location]bool
}

//NewMap returns a newly constructed blank map.
func NewMap(Rows, Cols int) *Map {
	m := &Map{
		Rows:     Rows,
		Cols:     Cols,
		Water:    make(map[grid.Location]bool),
		itemGrid: make([]Item, Rows*Cols),
	}
	m.Reset()
	return m
}

//String returns an ascii diagram of the map.
func (m *Map) String() string {
	str := ""
	for row := 0; row < m.Rows; row++ {
		for col := 0; col < m.Cols; col++ {
			s := m.itemGrid[row*m.Cols+col].Symbol()
			str += string([]byte{s}) + " "
		}
		str += "\n"
	}
	return str
}

//Reset clears the map (except for water) for the next turn
func (m *Map) Reset() {
	for i := range m.itemGrid {
		m.itemGrid[i] = UNKNOWN
	}
	for i, val := range m.Water {
		if val {
			m.itemGrid[i] = WATER
		}
	}
	m.Ants = make(map[grid.Location]Item)
	m.Hills = make(map[grid.Location]Item)
	m.Dead = make(map[grid.Location]Item)
	m.Food = make(map[grid.Location]bool)
	m.Destinations = make(map[grid.Location]bool)
}

//Item returns the item at a given location
func (m *Map) Item(loc grid.Location) Item {
	return m.itemGrid[loc]
}

//AddWater adds water to the map.
func (m *Map) AddWater(loc grid.Location) {
	m.Water[loc] = true
	m.itemGrid[loc] = WATER
}

//AddAnt adds an ant to the map. It can also accept an occupied ant hill.
func (m *Map) AddAnt(loc grid.Location, ant Item) {
	m.Ants[loc] = ant.ToAnt()
	if ant.IsOccupied() {
		m.Hills[loc] = ant.ToUnoccupied()
	}
	if m.Hills[loc] == ant.ToUnoccupied() {
		ant = ant.ToOccupied() //be sure to record the right thing in the itemGrid
	}
	m.itemGrid[loc] = ant
}

//AddHill takes an unoccupied ant hill and adds it to the map.
func (m *Map) AddHill(loc grid.Location, hill Item) {
	m.Hills[loc] = hill.ToUnoccupied()
	if m.Ants[loc] == hill.ToAnt() {
		hill = hill.ToOccupied() //an ant has already been added here!
	}
	m.itemGrid[loc] = hill
}

//AddLand adds a circle of land centered on the given location
func (m *Map) AddLand(center grid.Location, viewrad2 int) {
	m.DoInRad(center, viewrad2, func(row, col int) {
		loc := m.FromRowCol(row, col)
		if m.itemGrid[loc] == UNKNOWN {
			m.itemGrid[loc] = LAND
		}
	})
}

//DoInRad performs the given action for every square within the given circle.
func (m *Map) DoInRad(center grid.Location, rad2 int, Action func(row, col int)) {
	row1, col1 := m.FromLocation(center)
	for row := row1 - m.Rows/2; row < row1+m.Rows/2; row++ {
		for col := col1 - m.Cols/2; col < col1+m.Cols/2; col++ {
			row_delta := row - row1
			col_delta := col - col1
			if row_delta*row_delta+col_delta*col_delta < rad2 {
				Action(row, col)
			}
		}
	}
}

func (m *Map) AddDeadAnt(loc grid.Location, ant Item) {
	m.Dead[loc] = ant
	m.itemGrid[loc] = DEAD
}

func (m *Map) AddFood(loc grid.Location) {
	m.Food[loc] = true
	m.itemGrid[loc] = FOOD
}

func (m *Map) AddDestination(loc grid.Location) {
	if m.Destinations[loc] {
		log.Panicf("Already have something at that destination!")
	}
	m.Destinations[loc] = true
}

func (m *Map) RemoveDestination(loc grid.Location) {
	delete(m.Destinations, loc)
}

//SafeDestination will tell you if the given location is a 
//safe place to dispatch an ant. It considers water and both
//ants that have already sent an order and those that have not.
func (m *Map) SafeDestination(loc grid.Location) bool {
	if _, exists := m.Water[loc]; exists {
		return false
	}
	if occupied := m.Destinations[loc]; occupied {
		return false
	}
	return true
}

//FromRowCol returns a Location given an (Row, Col) pair
func (m *Map) FromRowCol(Row, Col int) grid.Location {
	for Row < 0 {
		Row += m.Rows
	}
	for Row >= m.Rows {
		Row -= m.Rows
	}
	for Col < 0 {
		Col += m.Cols
	}
	for Col >= m.Cols {
		Col -= m.Cols
	}

	return grid.Location(Row*m.Cols + Col)
}

//FromLocation returns an (Row, Col) pair given a Location
func (m *Map) FromLocation(loc grid.Location) (Row, Col int) {
	Row = int(loc) / m.Cols
	Col = int(loc) % m.Cols
	return
}

//Direction represents the direction concept for issuing orders.
type Direction int

const (
	North Direction = iota
	East
	South
	West

	NoMovement
)

func (d Direction) String() string {
	switch d {
	case North:
		return "n"
	case South:
		return "s"
	case West:
		return "w"
	case East:
		return "e"
	case NoMovement:
		return "-"
	}
	log.Panicf("%v is not a valid direction", d)
	return ""
}

//Move returns a new location which is one step in the specified direction from the specified location.
func (m *Map) Move(loc grid.Location, d Direction) grid.Location {
	// TODO: this could be changed to be Move(from x,y  to x,y)
	Row, Col := m.FromLocation(loc)
	switch d {
	case North:
		Row -= 1
	case South:
		Row += 1
	case West:
		Col -= 1
	case East:
		Col += 1
	case NoMovement: //do nothing
	default:
		log.Panicf("%v is not a valid direction", d)
	}
	return m.FromRowCol(Row, Col) //this will handle wrapping out-of-bounds numbers
}
