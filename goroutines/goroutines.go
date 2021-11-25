package goroutines

import (
	"io"
	"io/ioutil"
	"log"
	"sync"
	"sync/atomic"
)

func oneTimeRoutines(reader io.Reader, n int64) {
	var num int64
	s := make([]byte, 1)
	for {
		if atomic.LoadInt64(&num) > n {
			continue
		}

		nb, err := reader.Read(s)
		if err != nil {
			panic(err)
		}

		if nb == 0 {
			break
		}

		atomic.AddInt64(&num, 1)

		go func(s []byte) {
			_, err := ioutil.ReadFile("goroutines_bench_test.go")
			if err != nil {
				log.Println(err)
			}

			atomic.AddInt64(&num, -1)
		}(s)
	}

	for num > 0 {
	}
}

func longRoutines(reader io.Reader, n int64) {
	c := make(chan []byte)
	var wg sync.WaitGroup
	for i := int64(0); i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range c {
				if v == nil {
					break
				}
				_, err := ioutil.ReadFile("goroutines_bench_test.go")
				if err != nil {
					log.Println(err)
				}
			}
		}()
	}

	s := make([]byte, 1)

	for {
		nb, err := reader.Read(s)
		if err != nil {
			panic(err)
		}

		if nb == 0 {
			break
		}
		c <- s
	}

	close(c)

	wg.Wait()
}
