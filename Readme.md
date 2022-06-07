# NUVOLED 

## INSTALL and Start

### install MQTT

```bash
git clone https://github.com/hias222/nuvoled.git
go get github.com/eclipse/paho.mqtt.golang
```

```bash
go run main.go
```

## Test Send UDP

```bash
/usr/local/bin/
go test -timeout 30s -run ^TestUdpMulticast$ swimdata.de/nuvoled/udpclient -count=1
```

## start MQTT local

### MAC

```bash
brew install mosquitto
/usr/local/sbin/mosquitto -c /usr/local/etc/mosquitto/mosquitto.conf
```

### Docker

ToDo
