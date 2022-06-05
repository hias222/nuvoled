package udpclient_test

import (
	"testing"
	"time"

	"swimdata.de/nuvoled/udpclient"
	"swimdata.de/nuvoled/udpmessages"
)

func TestUdpUnicast(t *testing.T) {

	data := udpmessages.CreateInitMessage()

	if udpclient.SendClientInit(data) != "success" {
		t.Fatal("error unicast\n")
	}

}

func TestUdpMulticast(t *testing.T) {

	//SERVER := "192.168.178.255"
	SERVER := "10.10.10.255"
	//SERVER := "169.254.255.255"

	var testdata = "fail"

	data := udpmessages.CreateInitMessage()

	for i := 1; i <= 10; i++ {
		testdata = udpclient.SendClientBCInit(data, SERVER)
		time.Sleep(1 * time.Second)
	}

	if testdata != "success" {
		t.Fatal("error multicast \n", testdata, "\n")
	}

}

/*
GUI Commands

Panel
36 36 15 0 74 23 49 80 52 83 32 8 8 8 8

36 36 120 2 32 8 8 1 23 49 74 8 8 0 0

Auswahl
36 36 160 23 49 74

Reset+Refresh
36 36 160 0 0 0
36 36 130 0

Datenpacket
Antwort Klick auf Panel

36 36 20 2 10 0 0 0 35 45 \ 19 2 0 23 2 0 22 2 0 21 2 0 22 2 0 22 2 0 22 2 0 22 2 0 22 2 0 23 2 0 24 2 0 24 3 0 24 3 0 27 3 0 28 3 1 26 3 0 27 3 0 28 3 0 29 3 1 30 4 1 30 4 1 30 4 1 32 4 1 34 5 1 37 5 1 38 5 1 39 6 1 44 6 1 47 6 1 64 11 2 38 12 1 28 8 2 28 8 2 28 8 2 29 9 2 30 9 2 30 9 2 31 10 2 31 10 2 32 10 2 32 10 3 33 11 3 34 11 3 35 12 3 36 13 3 37 13 3 39 13 3 40 14 3 41 14 4 44 15 4 44 16 4 45 17 4 48 18 5 49 19 5 53 20 5 55 21 5 58 23 6 60 24 7 63 27 7 67 30 9 167 84 14 0 15 4 0 17 5 0 17 5 0 18 5 0 19 5 0 19 5 0 20 5 0 21 6 0 20 5 0 21 6 0 21 6 0 23 6 0 24 6 0 24 6 0 25 6 0 26 6 0 27 7 1 28 7 1 30 8 1 30 8 1 32 9 1 33 8 1 35 9 1 35 9 1 38 10 1 39 10 1 40 11 1 42 12 1 44 12 1 47 11 0 44 9 0 69 20 0 36 13 1 38 16 1 39 16 1 39 16 1 39 16 1 41 17 1 44 19 1 45 18 1 46 19 1 48 19 1 49 20 1 50 20 1 52 21 1 53 22 2 55 23 1 57 24 1 59 24 1 60 26 2 61 26 2 62 27 2 65 28 2 66 29 2 69 31 2 70 32 2 72 32 3 74 35 1 79 37 5 141 93 4 126 94 4 132 93 4 129 97 3 134 99 3 137 102 4 139 104 4 141 105 84 1 0 15 2 0 19 2 0 21 2 0 21 2 0 21 2 0 21 2 0 21 2 0 21 2 0 22 2 0 23 2 0 24 2 0 24 2 0 26 2 0 26 3 0 26 3 0 26 3 0 27 3 0 28 3 0 28 3 0 29 3 0 30 4 1 31 4 1 32 5 1 35 5 1 37 5 1 38 5 1 43 6 1 44 6 1 98 19 4 29 9 1 27 8 2 28 8 2 28 8 2 28 8 2 29 9 2 29 9 2 30 9 2 31 10 2 31 10 2 32 10 2 32 10 2 32 10 2 34 11 3 35 12 3 36 13 3 38 13 3 39 13 3 39 14 3 43 15 4 44 15 4 44 17 4 47 18 5 48 18 5 52 20 5 54 21 5 57 22 6 59 24 6 62 26 6 65 28 9 199 131 38 0 14 3 0 16 4 0 16 4 0 18 4 0 18 5 0 19 5 0 20 5 0 20 5 0 20 5 0 21 6 0 21 6 0 21 5 0 23 6 0 24 6 0 24 6 0 25 6 0 26 7 1 28 7 1 29 7 1 30 8 1 32 8 0 32 8 1 34 8 1 35 9 1 36 9 1 39 10 1 39 10 1 41 11 1 43 12 1 46 11 1 47 11 9 118 44 0 37 13 1 35 14 1 38 15 1 39 16 1 39 16 1 40 17 1 42 18 1 44 18 1 45 18 1 46 19 1 48 19 1 49 20 1 51 21 1 52 21 1 54 22 1 56 23 1 57 24 1 59 25 1 60 26 2 61 26 2 64 28 2 65 28 2 68 30 2 69 31 2 70 32 2 72 34 2 68 32 2 126 72 1 118 80 3 121 85 4 129 94 3 132 97 3 134 98 3 136 101 4 139 104 108 0 0 21 1 0 21 1 0 20 2 0 21 2 0 21 2 0 21 2 0 21 2 0 21 2 0 22 2 0 23 2 0 23 2 0 23 2 0 23 2 0 24 3 0 26 3 0 26 3 0 27 3 0 28 3 0 28 3 0 28 3 0 29 3 0 30 4 1 30 4 1 32 5 1 35 4 1 37 5 1 42 5 1 40 6 1 164 37 11 20 6 2 27 8 2 27 8 2 28 8 1 28 8 2 28 8 2 28 8 2 29 9 2 30 9 2 30 9 2 31 10 2 32 10 2 32 10 2 32 11 2 33 11 2 35 12 3 35 12 3 36 13 3 37 13 3 41 14 3 44 15 4 44 16 4 45 17 4 47 18 5 51 19 4 53 20 5 56 21 6 58 23 6 60 26 6 56 23 7 144 118 35 0 14 4 0 16 4 0 16 4 0 17 4 0 18 4 0 18 4 0 19 5 0 19 5 0 20 5 0 20 5 0 21 6 0 21 5 0 21 5 0 23 5 0 24 6 0 24 6 0 26 6 0 27 7 1 28 7 0 29 7 0 31 8

grüner hacken

<-- 36 36 130 0
<-- 36 36 160 0 0 0 (server init)
--> 36 36 15 0 74 23 49 80 52 83 32 8 8 8 8
<-- 36 36 120 2 0 8 8 1 23 49 74 8 8 0 0
<-- 36 36 160 23 49 74 (aktiviert)
<-- 36 36 120 2 32 8 8 1 23 49 74 8 8 0 0 (Hacken)
<-- 36 36 155 0 (save config)

*/
