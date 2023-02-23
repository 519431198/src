package redis_test

import (
	"fmt"
)

func Example_instrumentation() {
	redisdb := NewClient(&Options{
		Addr: ":6379",
	})
	redisdb.WrapProcess(func(old func(cmd Cmder) error) func(cmd Cmder) error {
		return func(cmd Cmder) error {
			fmt.Printf("starting processing: <%s>\n", cmd)
			err := old(cmd)
			fmt.Printf("finished processing: <%s>\n", cmd)
			return err
		}
	})

	redisdb.Ping()
	// Output: starting processing: <ping: >
	// finished processing: <ping: PONG>
}

func ExamplePipeline_instrumentation() {
	redisdb := NewClient(&Options{
		Addr: ":6379",
	})

	redisdb.WrapProcessPipeline(func(old func([]Cmder) error) func([]Cmder) error {
		return func(cmds []Cmder) error {
			fmt.Printf("pipeline starting processing: %v\n", cmds)
			err := old(cmds)
			fmt.Printf("pipeline finished processing: %v\n", cmds)
			return err
		}
	})

	redisdb.Pipelined(func(pipe Pipeliner) error {
		pipe.Ping()
		pipe.Ping()
		return nil
	})
	// Output: pipeline starting processing: [ping:  ping: ]
	// pipeline finished processing: [ping: PONG ping: PONG]
}
