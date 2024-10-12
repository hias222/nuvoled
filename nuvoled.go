package main

import (
	"flag"
	"os"

	"swimdata.de/nuvoled/logging"
	"swimdata.de/nuvoled/mqttclient"
	"swimdata.de/nuvoled/sendclock"
	"swimdata.de/nuvoled/traffic"
	"swimdata.de/nuvoled/udpmessages"
	"swimdata.de/nuvoled/udpserver"
)

func main() {

	bcPtr := flag.String("bc", "169.254.255.255", "broadcast address")
	ipPtr := flag.Bool("ip", true, "local ip address")
	regPtr := flag.Bool("reg", false, "broadcast address")
	mqttSrv := flag.String("mqtt", "localhost", "mqtt server name")

	var recvMode = flag.Bool("recv", false, "receive check")

	flag.Parse()

	var logger = logging.GetLogger()

	sendclock.InitParam()

	/*
		appEnv := os.Getenv("APP_ENV")

		opts := &slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelDebug,
		}

		var handler slog.Handler = slog.NewTextHandler(os.Stdout, opts)

		if appEnv == "production" {
			opts := &slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelInfo,
			}
			handler = slog.NewJSONHandler(os.Stdout, opts)
		}

		logger := slog.New(handler)

		logger.Info("environem set APP_ENV " + appEnv)
	*/
	// base parameter
	udpmessages.SetParameter(false, false, false)

	if *recvMode {
		logger.Info("Receive Mode")
		udpmessages.SetParameter(false, false, true)

		traffic.Read(true)
		os.Exit(0)
	}

	logger.Info("Normal Mode")

	//Start
	logger.Info(udpserver.StartMessage(*bcPtr, *regPtr, *ipPtr, *mqttSrv))

	var locallistenAddr = ""
	if *ipPtr {
		selfaddress := udpserver.GetIpaddress()

		if selfaddress == "" {
			logger.Debug("no self assigned")
			selfaddress = udpserver.GetLocaladdress()
			if selfaddress == "" {
				logger.Debug("nothing found (192 vs 169")
			}
		}
		logger.Info(selfaddress)
		locallistenAddr = selfaddress
	}

	// start MQTT

	for i := 0; i < 30; i++ {
		logger.Info("Init Mqtt ... ")
		var c, err = mqttclient.IntClientMqtt(*mqttSrv)
		if err != nil {
			logger.Warn(err.Error())
			logger.Warn("Waiting 30s")
		} else {
			logger.Info("Waiting on mqtt ... ")
			mqttclient.StartCLientMqtt(c)
			logger.Info("MQTT connected ")
			break
		}
	}

	//UDP
	//go mqtttoudpclient.SendUDPStartMessage()

	go mqttclient.StartClock()

	logger.Info("Starting UDP")
	udpserver.InitLocalUdpAdress(*bcPtr, *regPtr, locallistenAddr, *logger)
	udpserver.StartServer()

}
