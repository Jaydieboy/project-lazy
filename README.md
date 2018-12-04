# Project Lazy

This is a web application that creates a network of data for better lookup and to ultimately create the "button" that does all the things.

Note: Sometimes use golint.

Do this: https://tour.golang.org/concurrency/10

## Steps for Development
1.	Get the [server mux](https://godoc.org/github.com/julienschmidt/httprouter) up and running.
2.	Serve `/templates/index.gohtml`. If you want to go to another page then use `<a href="/pagename">pagename</a>`.

3.	Read through [Passing Data](https://github.com/GoesToEleven/golang-web-dev/tree/master/027_passing-data) setion.
4.	Create a generic header and footer 

## Built With

- The backend uses the Go programming language
  - The server mux employs the [julienschmidt/httprouter](https://godoc.org/github.com/julienschmidt/httprouter)

## Microservices
Go API will pass jobs to microservices mostly written in Go. Data to be sent between some services should be in the form of binary json until I know better. See more [here](https://medium.com/@nathankpeck/microservice-principles-smart-endpoints-and-dumb-pipes-5691d410700f) Look into managing multiple databases using event-driven architecture. Look into it more [here](https://microservices.io/patterns/data/database-per-service.html) I will be using same database. Here is something on creating roles for users [neo4j docs](https://neo4j.com/docs/operations-manual/current/authentication-authorization/subgraph-access-control/).

Look at how to get a microservice architectured product off the ground [here](https://www.devbridge.com/articles/a-6-point-plan-for-implementing-a-scalable-microservices-architecture/).

Try this for [microservice](https://medium.com/@shijuvar/building-microservices-with-event-sourcing-cqrs-in-go-using-grpc-nats-streaming-and-cockroachdb-983f650452aa) with go.

[Just Watch This](https://www.youtube.com/watch?v=j6ow-UemzBc)

### Developed

### To be Developed
- Look at pdf as image and see if there is a match stored in the database with all the useful content or parse pdf content if possible. Output all stored content and give the ability to add content. Also match with project. (Priority: Low)
- Match spec to cutsheet and give recommendations of what to look for on drawing if required (schedule or location on plans). (Priority: Low)
- Link Revit model in some way (Priority: Low)
- Store user information in database (Priority: High)
- Hash password (Priority: High)
- Send email to verify email (don't see the use in this at this point) (Priority: Low)
## To-Do
- [ ] Create an Ubuntu docker container to set up the [Neo4j Bolt Driver for Go](https://github.com/neo4j/neo4j-go-driver)
- [ ] Build REST-ful API in Go with help from [Todd McLeod's Golang Web Development Course](https://github.com/GoesToEleven/golang-web-dev)
- [ ] Connect backend to a Neo4j server using the [Official Neo4j Go Driver](https://github.com/neo4j/neo4j-go-driver)
- [ ] Add image recognition of some sort using [GoCV](https://gocv.io/) to check for similar documents.
- [ ] Parse word documents for formating and contents, eventially adding a searchable master spec system
- [ ] Implement [pdf parser](https://godoc.org/github.com/unidoc/unidoc/pdf/core#PdfObjectString)
- [ ] Implement a [load balancer](https://godoc.org/github.com/nienie/marathon) for microservices

## Just Notes

Cite code snippets as follows:
```
/***************************************************************************************
*    Title: <title of program/source code>
*    Author: <author(s) names>
*    Date: <date>
*    Code version: <code version>
*    Availability: <where it's located>
*
***************************************************************************************/
```

Does this need to happen of I have adapted a concept for my code from someone else?

### Here is Example code for connecting to neo4j

<details><summary>Go Code</summary>
<p>

```golang
package main

import (
  "flag"
  "fmt"
  "log"
  "os"
  "strings"

  "github.com/neo4j/neo4j-go-driver/neo4j"
)

var (
  uri      string
  username string
  password string
  query    string
)

// Simple header printing logic, open to improvements
func processHeaders(result neo4j.Result) {
  if keys, err := result.Keys(); err == nil {
    for index, key := range keys {
      if index > 0 {
        fmt.Print("\t")
      }
      fmt.Printf("%-10s", key)
    }
    fmt.Print("\n")

    for index := range keys {
      if index > 0 {
        fmt.Print("\t")
      }
      fmt.Print(strings.Repeat("=", 10))
    }
    fmt.Print("\n")
  }
}

// Simple record values printing logic, open to improvements
func processRecord(record neo4j.Record) {
  for index, value := range record.Values() {
    if index > 0 {
      fmt.Print("\t")
    }
    fmt.Printf("%-10v", value)
  }
  fmt.Print("\n")
}

// Transaction function
func executeQuery(tx neo4j.Transaction) (interface{}, error) {
  var (
    counter int
    result  neo4j.Result
    err     error
  )

  // Execute the query on the provided transaction
  if result, err = tx.Run(query, nil); err != nil {
    return nil, err
  }

  // Print headers
  processHeaders(result)

  // Loop through record stream until EOF or error
  for result.Next() {
    processRecord(result.Record())
    counter++
  }
  // Check if we encountered any error during record streaming
  if err = result.Err(); err != nil {
    return nil, err
  }

  // Return counter
  return counter, nil
}

func run() error {
  var (
    driver           neo4j.Driver
    session          neo4j.Session
    recordsProcessed interface{}
    err              error
  )

  // Construct a new driver
  if driver, err = neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""), func(config *neo4j.Config) {
    config.Log = neo4j.ConsoleLogger(neo4j.ERROR)
  }); err != nil {
    return err
  }
  defer driver.Close()

  // Acquire a session
  if session, err = driver.Session(neo4j.AccessModeRead); err != nil {
    return err
  }
  defer session.Close()

  // Execute the transaction function
  if recordsProcessed, err = session.ReadTransaction(executeQuery); err != nil {
    return err
  }

  fmt.Printf("\n%d records processed\n", recordsProcessed)

  return nil
}

func parseAndVerifyFlags() bool {
  flag.Parse()

  if uri == "" || username == "" || password == "" || query == "" {
    flag.Usage()
    return false
  }

  return true
}

func main() {
  if !parseAndVerifyFlags() {
    os.Exit(-1)
  }

  if err := run(); err != nil {
    log.Fatal(err)
    os.Exit(1)
  }

  os.Exit(0)
}

func init() {
  flag.StringVar(&uri, "uri", "bolt://localhost", "the bolt uri to connect to")
  flag.StringVar(&username, "username", "neo4j", "the database username")
  flag.StringVar(&password, "password", "", "the database password")
  flag.StringVar(&query, "query", "", "the query to execute")
}
```

</p>
</details>

## Front End
### HTML
Don't got nothing yet.
### CSS (I would love to use scss)
- [ ] Create a CSS reset document `projectlazy-reset.css`.
