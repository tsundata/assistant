
<a name="v0.2.1"></a>
## [v0.2.1](https://github.com/tsundata/assistant/compare/v0.2...v0.2.1) (2021-12-24)

### Bug Fixes

* go mod version
* sql uuid field length
* wire gen
* create message service
* create message rpc

### Code Refactoring

* bot
* chatbot handle
* repository return
* group rpc
* proto
* sql
* bot dir
* resty SetBaseURL
* todo service get user id
* model with user_id
* protoc grpc require_unimplemented_servers=false
* create group
*  message repository
* group service
* user, chatbot service
* rollbar error
* remove agent

### Features

* health check
* trace request id
* grpc auth interceptor
* group tag repository
* chatbot group setting repository
* message sequence
* global transaction
* update go 1.17
* grpc metadata auth id
* buf
* global lock
* id service
* global id
* swagger github action
* swagger docs
* user login rpc

### Pull Requests

* Merge pull request [#155](https://github.com/tsundata/assistant/issues/155) from tsundata/dependabot/go_modules/google.golang.org/grpc-1.43.0
* Merge pull request [#154](https://github.com/tsundata/assistant/issues/154) from tsundata/dependabot/go_modules/github.com/spf13/viper-1.10.1
* Merge pull request [#156](https://github.com/tsundata/assistant/issues/156) from tsundata/dependabot/go_modules/github.com/spf13/cobra-1.3.0
* Merge pull request [#153](https://github.com/tsundata/assistant/issues/153) from tsundata/dependabot/go_modules/github.com/spf13/viper-1.10.0
* Merge pull request [#152](https://github.com/tsundata/assistant/issues/152) from tsundata/dependabot/go_modules/github.com/mozillazg/go-pinyin-0.19.0
* Merge pull request [#151](https://github.com/tsundata/assistant/issues/151) from tsundata/dependabot/go_modules/github.com/uber/jaeger-client-go-2.30.0incompatible
* Merge pull request [#149](https://github.com/tsundata/assistant/issues/149) from tsundata/dependabot/go_modules/github.com/gofiber/websocket/v2-2.0.14
* Merge pull request [#150](https://github.com/tsundata/assistant/issues/150) from tsundata/dependabot/go_modules/github.com/newrelic/go-agent/v3-3.15.2
* Merge pull request [#148](https://github.com/tsundata/assistant/issues/148) from tsundata/dependabot/go_modules/github.com/gofiber/fiber/v2-2.23.0
* Merge pull request [#147](https://github.com/tsundata/assistant/issues/147) from tsundata/dependabot/go_modules/github.com/golang-jwt/jwt/v4-4.2.0
* Merge pull request [#146](https://github.com/tsundata/assistant/issues/146) from tsundata/dependabot/go_modules/github.com/tidwall/gjson-1.12.1
* Merge pull request [#144](https://github.com/tsundata/assistant/issues/144) from tsundata/dependabot/go_modules/gorm.io/driver/mysql-1.2.1
* Merge pull request [#143](https://github.com/tsundata/assistant/issues/143) from tsundata/dependabot/go_modules/github.com/swaggo/swag-1.7.6
* Merge pull request [#140](https://github.com/tsundata/assistant/issues/140) from tsundata/dependabot/go_modules/github.com/swaggo/swag-1.7.5
* Merge pull request [#142](https://github.com/tsundata/assistant/issues/142) from tsundata/dependabot/go_modules/github.com/influxdata/influxdb-client-go/v2-2.6.0
* Merge pull request [#141](https://github.com/tsundata/assistant/issues/141) from tsundata/dependabot/go_modules/github.com/tidwall/gjson-1.12.0
* Merge pull request [#139](https://github.com/tsundata/assistant/issues/139) from tsundata/dependabot/go_modules/gorm.io/gorm-1.22.3
* Merge pull request [#138](https://github.com/tsundata/assistant/issues/138) from tsundata/dependabot/go_modules/github.com/gofiber/websocket/v2-2.0.13
* Merge pull request [#137](https://github.com/tsundata/assistant/issues/137) from tsundata/dependabot/go_modules/github.com/gofiber/fiber/v2-2.22.0
* Merge pull request [#136](https://github.com/tsundata/assistant/issues/136) from tsundata/dependabot/go_modules/gorm.io/driver/mysql-1.2.0
* Merge pull request [#135](https://github.com/tsundata/assistant/issues/135) from tsundata/dependabot/go_modules/github.com/yuin/goldmark-1.4.4
* Merge pull request [#134](https://github.com/tsundata/assistant/issues/134) from tsundata/dependabot/go_modules/github.com/nats-io/nats.go-1.13.0
* Merge pull request [#133](https://github.com/tsundata/assistant/issues/133) from tsundata/dependabot/go_modules/github.com/yuin/goldmark-1.4.3
* Merge pull request [#132](https://github.com/tsundata/assistant/issues/132) from tsundata/dependabot/go_modules/gorm.io/gorm-1.22.2
* Merge pull request [#131](https://github.com/tsundata/assistant/issues/131) from tsundata/dependabot/go_modules/github.com/newrelic/go-agent/v3-3.15.1
* Merge pull request [#130](https://github.com/tsundata/assistant/issues/130) from tsundata/dependabot/go_modules/github.com/nats-io/nats.go-1.13.0
* Merge pull request [#129](https://github.com/tsundata/assistant/issues/129) from tsundata/dependabot/go_modules/github.com/golang-jwt/jwt/v4-4.1.0
* Merge pull request [#128](https://github.com/tsundata/assistant/issues/128) from tsundata/dependabot/go_modules/github.com/swaggo/swag-1.7.4
* Merge pull request [#127](https://github.com/tsundata/assistant/issues/127) from tsundata/dependabot/go_modules/github.com/go-redis/redis/v8-8.11.4
* Merge pull request [#126](https://github.com/tsundata/assistant/issues/126) from tsundata/dependabot/go_modules/gorm.io/driver/mysql-1.1.3
* Merge pull request [#125](https://github.com/tsundata/assistant/issues/125) from tsundata/dependabot/go_modules/github.com/PuerkitoBio/goquery-1.8.0
* Merge pull request [#123](https://github.com/tsundata/assistant/issues/123) from tsundata/dependabot/go_modules/github.com/gofiber/websocket/v2-2.0.12
* Merge pull request [#122](https://github.com/tsundata/assistant/issues/122) from tsundata/dependabot/go_modules/github.com/go-ego/gse-0.69.15
* Merge pull request [#120](https://github.com/tsundata/assistant/issues/120) from tsundata/dependabot/go_modules/google.golang.org/grpc-1.42.0
* Merge pull request [#124](https://github.com/tsundata/assistant/issues/124) from tsundata/dependabot/go_modules/github.com/gofiber/fiber/v2-2.21.0
* Merge pull request [#118](https://github.com/tsundata/assistant/issues/118) from tsundata/dependabot/go_modules/github.com/go-resty/resty/v2-2.7.0
* Merge pull request [#115](https://github.com/tsundata/assistant/issues/115) from tsundata/dependabot/go_modules/github.com/rollbar/rollbar-go-1.4.2
* Merge pull request [#117](https://github.com/tsundata/assistant/issues/117) from tsundata/dependabot/go_modules/github.com/arsmn/fiber-swagger/v2-2.20.0
* Merge pull request [#116](https://github.com/tsundata/assistant/issues/116) from tsundata/dependabot/go_modules/go.etcd.io/etcd/client/v3-3.5.1
* Merge pull request [#119](https://github.com/tsundata/assistant/issues/119) from tsundata/dependabot/go_modules/github.com/tidwall/gjson-1.11.0


<a name="v0.2"></a>
## [v0.2](https://github.com/tsundata/assistant/compare/v0.1.1...v0.2) (2021-09-26)

### Bug Fixes

* ineffectual assignment to err

### Code Refactoring

* repository
* uuid
* print log
* ws router

### Features

* ws handle message
* room chat
* ws controller
* push notification
* user auth jwt

### Pull Requests

* Merge pull request [#113](https://github.com/tsundata/assistant/issues/113) from tsundata/dependabot/go_modules/github.com/nats-io/nats.go-1.12.3
* Merge pull request [#110](https://github.com/tsundata/assistant/issues/110) from tsundata/dependabot/go_modules/github.com/gofiber/fiber/v2-2.19.0
* Merge pull request [#111](https://github.com/tsundata/assistant/issues/111) from tsundata/dependabot/go_modules/github.com/spf13/viper-1.9.0
* Merge pull request [#112](https://github.com/tsundata/assistant/issues/112) from tsundata/dependabot/go_modules/github.com/nats-io/nats.go-1.12.2
* Merge pull request [#109](https://github.com/tsundata/assistant/issues/109) from tsundata/dependabot/go_modules/github.com/go-ego/gse-0.69.5
* Merge pull request [#108](https://github.com/tsundata/assistant/issues/108) from tsundata/dependabot/go_modules/github.com/influxdata/influxdb-client-go/v2-2.5.1
* Merge pull request [#107](https://github.com/tsundata/assistant/issues/107) from tsundata/dependabot/go_modules/github.com/slack-go/slack-0.9.5
* Merge pull request [#106](https://github.com/tsundata/assistant/issues/106) from tsundata/dependabot/go_modules/github.com/tidwall/gjson-1.9.1
* Merge pull request [#105](https://github.com/tsundata/assistant/issues/105) from tsundata/dependabot/go_modules/github.com/yuin/goldmark-1.4.1
* Merge pull request [#103](https://github.com/tsundata/assistant/issues/103) from tsundata/dependabot/go_modules/github.com/valyala/fasthttp-1.30.0
* Merge pull request [#104](https://github.com/tsundata/assistant/issues/104) from tsundata/dependabot/go_modules/github.com/go-ego/gse-0.69.3
* Merge pull request [#102](https://github.com/tsundata/assistant/issues/102) from tsundata/dependabot/go_modules/go.uber.org/zap-1.19.1


<a name="v0.1.1"></a>
## [v0.1.1](https://github.com/tsundata/assistant/compare/v0.1...v0.1.1) (2021-09-08)

### Bug Fixes

* github pocket id filter


<a name="v0.1"></a>
## [v0.1](https://github.com/tsundata/assistant/compare/v0.0.27...v0.1) (2021-09-08)

### Bug Fixes

* gse new
* file path provided as taint input
* weak cryptographic primitive
* weak random generator
* message trigger
* delete workflow message
* svc addr select
* use crypto/rand
* duplicate struct tag "json"
* service connection refused

### Code Refactoring

* list cron rpc
* remove subscribe service
* cron log

### Features

* command support string
* todo list command
* webhook list command
* sort stats result
* update go version 1.17
* fund chart
* fund, stock detail service
* doctorxiong vendor api
* org command
* chart
* delete message command
* test command
* middle tag service
* org repository
* org service
* cron command
* cloudcone billing

### Pull Requests

* Merge pull request [#101](https://github.com/tsundata/assistant/issues/101) from tsundata/dependabot/go_modules/github.com/nats-io/nats.go-1.12.1
* Merge pull request [#99](https://github.com/tsundata/assistant/issues/99) from tsundata/dependabot/go_modules/github.com/go-ego/gse-0.69.2
* Merge pull request [#96](https://github.com/tsundata/assistant/issues/96) from tsundata/dependabot/go_modules/github.com/tidwall/gjson-1.9.0
* Merge pull request [#100](https://github.com/tsundata/assistant/issues/100) from tsundata/dependabot/go_modules/github.com/newrelic/go-agent/v3-3.15.0
* Merge pull request [#98](https://github.com/tsundata/assistant/issues/98) from tsundata/dependabot/go_modules/github.com/shirou/gopsutil/v3-3.21.8
* Merge pull request [#94](https://github.com/tsundata/assistant/issues/94) from tsundata/dependabot/go_modules/github.com/nats-io/nats.go-1.12.0
* Merge pull request [#95](https://github.com/tsundata/assistant/issues/95) from tsundata/dependabot/go_modules/github.com/gofiber/fiber/v2-2.18.0
* Merge pull request [#93](https://github.com/tsundata/assistant/issues/93) from tsundata/dependabot/go_modules/github.com/influxdata/influxdb-client-go/v2-2.5.0
* Merge pull request [#91](https://github.com/tsundata/assistant/issues/91) from tsundata/dependabot/go_modules/google.golang.org/grpc-1.40.0
* Merge pull request [#92](https://github.com/tsundata/assistant/issues/92) from tsundata/dependabot/go_modules/github.com/valyala/fasthttp-1.29.0
* Merge pull request [#89](https://github.com/tsundata/assistant/issues/89) from tsundata/dependabot/go_modules/github.com/go-redis/redis/v8-8.11.3
* Merge pull request [#90](https://github.com/tsundata/assistant/issues/90) from tsundata/dependabot/go_modules/github.com/gofiber/fiber/v2-2.17.0
* Merge pull request [#88](https://github.com/tsundata/assistant/issues/88) from tsundata/dependabot/go_modules/go.uber.org/zap-1.19.0
* Merge pull request [#87](https://github.com/tsundata/assistant/issues/87) from tsundata/dependabot/go_modules/google.golang.org/grpc-1.39.1
* Merge pull request [#86](https://github.com/tsundata/assistant/issues/86) from tsundata/dependabot/go_modules/github.com/go-redis/redis/v8-8.11.2
* Merge pull request [#84](https://github.com/tsundata/assistant/issues/84) from tsundata/dependabot/go_modules/github.com/shirou/gopsutil/v3-3.21.7
* Merge pull request [#85](https://github.com/tsundata/assistant/issues/85) from tsundata/dependabot/go_modules/github.com/gofiber/fiber/v2-2.16.0
* Merge pull request [#83](https://github.com/tsundata/assistant/issues/83) from tsundata/dependabot/go_modules/github.com/slack-go/slack-0.9.4
* Merge pull request [#82](https://github.com/tsundata/assistant/issues/82) from tsundata/dependabot/go_modules/github.com/go-redis/redis/v8-8.11.1


<a name="v0.0.27"></a>
## [v0.0.27](https://github.com/tsundata/assistant/compare/v0.0.26...v0.0.27) (2021-07-27)

### Bug Fixes

* subscribe register
* gateway request limit number
* json.Unmarshal
* GetAvailableApp
* sdk error message
* webhook action trigger

### Code Refactoring

* remove consul, update go.mod
* config use etcd

### Features

* etcd config file
* update zap, gjson, grpc, slack, gopsutil version


<a name="v0.0.26"></a>
## [v0.0.26](https://github.com/tsundata/assistant/compare/v0.0.25...v0.0.26) (2021-07-22)

### Code Refactoring

* rpc discovery, enum app


<a name="v0.0.25"></a>
## [v0.0.25](https://github.com/tsundata/assistant/compare/v0.0.24...v0.0.25) (2021-07-22)

### Bug Fixes

* parse command check empty string
* workflow trigger
* sdk
* config watch
* time tick
* go.sum
* complete todo sql

### Code Refactoring

* rule parse func
* regex bot rule
* task enum
* remove mysql
* middle repository
* message repository
* user repository
* todo repository
* svc pb
* message pb
* middle pb
* workflow pb
* user pb
* todo pb
* model dir
* use rqlite instead of mysql
* bus, log
* context
* wire_gen
* config watch, rpc client timeout
* todo remind check
* rename args
* created_at column
* get role image

### Features

* tag doc
* newrelic
* grpc log ServerInterceptor
* app id
* cli app
* user trigger
* user rpc service
* chatbot service
* filesystem
* subscribe send with channel
* send message with channel
* todo remind


<a name="v0.0.24"></a>
## [v0.0.24](https://github.com/tsundata/assistant/compare/v0.0.23...v0.0.24) (2021-06-29)

### Bug Fixes

* clasifier load rule
* create message with default time
* go.sum
* gateway, web router
* wire gen
* go mod
* go.sum
* gateway, web router
* wire gen
* go mod
* lint
* lint issues

### Code Refactoring

* todo Complete
* rpc client timeout
* rulebot load options
* rpc client
* rpc resolver
* grpc discovery
* context
* logger
* spider read remote config
* rpc client name
* app id, app config
* rulebot.Context
* remove worker service
* rename transports package
* rename tasks package
* rename components package
* rename controllers package
* rename rules package
* rename utils package
* context
* logger
* spider read remote config
* rpc client name
* app id, app config
* rulebot.Context
* remove worker service
* rename transports package
* rename tasks package
* rename components package
* rename controllers package
* rename rules package
* rename utils package

### Features

* classifier do rule
* classifier load rule
* nlp segmentation
* nlp pinyin conversion
* nlp service
* AppConfig get/set method
* send message event
* role rpc
* todo rpc
* debug event
* cleassifier
* nlp segmentation
* nlp pinyin conversion
* nlp service
* AppConfig get/set method
* send message event
* role rpc
* todo rpc
* debug event
* consul env
* consul discovery
* consul
* todo app
* user app
* finance app
* event listener
* event bus
* nats
* middleware dir
* gateway sdk
* push ghcr.io image


<a name="v0.0.23"></a>
## [v0.0.23](https://github.com/tsundata/assistant/compare/v0.0.22...v0.0.23) (2021-06-03)

### Bug Fixes

* build image tag


<a name="v0.0.22"></a>
## [v0.0.22](https://github.com/tsundata/assistant/compare/v0.0.21...v0.0.22) (2021-06-02)

### Bug Fixes

* Dockerfile
* Dockerfile
* agent cmd
* lint
* run web action

### Features

* gateway k8s
* config center
* workflow repository
* middle repository
* message repository
* storage service
* webhook auth
* amqp
* rabbitmq docker-compose
* stats
* telegram incoming

### Pull Requests

* Merge pull request [#44](https://github.com/tsundata/assistant/issues/44) from tsundata/dependabot/go_modules/github.com/uber/jaeger-client-go-2.27.0incompatible
* Merge pull request [#41](https://github.com/tsundata/assistant/issues/41) from tsundata/dependabot/go_modules/github.com/uber/jaeger-client-go-2.26.0incompatible
* Merge pull request [#40](https://github.com/tsundata/assistant/issues/40) from tsundata/dependabot/go_modules/github.com/slack-go/slack-0.9.0
* Merge pull request [#39](https://github.com/tsundata/assistant/issues/39) from tsundata/dependabot/go_modules/github.com/go-redis/redis/v8-8.8.2


<a name="v0.0.21"></a>
## [v0.0.21](https://github.com/tsundata/assistant/compare/v0.0.20...v0.0.21) (2021-03-24)

### Bug Fixes

* subscribe name
* create script url

### Features

* cloudflare cron
* env opcode
* secret opcode
* pipeline, stage


<a name="v0.0.20"></a>
## [v0.0.20](https://github.com/tsundata/assistant/compare/v0.0.19...v0.0.20) (2021-03-20)

### Bug Fixes

* rand int security

### Features

* url trigger
*  email trigger
* todo, issue tag trigger
* wire di
* tag trigger
* delay task


<a name="v0.0.19"></a>
## [v0.0.19](https://github.com/tsundata/assistant/compare/v0.0.18...v0.0.19) (2021-03-17)

### Bug Fixes

* spider todo
* action parser eof error

### Features

* trigger
* dedupe opcode
* query opcode
* message opcode
* opcode doc
* status opcode
* profile duration function
* set opcode, if opcode, else opcode
* json opcode

### Pull Requests

* Merge pull request [#20](https://github.com/tsundata/assistant/issues/20) from tsundata/dependabot/go_modules/github.com/gofiber/fiber/v2-2.6.0


<a name="v0.0.18"></a>
## [v0.0.18](https://github.com/tsundata/assistant/compare/v0.0.17...v0.0.18) (2021-03-15)

### Bug Fixes

* lint

### Features

* round function
* deubg action
* task action
* delete workflow message
* cron action
* webhook action
* action semantic analyzer


<a name="v0.0.17"></a>
## [v0.0.17](https://github.com/tsundata/assistant/compare/v0.0.16...v0.0.17) (2021-03-11)

### Features

* backup
* dropbox oauth
* rollbar interceptor
* rollbar, logger
* machinery
* task, worker
* cache auth
* fiber http server
* ping

### Pull Requests

* Merge pull request [#19](https://github.com/tsundata/assistant/issues/19) from tsundata/dependabot/go_modules/github.com/go-redis/redis/v8-8.7.1


<a name="v0.0.16"></a>
## [v0.0.16](https://github.com/tsundata/assistant/compare/v0.0.15...v0.0.16) (2021-03-03)

### Bug Fixes

* lint

### Features

* golang metrics
* run action
* action interpreter


<a name="v0.0.15"></a>
## [v0.0.15](https://github.com/tsundata/assistant/compare/v0.0.14...v0.0.15) (2021-02-26)

### Bug Fixes

* lint
* lint
* extractUUID func

### Features

* interpreter debug
* script ui
* run script, pushover node
* architecture image
* github fetch starred
* github oauth
* cron filter and send
* fetch pocket
* pocket oauth
* pocket agent
* cloc github action
* docker-compose mysql

### Pull Requests

* Merge pull request [#14](https://github.com/tsundata/assistant/issues/14) from tsundata/dependabot/go_modules/github.com/valyala/fasthttp-1.21.0
* Merge pull request [#13](https://github.com/tsundata/assistant/issues/13) from tsundata/dependabot/go_modules/github.com/slack-go/slack-0.8.1
* Merge pull request [#11](https://github.com/tsundata/assistant/issues/11) from tsundata/dependabot/go_modules/github.com/yuin/goldmark-1.3.2


<a name="v0.0.14"></a>
## [v0.0.14](https://github.com/tsundata/assistant/compare/v0.0.13...v0.0.14) (2021-02-07)

### Features

* show version


<a name="v0.0.13"></a>
## [v0.0.13](https://github.com/tsundata/assistant/compare/v0.0.12...v0.0.13) (2021-02-07)

### Bug Fixes

* goreleaser.yml ldflags
* goreleaser.yml
* goreleaser.yml
* goreleaser action
* build script path

### Features

* bin version
* credentials create js
* webhook, execute node

### Pull Requests

* Merge pull request [#10](https://github.com/tsundata/assistant/issues/10) from tsundata/dependabot/go_modules/github.com/go-redis/redis/v8-8.5.0


<a name="v0.0.12"></a>
## [v0.0.12](https://github.com/tsundata/assistant/compare/v0.0.11...v0.0.12) (2021-02-02)

### Features

* markdown convert


<a name="v0.0.11"></a>
## [v0.0.11](https://github.com/tsundata/assistant/compare/v0.0.10...v0.0.11) (2021-02-02)

### Features

* etcd auth

### Pull Requests

* Merge pull request [#8](https://github.com/tsundata/assistant/issues/8) from tsundata/dependabot/go_modules/github.com/shirou/gopsutil/v3-3.21.1
* Merge pull request [#9](https://github.com/tsundata/assistant/issues/9) from tsundata/dependabot/go_modules/github.com/tidwall/gjson-1.6.8


<a name="v0.0.10"></a>
## [v0.0.10](https://github.com/tsundata/assistant/compare/v0.0.9...v0.0.10) (2021-01-30)

### Bug Fixes

* spider

### Pull Requests

* Merge pull request [#6](https://github.com/tsundata/assistant/issues/6) from tsundata/dependabot/go_modules/github.com/influxdata/influxdb-client-go/v2-2.2.2
* Merge pull request [#7](https://github.com/tsundata/assistant/issues/7) from tsundata/dependabot/go_modules/github.com/gogo/protobuf-1.3.2


<a name="v0.0.9"></a>
## [v0.0.9](https://github.com/tsundata/assistant/compare/v0.0.8...v0.0.9) (2021-01-30)

### Bug Fixes

* tests
* github action script

### Features

* memo template
* http node
* setting, credentials template
* apps template
* workflow service
* flowscript


<a name="v0.0.8"></a>
## [v0.0.8](https://github.com/tsundata/assistant/compare/v0.0.7...v0.0.8) (2021-01-26)

### Bug Fixes

* tests


<a name="v0.0.7"></a>
## [v0.0.7](https://github.com/tsundata/assistant/compare/v0.0.6...v0.0.7) (2021-01-19)

### Bug Fixes

* a element


<a name="v0.0.6"></a>
## [v0.0.6](https://github.com/tsundata/assistant/compare/v0.0.5...v0.0.6) (2021-01-19)

### Bug Fixes

* lint

### Features

* redis agent


<a name="v0.0.5"></a>
## [v0.0.5](https://github.com/tsundata/assistant/compare/v0.0.4...v0.0.5) (2021-01-18)

### Features

* influx interceptor


<a name="v0.0.4"></a>
## [v0.0.4](https://github.com/tsundata/assistant/compare/v0.0.3...v0.0.4) (2021-01-18)

### Bug Fixes

* web route


<a name="v0.0.3"></a>
## [v0.0.3](https://github.com/tsundata/assistant/compare/v0.0.2...v0.0.3) (2021-01-15)

### Bug Fixes

* rpc host check
* error log

### Features

* spider script
* bbolt, etcd config


<a name="v0.0.2"></a>
## [v0.0.2](https://github.com/tsundata/assistant/compare/v0.0.1...v0.0.2) (2021-01-14)

### Bug Fixes

* .goreleaser.yml

### Features

* download script


<a name="v0.0.1"></a>
## v0.0.1 (2021-01-14)

### Bug Fixes

* spider action
* registry time
* rename dir
* update rpc client
* service instances
* log print
* opt

### Features

* goreleaser
* server agent
* influxdb, remove prometheus
* cron
* spider
* prometheus
* jaeger
* archive production artifacts
* grpc middleware
* qr command
* pwd, ut, rand command
* subs command
* middle, page service
* web components
* subscribe cron
* bot cron
* bot, plugin
* run flowscript
* message const
* map, filter, reduce
* list, dict
* package
* builtin func
* undefined function error
* function return
* print
* string, boolean
* while grammar
* if grammar
* grammar file
* nested procedure calls
* executing procedure calls
* callstack
* procedure call
* interpreter error
* scope
* collection
* interpreter
* symbol
* interpareter
* UnaryOp
* parser
* interpreter
* lexer
* interpreter
* interpreter, token
* run message
* registry tcp
* fasthttp
* logging
* message event
* slack event
* rpcx, gin
* client auth
* RPC broadcast func
* app name
* proto
* utils
* gitkeep
* cron
* subscribe service
* generator
* Slack Incoming Message
* update Dockerfile
* read config.yml
* register rpc service
* rpc registry
* base framework
* Update README

