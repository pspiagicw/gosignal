package main

import (
	"testing"
	"time"

	"github.com/pspiagicw/gosignal/pkg/controller"
)

type TestProvider struct {
	charge int
	status bool
}
func (t *TestProvider) GetBatteryCharge() int {
	return t.charge
}
func (t *TestProvider) GetBatteryStatus() bool {
	return t.status
}

type TestSleeper struct {
	slept int
}
func (t *TestSleeper) Sleep(d time.Duration) {
	t.slept++
}
type TestNotifier struct {
	notified int
}
func (t *TestNotifier) Notify(header string , body string , icon string , duration int) {
	t.notified++
}
