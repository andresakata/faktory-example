# Faktory Example

This repository is a test of the Faktory tool.

## How To Run

Ruby the project:

```
docker compose run node-app npm install
docker compose up
```

### Use Case 1) Retry

Enqueue jobs to `ruby-app`:

```
docker compose run ruby-app ruby enqueue.rb
```

### Use Case 2) Schedule

Enqueue jobs to `node-app`:

```
docker compose run node-app node enqueue.js push
```

It is going to schedule a job to run one minute ahead. You can discard the job by running the following command:

```
docker compose run node-app node enqueue.js discard <jid>
```

### Use Case 3) Retry configuration.

Enqueue jobs to `golang-app`:

```
docker compose run golang-app enqueue
```

It is going to push 6 jobs to run asynchronously and will schedule 1 job in 30 seconds in the future.
