package ent

import "fmt"

func prepareSpace(spaceName string) error {
	fmt.Println("prepare space", spaceName)
	return nil
}

func dropSpace(spaceName string) {
}
