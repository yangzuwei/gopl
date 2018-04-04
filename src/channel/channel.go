// channel project channel.go
package main

import (
	"log"
	"os"
	"sync"

	"gopl.io/ch8/thumbnail"
)

//这种方式 主线程已经退出了 它们还没干活
func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f)
	}
}

func makeThumbnails1(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(f)
	}

	for range filenames {
		<-ch
	}
}

func makeThumbnails2(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	//closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

func main() {
	files := []string{"img/a.jpg", "img/b.jpg", "img/c.jpg"}
	files2 := make(chan string)
	for _, f := range files {
		files2 <- f
		makeThumbnails2(files2)
	}
}
