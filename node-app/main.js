const faktory = require('faktory-worker');

faktory.register('NodeWorker', async (id) => {
  console.log("Hello from NodeWorker. I'm processing " + id);
});

faktory.work({ queues: ["node-app"] });
