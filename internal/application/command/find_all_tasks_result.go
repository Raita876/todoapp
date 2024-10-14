package command

type FindAllTasksCommand struct {
	ContainsForName string
	FilterStatusId  int
	SortBy          string
	OrderIsAsc      bool
}
