package room

import (
	"fmt"
	"server/face"
	"server/pb"
	"time"
)

// import "server/face"

type RoomStateConfirm struct {
	// *RoomStateBase
	room       face.IRoom
	confirmArr []bool
}

func (state *RoomStateConfirm) Enter() {
	for i := 0; i < state.room.GetRoomPlayerCount(); i++ {
		state.confirmArr[i] = false
	}

	state.room.ChangePlayersRoomId()
	mes := pb.MakeRoomConfirmMessage()
	//向玩家发送进入房间命令，和房间成员信息

	state.room.Broadcast(mes)
	state.room.SendIndex()


	state.CheckTimeLimit()

}
func (state *RoomStateConfirm) Dismiss() {

	mes := pb.MakeRoomDismissMes()
	state.room.Broadcast(mes)
	state.room.Delete()

}
func (state *RoomStateConfirm) Exit() {

}

func (state *RoomStateConfirm) Update(sid uint32, mes *pb.PbMessage) {
	index := state.room.GetPlayerIndex(sid)
	state.confirmArr[index] = true

	if state.CheckAllConfirm() {

		state.room.ChangeRoomState(int(roomStateSelect))
	}
}

func (state *RoomStateConfirm) CheckTimeLimit() {

	select {
	case <-time.After(time.Second * 999):
		fmt.Println("room confirm reachtime ")
		state.Dismiss()
		return
	}

}

func (state *RoomStateConfirm) CheckAllConfirm() bool {

	for _, i := range state.confirmArr {
		if !i {
			return false
		}
	}
	fmt.Println("全部确认")
	return true

}
