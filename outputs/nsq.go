package outputs

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"github.com/Acey9/apacket/logp"
	"github.com/Acey9/apacket/utils"
	"github.com/nsqio/go-nsq"
	"io/ioutil"
	"log"
	"time"
)

var nullLogger = log.New(ioutil.Discard, "", log.LstdFlags)

type NSQOutputer struct {
	Addr      utils.StringArray
	Topic     string
	msgQueue  chan []byte
	producers map[string]*nsq.Producer
}

func NewNSQOutputer(serverAddrs utils.StringArray, topic string) (*NSQOutputer, error) {
	sn := &NSQOutputer{
		Addr:     serverAddrs,
		Topic:    topic,
		msgQueue: make(chan []byte, 1024),
	}
	err := sn.Init()
	if err != nil {
		return nil, err
	}
	go sn.Start()
	return sn, nil
}

func (this *NSQOutputer) Close() {
	for _, producer := range this.producers {
		producer.Stop()
	}
}

func (this *NSQOutputer) Init() error {
	cfg := nsq.NewConfig()
	cfg.UserAgent = fmt.Sprintf("apacket-go-nsq/%s", nsq.VERSION)
	cfg.DialTimeout = time.Millisecond * time.Duration(3000)

	this.producers = make(map[string]*nsq.Producer)
	for _, addr := range this.Addr {
		producer, err := nsq.NewProducer(addr, cfg)
		if err != nil {
			logp.Err("new Producer error: %v", err)
			return err
		}
		producer.SetLogger(nullLogger, nsq.LogLevelInfo)
		this.producers[addr] = producer
	}
	return nil
}

func (this *NSQOutputer) Output(msg []byte) {
	logp.Info("pkt %s", msg)

	var buf bytes.Buffer
	w := zlib.NewWriter(&buf)
	w.Write(msg)
	w.Close()

	this.msgQueue <- buf.Bytes()
}

func (this *NSQOutputer) Send(msg []byte) {
	for _, producer := range this.producers {
		err := producer.Publish(this.Topic, msg)
		if err != nil {
			logp.Err("nsq publish: %v", err)
			continue
		}
		break
	}
}

func (this *NSQOutputer) Start() {
	counter := 0
	defer this.Close()
	for {
		select {
		case msg := <-this.msgQueue:
			counter++
			this.Send(msg)
			if counter%1024 == 0 {
				logp.Debug("nsq", "nsq packet number: %d", counter)
			}
		}
	}
}
