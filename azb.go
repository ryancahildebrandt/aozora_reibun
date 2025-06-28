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

func constructQuery(v []string, c ConfigFields) string {
	var b strings.Builder

	b.WriteString("WITH s AS (SELECT UNNEST(STR_SPLIT_REGEX(本文, '\\.|。|？|\\?|!|！')) AS sentences FROM texts) SELECT sentences FROM s WHERE CHAR_LENGTH(sentences) BETWEEN ")
	b.WriteString(fmt.Sprintf("%v AND %v", c.MinLen, c.MaxLen))
	if len(v) != 0 {
		b.WriteString(" AND")
	}
	for i, w := range v {
		b.WriteString(fmt.Sprintf(" CONTAINS(sentences, '%s')", w))
		if i != len(v)-1 {
			b.WriteString(" OR")
		}
	}
	b.WriteString(" LIMIT 100000")

	return b.String()
}

func getSentences(db *sql.DB, q string) ([]string, error) {
	var res []string
	var err error
	var text string

	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&text)
		res = append(res, text)
	}

	return res, err
}
