Use-Cases
1. To quickly build POC backend systems. Consider, you have an interview and you were asked to build a backend system, in a time frame.
Instead of spending time on creating folder structure and making sure that db is up and running.
You can do:
```
git clone https://github.com/raunakjodhawat/go-docker-backend.git
```
and then run:
```
docker build -t yourname\interview && docker run -p 8080:8080 yourname\interview
```


In future, you can also use, golang tests, to run in a separate docker container

Currently, it only supports mongodb. I plan to include an array of db options.



Use-cases to quickly build scalable proof of concept backends

Goal is to give full set of CRUD feature
