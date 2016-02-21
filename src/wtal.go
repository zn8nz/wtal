/*
Package main is a stand-alone program to keep a tally of words from a text file.
It prints a list of unique words and their count. One pair per line. Format on each line is:
	99999999: word

There are several options:
	-c the minimum tally for a word to be printed to output
	-w the minimum number of letters for the word to be printed to output
	-a sort ascending (default descending) on tally, word
	-i ignore case
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	minLength   int
	occurrences int
	ascending   bool
	ignoreCase  bool
	prefix      string
)

func init() {
	flag.IntVar(&minLength, "w", 1, "Minimum letters per word to keep tally for")
	flag.IntVar(&occurrences, "c", 1, "Minimum number of occurences")
	flag.BoolVar(&ascending, "a", false, "Sort ascending")
	flag.BoolVar(&ignoreCase, "i", false, "Ignore case")
	flag.StringVar(&prefix, "p", "", "Starts with these letters")
}

func main() {
	flag.Parse()
	
	tally := make(map[string]int)

	filename := flag.Arg(0)
	if filename == "" {
		fmt.Println("File name missing")
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	
	if prefix != "" {
		minLength -= len([]rune(prefix))
		if minLength < 0 {
			minLength = 0
		}
		if ignoreCase {
			prefix = "(?i)" + prefix // (?i) in regexp indicates: ignore case from here
		}		
	}
	
	// Find all letters with \pL and dash and apostrophe; note \w will only find ASCII
	// The minLength is inserted and works on unicode code points, not on ASCII, so 
	// will count non-ASCII characters correctly.
	search := `\b` + prefix + `[\pL'’-]{` + strconv.Itoa(minLength) + `,}`
	rx := regexp.MustCompile(search)

	// Open file, read line by line, find matching words on a line and process these words.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words := rx.FindAllString(line, -1)
		for _, w := range words {
			if ignoreCase {
				w = strings.ToLower(w)
			}
			if _, found := tally[w]; found {
				tally[w]++
			} else {
				tally[w] = 1
			}
		}
	}

	// create a slice with an underlying array big enough to hold all elements
	// but put only the elements in it with enough occurrences.
	s := make([]string, len(tally))

	k := 0
	for w, t := range tally {
		if t < occurrences {
			continue // filter out shorter words as per given option
		}
		s[k] = fmt.Sprintf("%8d: %s", t, w)
		k++
	}
	sk := s[:k]

	// we can sort this slice, as the tallies are right aligned.
	sort.Strings(sk)
	if !ascending {
		// descending is default, so reverse the array
		for i, j := 0, k-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}

	// output
	for _, s := range sk {
		fmt.Println(s)
	}
}
