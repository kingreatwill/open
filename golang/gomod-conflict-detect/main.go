package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var report struct {
	rootNode *node
	// pkgname+ver -> node
	nodeMap map[string]*node

	// pkgname -> node array
	pkgnameMap map[string][]*node

	conflictPaths map[string][][]string
}

func init() {
	report.nodeMap = make(map[string]*node)
	report.pkgnameMap = make(map[string][]*node)

	report.conflictPaths = make(map[string][][]string)
}

// node represent a unique lib appeared in go mod graph
type node struct {
	pkgNameWithVersion string // package name with version
	parent             []*node
	children           []*node
	path               []string
}

func recordPath(n *node, parentPath []string) {
	if n == nil {
		return
	}

	curPath := make([]string, len(parentPath))
	copy(curPath, parentPath)

	curPath = append(curPath, n.pkgNameWithVersion)
	n.path = curPath

	for _, subNode := range n.children {
		recordPath(subNode, curPath)
	}
}

func recordPkg2Nodes(n *node) {
	if n == nil {
		return
	}
	nameAndVer := strings.Split(n.pkgNameWithVersion, "@")
	var name = nameAndVer[0]
	report.pkgnameMap[name] = append(report.pkgnameMap[name], n)

	for _, subNode := range n.children {
		recordPkg2Nodes(subNode)
	}
}

func findConflict() {
	conflictInSomePkg := false
	for pkgname, nodes := range report.pkgnameMap {
		if len(nodes) > 1 {
			// may have conflict
			var hasConflict = false
			for i := 1; i < len(nodes); i++ {
				if nodes[i].pkgNameWithVersion != nodes[i-1].pkgNameWithVersion {
					hasConflict = true
					conflictInSomePkg = true
					break
				}
			}
			if hasConflict {
				fmt.Println("Conflict in pkg", pkgname, "paths are: ")
				for _, n := range nodes {
					fmt.Println("\r", strings.Join(n.path, " -> "))
					report.conflictPaths[pkgname] = append(report.conflictPaths[pkgname], n.path)
				}
			}
		}
	}
	if !conflictInSomePkg {
		fmt.Println("there is no conflict in your project dependencies")
	}
}

func readInput() string {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | gocowsay")
		return ""
	}

	reader := bufio.NewReader(os.Stdin)
	var output []byte

	for {
		input, err := reader.ReadByte()
		if err != nil && err == io.EOF {
			break
		}
		output = append(output, input)
	}

	for j := 0; j < len(output); j++ {
		fmt.Printf("%c", output[j])
	}
	return string(output)
}

// go mod graph | this lib
func main() {
	var lines = strings.Split(readInput(), "\n")
	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		pkgDep := strings.Split(line, " ")
		if len(pkgDep) != 2 {
			fmt.Println("wrong format", pkgDep)
			return
		}

		// github.com/cch123@v1.0.1    github.com/google@v2.3.4
		pkg, dep := pkgDep[0], pkgDep[1]
		pkgNode := &node{pkgNameWithVersion: pkg}
		depNode := &node{pkgNameWithVersion: dep}
		if _, ok := report.nodeMap[pkg]; !ok {
			report.nodeMap[pkg] = pkgNode
		}

		report.nodeMap[pkg].children = append(report.nodeMap[pkg].children, depNode)

		if _, ok := report.nodeMap[dep]; !ok {
			report.nodeMap[dep] = depNode
		}
		report.nodeMap[dep].parent = append(report.nodeMap[dep].parent, pkgNode)

		if report.rootNode == nil {
			// root pkg
			report.rootNode = pkgNode
		}
	}

	recordPath(report.rootNode, []string{})
	recordPkg2Nodes(report.rootNode)
	findConflict()
}
