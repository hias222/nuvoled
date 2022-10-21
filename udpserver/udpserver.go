package udpserver

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"swimdata.de/nuvoled/udpmessages"
)

var udpSource *net.UDPAddr
var udpDestination *net.UDPAddr
var listenConnection *net.UDPConn
var working bool
var register bool

func StartMessage(bc string, reg bool, ip bool, mqttSrv string) string {
	var Message = "Start laufanzeiger .."

	if reg {
		Message = Message + "\n # Registration of panles is on "
	} else {
		Message = Message + "\n # Registration is off "
	}

	Message = Message + "\n # udp send to " + bc
	Message = Message + "\n # local ip (empty listen an all) " + strconv.FormatBool(ip)
	Message = Message + "\n # usage: ./laufanzeiger -bc 169.254.255.255 -reg=false/true -ip=true/false -mqtt localhost "
	Message = Message + "\n # usage: ./laufanzeiger -recv=false/true "
	Message = Message + "\n\n # to work with multiple interfaces use local IP - reg is not possible with IP " + strconv.FormatBool(ip)
	Message = Message + "\n"
	return Message
}

func InitLocalUdpAdress(broadcast string, reg bool, localip string) {
	PORT := localip + ":2000"

	working = false
	register = reg

	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
	}

	udpSource = s

	fmt.Println("Local Listener Address: ", s.String())

	//SENDERPORT := "169.254.255.255:2000"
	SENDERPORT := broadcast + ":2000"

	sender, err := net.ResolveUDPAddr("udp4", SENDERPORT)
	if err != nil {
		fmt.Println(err)
	}
	udpDestination = sender
	fmt.Println("UDP Detination Address: ", sender.String())

}

// ?????

func SendUDPMessage(data []byte) {
	c, err := net.DialUDP("udp4", nil, udpDestination)
	if err != nil {
		fmt.Println(err)
	}
	defer c.Close()
	_, err = c.Write(data)
}

func SendUDPListenMessage(data []byte) {

	_, error := listenConnection.WriteToUDP(data, udpDestination)

	if error != nil {
		fmt.Println("Error sending ")
	}

}

func handleBufferData(buffer []byte, n int, addr net.Addr, sendaction bool) {
	udpmessages.BufferToString(buffer, n)
	if working {
		return
	}

	working = true

	if strings.TrimSpace((string(buffer[0 : n-1]))) == "STOP" {
		fmt.Println("Exiting UDP server")
		working = false
		return
	}

	if n > 3 && buffer[2] == 15 && sendaction {
		fmt.Println("Send Messages to panel ")
		fmt.Print("-> ", string(addr.String()), "\n")
		if register {
			SendUDPMessage(udpmessages.ResetPanles())
			time.Sleep(2 * time.Second)
		}

		SendUDPMessage(udpmessages.RefreshPanles())
		time.Sleep(1 * time.Second)

		if register {
			SendUDPMessage(udpmessages.CreateRegisterMessage(buffer))
			time.Sleep(1 * time.Second)
			SendUDPMessage(udpmessages.ActivatePanles(buffer))
			time.Sleep(1 * time.Second)
			SendUDPMessage(udpmessages.SavePanles(buffer))
			time.Sleep(1 * time.Second)
			fmt.Println("Finish Registration")
			time.Sleep(1 * time.Second)
		}
		//SendUDPMessage(udpmessages.TurnOnPanles(buffer))
		//time.Sleep(1 * time.Second)

	}
	working = false
}

func StartServer() {

	udpSource2, err := net.ResolveUDPAddr("udp", ":2000")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	Connection, err := net.ListenUDP("udp4", udpSource2)
	fmt.Println(udpSource2)

	listenConnection = Connection
	listenConnection.SetReadBuffer(1048576)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	//defer listenConnection.Close()
	buffer := make([]byte, 2048)
	rand.Seed(time.Now().Unix())

	fmt.Println("Start listening")

	for {
		n, addr, err := listenConnection.ReadFromUDP(buffer)

		if err != nil {
			fmt.Println(err)
			return
		}

		if listenConnection == nil {
			fmt.Println("broken")
		}

		go handleBufferData(buffer, n, addr, false)
	}
}
