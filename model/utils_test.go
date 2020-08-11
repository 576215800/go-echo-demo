package model

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerateTime(t *testing.T) {
	Time := time.Now()
	fmt.Printf("%+v\n", Time)
}
func TestGenerateID(t *testing.T) {
	id, _ := GetNewID()
	fmt.Println(id)
}
