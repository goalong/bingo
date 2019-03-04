module github.com/goalong/bingo

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/Unknwon/com v0.0.0-20190214221849-2d12a219ccaf
	github.com/gin-contrib/sse v0.0.0-20190301062529-5545eab6dad3 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.0 // indirect
	github.com/jinzhu/configor v1.0.0
	github.com/jinzhu/gorm v1.9.2
	github.com/jinzhu/inflection v0.0.0-20180308033659-04140366298a // indirect
	github.com/joho/godotenv v1.3.0
	github.com/mattn/go-isatty v0.0.6 // indirect
	github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gopkg.in/yaml.v2 v2.2.2 // indirect
)

replace (
	example.com/routers v0.0.0 => ./routers
	golang.org/x/net v0.0.0-20180906233101-161cd47e91fd => github.com/golang/net v0.0.0-20180906233101-161cd47e91fd
	golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f => github.com/golang/sync v0.0.0-20180314180146-1d60e4601c6f
	golang.org/x/sys v0.0.0-20190222072716-a9d3bda3a223 => github.com/golang/sys v0.0.0-20190222072716-a9d3bda3a223
	google.golang.org/genproto v0.0.0-20180831171423-11092d34479b => github.com/google/go-genproto v0.0.0-20180831171423-11092d34479b

)
