package jcba

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/sacOO7/gowebsocket"
)

type jcbaInfo struct {
	Code     int    `json:"code"`
	Location string `json:"location"`
	Token    string `json:"token"`
}

func Main(station string, durationSec int) {
	log.Println("station:", station)
	log.Println("sleep:", durationSec)
	// token, location取得
	info, err := getInfo(station)
	if err != nil {
		log.Println(err)
		return
	}

	//ws開始
	wsReciever(info, durationSec)
	return
}

func getInfo(station string) (jcbaInfo, error) {
	var info jcbaInfo

	url := "https://api.radimo.smen.biz/api/v1/select_stream?station=" + station + "&channel=0&quality=high&burst=5"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Origin", "https://www.jcbasimul.com")

	client := new(http.Client)
	resp, err := client.Do(req)

	if err != nil {
		return info, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return info, err
	}
	if err := json.Unmarshal(body, &info); err != nil {
		return info, err
	}
	return info, nil

}

func wsReciever(info jcbaInfo, durationSec int) error {

	time.AfterFunc(time.Duration(durationSec)*time.Second, func() {
		log.Println("Stream finish")
		os.Exit(0)
	})

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	socket := gowebsocket.New(info.Location)

	//socket.RequestHeader.Set("user-agent", UserAgent)
	socket.RequestHeader.Set("sec-websocket-protocol", "listener.fmplapla.com")

	socket.OnConnected = func(socket gowebsocket.Socket) {
		log.Println("Connected to server")
	}

	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		log.Println("Received connect error ", err)
	}

	socket.OnBinaryMessage = func(data []byte, socket gowebsocket.Socket) {
		os.Stdout.Write(data)
	}

	socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		log.Println("Disconnected from server ")
		return
	}

	socket.Connect()

	socket.SendText(info.Token)

	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			socket.Close()
			return nil
		}
	}

	return nil

}
