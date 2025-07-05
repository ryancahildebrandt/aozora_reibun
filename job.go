// -*- coding: utf-8 -*-

// Created on Thu Jun 12 11:00:24 AM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"log"
	"strings"

	"github.com/ikawaha/kagome/v2/tokenizer"
)

func job(sq SentenceQueue, t *tokenizer.Tokenizer, c ConfigFields) error {
	var (
		err error
		b   strings.Builder
		ss  []string
		sss []SentenceLookups
		vv  []string
		l   SentenceLookups
	)

	vv, ss, err = sampleSentences(sq, c.NVocab, c.NExamples)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for _, sent := range ss {
		l, err = lookupSentence(sent, t)
		if err != nil {
			log.Fatal(err)
			return err
		}
		sss = append(sss, l)
	}

	b, err = renderEmail(sss, &b)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = SendEmail(vv, b.String(), c)

	return err
}
