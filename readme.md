# Distributed Web Crawler (Go)

## How to run
1. Start docker container: elastic search (default port 9200)
2. Start persistence service (default port 1234)
3. Start worker services 
3. Start crawler engine

~~~bash
dock-compose up -d
go run crawler_distributed/persist/server/persist_start.go 
go run worker_start.go --port=9000
go run worker_start.go --port=9001
go run crawler_distributed/main.go --worker_hosts=9000,9001
~~~


## Single Task Version
Only one worker is fetching data.

## Concurrent Version
### Simple Scheduler
All workers share one IN channel and one OUT Channel.
### Queued Scheduler
Worker have their own IN channel. When worker is free, they will put their IN channel to the scheduler Worker Queue.
Scheduler also have a queue for new Request, use "select" to connect them together.

## Distributed Version
Using jsonrpc to communicate between services.

###Notes
docker run -it(interactive tty)
docker run -d(detached daemon) -p(port)

elastic search url structure
* index(database)/type(table)/id
* use Put/Post to add data, post: id is optional
* Get: index/type/_search?q=  