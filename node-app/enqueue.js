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

async function deleteJob(jid) {
  const client = await faktory.connect();
  client.scheduled.withJids(jid).discard();
  await client.close();
  console.log("Job", jid, "deleted");
}

if (process.argv[2] == "push") {
  var scheduleAt = new Date(new Date().getTime() + (1 * 60 * 1000));
  createJob(scheduleAt);
}

if (process.argv[2] == "discard") {
  deleteJob(process.argv[3]);
}
