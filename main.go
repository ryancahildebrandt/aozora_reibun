// -*- coding: utf-8 -*-

// Created on Thu Jun 12 10:59:39 AM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-co-op/gocron/v2"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
	_ "github.com/marcboeker/go-duckdb/v2"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	voc, err := ReadVocab()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("read vocab successfully %s", voc)

	c, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("read config successfully %+v\n", c)

	query := constructQuery(voc, c)
	db, err := sql.Open("duckdb", "./data/aozora_corpus.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sentences, err := getSentences(db, query)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("read sentences successfully")

	sq := enqueueSentences(sentences, voc)
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("enqueued sentences successfully")

	sch, err := gocron.NewScheduler()
	if err != nil {
		log.Fatal(err)
	}
	defer sch.Shutdown()

	_, err = sch.NewJob(
		gocron.CronJob(c.Crontab, false),
		gocron.NewTask(func() { job(sq, t, c) }),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("loaded cronjob successfully")

	err = SendEmail(fmt.Sprintf("<p>This email confirms that Aozora Reibun has started successfully with config %+v\n </p>", c), c)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("sent confirmation email successfully")

	log.Printf("starting scheduler")
	sch.Start()
	select {}
}
