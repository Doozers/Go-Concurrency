package main

import "sync"

type node struct {
	value    string
	children []*node
}
type mockData struct {
	value    string
	children []*mockData
}

func getDefinition(oldNode *node, definitions *mockData, wait *sync.WaitGroup, mutex *sync.Mutex) {
	defNode := &node{}

	mutex.Lock()
	oldNode.children = append(oldNode.children, defNode)
	mutex.Unlock()

	defNode.value = findDefinition(definitions)
	NewDefinitions := searchDefinition(definitions)

	wait.Add(len(NewDefinitions))
	addDefinitionToDag(defNode, NewDefinitions, wait, mutex)
	wait.Done()
}

func addDefinitionToDag(oldNode *node, definitions []*mockData, wait *sync.WaitGroup, mutex *sync.Mutex) {
	for _, definition := range definitions {
		go getDefinition(oldNode, definition, wait, mutex)
	}
}

func createDag(data *mockData) *node {
	root := &node{"root", nil}
	definitions := searchDefinition(data)

	mutex := sync.Mutex{}

	wait := sync.WaitGroup{}
	wait.Add(len(definitions))

	addDefinitionToDag(root, definitions, &wait, &mutex)
	wait.Wait()
	return root
}

func printDag(root *node, depth int) {
	for i := 0; i < depth; i++ {
		print("  ")
	}

	println(root.value)

	for _, child := range root.children {
		printDag(child, depth+1)
	}
}

func main() {
	data := &mockData{
		value: "root",
		children: []*mockData{
			{
				value: "child1",
				children: []*mockData{
					{value: "child1.1"},
					{value: "child1.2"},
				},
			},
			{value: "child2",
				children: []*mockData{
					{value: "child2.1"},
					{value: "child2.2"},
				},
			},
			{value: "child3",
				children: []*mockData{
					{value: "child3.1"},
					{value: "child3.2"},
					{value: "child3.3"},
				},
			},
		},
	}
	dag := createDag(data)
	printDag(dag, 0)
}
