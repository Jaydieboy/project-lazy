# Project Lazy

This is a web application that creates a network of data for better lookup and to ultimately create the "button" that does all the things.

## Steps for Development
1.  Get the server mux up and running.
    1. Server `index.gohtml`

## Built With

- The backend uses the Go programming language
  - The server mux employs the [julienschmidt/httprouter](https://godoc.org/github.com/julienschmidt/httprouter)
- SPA front end uses Angular

## Microservices
Go API will pass jobs to microservices mostly written in Go. Data to be sent between some services should be in the form of binary json until I know better. See more [here](https://medium.com/@nathankpeck/microservice-principles-smart-endpoints-and-dumb-pipes-5691d410700f) Look into managing multiple databases using event-driven architecture. Look into it more [here](https://microservices.io/patterns/data/database-per-service.html) I will be using same database. Here is something on creating roles for users [neo4j docs](https://neo4j.com/docs/operations-manual/current/authentication-authorization/subgraph-access-control/).

Look at how to get a microservice architectured product off the ground [here](https://www.devbridge.com/articles/a-6-point-plan-for-implementing-a-scalable-microservices-architecture/).
### Developed

### To be Developed
- Look at pdf as image and see if there is a match stored in the database with all the useful content or parse pdf content if possible. Output all stored content and give the ability to add content. Also match with project. (Priority: Low)
- Match spec to cutsheet and give recommendations of what to look for on drawing if required (schedule or location on plans). (Priority: Low)
- Link Revit model in some way (Priority: Low)
- Store user information in database (Priority: High)
- Hash password (Priority: High)
- Send email to verify email (don't see the use in this at this point) (Priority: Low)
## To-Do
- [ ] Create an Ubuntu docker container to set up the Neo4j Go Driver
- [ ] Build REST-ful API in Go with help from [Todd McLeod's Golang Web Development Course](https://github.com/GoesToEleven/golang-web-dev)
- [ ] Connect backend to a Neo4j server using the [Official Neo4j Go Driver](https://github.com/neo4j/neo4j-go-driver)
- [ ] Add image recognition of some sort using [GoCV](https://gocv.io/) to check for similar documents.
- [ ] Parse word documents for formating and contents, eventially adding a searchable master spec system
- [ ] Implement [pdf parser](https://godoc.org/github.com/unidoc/unidoc/pdf/core#PdfObjectString)

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
