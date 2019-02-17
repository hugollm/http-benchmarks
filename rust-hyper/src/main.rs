extern crate hyper;

use std::{thread, time};
use std::fs;

use hyper::{Body, Request, Response, Server};
use hyper::service::service_fn_ok;
use hyper::rt::{self, Future};


fn main() {
    let addr = ([0, 0, 0, 0], 8000).into();
    rt::run(Server::bind(&addr).serve(|| {
        return service_fn_ok(handler);
    }).map_err(|e| eprintln!("server error: {}", e)));
}

fn handler(request: Request<Body>) -> Response<Body> {
    let path = request.uri();
    if path == "/hello" {
        return hello(request)
    }
    else if path == "/file" {
        return file(request);
    }
    else if path == "/sleep/10ms" {
        return sleep10ms(request);
    }
    else if path == "/sleep/100ms" {
        return sleep100ms(request);
    }
    else if path == "/sleep/1s" {
        return sleep1s(request);
    }
    else if path == "/sleep/5s" {
        return sleep5s(request);
    }
    else {
        return Response::new(Body::from("not found"));
    }
}

fn hello(_request: Request<Body>) -> Response<Body> {
    return Response::new(Body::from("hello"));
}

fn file(_request: Request<Body>) -> Response<Body> {
    return Response::new(Body::from(fs::read_to_string("file.txt").unwrap()));
}

fn sleep10ms(_request: Request<Body>) -> Response<Body> {
    thread::sleep(time::Duration::from_millis(10));
    return Response::new(Body::from("ok"));
}

fn sleep100ms(_request: Request<Body>) -> Response<Body> {
    thread::sleep(time::Duration::from_millis(100));
    return Response::new(Body::from("ok"));
}

fn sleep1s(_request: Request<Body>) -> Response<Body> {
    thread::sleep(time::Duration::from_millis(1000));
    return Response::new(Body::from("ok"));
}

fn sleep5s(_request: Request<Body>) -> Response<Body> {
    thread::sleep(time::Duration::from_millis(5000));
    return Response::new(Body::from("ok"));
}
