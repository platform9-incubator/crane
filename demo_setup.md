Running the demo setup

I have a base node setup at 

https://github.com/roopakparikh/crane-node

clone it locally


git clone https://github.com/roopakparikh/crane-node

cd crane-node

Add  the following at /opt/pf9/hello_world.js

    #!/usr/local/bin/node
    // Load the http module to create an http server.
    var http = require('http');

    // Configure our HTTP server to respond with Hello World to all requests.
    var server = http.createServer(function (request, response) {
        response.writeHead(200, {"Content-Type": "text/plain"});
        response.end("Hello World\n");
    });

    // Listen on port 8000, IP defaults to 127.0.0.1
    server.listen(8000);

    // Put a friendly message on the terminal
     console.log("Server running at http://127.0.0.1:8000/");

echo 'cmd="/opt/pf9/hello_world.js"' > ./.crane.env

And run
   
   crane run --local <repo url> --port 8000:8000
