package core

import (
	"fmt"
	"github.com/gorilla/websocket"
	"ninety/model"
)

// 处理信息函数
func HandleMes(clientMes model.Message, conn *websocket.Conn) {
	fmt.Println("HandleMes...")
	//根据mesType进入下一步流程
	switch clientMes.MesType {
	case model.SIGNAL_TYPE_JOIN:
		fmt.Println("MesType-join...")
		//调用userJoin函数
		UserJoin(clientMes, conn)
	case model.SIGNAL_TYPE_CANCEL:
		fmt.Println("MesType-cancel...")
		//调用userCancel函数
		UserCancel(clientMes)
	case model.SIGNAL_TYPE_LEAVE:
		fmt.Println("MesType-leave...")
		//调用userLeave函数
		//userLeave(mes, conn)
	case model.SIGNAL_TYPE_OFFER:
		fmt.Println("MesType-offer...")
		//调用tansoffer函数
		TransOffer(clientMes, conn)
	case model.SIGNAL_TYPE_ANSWER:
		fmt.Println("MesType-answer...")
		//调用tansanswer函数
		TransAnswer(clientMes, conn)
	case model.SIGNAL_TYPE_CANDIDATE:
		fmt.Println("MesType-candidate...")
		//调用tanscandidate函数
		TransCandidate(clientMes, conn)
	default:
		fmt.Println("MesType-unknow...")
		return
	}
}
