// -*- coding: utf-8 -*-

// Created on Sun Jun 22 09:34:16 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"errors"
	"log"
	"math/rand"
	"strings"
)

type SentenceQueue = map[string][]string

func sampleSentences(s SentenceQueue, v int, e int) ([]string, error) {
	var (
		err  error
		ss   []string
		keys []string
	)

	if v == 0 || e == 0 {
		err = errors.New("either v or e are 0 in call to sampleSentences")
		return ss, err
	}

	for k := range s {
		if len(keys) == v {
			break
		}
		keys = append(keys, k)
	}

	if len(keys) < v {
		err = errors.New("call to sampleSentences selects more vocabulary words v than exist in sentence queue")
		return ss, err
	}

	for _, k := range keys {
		if e > len(k) {
			log.Printf("not enough example sentences for vocab %s", k)
			ss = s[k][:len(k)]
			return ss, err
		}
		rand.Shuffle(len(s[k]), func(i, j int) { s[k][i], s[k][j] = s[k][j], s[k][i] })
		ss = append(ss, s[k][:e]...)
	}

	return ss, err
}

func enqueueSentences(s []string, v []string) SentenceQueue {
	var m SentenceQueue = make(map[string][]string)

	for _, voc := range v {
		for _, sent := range s {
			if strings.Contains(sent, voc) {
				m[voc] = append(m[voc], sent)
			}
		}
		if len(m[voc]) == 0 {
			log.Printf("no queried sentences contain vocab %v, so it will not be included in the sentence queue", voc)
			delete(m, voc)
		}
	}
	return m
}
