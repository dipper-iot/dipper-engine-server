module github.com/dipper-iot/dipper-engine-server

go 1.19

require (
	entgo.io/ent v0.11.3
	github.com/99designs/gqlgen v0.17.20
	github.com/dipper-iot/dipper-engine v0.0.4
	github.com/go-chi/chi/v5 v5.0.7
	github.com/go-redis/redis/v9 v9.0.0-rc.1
	github.com/gorilla/websocket v1.5.0
	github.com/lib/pq v1.10.7
	github.com/mattn/go-sqlite3 v1.14.15
	github.com/sirupsen/logrus v1.9.0
	github.com/sony/sonyflake v1.1.0
	github.com/vektah/gqlparser/v2 v2.5.1
)

require (
	ariga.io/atlas v0.7.2-0.20220927111110-867ee0cca56a // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/agnivade/levenshtein v1.1.1 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/asaskevich/EventBus v0.0.0-20200907212545-49d423059eef // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-chi/cors v1.2.1 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl/v2 v2.13.0 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/urfave/cli/v2 v2.11.2 // indirect
	github.com/xrash/smetrics v0.0.0-20201216005158-039620a65673 // indirect
	github.com/zclconf/go-cty v1.8.0 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/sys v0.0.0-20220811171246-fbc7d0a398ab // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/tools v0.1.13-0.20220804200503-81c7dc4e4efa // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/dipper-iot/dipper-engine => ../rule-engine-service
