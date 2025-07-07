// -*- coding: utf-8 -*-

// Created on Sun Jun 22 09:34:16 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"errors"
	"log"
	"math/rand"
)

type SentenceQueue = map[string][]string

func sampleSentences(s SentenceQueue, v int, e int) ([]string, []string, error) {
	var (
		err  error
		ss   []string
		vv   []string
		keys []string
	)

	if v == 0 || e == 0 {
		err = errors.New("either v or e are 0 in call to sampleSentences")
		return vv, ss, err
	}

	for k := range s {
		if len(keys) == v {
			break
		}
		keys = append(keys, k)
	}

	if len(keys) < v {
		err = errors.New("call to sampleSentences selects more vocabulary words v than exist in sentence queue")
		return vv, ss, err
	}

	for _, k := range keys {
		if e > len(k) {
			log.Printf("not enough example sentences for vocab %s", k)
			ss = s[k][:len(k)]
			return vv, ss, err
		}
		rand.Shuffle(len(s[k]), func(i, j int) { s[k][i], s[k][j] = s[k][j], s[k][i] })
		vv = append(vv, k)
		ss = append(ss, s[k][:e]...)
	}

	return vv, ss, err
}
