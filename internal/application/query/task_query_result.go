package query

import "github.com/raita876/todoapp/internal/application/common"

type TaskQueryResult struct {
	Result *common.TaskResult
}

type TaskQueryListResult struct {
	Result []*common.TaskResult
}
