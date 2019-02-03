const http = require('http');

let server = http.createServer(handler);
server.listen(8000, '0.0.0.0');

function handler(req, res) {
    setTimeout(() => {
        res.statusCode = 200;
        res.end('OK');
    }, 5000);
}
