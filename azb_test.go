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
		v []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{c: ConfigFields{Crontab: "", NVocab: 0, NExamples: 0, MinLen: 0, MaxLen: 0}, v: []string{"a", "b", "c"}}, want: "SELECT text FROM vtexts WHERE text MATCH 'a OR b OR c' AND LENGTH(text) BETWEEN 0 AND 0 ORDER BY RANDOM();"},
		{name: "", args: args{c: ConfigFields{Crontab: "", NVocab: 0, NExamples: 0, MinLen: 0, MaxLen: 100}, v: []string{""}}, want: "SELECT text FROM vtexts WHERE text MATCH '' AND LENGTH(text) BETWEEN 0 AND 100 ORDER BY RANDOM();"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructQuery(tt.args.c, tt.args.v); got != tt.want {
				t.Errorf("constructQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
