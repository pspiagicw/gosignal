package battery

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

func GetBatteryCharge() string {
	battery := getBatteryCapacityFile()
	contents , err  := ioutil.ReadFile(battery)
	if err != nil {
		log.Fatalf("Error reading battery status , %v" , err)
	}
	return string(contents)
}
func getBatteryName() string {
	return filepath.Join(POWER_SUPPLY_SUBSYSTEM , BATTERY_NAME)
}
func getBatteryCapacityFile() string {
	return filepath.Join(getBatteryName() , "capacity" )
}
func getBatteryStatusFile() string {
	return filepath.Join(getBatteryName() , "status")
}

func GetBatteryStatus() string {
	battery := getBatteryStatusFile()
	contents , err := ioutil.ReadFile(battery)

	if err != nil {
		log.Fatalf("Error reading status file , %v" , err)
	}
	return string(contents)
	
}

