// -*- coding: utf-8 -*-

// Created on Thu Jun 12 11:01:25 AM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

func constructQuery(c ConfigFields, v []string) string {
	return fmt.Sprintf("SELECT text FROM vtexts WHERE text MATCH '%s' AND LENGTH(text) BETWEEN %v AND %v ORDER BY RANDOM();", strings.Join(v, " OR "), c.MinLen, c.MaxLen)
}

func getSentences(db *sql.DB, q string, v []string) (SentenceQueue, error) {
	var (
		res SentenceQueue = make(map[string][]string)
		err error
		s   string
	)

	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&s)
		for _, word := range v {
			if strings.Contains(s, word) {
				res[word] = append(res[word], s)
			}
		}
	}

	return res, err
}
