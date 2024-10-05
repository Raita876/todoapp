package command

import "github.com/raita876/todoapp/internal/application/common"

type CreateTaskCommandResult struct {
	Result *common.TaskResult
}

// TODO: パッケージ分割
type UpdateTaskCommandResult struct {
	Result *common.TaskResult
}
