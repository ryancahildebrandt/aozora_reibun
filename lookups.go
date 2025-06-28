// -*- coding: utf-8 -*-

// Created on Tue Jun 17 11:06:54 AM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"slices"
	"strings"

	"github.com/ikawaha/kagome/v2/tokenizer"
)

type RequestBody struct {
	Query     string `json:"query"`
	Language  string `json:"language"`
	NoEnglish bool   `json:"no_english"`
}

type WordResponseBody struct {
	Words []ResponseWord `json:"words"`
}

type KanjiResponseBody struct {
	Kanji []ResponseKanji `json:"kanji"`
}

type ResponseKanji struct {
	Literal     string   `json:"literal"`
	Meanings    []string `json:"meanings"`
	Grade       int      `json:"grade"`
	Frequency   int      `json:"frequency"`
	JLPT        int      `json:"jlpt"`
	Onyomi      []string `json:"onyomi"`
	Kunyomi     []string `json:"kunyomi"`
	StrokeCount int      `json:"stroke_count"`
	Parts       []string `json:"parts"`
	Radical     string   `json:"radical"`
}

type ResponseWord struct {
	Reading ResponseReading  `json:"reading"`
	Senses  []ResponseSenses `json:"senses"`
}

type ResponseReading struct {
	Kana     string `json:"kana"`
	Kanji    string `json:"kanji"`
	Furigana string `json:"furigana"`
}

type ResponseSenses struct {
	Glosses  []string `json:"glosses"`
	POS      any      `json:"pos"`
	Language string   `json:"language"`
}

type WordLookup struct {
	Query string
	WordResponseBody
	KanjiResponseBody
}

type SentenceLookups struct {
	Sentence string
	Lookups  []WordLookup
}

func queryWord(w string) (WordResponseBody, error) {
	b, err := json.Marshal(RequestBody{w, "English", false})
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodPost, "https://jotoba.de/api/search/words", bytes.NewBuffer(b))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	out, err := unmarshalWordResponse(res, w)

	return out, err
}

func queryKanji(w string) (KanjiResponseBody, error) {
	b, err := json.Marshal(RequestBody{w, "English", false})
	req, err := http.NewRequest(http.MethodPost, "https://jotoba.de/api/search/kanji", bytes.NewBuffer(b))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	out, err := unmarshalKanjiResponse(res, w)

	return out, err
}

func unmarshalWordResponse(r *http.Response, w string) (WordResponseBody, error) {
	var out WordResponseBody
	var err error

	j := json.NewDecoder(r.Body)
	err = j.Decode(&out)
	if err != nil {
		return out, err
	}
	out.Words = slices.DeleteFunc(out.Words, func(t ResponseWord) bool { return t.Reading.Kanji != w })

	return out, err
}

func unmarshalKanjiResponse(r *http.Response, w string) (KanjiResponseBody, error) {
	var out KanjiResponseBody
	var err error

	j := json.NewDecoder(r.Body)
	err = j.Decode(&out)
	if err != nil {
		return out, err
	}
	out.Kanji = slices.DeleteFunc(out.Kanji, func(t ResponseKanji) bool { return !strings.Contains(w, t.Literal) })
	return out, err
}

func lookupSentence(s string, t *tokenizer.Tokenizer) (SentenceLookups, error) {
	var (
		sl  SentenceLookups
		err error
		r   *regexp.Regexp
		w   WordResponseBody
		k   KanjiResponseBody
		l   WordLookup
	)

	r, err = regexp.Compile("([一-龯])")
	if err != nil {
		return sl, err
	}
	sl.Sentence = s
	for _, t := range t.Wakati(s) {
		if r.MatchString(t) {
			w, err = queryWord(t)
			if err != nil {
				return sl, err
			}
			k, err = queryKanji(t)
			if err != nil {
				return sl, err
			}
			l = WordLookup{t, w, k}
			sl.Lookups = append(sl.Lookups, l)
		}
	}

	return sl, err
}
