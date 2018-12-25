## Web Crawler (Go)

### Simple Scheduler
All workers share one In channel and one Out Channel
### Queued Scheduler
Worker have their own In channel. When worker is free, they will put their In channel to the scheduler WorkerQueue
Scheduler also have a queue for new Request, use "select" to connect them together

####Notes
docker run -it(interactive tty)
docker run -d(detached daemon) -p(port)

elastic search url structure
* index(database)/type(table)/id
* use Put/Post to add data, post: id is optional
* Get: index/type/_search?q=  