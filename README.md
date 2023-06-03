# Faktory Example

TODO
- Build a diagram explaining how it works
- Prepare second user case
- Prepare third use case
- Prepare fourth use case
- Prepare documentation

USE CASES
1) When there is an error, it should retry
2) Schedule job
3) Schedule and remove job
4) Specify number of retries and queue

PRESENTATION
- Problems we have today
- Architecture
- Use cases
- Dashboard
- Attention to idempotency and concurrency

## Diagram

TODO

## How To Run

Ruby the project:

```
docker compose up
```

Enqueue jobs to `ruby-worker`:

```
docker compose run ruby-worker ruby enqueue.rb
```
