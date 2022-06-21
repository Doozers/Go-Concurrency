package main

func searchDefinition(data *mockData) []*mockData {
	return data.children
}

func findDefinition(data *mockData) string {
	return data.value
}
