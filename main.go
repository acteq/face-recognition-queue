package main

import (
	"github.com/prep/beanstalk"
	"context"
	"log"
	"fmt"
	"encoding/json"
	"bytes"
	"io"
)


func main() {

	consumer, err := beanstalk.NewConsumer([]string{"localhost:11300"}, []string{"face-recognize"}, beanstalk.Config{
		// Multiply the list of URIs to create a larger pool of connections.
		Multiply: 3,
		// NumGoroutines is the number of goroutines that the Receive method will
		// spin up to process jobs concurrently.
		NumGoroutines: 30,
	})
	if err != nil {
		// handle error
		log.Fatal(err)
		return
	}

	type Message struct {
		Filepath string
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for {
		consumer.Receive(ctx, func(ctx context.Context, job *beanstalk.Job) {
			// handle job
			fmt.Println("job", job.ID)
			dec := json.NewDecoder(bytes.NewReader(job.Body))
			var m Message
			for {
				if err := dec.Decode(&m); err == io.EOF {
					break
				} else if err != nil {
					log.Println(err)
					break
				}
				fmt.Printf("%s\n", m.Filepath)
			}
		
			if err := job.Delete(ctx); err != nil {
				log.Println(err)
				// log.Fatal(err)
			}
		})	
	}
}


