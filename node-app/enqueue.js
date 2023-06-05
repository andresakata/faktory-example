const faktory = require("faktory-worker");
const { v4: uuidv4 } = require("uuid");

async function createJob(scheduleAt) {
  const client = await faktory.connect();
  const job = await client.job("NodeWorker", uuidv4());
  job.queue = "node-app";
  job.at = scheduleAt;
  const jid = job.push();
  await client.close();
  console.log("Job", jid);
}

var scheduleAt = new Date(new Date().getTime() + 60 * 1000);

createJob(scheduleAt);
