package ent

import (
	"log"
	"time"

	nebula "github.com/vesoft-inc/nebula-go/v3"
)

const (
	address  = "127.0.0.1"
	port     = 3699
	username = "root"
	password = "nebula"
)

func prepareSpace(spaceName string) error {
	config := nebula.GetDefaultConf()
	host := nebula.HostAddress{
		Host: address,
		Port: port,
	}

	pool, err := nebula.NewConnectionPool(
		[]nebula.HostAddress{host},
		config,
		nebula.DefaultLogger{},
	)
	if err != nil {
		return err
	}

	sess, err := pool.GetSession(username, password)
	if err != nil {
		return err
	}

	conf := nebula.SpaceConf{
		Name:      spaceName,
		Partition: 1,
		Replica:   1,
		VidType:   "FIXED_STRING(12)",
	}

	_, err = sess.CreateSpace(conf)
	if err != nil {
		return err
	}

	log.Println("Space created successfully, waiting for 5 seconds ...")
	time.Sleep(5 * time.Second)

	return nil
}

func newSessionPool(spaceName string) (*nebula.SessionPool, error) {
	err := prepareSpace(spaceName)
	if err != nil {
		return nil, err
	}

	hostAddr := nebula.HostAddress{Host: address, Port: port}
	conf, err := nebula.NewSessionPoolConf(
		username,
		password,
		[]nebula.HostAddress{hostAddr},
		spaceName,
	)
	if err != nil {
		return nil, err
	}

	sessionPool, err := nebula.NewSessionPool(*conf, nebula.DefaultLogger{})
	if err != nil {
		return nil, err
	}

	return sessionPool, nil
}
