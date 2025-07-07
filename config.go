// -*- coding: utf-8 -*-

// Created on Thu Jun 12 11:00:17 AM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"encoding/json"
	"io"
	"os"
)

type ConfigFields struct {
	Crontab    string   `json:"crontab"`
	NVocab     int      `json:"n_vocab"`
	NExamples  int      `json:"n_examples"`
	MinLen     int      `json:"min_len"`
	MaxLen     int      `json:"max_len"`
	Recipients []string `json:"recipients"`
}

func ReadConfig() (ConfigFields, error) {
	var (
		fields ConfigFields
		file   io.Reader
		err    error
	)

	file, err = os.Open("./config.json")
	if err != nil {
		return fields, err
	}
	err = json.NewDecoder(file).Decode(&fields)

	return fields, err
}
