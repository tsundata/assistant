# Assistant Bot

![Build](https://github.com/tsundata/assistant/workflows/Build/badge.svg)
![CodeQL](https://github.com/tsundata/assistant/workflows/CodeQL/badge.svg)
![Lint](https://github.com/tsundata/assistant/workflows/Lint/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/tsundata/assistant)](https://goreportcard.com/report/github.com/tsundata/assistant)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/tsundata/assistant)
![GitHub](https://img.shields.io/github/license/tsundata/assistant)

Assistant Bot is a workflow engine for chatbot

## Features

- Chat bot
- Message Publish/Subscribe Hub
- Message Cron, Trigger, Task
- Workflow Action ([Syntax](./doc/action_syntax.md))
- Workflow Script ([Syntax](./doc/script_syntax.md))

## Architecture

<img src="./doc/architecture.png" alt="Architecture" align="center" width="100%" /> 

## Applications used

- Github
- Pocket
- Pushover
- Dropbox
- Slack
- Rollbar
- Email

## Requirements

This project requires Go 1.16 or newer

## Installation

- install MySQL, Redis, etcd, influx
- install jaeger (optional)

- import sql files

- Configuration
```
cp ./configs/*.yml.example ./config/*.yml
```

- Build binary
```
make build
```

- Run
```
./dist/* -f configs/*.yml
```

# License

Assistant Bot is licensed under the [MIT license](https://opensource.org/licenses/MIT).