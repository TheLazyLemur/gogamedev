package main

import (
	"gamedev/internal/graph"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func isPointInRectangle(point rl.Vector2, rec rl.Rectangle) bool {
	return point.X >= rec.X && point.X <= rec.X+rec.Width &&
		point.Y >= rec.Y && point.Y <= rec.Y+rec.Height
}

func main() {
	nodeA := graph.CreateNode("NodeA", rl.Vector2{X: 0, Y: 0}, rl.Vector2{X: 100, Y: 100})
	nodeB := graph.CreateNode("NodeB", rl.Vector2{X: 1, Y: 0}, rl.Vector2{X: 100, Y: 100})
	nodeC := graph.CreateNode("NodeC", rl.Vector2{X: 2, Y: 0}, rl.Vector2{X: 100, Y: 100})
	nodeD := graph.CreateNode("NodeD", rl.Vector2{X: 0, Y: 1}, rl.Vector2{X: 100, Y: 100})
	nodeE := graph.CreateNode("NodeE", rl.Vector2{X: 1, Y: 1}, rl.Vector2{X: 100, Y: 100})
	nodeF := graph.CreateNode("NodeF", rl.Vector2{X: 2, Y: 1}, rl.Vector2{X: 100, Y: 100})
	nodeG := graph.CreateNode("NodeG", rl.Vector2{X: 0, Y: 2}, rl.Vector2{X: 100, Y: 100})

	g := &graph.Graph{}
	graph.AddNode(g, nodeA)
	graph.AddNode(g, nodeB)
	graph.AddNode(g, nodeC)
	graph.AddNode(g, nodeD)
	graph.AddNode(g, nodeE)
	graph.AddNode(g, nodeF)
	graph.AddNode(g, nodeG)

	graph.GetNeighbors(nodeA, g, 3, 2)
	graph.GetNeighbors(nodeB, g, 3, 2)
	graph.GetNeighbors(nodeC, g, 3, 2)
	graph.GetNeighbors(nodeD, g, 3, 2)
	graph.GetNeighbors(nodeE, g, 3, 2)
	graph.GetNeighbors(nodeF, g, 3, 2)
	graph.GetNeighbors(nodeG, g, 3, 2)

	graph.PrintGraph(g)
	// return

	rl.InitWindow(800, 450, "AStar")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		selected := ""
		for i, n := range graph.GetNodes(g) {
			rec := rl.Rectangle{
				X:      graph.GetNodeSize(n).X * graph.GetNodePos(n).X,
				Y:      graph.GetNodeSize(n).Y * graph.GetNodePos(n).Y,
				Width:  graph.GetNodeSize(n).X,
				Height: graph.GetNodeSize(n).Y,
			}
			if i%2 == 0 {
				rl.DrawRectangleRec(rec, rl.Green)
			} else {
				rl.DrawRectangleRec(rec, rl.Red)
			}

			mousePos := rl.GetMousePosition()
			if isPointInRectangle(mousePos, rec) {
				selected = graph.GetName(n)
			}
		}

		for _, n := range graph.GetNodes(g) {
			if graph.GetName(n) == selected {
				ncx := graph.GetNodePos(n).X + graph.GetNodePos(n).X*100 + 100/2
				ncy := graph.GetNodeSize(n).Y*graph.GetNodePos(n).Y + graph.GetNodeSize(n).Y/2

				rl.DrawRectangle(int32(ncx), int32(ncy), 10, 10, rl.Blue)

				for _, nn := range graph.GetEdges(n) {
					nncx := graph.GetNodePos(nn).X + graph.GetNodePos(nn).X*100 + 100/2
					nncy := graph.GetNodeSize(nn).Y*graph.GetNodePos(nn).Y + graph.GetNodeSize(nn).Y/2
					rl.DrawRectangle(int32(nncx), int32(nncy), 10, 10, rl.Yellow)

					rl.DrawLine(int32(ncx), int32(ncy), int32(nncx), int32(nncy), rl.Purple)
				}
			}
		}

		rl.EndDrawing()
	}
}
