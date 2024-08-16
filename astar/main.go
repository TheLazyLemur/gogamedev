package main

import (
	"gamedev/internal/graph"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	nodeA := graph.CreateNode("NodeA", rl.Vector2{X: 0, Y: 0})
	nodeB := graph.CreateNode("NodeB", rl.Vector2{X: 1, Y: 0})
	nodeC := graph.CreateNode("NodeC", rl.Vector2{X: 2, Y: 0})
	nodeD := graph.CreateNode("NodeD", rl.Vector2{X: 0, Y: 1})
	nodeE := graph.CreateNode("NodeE", rl.Vector2{X: 1, Y: 1})
	nodeF := graph.CreateNode("NodeF", rl.Vector2{X: 2, Y: 1})

	g := &graph.Graph{}
	graph.AddNode(g, nodeA)
	graph.AddNode(g, nodeB)
	graph.AddNode(g, nodeC)
	graph.AddNode(g, nodeD)
	graph.AddNode(g, nodeE)
	graph.AddNode(g, nodeF)

	graph.GetNeighbors(nodeA, g, 3, 2)
	graph.GetNeighbors(nodeB, g, 3, 2)
	graph.GetNeighbors(nodeC, g, 3, 2)
	graph.GetNeighbors(nodeD, g, 3, 2)
	graph.GetNeighbors(nodeE, g, 3, 2)
	graph.GetNeighbors(nodeF, g, 3, 2)

	graph.PrintGraph(g)

	rl.InitWindow(800, 450, "AStar")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		for i, n := range graph.GetNodes(g) {
			rec := rl.Rectangle{
				X:      200 * graph.GetNodePos(n).X,
				Y:      200 * graph.GetNodePos(n).Y,
				Width:  200,
				Height: 200,
			}
			if i%2 == 0 {
				rl.DrawRectangleRec(rec, rl.Green)
			} else {
				rl.DrawRectangleRec(rec, rl.Red)
			}
		}

		rl.EndDrawing()
	}
}
