const http = require('http');
const fs = require('fs');

let server = http.createServer(handler);
server.listen(8000, '0.0.0.0');

routes = {
    '/hello': hello,
    '/file': file,
    '/sleep/10ms': sleep10ms,
    '/sleep/100ms': sleep100ms,
    '/sleep/1s': sleep1s,
    '/sleep/5s': sleep5s,
};

function handler(request, response) {
    routes[request.url](response);
}

function hello(response) {
    response.end('hello');
}

function file(response) {
    var readStream = fs.createReadStream('file.txt');
    readStream.pipe(response);
}

function sleep10ms(response) {
    setTimeout(() => {
        response.end('ok');
    }, 10);
}

function sleep100ms(response) {
    setTimeout(() => {
        response.end('ok');
    }, 100);
}

function sleep1s(response) {
    setTimeout(() => {
        response.end('ok');
    }, 1000);
}

function sleep5s(response) {
    setTimeout(() => {
        response.end('ok');
    }, 5000);
}
