package main

import (
	"github.com/prep/beanstalk"
	"context"
	"log"
	"fmt"
	"encoding/json"
	"bytes"
	"io"
	"time"
	"os"
	"strings"
)


func main() {
	pool := NewEnginePool(10)

	env := getEnvironment()
	host := env["beanstalkd_host"] 
	if host == "" {
		host = "localhost"
	}
	port := env["beanstalkd_port"] 
	if port == "" {
		port = "11300"
	}
	tuneName := env["beanstalkd_tube"] 
	if tuneName == "" {
		tuneName =  "face-recognize"
	}
	
	consumer, err := beanstalk.NewConsumer([]string{host + ":" + port}, []string{tuneName}, beanstalk.Config{
		// Multiply the list of URIs to create a larger pool of connections.
		Multiply: 3,
		// NumGoroutines is the number of goroutines that the Receive method will
		// spin up to process jobs concurrently.
		NumGoroutines: 30,
	})
	if err != nil {
		// handle error
		log.Fatal(err)
	}

	producer, err := beanstalk.NewProducer([]string{host + ":" + port}, beanstalk.Config{
		// Multiply the list of URIs to create a larger pool of connections.
		Multiply: 3,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer producer.Stop()
	params := beanstalk.PutParams{Priority: 1024, Delay: 0, TTR: 30 * time.Second}

	type Message struct {
		id int
		filepath string
	}
	type Result struct {
		id int
		feature []byte
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	for {
		consumer.Receive(ctx, func(ctx context.Context, job *beanstalk.Job) {
			// handle job
			dec := json.NewDecoder(bytes.NewReader(job.Body))
			var m Message
			for {
				if err := dec.Decode(&m); err == io.EOF {
					break
				} else if err != nil {
					log.Println(err)
					break
				}
				log.Println(job.ID,  m.filepath)
				handle := pool.Get()
				featureBytes := extract(m.filepath, handle)
				//convert 
				result := Result{id: m.id, feature: featureBytes}
		
				resultStr, err := json.Marshal(result)
				id, err := producer.Put(ctx, tuneName, resultStr, params)
				log.Println("send result back", id)
				if err != nil {
					log.Println(err)
				}
				pool.Put(handle)
			}

			if err := job.Delete(ctx); err != nil {
				log.Println(err)
			}
		})
		//runtime.Gosched()	
	}
}

func Dec2Hex(n int64) string {
	if n < 0 {
	   log.Println("Decimal to hexadecimal error: the argument must be greater than zero.")
	   return ""
	}
	if n == 0 {
	   return "00"
	}
	hex := map[int64]int64{10: 65, 11: 66, 12: 67, 13: 68, 14: 69, 15: 70}
	s := ""
	for q := n; q > 0; q = q / 16 {
	   m := q % 16
	   if m > 9 && m < 16 {
		  m = hex[m]
		  s = fmt.Sprintf("%v%v", string(m), s)
		  continue
	   }
	   s = fmt.Sprintf("%v%v", m, s)
	}
	return s
 }

 func getEnvironment() map[string]string{
	var env = map[string]string{}
	for _, line:= range os.Environ() {
		pair := strings.Split(line, "=")
		if len(pair) > 1 {
			key := strings.Trim(pair[0], " ")
			if(len(key) > 0) {
				env[key] = strings.Trim(pair[1], " ")
			}
		}
	}
	return env
}