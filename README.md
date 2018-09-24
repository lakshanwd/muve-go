# muve-go ![alt text][build_status]

## assumptions
* you have go1.10.1 installed
* you have correcly setup GOPATH variable
* you have installed mysql and you have mysql username/password

## setup instructions
* open terminal and execute following commands
```shell
$ go get -u github.com/lakshanwd/muve-go
$ cd $GOPATH/github.com/lakshanwd/muve-go
$ mysql -u <mysql-user> -p < backup.sql
$ cd go-crud
$ go get -v ./...
```
* open a text editor and configure sql credentials in [config.json][config.json]
* and finally `$ go run server.go`

There is the JWT authenticated golang api

## testing
* Open google chrome and install [Restlet][restlet], 
sign in and click import and Select **Restlet Client Repository** and browse for the [muve-go.json][muve-go.json]

### database structure
![alt text][db_structure]

### docker
* install docker-ce
* install docker-compose
* run `$ docker-compose -f muve-go/docker-compose.yml up -d`

[build_status]: https://travis-ci.org/lakshanwd/muve-go.svg?branch=master "Travis Build Status"
[db_structure]: ../master/db-structure.png "Database Structure"
[restlet]: https://chrome.google.com/webstore/detail/restlet-client-rest-api-t/aejoelaoggembcahagimdiliamlcdmfm "Restlet client"
[config.json]: https://github.com/lakshanwd/muve-go/blob/master/go-crud/config.json "config.json"
[muve-go.json]: https://github.com/lakshanwd/muve-go/blob/master/muve-go.json "muve-go.json"