# Faktory Example

USE CASES
1) When there is an error, it should retry
2) Schedule job
3) Schedule and remove job
4) Specify number of retries and queue

## Diagram

TODO

## How To Run

Ruby the project:

```
docker compose up
```

***

Enqueue jobs to `ruby-app`:

```
docker compose run ruby-app ruby enqueue.rb
```

***

Enqueue jobs to `node-app`:

```
docker compose run node-app node enqueue.js push
```

It is going to schedule a job to run one minute ahead. You can discard the job by running the following command:

```
docker compose run node-app node enqueue.js discard <jid>
```
