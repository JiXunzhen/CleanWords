package main

import (
	"Trie"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dir := os.Args[1]
	size := getFileCount(dir + "/files")
	root := buildTrie("./sensitive_words/word.csv")

	ch := make(chan int, 1)
	exit := make(chan int, 20)

	ch <- 0

	for i := 0; i < 20; i++ {
		go run(ch, exit, root, size, dir)
	}

	for i := 0; i < 20; i++ {
		<-exit
	}
	<-ch

	fmt.Println("over")
}

func run(ch chan int, exit chan int, root *Trie.Trie, size int, dir string) {
	for {
		cur := <-ch
		ch <- (cur + 1)

		if cur >= size {
			exit <- 1
			break
		}

		fmt.Println(cur)

		fileName := fmt.Sprintf(dir+"/files/words%d.csv", cur)
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			panic(err)
		}
		str := string(data)
		words := strings.Split(str, "\n")

		resArr := filterFile(words, root)
		resStr := strings.Join(resArr, "\n")
		ioutil.WriteFile(fmt.Sprintf(dir+"/res/res%d.csv", cur), []byte(resStr), os.ModePerm)

		fmt.Printf("file %d over.\n", cur)
	}
}

func filterFile(rows []string, root *Trie.Trie) []string {
	res := make([]string, 0)
	for _, row := range rows {
		word := strings.Split(row, "&")[0]

		if filterWord(word, root) {
			res = append(res, row)
		}
	}

	return res
}

func filterWord(word string, root *Trie.Trie) bool {
	// 符号过滤
	symbols := []string{"，", "；", ":", "&", "=", ",", ";", "*", "#", ".", "-", "‘", "+"}
	for _, symbol := range symbols {
		if strings.HasPrefix(word, symbol) {
			return true
		}
	}

	// 单字过滤
	chword := ([]rune)(word)
	if len(chword) == 1 {
		return true
	}

	// 敏感词过滤
	return root.IsExist(([]rune)(word))
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
	if fileName != "" {
		data, err := ioutil.ReadFile(fileName)

		if err != nil {
			panic(err)
		}
		str := string(data)

		rows := strings.Split(str, "\n")
		for _, row := range rows {
			root.Insert(([]rune)(row))
		}
	}

	root.BuildFailPointer()
	return root
}
