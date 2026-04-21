package main

import (
	"context"
	"strings"
	"unicode"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Text string `json:"text"`
}

type Response struct {
	WordCount        int            `json:"word_count"`
	CharCount        map[string]int `json:"char_count"`
	MostFrequentWord string         `json:"most_frequent_word"`
	ReadingTime      int            `json:"reading_time_minutes"`
}

// normalize is used as the first step to normalize the given text and clean it up. It removes extra spaces, converts to lowercase.
func normalize(s string) string {
	s = strings.ToLower(s)
	var b strings.Builder
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsDigit(c) || unicode.IsSpace(c) {
			b.WriteRune(c)
		}

	}
	cleaned := b.String()
	words := strings.Fields(cleaned)

	return strings.Join(words, " ")
}

// wordCount returns the number of words in the given text
func wordCount(s string) int {
	words := strings.Fields(s)
	return len(words)
}

// charCount returns the number of characters (excluding spaces) from the given text
func charCount(s string) map[string]int {
	result := make(map[string]int)
	for _, c := range s {
		if string(c) != " " {

			result[string(c)]++
		}

	}
	return result
}

// mostFrequentWord will return a string of the most most frequently appearing word in the given input
func mostFrequentWord(s string) string {
	m := make(map[string]int)
	max_count := 0
	result := ""
	words := strings.FieldsSeq(s)

	for word := range words {
		m[word] += 1
		if m[word] > max_count {
			max_count = m[word]
			result = word
		}
	}
	return result
}

// approxReadTime will return the approximate time it takes to read the passage, given your speed is 200 words per minute (considered standard)
func approxReadTime(s string) int {
	num_words := wordCount(s)
	return (num_words + 199) / 200
}

// IntelligentReader() will take an english paragraph and compute the following:
// 1. word count
// 2. character count
// 3. most frequent word
// 4. approximate reading time
func IntelligentReader(ctx context.Context, req Request) (Response, error) {

	answer := Response{}
	s := normalize(req.Text)
	answer.WordCount = wordCount(s)
	answer.CharCount = charCount(s)
	answer.MostFrequentWord = mostFrequentWord(s)
	answer.ReadingTime = approxReadTime(s)
	return answer, nil
}

func main() {
	lambda.Start(IntelligentReader)
}
