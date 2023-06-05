# graphQL_template
graphQL_template golang

Chi + 99designs/gqlgen

### tree

``` 
.
├── cmd
│   └── main.go
├── configs
│   ├── config.yaml
│   └── gqlgen.yml
├── gen.go
├── go.mod
├── go.sum
├── internal
│   ├── conf
│   │   └── config.go
│   ├── generated
│   │   ├── generated.go
│   │   ├── models_gen.go
│   │   ├── README.md
│   │   └── resolver.go
│   ├── graphql
│   │   ├── common
│   │   │   └── common.graphql
│   │   └── user
│   │       └── auth.graphql
│   ├── middlewares
│   │   ├── auth.go
│   │   ├── context.go
│   │   └── cors.go
│   ├── pkg
│   │   ├── enum
│   │   │   ├── auth.go
│   │   │   └── middleware.go
│   │   └── models
│   │       └── middleware.go
│   ├── resolvers
│   │   ├── resolver.go
│   │   └── user.go
│   ├── storage
│   │   ├── interface.go
│   │   └── simple
│   │       └── simple.go
│   └── utils
│       ├── base.go
│       ├── jwt.go
│       └── logger.go
├── LICENSE
├── README.md
└── tools.go
```

#### example payload

``` 
# Welcome to GraphiQL
#
# GraphiQL is an in-browser tool for writing, validating, and
# testing GraphQL queries.
#
# Type queries into this side of the screen, and you will see intelligent
# typeaheads aware of the current GraphQL type schema and live syntax and
# validation errors highlighted within the text.
#
# GraphQL queries typically start with a "{" character. Lines that start
# with a # are ignored.
#
# An example GraphQL query might look like:
#
#     {
#       field(arg: "value") {
#         subField
#       }
#     }
#
# Keyboard shortcuts:
#
#  Prettify Query:  Shift-Ctrl-P (or press the prettify button above)
#
#     Merge Query:  Shift-Ctrl-M (or press the merge button above)
#
#       Run Query:  Ctrl-Enter (or press the play button above)
#
#   Auto Complete:  Ctrl-Space (or just start typing)
#


query captcha {
  captcha{
    captchaId
    base64Captcha
  }
}

query userInfo{
  user {
    role
    account
    accountId
    accountName
  }
}

mutation registry {
  registry(input: {
    name: "dollarkiller",
    account: "xxx@xxx.com",
    password:"123456",
    captchaID: "JA6A2L",
    captchaCode: "tnka"
  })
}


mutation login {
  loginByPassword(input: {account:"xxx@xxx.com",password:"123456",captchaID: "0BW0PS",captchaCode:"hk5w"}) {
    userID
    accessTokenString
  }
}
```