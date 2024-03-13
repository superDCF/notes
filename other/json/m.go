package main

import (
	"encoding/json"
	"log"
)

type FeedbackRsp struct {
	Score              string   `json:"score"`
	ImprovedVersion    string   `json:"improved_version"`
	GrammarMistakes    []string `json:"grammar_mistakes"`
	VocabularyMistakes []string `json:"vocabulary_mistakes"`
	RawMsgs            []string `json:"raw_msgs"`
}

func main() {
	var s = "{\n    \"score\": \"Can be Improved\",\n    \"improved_version\": \"我想要一个南面的房间。\",\n    \"grammar_mistakes\": [\n        \"Missing a word to make the sentence more complete and precise ('要' is needed after '想' to express 'would like to have').\",\n        \"The direction '南部' is more commonly used to describe a larger geographical area rather than a direction in this context. '南面' is more appropriate when referring to the orientation of a room.\"\n    ],\n    \"vocabulary_mistakes\": []\n}"
	var b FeedbackRsp
	err := json.Unmarshal([]byte(s), &b)
	log.Printf("b%+v err=%v", b, err)
}
