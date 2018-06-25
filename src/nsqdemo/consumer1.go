package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/pquerna/ffjson/ffjson"
	"log"
	"nsqdemo/models"
	"sync"
)

type MsgRow struct {
}
type NSQHandler struct {
	msglist chan interface{}
}

func (this *NSQHandler) HandleMessage(message *nsq.Message) error {

	//result := make(map[int][string]interface{})

	//var
	//result :=[]makemap[string]interface{}
	result := make(map[string]interface{})

	msgcontent := string(message.Body)

	//var err error
	ffjson.Unmarshal([]byte(msgcontent), &result)

	logmodel := models.GqLog{}

	log.Println("consumer", logmodel)
	//log.Println(result)
	//log.Println(err)

	tm := result["time"].(string)
	logmsg := models.GqLog{FAddtime: tm, FContent: result["msg"].(string)}

	//log.Println(logmsg)
	logmodel.Add(logmsg)
	//e := logmodel.Add(logmsg)
	//log.Println(e)

	//if err == nil {
	//	tm := result["time"].(string)
	//
	//	logmsg := models.Log{FAddtime: tm, FContent: result["msg"].(string)}
	//
	//	e := logmodel.Add(logmsg)
	//	log.Println(e)
	//} else {
	//	log.Println(err)
	//}

	log.Println("recv:", msgcontent)
	return nil
}

type nsqLogHook struct {
}

func (h nsqLogHook) Output(c int, s string) error {
	fmt.Println(s)
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

		var loghook nsqLogHook
		consumer.AddHandler(&NSQHandler{})
		consumer.SetLogger(loghook, nsq.LogLevelInfo)
		err = consumer.ConnectToNSQD("127.0.0.1:4150")
		if nil != err {
			log.Println(err)
			fmt.Println(err)
			return
		}
		//tick := time.NewTicker(time.Second*1)

		//for {
		//	select {
		//	case <-tick.C:
		//
		//	}
		//}

		select {}
	}()

	waiter.Wait()
}

func main() {
	consumNsq()
}
