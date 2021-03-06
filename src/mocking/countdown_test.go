package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("3,2,1 Go! with sleep interrupts", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpySleeper{}

		Countdown(buffer, sleeper)

		got := buffer.String()
		want := "3\n2\n1\nGo!\n"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
		if sleeper.Calls != 4 {
			t.Errorf("not enough calls to sleeper, "+
				"want 4 got %d", sleeper.Calls)
		}
	})

	t.Run("Sleep before every countdown call", func(t *testing.T) {
		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		sleeper := &CountdownOperationSpy{}

		Countdown(sleeper, sleeper)

		if !reflect.DeepEqual(want, sleeper.CallOrder) {
			t.Errorf("Call order does not match. "+
				"Got %v, want %v", sleeper.CallOrder, want)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
