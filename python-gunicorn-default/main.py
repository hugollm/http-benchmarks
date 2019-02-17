from time import sleep


def application(env, start_response):
    return handler(env, start_response)


def handler(env, start_response):
    return routes[env['PATH_INFO']](env, start_response)


def hello(env, start_response):
    start_response('200 OK', ())
    return (b'hello',)


def file(env, start_response):
    start_response('200 OK', ())
    return _serve_file('file.txt')


def _serve_file(name):
    with open(name, 'rb') as f:
        yield f.read(1024 ** 2)


def sleep10ms(env, start_response):
    sleep(0.01)
    start_response('200 OK', ())
    return (b'ok',)


def sleep100ms(env, start_response):
    sleep(0.1)
    start_response('200 OK', ())
    return (b'ok',)


def sleep1s(env, start_response):
    sleep(1)
    start_response('200 OK', ())
    return (b'ok',)


def sleep5s(env, start_response):
    sleep(5)
    start_response('200 OK', ())
    return (b'ok',)


routes = {
    '/hello': hello,
    '/file': file,
    '/sleep/10ms': sleep10ms,
    '/sleep/100ms': sleep100ms,
    '/sleep/1s': sleep1s,
    '/sleep/5s': sleep5s,
}
