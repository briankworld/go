package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "sync"
)

func countWords(filename string, ch chan<-map[string]int, wg *sync.WaitGroup) {
  defer wg.Done()

  wordCount := make(map[string]int)

  file, err := os.Open(filename)
  if err != nil {
		fmt.Println("Error opening file:", filename, err)
		return
	}
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
      line := scanner.Text()
      words := strings.Fields(line)
      for _, word := range words {
        wordCount[word]++
      }
  }

  ch <-wordCount
}

func merge(totalCounts map[string]int, partialCounts map[string]int) {
  for word, count := range partialCounts {
    totalCounts[word] += count
  }
}


func main() {
  files := []string { "filename1.txt", "filename2.txt" }

  ch := make(chan map[string]int)

  var wg wg.WaitGroup

  for _, file := range files {
    wg.Add(1)
    go countWords(file, ch, &wg)
  }

  go func() {
		wg.Wait()
		close(ch)
	}()

  totalCounts := make(map[string]int)
  
  for partialCounts := range ch {
    merge(totalCounts, partialCounts)
  }

	fmt.Println("Word Frequencies:")
	for word, count := range totalCounts {
		fmt.Printf("%s: %d\n", word, count)
	}
}
