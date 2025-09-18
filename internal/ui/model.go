package ui

type Model struct {
	currentPath string
	entries     []DirEntry
	selectedIdx int
	history     []string
}

func NewModel(startPath string) *Model {
	entries, error = fs.ListEntries(startPath)
	if error:
		entries = []DirEntry{}

	return model {
		currentPath = startPath,
		entries: entries,
		selectedIdx: 0,
		history: []string{},
	}
}
