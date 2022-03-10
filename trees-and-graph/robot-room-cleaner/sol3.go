package robotRoomCleaner

// type Robot interface {
// 	move() bool
// 	turnLeft()
// 	turnRight()
// 	clean()
// }

// directions : 0: 'up', 1: 'right', 2: 'down', 3: 'left'
var newDirRight = map[int]int{
	0: 1,
	1: 2,
	2: 3,
	3: 0,
}

func cleanRoom3(robot *Robot) {
	visited := map[[2]int]bool{}
	backtrack4(robot, visited, [2]int{0, 0}, 0)
}

func backtrack4(robot *Robot, visited map[[2]int]bool, pos [2]int, dir int) {
	robot.Clean()
	visited[pos] = true
	for i := 0; i < 4; i++ {
		// get next position if we move in the current direction
		newPos := getNewPos(pos, dir)
		if !visited[newPos] && robot.Move() {
			// try to
			backtrack4(robot, visited, newPos, dir)
			goBack(robot)
		}
		robot.TurnRight()
		dir = newDirRight[dir]
	}
}

func getNewPos(pos [2]int, dir int) [2]int {
	switch dir {
	case 0:
		return [2]int{pos[0], pos[1] + 1}
	case 1:
		return [2]int{pos[0] + 1, pos[1]}
	case 2:
		return [2]int{pos[0], pos[1] - 1}
	case 3:
		return [2]int{pos[0] - 1, pos[1]}
	}
	return [2]int{}
}

// Robot finish in the same direction, as turned 4 times
func goBack4(robot *Robot) {
	// Turn 2 times, poiting to the opposite
	robot.TurnRight()
	robot.TurnRight()
	// Move backwards
	robot.Move()
	// Turn 2 times, to point to the same direction that at the beggining
	robot.TurnRight()
	robot.TurnRight()
}
