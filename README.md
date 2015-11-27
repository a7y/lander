# lander
An extremely tiny RESTful-ish backend for landing pages.

## API

### HTTP
| method |    URI     | (form) data |
| ------ | ---------- | ----------- |
|  POST  | /subscribe |    email    |

### CSV format
`email, host, when`

|  key  |            description            |                 sample                  |
| ----- | --------------------------------- | --------------------------------------- |
| email | email address provided by client  |             me@mailbox.com              |
| host  |     IP address of the client      |              192.168.1.1                |
| when  | date/time the email was submitted | 1970-01-01 00:00:00.000000000 -0600 CST |

----

## About

### Why
Launching a couple of web-based projects, I needed a simple backend to collect email addresses from landing pages.

### Golang
Despite having never written in Go before, I couldn't ignore the fact that it was the perfect language for this. A server capable of serving your static frontend, along with the single Lander binary, is all you need.

### CSV
I wanted my landing page instances to be as lightweight possible so I decided against any kind of storage backend. For data processing, you can easily import the CSV to the DBMS of your choice once you launch your project.

### Rate Limiting
There is no rate limiting implementation included. If you are using Nginx, my suggestion would be to look at ngx_http_limit_req_module.

#### FYI
This is my first time writing something in Go. The code isn't horrendous, but it definitely isn't great.

----

## License
MIT; go hog-wild. :)
