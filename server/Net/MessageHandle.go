package mnet

import (
	"fmt"
	"net"
	"server/face"
)

type MessageHandle struct {
	WorkerPoolSize uint32
	TaskQueue      []chan face.IRequest
}

func (messageHandle *MessageHandle) DoMessageHandler(request face.IRequest) {
	request.GetSession().SendMessage(request.GetMessage())

}
func (messageHandle *MessageHandle) StartWorkerPool() {
	for i := 0; i < int(messageHandle.WorkerPoolSize); i++ {
		messageHandle.TaskQueue[i] = make(chan face.IRequest)
		go messageHandle.StartOneWorker(i, messageHandle.TaskQueue[i])
	}

}
func (messageHandle *MessageHandle) StartOneWorker(workerID int, taskQueue chan face.IRequest) {
	fmt.Println("WorkerId = ", workerID, "  Start")
	for {
		select {
		case request := <-taskQueue:
			messageHandle.DoMessageHandler(request)
		}
	}
}
func (messageHandle *MessageHandle) SendMessageToTaskQueue(request face.IRequest) {
	workerID := request.GetSession().GetSid() % 10
	messageHandle.TaskQueue[workerID] <- request

}

func (messageHandle *MessageHandle) LobbyMessageHandle(conn net.Conn, bytes []byte, length int) {

	// return nil
}

func (messageHandle *MessageHandle) GameMessageHandle(roomId uint32, bytes []byte, length int) {

}