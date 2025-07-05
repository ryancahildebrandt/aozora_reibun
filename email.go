// -*- coding: utf-8 -*-

// Created on Thu Jun 12 11:00:47 AM EDT 2025
// author: Ryan Hildebrandt, github.com/ryancahildebrandt

package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/joho/godotenv"
)

func loadCredentials() (string, string, error) {
	var (
		err error
		f   string
		p   string
	)

	err = godotenv.Load()
	if err != nil {
		return f, p, err
	}
	f = os.Getenv("FROM")
	p = os.Getenv("PASSWORD")

	return f, p, err
}

func SendEmail(v []string, b string, c ConfigFields) error {
	var (
		from     string
		password string
		err      error
		auth     smtp.Auth
		msg      []byte
	)

	from, password, err = loadCredentials()
	if err != nil {
		return err
	}

	auth = smtp.PlainAuth("", from, password, "smtp.gmail.com")
	msg = fmt.Appendf([]byte{}, "%s", fmt.Sprintf("Subject: 青空例文 %s %s\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s", time.Now().Local().Format("2006.01.02"), v, b))

	err = smtp.SendMail("smtp.gmail.com:587", auth, from, c.Recipients, msg)
	if err != nil {
		return err
	}

	log.Printf("Successfully sent mail to %s", c.Recipients)
	return err
}

func renderEmail(s []SentenceLookups, b *strings.Builder) (strings.Builder, error) {
	tmpl := template.Must(template.ParseGlob("./templates/*.gohtml"))
	err := tmpl.ExecuteTemplate(b, "main", s)
	return *b, err
}
