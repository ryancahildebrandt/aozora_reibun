# 青空例文

## _Targeted Example Sentences from the Aozora Bunko Corpus_

---

[![Open in gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/ryancahildebrandt/aozora_reibun)
[![This project contains 0% LLM-generated content](https://brainmade.org/88x31-dark.png)](https://brainmade.org/)

## Purpose

This is a pretty barebones language practice service that:

- Reads in a provided vocabulary list
- Samples sentences from the Aozora Bunko Corpus containing the targeted vocabulary words
- Looks up words and kanji in the sentences for readings, definitions, etc.
- Formats the sentences and lookup information into a simple HTML email
- Sends the email to your inbox with specified frequency

---

## Dataset

The dataset used for the current project was pulled from the following:

- [Aozora Bunko Corpus](https://www.kaggle.com/datasets/ryancahildebrandt/azbcorpus)
- [Jotoba API](https://jotoba.de)

---

## Usage

Setup is handled via a few key files:

- vocab.txt, for the vocab you want to study/get sentences for. Empty file is included in the repo as .vocab.txt, so change the name to vocab.txt on your first run.

```
昨日
すき焼き
を
食べました
...
```

- config.json, for setting email frequency and content. Empty file is included in the repo as .config.json, so change the name to config.json on your first run.

```json
{
  "crontab": "* * * * *", // Cron expression for email frequency
  "n_vocab": 2, // number of vocab words per email
  "n_examples": 2, // number of example sentences per vocab word per email
  "min_len": 0, // shortest allowed sentence
  "max_len": 100, // longest allowed sentence
  "recipients": ["<recipient email 1>", "<recipient email 2>", ...] //
}
```

- .env, for your super secret email and incredibly secure password

```
FROM=<email login>
PASSWORD=<email password>
```

Notes:

- I recommend running the go binary (./aozora_reibun) as a systemd service using the provided example service file.
- It will re-query the database on each restart of the service.
- You can update the vocab and config at any time but you'll need to restart the service to have the changes reflected.
- Email limits and security concerns will come down to and be handled by your email provider/service.
