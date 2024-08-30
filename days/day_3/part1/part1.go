package part1

type Visits struct {
	x int
	y int
}

func Houses(s string) int {
	// visit := make(map[[2]int]bool)
	visited := make(map[Visits]bool)
	start := Visits{x: 0, y: 0}
	visited[start] = true

	/*
		> East
		< West
		^ North
		v South
	*/

	for _, house := range s {
		if house == '>' {
			visited[Visits{x: start.x + 1, y: start.y}] = true
			start.x += 1
		} else if house == '<' {
			visited[Visits{x: start.x - 1, y: start.y}] = true
			start.x -= 1
		} else if house == '^' {
			visited[Visits{x: start.x, y: start.y + 1}] = true
			start.y += 1
		} else if house == 'v' {
			visited[Visits{x: start.x, y: start.y - 1}] = true
			start.y -= 1
		}
		visited[start] = true
	}
	return len(visited)
}
