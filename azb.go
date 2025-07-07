// -*- coding: utf-8 -*-

// Created on Thu Jun 12 11:01:25 AM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"database/sql"
	"fmt"
	"log"
)

func constructQuery(c ConfigFields) string {
	return fmt.Sprintf("WITH a AS (SELECT UNNEST(STR_SPLIT_REGEX(本文, '\\.|。|？|\\?|!|！')) AS sentences FROM texts), b AS (SELECT sentences FROM a WHERE LENGTH(sentences) BETWEEN %v AND %v), c AS (SELECT * AS vocab FROM read_csv('vocab.txt', header=false) ORDER BY RANDOM() LIMIT 100), d AS (SELECT c.vocab, b.sentences, ROW_NUMBER() OVER (PARTITION BY c.vocab ORDER BY RANDOM()) as rownum FROM b LEFT JOIN c ON (CONTAINS(b.sentences, c.vocab))) SELECT vocab, sentences FROM d WHERE rownum <= 100", c.MinLen, c.MaxLen)
}

func getSentences(db *sql.DB, q string) (SentenceQueue, error) {
	var (
		res SentenceQueue = make(map[string][]string)
		err error
		v   string
		s   string
	)

	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&v, &s)
		res[v] = append(res[v], s)
	}

	return res, err
}
