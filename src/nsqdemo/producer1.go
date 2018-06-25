package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/pquerna/ffjson/ffjson"
	"log"
	"strconv"
	"utils/comm"
	"utils/php"
)

func getProducer(address string) nsq.Producer {
	producer, err := nsq.NewProducer(address, nsq.NewConfig())

	if err != nil {
		panic(err)
	}

	return *producer
}

func main() {
	pro := getProducer("127.0.0.1:4150")
	fmt.Println(pro)

	var s1 string
	var msg string

	msgmap := make(map[string]interface{})

	for i := 0; i < 20; i++ {
		n1 := php.PHPrand(1000000, 9999999)

		s1 = strconv.Itoa(n1)

		msgmap["time"] = comm.NowTime()

		msg = "hello world good morning " + s1

		msgmap["msg"] = msg

		result, _ := ffjson.Marshal(msgmap)
		resultstr := string(result)

		re := pro.Publish("test", []byte(resultstr))
		fmt.Println(re)
		log.Println(msgmap)
	}

}
