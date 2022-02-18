package robotRoomCleaner

var directions = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func cleanRoom(robot *Robot) {
	// going clockwise : 0: 'up', 1: 'right', 2: 'down', 3: 'left'
	visited := map[[2]int]bool{}
	backtrack(robot, visited, 0, 0, 0)
}

func backtrack(robot *Robot, visited map[[2]int]bool, row, col, d int) {
	visited[[2]int{row, col}] = true
	robot.Clean()
	// going clockwise : 0: 'up', 1: 'right', 2: 'down', 3: 'left'
	for i := 0; i < 4; i++ {
		newD := (d + i) % 4
		newRow := row + directions[newD][0]
		newCol := col + directions[newD][1]

		if !visited[[2]int{newRow, newCol}] && robot.Move() {
			backtrack(robot, visited, newRow, newCol, newD)
			goBack(robot)
		}
		// turn the robot following chosen direction : clockwise
		robot.TurnRight()
	}
}

func goBack(robot *Robot) {
	robot.TurnRight()
	robot.TurnRight()
	robot.Move()
	robot.TurnRight()
	robot.TurnRight()
}
