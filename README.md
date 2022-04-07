# Assistant Bot

![Build](https://github.com/tsundata/assistant/workflows/Build/badge.svg)
![CodeQL](https://github.com/tsundata/assistant/workflows/CodeQL/badge.svg)
![Lint](https://github.com/tsundata/assistant/workflows/Lint/badge.svg)
[![codecov](https://codecov.io/gh/tsundata/assistant/branch/main/graph/badge.svg?token=ZDTxMN5H92)](https://codecov.io/gh/tsundata/assistant)
[![Go Report Card](https://goreportcard.com/badge/github.com/tsundata/assistant)](https://goreportcard.com/report/github.com/tsundata/assistant)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/tsundata/assistant)
![GitHub](https://img.shields.io/github/license/tsundata/assistant)

Assistant Bot is a workflow engine for chatbot

## Features

- Chat bot
- Message Publish/Subscribe Hub
- Message Cron, Trigger, Task, Pipeline
- Workflow Action ([Syntax](./docs/syntax.md))

## Architecture

<img src="./docs/architecture.png" alt="Architecture" align="center" width="100%" /> 

## Applications used

- etcd
- influx
- jaeger
- mysql
- rabbitmq
- redis

## Requirements

This project requires Go 1.17 or newer

## Installation

1. Install MySQL, Redis, influx, jaeger, rabbitmq, etcd

2. Import Configuration to etcd

3. Database migrate

4. Set Environment

   See [docs/env.md](/docs/env.md)

5. Build binary
   ```
   task build
   ```

6. Run App binary

# License

Assistant Bot is licensed under the [MIT license](https://opensource.org/licenses/MIT).
