package main

import (
	"flag"
	"github.com/anacrolix/sync"
	"log"
	"net/rpc"
	"os"
	"strconv"
	"time"
	"math/rand"
	"sharing/rpc/lib"
)


var (
	ThreadCount = 500
	Repeat      = 1000
)

func init() {
	println("Usage Sample:")
	println(os.Args[0] + " -thread 100  -repeat 5000")
	flag.IntVar(&ThreadCount, "thread", 1000, "how many threads(goroutine) running in client side")
	flag.IntVar(&Repeat, "repeat", 500, "repeat count within one thread")
	flag.Parse()
	println("your input: -thread:" + strconv.Itoa(ThreadCount) + " -repeat:" + strconv.Itoa(Repeat))
}

func main() {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	_ = rd

	before := time.Now()
	var wg sync.WaitGroup
	wg.Add(ThreadCount)
	for r := 0; r < ThreadCount; r++ {
		conn, err := rpc.Dial("tcp", "127.0.0.1:8095")
		if err != nil {
			log.Fatalln("dailing error: ", err)
		}
		go func() {
			for i := 0; i < Repeat; i++ {
				req := lib.ArithRequest{i / 2, i}
				var res lib.ArithResponse

				err = conn.Call("Arith.Multiply", req, &res)
				if err != nil {
					log.Fatalln("arith error: ", err)
				}

			}
			wg.Done()
		}()
	}
	wg.Wait()
	t := time.Now().Sub(before).Seconds()
	log.Println("total time:", t)
	log.Println("QPS:", float64(Repeat*ThreadCount)/t)

}
