package main

import (
	"log"
	"strings"
	"unicode/utf8"
)

func main() {
	s1 := "An eye examination is necessary to check your vision and screen for eye conditions."
	s2 := "我是一名开发工程师1程师1程师123"
	ss1 := splitSentence(s1)
	log.Printf("%#v %v", ss1, len(ss1))
	ss2 := splitSentence(s2)
	log.Printf("%#v %v", ss2, len(ss2))
}

func splitSentence(sentence string) []string {
	min, max := 2, 4
	subSens := []string{}
	charLen := utf8.RuneCountInString(sentence)
	log.Println("charLen", charLen, len(sentence))
	subLen := min
	sep := ""
	if isEnglishWord(sentence) {
		sep = " "
	}
	sentences := strings.Split(sentence, sep)
	if len(sentences) > 3*max {
		subLen = 4
	} else if len(sentences) > 2*max {
		subLen = 3
	}
	log.Printf("subLen=%v, sentences=%#v", subLen, sentences)
	sub := len(sentences) / subLen
	for len(sentences) >= subLen*2 {
		log.Println("len(sentences)=", len(sentences))
		subSens = append(subSens, strings.Join(sentences[:sub], sep))
		sentences = sentences[sub:]
	}
	subSens = append(subSens, strings.Join(sentences, sep))
	return subSens
}

func isEnglishWord(word string) bool {
	if len(word) == 0 {
		return false
	}
	if len(word) == 1 {
		return (word[0] >= 'A' && word[0] <= 'z')
	}
	if (word[0] >= 'A' && word[0] <= 'z') && (word[len(word)-2] >= 'A' && word[len(word)-2] <= 'z') {
		return true
	}
	return false
}
