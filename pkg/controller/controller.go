package controller

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"github.com/pspiagicw/gosignal/pkg/battery"
)

const (
	LOW_STATUS = 1
	VERY_LOW_STATUS = 2
	CONNECTED_STATUS = 3
	FULL_STATUS = 4
)


type Provider interface {
	GetBatteryCharge() int
	GetBatteryStatus() bool
}
type Sleeper interface {
	Sleep(t time.Duration)
}
type Notifier interface {
	Notify(header string , body string , icon string , duration int)
}

type DefaultProvider struct {
}
func (d *DefaultProvider) GetBatteryCharge() int {
	charge := strings.Trim(battery.GetBatteryCharge() , "\n")
	integer , err := strconv.Atoi(charge)
	if err != nil {
		log.Fatalf("Error in converting battery status, %v" , err)
	}
	return integer
}

func (d *DefaultProvider) GetBatteryStatus() bool {
	status := strings.Trim(battery.GetBatteryStatus() , "\n")
	if status == "Charging" {
		return true
	}
	return false
}

func mainLoop(provider Provider , sleeper Sleeper , notifier Notifier, stop chan bool) {
	charge_notify := false
	battery_state := 0
	for true {
		charge := provider.GetBatteryCharge()
		status := provider.GetBatteryStatus()

		if !status {
			if charge_notify {
				charge_notify = false
				fmt.Println("Battery Disconnected")
			}
			if charge <= 20 && battery_state != LOW_STATUS{
				battery_state = LOW_STATUS
				fmt.Println("Battery Low")
				
			} else if charge <= 10 && battery_state != VERY_LOW_STATUS{
				battery_state = VERY_LOW_STATUS
				fmt.Println("Battery Very Low!")
				
			}
		} else {
			battery_state = CONNECTED_STATUS
			if !charge_notify {
				charge_notify = true
				fmt.Println("Battery Disconnected")
			}
			charge_notify = true
			if charge >= 75 && battery_state == FULL_STATUS {
				battery_state = FULL_STATUS
				fmt.Println("Battery Full!")
				
			}
		}

		sleeper.Sleep(60 * time.Second)
	}

}
