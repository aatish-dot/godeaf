package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	wordsresponse, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")

	if err != nil {
		log.Fatal("unable to fetch words from the website", err)
	}

	//bytstringwords, _ := ioutil.ReadAll(wordsresponse.Body)
	//stringofwords := string(bytstringwords)

	//fmt.Println(stringofwords)

	//words := make(map[string]string)

	sc := bufio.NewScanner(wordsresponse.Body)

	defer wordsresponse.Body.Close()
	sc.Split(bufio.ScanWords)

	buckets := make([][]string, 12)

	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}

	for sc.Scan() {
		word := sc.Text()
		n := hashbucket(word, 12)
		buckets[n] = append(buckets[n], word)
	}

	// i := 100

	// for k := range words {
	// 	fmt.Println(k)
	// 	if i == 200 {
	// 		break
	// 	}
	// 	i++
	// }

	for i := 0; i < 12; i++ {
		fmt.Println(i, "-", len(buckets[i]))

	}

}

func hashbucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
