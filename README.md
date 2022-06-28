# Overview

- Library `cobra` - for cli app.
- Sending email via smtp.
- Clean architecture reference.

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

# Unit test

- I just write unit test for 
  - Valid customer.
  - Generate mail from template.
  - Parse email from text (extract email in from data of template for SMTP header).
  - Valid email.

# Note

- Valid email (package `net/mail` - **RFC 5322**)

```
https://stackoverflow.com/questions/66624011/how-to-validate-an-email-address-in-go
```

- As in doc for **RFC 5322** `abcd@efgh` is a valid email because `efgh` may be a valid local domain name. So RFC 5322 does not check valid email strictly.
- I think in this application, we don't need valid emails strictly because we don't know where this email is from, we can't find valid email of that person.
- In many applications, we don't need to check a local domain name or special cases like RFC 5322 standard. This can cause many **incorrect email** addresses **in production.** 
- I think we can solve the problem with config:
  - If we just send email to the public domain or personal email - we can config valid email for that case to use regex to parse email more strictly.
  - If we need to send email to local or business email - we use package `net/mail`.
- In this case, I think we should use `net/mail` for easy to use.

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
