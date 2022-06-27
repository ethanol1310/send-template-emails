# Overview

- Library `cobra` - for cli app.
- Sending email via smtp.

 # Config and data (template/customer)

## Config

- Config file `config/esending.yaml`

```
smtp:
  host: "smtp.gmail.com"
  port: 587
  username: "quanhuynh1310@gmail.com"
  password: "password"
  tls_verify: true
```

-  Gmail -> generate and use app-password (https://support.google.com/mail/answer/185833?hl=en)

## Template & customer

```
data/customers/customers.csv
data/template/template.json
```

# Build and run local

```
# Download all dependencies.
go mod download 

# Build and run
go build -o main

./main send -c=data/customers/customers.csv -t=data/template/template.json -o=data/output.json -e=data/errors.csv 
```

# Build and run with docker

```
docker build -t esending .             
docker run esending
```
