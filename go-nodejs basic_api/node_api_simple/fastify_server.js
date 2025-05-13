// const fastify = require('fastify')({ logger: true });
const fastify = require('fastify')();

const cluster = require('cluster');
const os = require('os')

const numCPUs = os.cpus().length;

if (cluster.isMaster) {
  for (let i=0; i<numCPUs ; i++){
    cluster.fork();
  }

  cluster.on('exit', (worker, code, signal) => {
    console.log("Worker ${worker.process.pid} died. Forking a new worker...");
    cluster.fork();
  })
} else {

// Sample data
const personData = {
  "1": { Name: "John Doe", Age: 30 },
  "2": { Name: "Jane Doe", Age: 28 },
  "3": { Name: "Jack Doe", Age: 25 }
};

// Define the route
fastify.get('/person', async (request, reply) => {
  const id = request.query.id;

  if (!id) {
    reply.code(400).send('ID is missing');
    return;
  }

  const person = personData[id];

  if (!person) {
    reply.code(404).send('Person not found');
    return;
  }

  reply.send(person);
});

// Start the server
const start = async () => {
  try {
    await fastify.listen({port: 8080});
    fastify.log.info(`Server started on port 8080`);
  } catch (err) {
    fastify.log.error(err);
    process.exit(1);
  }
};
start();
}