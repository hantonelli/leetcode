package robotRoomCleaner

type Robot struct{}

func (robot *Robot) Move() bool { return false }
func (robot *Robot) TurnLeft()  {}
func (robot *Robot) TurnRight() {}
func (robot *Robot) Clean()     {}

type PlaceState int

const (
	PlaceStateUndefined PlaceState = iota
	PlaceStateWall
	PlaceStateEmpty
)

type Direction int

const (
	DirectionUndefined Direction = iota
	DirectionUp
	DirectionDown
	DirectionLeft
	DirectionRight
)

type Movement int

const (
	_ Movement = iota
	MovementForward
	MovementTurnLeft
	MovementTurnRight
	MovementClean
)

func cleanRoom2(robot Robot) {
	i, j := 0, 0
	var pos [2]int
	visited := map[[2]int]PlaceState{}
	toVisit := [][2]int{}
	// dir := DirectionUp
	// movement := MovementForward

	visited[[2]int{i, j}] = PlaceStateEmpty
	for 0 < len(toVisit) {
		pos = toVisit[0]
		i = pos[0]
		j = pos[1]
		visited[[2]int{i, j}] = PlaceStateEmpty

		toVisit = append(toVisit, [2]int{i, j - 1}, [2]int{i, j + 1}, [2]int{i - 1, j}, [2]int{i + 1, j})
	}
}

func addToVisit(i, j int, d Direction, toVisit *[][2]int, visited *[2]int) {
	// Add TurnLeft, TurnRight, Forward
	switch d {
	case DirectionUp:

	case DirectionDown:
	case DirectionLeft:
	case DirectionRight:
	}

}

func getNewDirection(d Direction, m Movement) Direction {
	switch d {
	case DirectionUp:
		switch m {
		case MovementForward:
			return d
		case MovementTurnLeft:
			return DirectionLeft
		case MovementTurnRight:
			return DirectionRight
		case MovementClean:
			return d
		}
	case DirectionDown:
		switch m {
		case MovementForward:
			return d
		case MovementTurnLeft:
			return DirectionRight
		case MovementTurnRight:
			return DirectionLeft
		case MovementClean:
			return d
		}
	case DirectionLeft:
		switch m {
		case MovementForward:
			return d
		case MovementTurnLeft:
			return DirectionDown
		case MovementTurnRight:
			return DirectionUp
		case MovementClean:
			return d
		}
	case DirectionRight:
		switch m {
		case MovementForward:
			return d
		case MovementTurnLeft:
			return DirectionUp
		case MovementTurnRight:
			return DirectionDown
		case MovementClean:
			return d
		}
	}
	return DirectionUndefined
}
