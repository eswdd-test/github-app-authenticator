# github-app-authenticator

Tool to generate temporary OAuth tokens for a Github App

## build

```
$ go build -o github-app-authenticator *.go
```

## options

```
$ ./github-app-authenticator -help
  Usage:
    github-app-authenticator [OPTIONS]
  
  Application Options:
        --installation-id= GitHub App installation id
        --application-id=  GitHub App id
        --private-key=     Path to private key PEM file
  
  Help Options:
    -h, --help             Show this help message
```

## run

```
$ ./github-app-authenticator --installation-id=7067981 --application-id=55935 --private-key=private-key.pem 
  v1.0af0d83e0df9e6644ae52e1253679fa5f7bed4c6
```