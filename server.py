from time import sleep


def application(env, start_response):
    sleep(5)
    start_response('200 OK', ())
    return (b'OK',)
