package main

import (
	"Trie"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println(getFileCount("./seowords/files"))
	size := getFileCount("./seowords/files")
	root := buildTrie("./sensitive_words/word.csv")

	ch := make(chan int, 1)
	exit := make(chan int, 20)

	ch <- 0

	for i := 0; i < 20; i++ {
		go run(ch, exit, root, size)
	}

	for i := 0; i < 20; i++ {
		<-exit
	}
	<-ch

	fmt.Println("over")
}

func run(ch chan int, exit chan int, root *Trie.Trie, size int) {
	for {
		cur := <-ch
		ch <- (cur + 1)

		if cur >= size {
			exit <- 1
			break
		}

		fmt.Println(cur)

		fileName := fmt.Sprintf("./seowords/files/words%d.csv", cur)
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			panic(err)
		}
		str := string(data)
		words := strings.Split(str, "\n")

		resArr := filter(words, root)
		resStr := strings.Join(resArr, "\n")
		ioutil.WriteFile(fmt.Sprintf("./seowords/res/res%d.csv", cur), []byte(resStr), os.ModePerm)

		fmt.Printf("file %d over.\n", cur)
	}
}

func filter(words []string, root *Trie.Trie) []string {
	res := make([]string, 0)
	for _, word := range words {
		if root.IsExist(([]rune)(word)) {
			res = append(res, word)
		}
	}

	return res
}

func getFileCount(path string) int {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	return len(files)
}

func buildTrie(fileName string) *Trie.Trie {
	root := Trie.NewTrie()

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Printf("read error %v\n", err)
	}
	str := string(data)

	rows := strings.Split(str, "\n")
	for _, row := range rows {
		root.Insert(([]rune)(row))
	}

	root.BuildFailPointer()
	return root
}
