package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "math"
  "net/http"
  "os"
  "strconv"
  "strings"
  "time"
)

const (
  filename = "p067_triangle.txt"
  url = "https://projecteuler.net/project/resources/p067_triangle.txt"
)

func Max(a int, b int) int {
  if a > b { return a }
  return b
}

func main() {
  // Download file if it doesn't already exists:
  if _, err := os.Stat(filename); os.IsNotExist(err) {
    resp, err := http.Get(url);
    if err != nil {
      log.Fatal(err)
    }
    defer resp.Body.Close()
    buf, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Fatal(err)
    }
    if err := ioutil.WriteFile(filename, buf, 0644); err != nil {
      log.Fatal(err)
    }
  }
  // Read file and parse:
  content, err := ioutil.ReadFile(filename)
  if err != nil {
    log.Fatal(err)
  }
  numbers := strings.Fields(string(content))
  // For testing:
  // numbers = []string{"3", "7", "4", "2", "4" , "6", "8", "5", "9", "3"}
  var sums = make([]int, len(numbers))
  for i := range numbers {
    sums[i], _ = strconv.Atoi(numbers[i]);
  }
  // Measure execution time:
  t0 := time.Now()
  // calculate number of rows
  // f(r) = r + f(r-1), f(0) = 0
  // => rows = (sqrt(8*numbers+1)-1)/2
  rows := (int)(math.Sqrt((float64)(8*len(numbers)+1))-1)/2
  // Start with the last item, at the bottom row:
  for row, offset := rows, len(sums)-1; row > 0; row-- {
    for i := row; i > 1; i-- {
      // Take the biggest of the numbers next to each other, and add to the above
      sums[offset-row] += Max(sums[offset], sums[offset-1])
      offset--;
    }
    offset--;
  }
  t1 := time.Now()
  fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
  fmt.Println(sums[0])
}

