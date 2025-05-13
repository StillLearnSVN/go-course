const express = require('express');
const cluster = require('cluster');
const os = require('os')

const numCPUs = os.cpus().length;
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
  const app = express();


  // Sample data
  const personData = {
    "1": { Name: "John Doe", Age: 30 },
    "2": { Name: "Jane Doe", Age: 28 },
    "3": { Name: "Jack Doe", Age: 25 }
  };
  
  // Handler function for the endpoint
  app.get('/person', (req, res) => {
    const id = req.query.id;
  
    if (!id) {
      return res.status(400).send('ID is missing');
    }
  
    const person = personData[id];
  
    if (!person) {
      return res.status(404).send('Person not found');
    }
  
    res.json(person);
  });
  
  // Start the server
  app.listen(port, () => {
    console.log(`Server started on port ${port}`);
  });
  
}

