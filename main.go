// -*- coding: utf-8 -*-

// Created on Thu Jun 12 10:59:39 AM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-co-op/gocron/v2"
	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
	_ "modernc.org/sqlite"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	c, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("read config successfully %+v\n", c)

	file, err := os.Open("./vocab.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	vocab := []string{}
	for scanner.Scan() {
		vocab = append(vocab, strings.TrimSpace(scanner.Text()))
	}
	log.Printf("read vocab successfully")

	dbfile := "./data/aozora_corpus.db"
	if c.UseTatoeba {
		dbfile = "./data/tatoeba.db"
	}
	db, err := sql.Open("sqlite", dbfile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query := constructQuery(c, vocab)
	sq, err := getSentences(db, query, vocab)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("read sentences successfully")

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

	err = SendEmail([]string{"Confirmation"}, fmt.Sprintf("<p>This email confirms that Aozora Reibun has started successfully with config %+v\n </p>", c), c)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("sent confirmation email successfully")

	log.Printf("starting scheduler")
	sch.Start()
	select {}
}
