# muve-go

## assumptions
* you have go1.10.1 installed
* you have correcly setup GOPATH variable
* you have installed mysql and you have mysql user/password

## setup instructions
* open terminal and execute following commands
```shell
$ go get -u github.com/lakshanwd/muve-go
$ cd $GOPATH/github.com/lakshanwd/muve-go
$ mysql -u <mysql-user> -p < backup.sql
$ cd go-crud
$ go get ./...
```
* open a text editor and replace `developer:123` with your `<mysql-user>:<mysql-password>` and password in [db/db.go](https://github.com/lakshanwd/muve-go/blob/master/go-crud/db/db.go)
* and finally `$ go run server.go`

There is the JWT authenticated golang api

## testing
* Open google chrome and install [Restlet](https://chrome.google.com/webstore/detail/restlet-client-rest-api-t/aejoelaoggembcahagimdiliamlcdmfm), 
sign in and click import and Select **Restlet Client Repository** and browse for the [muve-go.json](https://github.com/lakshanwd/muve-go/blob/master/muve-go.json)



### database structure
![db-structure](../master/db-structure.png)
