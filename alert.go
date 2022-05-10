package main

import (
	zmq "github.com/pebbe/zmq4"
)

func appAlert(message string) {
	if ZMQAlert {
		//  Socket to talk to server
		requester, _ := zmq.NewSocket(zmq.REQ)
		defer requester.Close()
		requester.Connect("tcp://localhost:5555")
		requester.Send(message, 0)
		_, _ = requester.Recv(0)
	}

}
