package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func generateGraphFile() {
	cmd := exec.Command("go", "mod", "graph")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error generating go.graph:", err)
		return
	}

	// Write the output to go.graph file
	err = os.WriteFile("go.graph", output, 0644)
	if err != nil {
		fmt.Println("Error writing go.graph file:", err)
		return
	}

	fmt.Println("go.graph file generated successfully.")
}

func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func main() {
	generateGraphFile()
	// Lesen der go.graph Datei
	file, err := os.Open("go.graph")
	if err != nil {
		fmt.Println("Fehler beim Öffnen der Datei:", err)
		return
	}
	defer file.Close()

	nodes := make(map[string]bool)
	var links []opts.GraphLink

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		if len(parts) == 2 {
			source := parts[0]
			target := parts[1]
			nodes[source] = true
			nodes[target] = true
			links = append(links, opts.GraphLink{Source: source, Target: target})
		}
	}

	var graphNodes []opts.GraphNode
	for node := range nodes {
		graphNodes = append(graphNodes, opts.GraphNode{
			Name:       node,
			X:          rand.Float32() * 100,
			Y:          rand.Float32() * 100,
			SymbolSize: 10,
		})
	}

	// Create a new graph instance
	graph := charts.NewGraph()

	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Go Dependencies",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "100%",
			Height: "1200px",
			Theme:  "light",
		}),
	)

	graph.AddSeries("graph", graphNodes, links,
		charts.WithGraphChartOpts(
			opts.GraphChart{
				Layout:             "force",
				Roam:               opts.Bool(true),
				FocusNodeAdjacency: opts.Bool(true),
				Force: &opts.GraphForce{
					Repulsion:  100,
					Gravity:    0.1,
					EdgeLength: 100,
				},
			},
		),
		charts.WithLabelOpts(opts.Label{
			Show:       opts.Bool(true),
			Position:   "right",
			Color:      "black",
			FontSize:   12,
			FontFamily: "Arial, Helvetica, sans-serif",
		}),
		charts.WithItemStyleOpts(opts.ItemStyle{
			BorderColor: "black",
			BorderWidth: 1,
			Color:       "lightblue",
		}),
	)

	// Save as HTML
	htmlFile := "go_dependencies.html"
	f, _ := os.Create(htmlFile)
	graph.Render(f)
	f.Close()

	fmt.Printf("Graph has been saved as %s\n", htmlFile)

	// Open the HTML file in the default browser if not in silent mode
	if !SilentMode {
		err = openBrowser(htmlFile)
		if err != nil {
			fmt.Printf("Error opening %s in browser: %v\n", htmlFile, err)
		} else {
			fmt.Printf("Opened %s in your default browser.\n", htmlFile)
		}
	} else {
		fmt.Println("Silent mode: browser not opened automatically.")
	}
}
