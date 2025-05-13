const http = require('http');
const url = require('url');
const cluster = require('cluster');
const os = require('os')

const numCPUs = os.cpus().length;

// Sample data
const personData = {
  "1": { Name: "John Doe", Age: 30 },
  "2": { Name: "Jane Doe", Age: 28 },
  "3": { Name: "Jack Doe", Age: 25 }
};

// Handler function for the endpoint
const requestHandler = (req, res) => {
  const parsedUrl = url.parse(req.url, true);
  const path = parsedUrl.pathname;
  const query = parsedUrl.query;

  if (path === '/person') {
    const id = query.id;

    if (!id) {
      res.statusCode = 400;
      res.setHeader('Content-Type', 'text/plain');
      res.end('ID is missing');
      return;
    }

    const person = personData[id];

    if (!person) {
      res.statusCode = 404;
      res.setHeader('Content-Type', 'text/plain');
      res.end('Person not found');
      return;
    }

    res.statusCode = 200;
    res.setHeader('Content-Type', 'application/json');
    res.end(JSON.stringify(person));
  } else {
    res.statusCode = 404;
    res.setHeader('Content-Type', 'text/plain');
    res.end('Not Found');
  }
};

// Define the port
const port = 8080;

if (cluster.isMaster) {
  for (let i=0; i<numCPUs ; i++){
    cluster.fork();
  }

  cluster.on('exit', (worker, code, signal) => {
    console.log("Worker ${worker.process.pid} died. Forking a new worker...");
    cluster.fork();
  })
} else {

// Create the server
const server = http.createServer(requestHandler);

// Start the server
server.listen(port, () => {
  console.log(`Server started on port ${port}`);
});
}