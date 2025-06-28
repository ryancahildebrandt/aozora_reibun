// -*- coding: utf-8 -*-

// Created on Thu Jun 12 11:01:25 AM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"testing"
)

func Test_constructQuery(t *testing.T) {
	type args struct {
		v []string
		c ConfigFields
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{v: []string{}, c: ConfigFields{Crontab: "", NVocab: 0, NExamples: 0, MinLen: 0, MaxLen: 0}}, want: "WITH s AS (SELECT UNNEST(STR_SPLIT_REGEX(本文, '\\.|。|？|\\?|!|！')) AS sentences FROM texts) SELECT sentences FROM s WHERE CHAR_LENGTH(sentences) BETWEEN 0 AND 0 LIMIT 100000"},
		{name: "", args: args{v: []string{""}, c: ConfigFields{Crontab: "", NVocab: 0, NExamples: 0, MinLen: 0, MaxLen: 0}}, want: "WITH s AS (SELECT UNNEST(STR_SPLIT_REGEX(本文, '\\.|。|？|\\?|!|！')) AS sentences FROM texts) SELECT sentences FROM s WHERE CHAR_LENGTH(sentences) BETWEEN 0 AND 0 AND CONTAINS(sentences, '') LIMIT 100000"},
		{name: "", args: args{v: []string{"a", ""}, c: ConfigFields{Crontab: "", NVocab: 0, NExamples: 0, MinLen: 0, MaxLen: 0}}, want: "WITH s AS (SELECT UNNEST(STR_SPLIT_REGEX(本文, '\\.|。|？|\\?|!|！')) AS sentences FROM texts) SELECT sentences FROM s WHERE CHAR_LENGTH(sentences) BETWEEN 0 AND 0 AND CONTAINS(sentences, 'a') OR CONTAINS(sentences, '') LIMIT 100000"},
		{name: "", args: args{v: []string{"a", "b", "c", "d"}, c: ConfigFields{Crontab: "", NVocab: 0, NExamples: 0, MinLen: 0, MaxLen: 0}}, want: "WITH s AS (SELECT UNNEST(STR_SPLIT_REGEX(本文, '\\.|。|？|\\?|!|！')) AS sentences FROM texts) SELECT sentences FROM s WHERE CHAR_LENGTH(sentences) BETWEEN 0 AND 0 AND CONTAINS(sentences, 'a') OR CONTAINS(sentences, 'b') OR CONTAINS(sentences, 'c') OR CONTAINS(sentences, 'd') LIMIT 100000"},

		{name: "", args: args{v: []string{}, c: ConfigFields{Crontab: "", NVocab: 0, NExamples: 0, MinLen: 0, MaxLen: 100}}, want: "WITH s AS (SELECT UNNEST(STR_SPLIT_REGEX(本文, '\\.|。|？|\\?|!|！')) AS sentences FROM texts) SELECT sentences FROM s WHERE CHAR_LENGTH(sentences) BETWEEN 0 AND 100 LIMIT 100000"},
		{name: "", args: args{v: []string{""}, c: ConfigFields{Crontab: "", NVocab: 0, NExamples: 0, MinLen: 0, MaxLen: 100}}, want: "WITH s AS (SELECT UNNEST(STR_SPLIT_REGEX(本文, '\\.|。|？|\\?|!|！')) AS sentences FROM texts) SELECT sentences FROM s WHERE CHAR_LENGTH(sentences) BETWEEN 0 AND 100 AND CONTAINS(sentences, '') LIMIT 100000"},
		{name: "", args: args{v: []string{"a", ""}, c: ConfigFields{Crontab: "", NVocab: 0, NExamples: 0, MinLen: 0, MaxLen: 100}}, want: "WITH s AS (SELECT UNNEST(STR_SPLIT_REGEX(本文, '\\.|。|？|\\?|!|！')) AS sentences FROM texts) SELECT sentences FROM s WHERE CHAR_LENGTH(sentences) BETWEEN 0 AND 100 AND CONTAINS(sentences, 'a') OR CONTAINS(sentences, '') LIMIT 100000"},
		{name: "", args: args{v: []string{"a", "b", "c", "d"}, c: ConfigFields{Crontab: "", NVocab: 0, NExamples: 0, MinLen: 0, MaxLen: 100}}, want: "WITH s AS (SELECT UNNEST(STR_SPLIT_REGEX(本文, '\\.|。|？|\\?|!|！')) AS sentences FROM texts) SELECT sentences FROM s WHERE CHAR_LENGTH(sentences) BETWEEN 0 AND 100 AND CONTAINS(sentences, 'a') OR CONTAINS(sentences, 'b') OR CONTAINS(sentences, 'c') OR CONTAINS(sentences, 'd') LIMIT 100000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructQuery(tt.args.v, tt.args.c); got != tt.want {
				t.Errorf("constructQuery(%v) = %v, want %v", tt.args.v, got, tt.want)
			}
		})
	}
}
