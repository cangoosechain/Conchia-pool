# Prerequisites:
- install docker
- install docker-compose
- install go
- install make (if not already installed)

# To run locally:

1. start redis:
   `docker-compose up -d redis`
2. run the web app
   `make run`

hit http://localhost:8080/farmers in your browser

# To build and run unit tests:

1. `make build-and-test`

# To load sample data into redis:

1. start redis
   `docker-compose up -d redis`
2. copy the sample data file to the container
   `docker ps # to get the container id for redis`
   `docker cp sample-redis-data.redis <container id>:/sample-redis-data.redis`
   `docker-compose exec redis bash`
   # now you are in the docker container environment
   `cd /`
   `cat sample-redis-data.redis | redis-cli --pipe`
   
# To use the redis cli to look at the data

1. start redis if not already running
2. `docker-compose exec redis redis-cli`
# now you can execute any redis commands, such as "keys *" to list all keys currently stored.

# Building for linux (optional)
If you are developing on a windows or mac machine, you will need to build an executable for linux before you can copy it to the server (assuming the chia pool server is running on linux)
To do this, run `make build-linux`

# To Do
## update chia pool reference code (python)
1. change to connect to redis and on first connect by a farmer, record their launcher_id in redis under `farmerInfo:<launcher id>` and other info
2. change to connect to redis and append a new partial with result to the `farmerPartials:<launcher id>` list.  The `RPUSH` command will append it to the end of the list.
## deploy pool web
1. since the pool webapp is written in Go, a simple `make build-and-test` will produce an executable file named `chia-pool-web`. This executable can be copied to the machine where you are running the chia pool and started there.
## security
1. the `/farmers` and `/farmers/:id/edit` endpoints should be secured and require a login. The `/farmer/:id` endpoint is meant for giving to the individual users of the pool so they can see their own stats (if you want to give them that)
## lots of other ideas!
