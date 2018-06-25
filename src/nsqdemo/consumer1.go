package main

import (
		"github.com/nsqio/go-nsq"
	"github.com/pquerna/ffjson/ffjson"
	"log"
	"nsqdemo/models"
	"sync"
)

type NSQHandler struct {
}

func (this *NSQHandler) HandleMessage(message *nsq.Message) error {

	result := make(map[string]interface{})

	msgcontent := string(message.Body)

	ffjson.Unmarshal([]byte(msgcontent), &result)

	logmodel := models.GqLog{}

	tm := result["time"].(string)
	logmsg := models.GqLog{FAddtime: tm, FContent: result["msg"].(string)}

	logmodel.Add(logmsg)

	log.Println("recv:", msgcontent)
	return nil
}

func consumNsq() {
	waiter := sync.WaitGroup{}
	waiter.Add(1)

	go func() {
		defer waiter.Done()

		consumer, err := nsq.NewConsumer("test", "1", nsq.NewConfig())
		if nil != err {
			log.Println(err)
			return
		}

		consumer.AddHandler(&NSQHandler{})

		err = consumer.ConnectToNSQD("127.0.0.1:4150")
		if nil != err {
			log.Println(err)
			return
		}

		select {}
	}()

	waiter.Wait()
}

func main() {
	consumNsq()
}
