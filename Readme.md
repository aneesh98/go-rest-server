A repository to try different web frameworks in golang on todo list data model.<br/>
Article used as reference for code implementation is [Rest Servers in Go](https://eli.thegreenplace.net/2021/rest-servers-in-go-part-1-standard-library/#)
<br/> Using mongodb as backend instead of in-memory go structs

**DEVELOPMENT SETUP**
1. Prerequisites:
   * Go
   * NodeJS
   * MongoDB local installation 
   * Make Utility
2. Install dependencies of node project (to populate fake data in mongo db) using `cd ./db_population_helper/mongo && npm install`
3. Setup Go Dependencies using `go mod tidy`
4. Populate MongoDB (this will create a database `ToDoListDev` and collection `todoLists` in it) using the following command:
   `make mongodb_fake_populate SIZE=50` (Can change 50 to whatever size of data you want to have)
