# demas/quzx-crawler

# Introduction

News fetcher for QUZX. Collect new posts from RSS, HackerNews, Twitter, YouTube, VK, Reddit and StackOverflow.

# Installation

Automated builds of the image available and is the recommended method of installation:

```bash
docker pull demas1251/quzx-crawler
```

Alternatively you can build the image locally.

```bash
docker build -t demas1251/quzx-crawler .
```

# Quick Start

Set the environment variables: `DBUSER`, `DBPASS`, `DBHOST`, `DBPORT` and `DBNAME` to establish connection to PostgreSQL database.
Set the environment variable `MONGODB` to establish connection to MongoDb.
