package main

type ActionsStruct struct {
	InsertionActions []string
	DeletionActions  []string
	CombinedActions  []string
}

type ActionArgs struct {
	Name           string
	InsertionCount int
	DeletionCount  int
	SI             string
	SD             string
}

var Actions = ActionsStruct{
	InsertionActions: []string{
		"{{.Name}} baked {{.InsertionCount}} muffin{{.SI}}",
		"{{.Name}} baked {{.InsertionCount}} cookie{{.SI}}",
		"{{.Name}} built a sandcastle with {{.InsertionCount}} tower{{.SI}} in it",
	},
	DeletionActions: []string{
		"{{.Name}} ate {{.DeletionCount}} muffin{{.SD}}",
		"{{.Name}} ate {{.DeletionCount}} cookie{{.SD}}",
		"{{.Name}} kicked {{.DeletionCount}} traffic cone{{.SD}}",
	},
	CombinedActions: []string{
		"{{.Name}} baked {{.InsertionCount}} muffin{{.SI}} and ate {{.DeletionCount}}{{ if neq .InsertionCount 0 }} of them{{end}}",
		"{{.Name}} baked {{.InsertionCount}} cookie{{.SI}} and ate {{.DeletionCount}}{{ if neq .InsertionCount 0 }} of them{{end}}",
	},
}