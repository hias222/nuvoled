# NUVOLED 

## Prepration

### install MQTT

```bash
git clone https://github.com/hias222/nuvoled.git
go get github.com/eclipse/paho.mqtt.golang
```

### MAC

```bash
brew install mosquitto
/usr/local/sbin/mosquitto -c /usr/local/etc/mosquitto/mosquitto.conf
```

## INSTALL and Start

```bash
go run main.go
```

## Test Send UDP

```bash
/usr/local/bin/
go test -timeout 30s -run ^TestUdpMulticast$ swimdata.de/nuvoled/udpclient -count=1
```


### Docker

ToDo
