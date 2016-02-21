# wtal
Word tally.

A stand-alone program that opens a given text file and counts occurrences of words.
It prints a list of unique words and their count. One pair per line. Format on each line is:

	99999999: word

## options
	-c the minimum tally for a word to be printed to output
	-w the minimum number of letters for the word to be printed to output
	-a sort ascending (default descending) on tally, word
	-i ignore case

## examples
	./wtal -i -c=50 -w=4 myfile.txt
	Print out words of four letters or more and their tally where the tally is 50 or more. Ignore case.

	`./wtal -a -w=8 mydoc`
	Print out words of 8 letters or more and their tally in ascending order of tally.
	