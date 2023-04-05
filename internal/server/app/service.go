package app

import (
	"github.com/AlexEkdahl/kango/internal/server/service"
	desc "github.com/AlexEkdahl/kango/pkg/contract"
)

type MicroserviceServer struct {
	desc.UnimplementedKanbanServer
	taskService service.TaskService
}

func NewMicroservice(
	taskService service.TaskService,
) *MicroserviceServer {
	return &MicroserviceServer{
		taskService: taskService,
	}
}
