package searcher

import (
	aho "Searcher/aho-corasick"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
)

type Searcher struct {
	strings []string
}

func New(strings []string) Searcher {
	return Searcher{strings}
}

func (searcher *Searcher) Contain(directoryName *string) bool {
	var wg sync.WaitGroup
	var result int32

	err := filepath.Walk(*directoryName, func(path string, info os.FileInfo, err error) error {
		if info == nil || info.IsDir() {
			return nil
		}

		wg.Add(1)

		go func(filename string) {
			defer wg.Done()

			file, err := os.Open(filename)
			if err != nil {
				return
			}

			defer file.Close()

			ahoCorasick := aho.New()
			for _, val := range searcher.strings {
				ahoCorasick.Add(&val)
			}

			if ahoCorasick.Contain(bufio.NewReader(file)) {
				atomic.StoreInt32(&result, 1)
			}
		}(path)

		return nil
	})

	if err != nil {
		fmt.Println("Directory does not exist.")
		return false
	}

	wg.Wait()

	return result == 1
}
