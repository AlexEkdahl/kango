package app

import (
	"github.com/AlexEkdahl/kango/internal/server/service"
	desc "github.com/AlexEkdahl/kango/pkg/contract"
)

type KangoBoardServer struct {
	desc.UnimplementedKanbanServer
	taskService service.TaskService
}

func NewMicroservice(
	taskService service.TaskService,
) *KangoBoardServer {
	return &KangoBoardServer{
		taskService: taskService,
	}
}
