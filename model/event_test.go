package model

import (
	"fmt"
	"testing"
)

func TestGetAllEvents(t *testing.T) {
	events := GetAllEvents()
	fmt.Println(events)

	if len(events) > 0 {
		t.Errorf("Events array should be empty")
	}
}
