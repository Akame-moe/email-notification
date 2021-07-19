# Email Notification

### 1.compile command line tool

```bash
go build -o sendemail cmd/command.go
```
###### Usage
```bash
./sendemail -s yoursubject -c yourcontent -a yourattachment1.jpg -a yourattachment2.png
```

### 2.compile command line tool

```bash
go build -o emailhttpserver httpserver/httpserver.go
```

###### Usage
```bash
./emailhttpserver -addr 0.0.0.0:1234 -path /email/secretpath
```
now you can send email notification in both GET or POST method, e.g.:
```bash
curl "http://<HOST>:<PORT>/email/secretpath?subject=hello&content=congratulations"
```