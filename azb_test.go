// -*- coding: utf-8 -*-

// Created on Thu Jun 12 11:01:25 AM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"testing"
)

func Test_constructQuery(t *testing.T) {
	type args struct {
		c ConfigFields
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{c: ConfigFields{Crontab: "", NVocab: 0, NExamples: 0, MinLen: 0, MaxLen: 0}}, want: "WITH a AS (SELECT UNNEST(STR_SPLIT_REGEX(本文, '\\.|。|？|\\?|!|！')) AS sentences FROM texts), b AS (SELECT sentences FROM a WHERE LENGTH(sentences) BETWEEN 0 AND 0), c AS (SELECT * AS vocab FROM read_csv('vocab.txt', header=false) ORDER BY RANDOM() LIMIT 100), d AS (SELECT c.vocab, b.sentences, ROW_NUMBER() OVER (PARTITION BY c.vocab ORDER BY RANDOM()) as rownum FROM b LEFT JOIN c ON (CONTAINS(b.sentences, c.vocab))) SELECT vocab, sentences FROM d WHERE rownum <= 100"},
		{name: "", args: args{c: ConfigFields{Crontab: "", NVocab: 0, NExamples: 0, MinLen: 0, MaxLen: 100}}, want: "WITH a AS (SELECT UNNEST(STR_SPLIT_REGEX(本文, '\\.|。|？|\\?|!|！')) AS sentences FROM texts), b AS (SELECT sentences FROM a WHERE LENGTH(sentences) BETWEEN 0 AND 100), c AS (SELECT * AS vocab FROM read_csv('vocab.txt', header=false) ORDER BY RANDOM() LIMIT 100), d AS (SELECT c.vocab, b.sentences, ROW_NUMBER() OVER (PARTITION BY c.vocab ORDER BY RANDOM()) as rownum FROM b LEFT JOIN c ON (CONTAINS(b.sentences, c.vocab))) SELECT vocab, sentences FROM d WHERE rownum <= 100"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructQuery(tt.args.c); got != tt.want {
				t.Errorf("constructQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
