package test

import (
	"fmt"
	"testing"

	"github.com/budimanlai/go-traccar"
)

var (
	tObject   *traccar.Traccar
	device_id []string
)

func TestMain(m *testing.M) {
	url := "http://tcmolis.stbku.com:8082/api"
	token := "c6HyvvE3eThCt3OHK7eJW9JTWRANExD4"

	tObject = traccar.NewTraccar(url, "admin", "adminjimatt", token)
	device_id = []string{
		"216",
		"40",
	}
	m.Run()
}

func TestTrips(t *testing.T) {
	from := "2022-09-23 00:00:00"
	to := "2022-09-23 23:59:59"

	resp, e := tObject.Trips(device_id, from, to, 1, 0, 25)
	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println("Count: ", len(resp))
	}
}

func TestSession(t *testing.T) {

	resp, e := tObject.GetSession()
	if e != nil {
		fmt.Println(e.Error())
	}

	fmt.Println("Cookie: ", resp)
}
