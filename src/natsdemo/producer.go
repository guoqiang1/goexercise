package main

import (
	"flag"
	"log"
	"github.com/nats-io/go-nats"
	"gqutils"
)

func main() {

	var natsurls = "nats://localhost:4222"
	var urls = flag.String("s", natsurls, "nats连接地址")

	log.SetFlags(0)
	flag.Parse()

	nc, err := nats.Connect(*urls)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	subj := "test"
	var msg string

	for i := 0; i < 1000; i++ {
		msg = gqutils.NowTime() + "\t" + gqutils.RandString(20)
		nc.Publish(subj, []byte(msg))
		nc.Flush()

		if err := nc.LastError(); err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Published [%s] : '%s'\n", subj, msg)
		}
	}

}
