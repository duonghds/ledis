#LEDIS
##Author: duonghds
This is a simple Redis service version named Ledis (lightweight redis)
and built with Golang. Using gin-gonic library to export Rest API.
In Front-end, I use jquery-terminal framework to fake redis-cli.
Both service deployed on Heroku.

###Source
- Github `https://github.com/duonghds/ledis`
- Github FE `https://github.com/duonghds/ledis-fe`
- Describe about ledis from `Holistics company`
- Heroku deploy domain `https://afternoon-wildwood-94295.herokuapp.com`
- Heroku Post API `https://afternoon-wildwood-94295.herokuapp.com/api/`
- Test with cli ` https://hidden-bayou-11572.herokuapp.com/`

###Pre-commit
run `go test ./..` to test file before commit            

###Support command
- String: [SET key], [GET key]
- List: [LLEN key], [RPUSH key value1 value2...], [LPOP key], [RPOP key], [LRANGE key start stop]
- Set: [SADD key value1 value2...], [SCARD key], [SMEMBERS key], [SREM key value1 value2...], [SINTER key1 key2 key3...]
- Data expiration: [KEYS], [DEL key], [FLUSHDB], [EXPIRE key seconds], [TTL]
- Snapshot: [SAVE], [RESTORE]

###Project Structure

[MemCache]<------>[GlobalService]<------>[List/Set/StringService]<----->[Request]

###Tech flow
1. Request from client
2. Validate and split string to get correct format and detect command
3. Call to List/Set/String service with each command
4. List/Set/String service handle logic command
5. List/Set/String service call to Global service to access mem cache
6. Handle mem cache to get and store data

###How to run
- Install and setup Golang on your device
- Clone project
- Run `go mod tidy` (download all the dependencies that are required in
source files and update go.mod file with that dependency)
- Run `go mod vendor` (constructs a directory named vendor in the main module's
root directory that contains copies of all packages needed to support builds
and tests of packages. It's the same with `node_modules`)
- Final run `go run main.go`
- To run all test `go test ./..`