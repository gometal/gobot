// +build example
//
// Do not build by default.

/*
 To run this example, pass the BLE addresses or BLE names as first param:

 go run examples/ble_multiple.go 2B-1234 2B-5678

 NOTE: sudo is required to use BLE in Linux
*/

package main

import (
	"fmt"
	"os"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/api"
	"gobot.io/x/gobot/platforms/ble"
)

func NewSwarmBot(port string) *gobot.Robot {
	bleAdaptor := ble.NewClientAdaptor(port)
	info := ble.NewDeviceInformationDriver(bleAdaptor)

	work := func() {
		fmt.Println("Model number:", info.GetModelNumber())
		fmt.Println("Firmware rev:", info.GetFirmwareRevision())
		fmt.Println("Hardware rev:", info.GetHardwareRevision())
		fmt.Println("Manufacturer name:", info.GetManufacturerName())
		fmt.Println("PnPId:", info.GetPnPId())
	}

	robot := gobot.NewRobot("bot "+port,
		[]gobot.Connection{bleAdaptor},
		[]gobot.Device{info},
		work,
	)

	return robot
}

func main() {
	master := gobot.NewMaster()
	api.NewAPI(master).Start()

	for _, port := range os.Args[1:] {
		bot := NewSwarmBot(port)
		master.AddRobot(bot)
	}

	master.Start()
}
