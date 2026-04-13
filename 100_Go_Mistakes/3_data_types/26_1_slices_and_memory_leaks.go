// Use a slice copy to avoid potential high memory consumption
func consumeMessages() {
	for {
		msg := receiveMessage()
		// Do something with msg
		storeMeessage(getMessageType(msg)) // the original capacity of the msg is stored, even though we are only returning a length of 5 
	}
}

func getMessageType(msg []byte) []byte {
	return msg[:5]
}

// correction 
func getMessageType(msg []byte) []byte {
	msgType := make([]byte, 5)
	copy(msgType, msg)
	return msgType
}



