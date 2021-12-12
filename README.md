#LEDIS
##Author: duonghds

###Pre-commit
run `go test ./..` to test file before commit            

###Support command
String: get/set
- List: llen, rpush, lpop, rpop, lrange
- Set: sadd, scard, smembers, srem, sinter
- Data expiration: keys, del, flushdb, expire, ttl
- Snapshot: save, restore

###Project Structure

[MemCache]<---Use--->[GlobalService]-----Injected----->[List/Set/StringService]

###Tech flow
1. Request from client
2. Validate and split string to get correct format and detect command
3. Call to List/Set/String service with each command
4. List/Set/String service handle logic command
5. List/Set/String service call to Global service to access mem cache
6. Handle mem cache to get and store data

###Source
- Github `https://github.com/duonghds/ledis`
- Describe about ledis `Holistics company`
- Heroku deploy domain `https://afternoon-wildwood-94295.herokuapp.com`
- Heroku Post API `https://afternoon-wildwood-94295.herokuapp.com/api/`