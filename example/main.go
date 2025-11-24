package main

import (
	"encoding/hex"
	"fmt"
	"gptp"
	"gptp/gptpMessage"
)

func main() {

	req, _ := hex.DecodeString("")
	message, err := gptp.Decoder(req)
	if err != nil {
		fmt.Printf("err : %v\n", err)
	}

	switch m := message.(type) {
	case *gptpMessage.SyncMessage:
		HandleSyncMessage(m)
	case *gptpMessage.PeerDelayReqMessage:
		HandlePeerDelayReqMessage(m)
	case *gptpMessage.PeerDelayRespMessage:
		HandlePeerDelayRespMessage(m)
	case *gptpMessage.FollowUpMessage:
		HandleFollowUpMessage(m)
	case *gptpMessage.PeerDelayRespFollowUpMessage:
		HandlePeerDelayRespFollowUpMessage(m)
	case *gptpMessage.AnnounceMessage:
		HandleAnnounceMessage(m)
	}

	result, err := gptp.Encoder(message)
	if err != nil {
		fmt.Printf("err : %v\n", err)
	}
	fmt.Printf("Message : %+v\n", hex.EncodeToString(req))
	fmt.Printf("Message : %+v\n", hex.EncodeToString(result))

}

func HandleSyncMessage(msg *gptpMessage.SyncMessage) {
	fmt.Printf("Sync Message : %+v\n", msg)
}

func HandlePeerDelayReqMessage(msg *gptpMessage.PeerDelayReqMessage) {
	fmt.Printf("Peer Delay Req Message : %+v\n", msg)
}

func HandlePeerDelayRespMessage(msg *gptpMessage.PeerDelayRespMessage) {
	fmt.Printf("Peer Delay Resp Message : %+v\n", msg)
}

func HandleFollowUpMessage(msg *gptpMessage.FollowUpMessage) {
	fmt.Printf("Follow Up Message : %+v\n", msg)
}

func HandlePeerDelayRespFollowUpMessage(msg *gptpMessage.PeerDelayRespFollowUpMessage) {
	fmt.Printf("Peer Delay Resp Follow Up Message : %+v\n", msg)
}

func HandleAnnounceMessage(msg *gptpMessage.AnnounceMessage) {
	fmt.Printf("Announce Message : %+v\n", msg)
}
