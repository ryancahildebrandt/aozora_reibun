// -*- coding: utf-8 -*-

// Created on Sun Jun 22 09:34:16 PM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"slices"
	"testing"
)

func Test_sampleSentences(t *testing.T) {
	type args struct {
		s SentenceQueue
		m int
		n int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{name: "", args: args{s: SentenceQueue{"": []string{""}}, m: 0, n: 0}, want: []string{}, wantErr: true},
		{name: "", args: args{s: SentenceQueue{"": []string{""}}, m: 1, n: 0}, want: []string{}, wantErr: true},
		{name: "", args: args{s: SentenceQueue{"": []string{""}}, m: 0, n: 1}, want: []string{}, wantErr: true},
		{name: "", args: args{s: SentenceQueue{"a": []string{"a1", "a2"}, "b": []string{"b1", "b2"}, "c": []string{"c1", "c2"}}, m: 4, n: 1}, want: []string{}, wantErr: true},

		{name: "", args: args{s: SentenceQueue{"": []string{""}}, m: 1, n: 1}, want: []string{}, wantErr: false},
		{name: "", args: args{s: SentenceQueue{"a": []string{""}, "b": []string{""}}, m: 1, n: 1}, want: []string{""}, wantErr: false},
		{name: "", args: args{s: SentenceQueue{"a": []string{"a1", "a2"}, "b": []string{"b1", "b2"}}, m: 1, n: 1}, want: []string{}, wantErr: false},
		{name: "", args: args{s: SentenceQueue{"a": []string{"a1", "a2"}, "b": []string{"b1", "b2"}, "c": []string{"c1", "c2"}}, m: 1, n: 4}, want: []string{}, wantErr: false},
		{name: "", args: args{s: SentenceQueue{"a": []string{"a1", "a2"}, "b": []string{"b1", "b2"}, "c": []string{"c1", "c2"}}, m: 1, n: 1}, want: []string{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got, err := sampleSentences(tt.args.s, tt.args.m, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("sampleSentences() error = %v, wantErr %v", err, tt.wantErr)
			}
			for _, g := range got {
				if len(g) == 0 {
					continue
				}
				key := string(g[0])
				if !slices.Contains(tt.args.s[key], g) {
					t.Errorf("sampleSentences() got unexpected %v, want %v", g, got)
				}
			}
		})
	}
}
