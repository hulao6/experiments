package zapkafka

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/Shopify/sarama"
	"github.com/Shopify/sarama/mocks"
	"go.uber.org/zap"
)

func TestWriteFailWithKafkaSyncer(t *testing.T) {
	config := sarama.NewConfig()
	p := mocks.NewAsyncProducer(t, config)

	var buf = make([]byte, 0, 256)
	w := bytes.NewBuffer(buf)
	w.Write([]byte("hello"))
	logger := New(NewKafkaSyncer(p, "test", NewFileSyncer(w)), 0)

	p.ExpectInputAndFail(errors.New("produce error"))
	p.ExpectInputAndFail(errors.New("produce error"))

	// all below will be written to the fallback sycner
	logger.Info("demo1", zap.String("status", "ok")) // write to the kafka syncer
	logger.Info("demo2", zap.String("status", "ok")) // write to the kafka syncer

	// make sure the goroutine which handles the error writes the log to the fallback syncer
	time.Sleep(2 * time.Second)

	s := string(w.Bytes())
	if !strings.Contains(s, "demo1") {
		t.Errorf("want true, actual false")
	}
	if !strings.Contains(s, "demo2") {
		t.Errorf("want true, actual false")
	}

	if err := p.Close(); err != nil {
		t.Error(err)
	}
}

func TestWriteOKWithKafkaSyncer(t *testing.T) {
	config := sarama.NewConfig()
	p := mocks.NewAsyncProducer(t, config)

	var buf = []byte{}
	w := bytes.NewBuffer(buf)

	logger := New(NewKafkaSyncer(p, "test", NewFileSyncer(w)), 0)

	messageChecker := func(msg *sarama.ProducerMessage) error {
		b, err := msg.Value.Encode()
		if err != nil {
			return err
		}

		var m = make(map[string]interface{})
		err = json.Unmarshal(b, &m)
		if err != nil {
			fmt.Printf("unmarshal error: %s\n", err)
			return err
		}
		//fmt.Printf("topic=%s, value=%s\n", msg.Topic, m)

		v, ok := m["msg"].(string)
		if !ok {
			err = errors.New("invalid msg")
			fmt.Printf("type assertion error: %s\n", err)
			return err
		}

		if v != "demo1" {
			err = errors.New("invalid msg value")
			fmt.Printf("type assertion error: %s\n", err)
			return err
		}

		v1, ok := m["status"].(string)
		if !ok {
			err = errors.New("invalid status")
			fmt.Printf("type assert error")
			return err
		}
		if v1 != "ok" {
			return errors.New("invalid status value")
		}
		if msg.Topic != "test" {
			return errors.New("invalid topic")
		}
		//fmt.Printf("topic=%s, value=%s\n", msg.Topic, m)
		return nil
	}

	p.ExpectInputWithMessageCheckerFunctionAndSucceed(mocks.MessageChecker(messageChecker))
	logger.Info("demo1", zap.String("status", "ok"))

	if err := p.Close(); err != nil {
		t.Error(err)
	}

	if len(w.Bytes()) != 0 {
		t.Errorf("want 0, actual %d", len(w.Bytes()))
	}
}
