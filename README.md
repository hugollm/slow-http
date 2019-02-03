# Slow HTTP

Test with different http stacks to experiment with slow endpoints.
All examples implement an HTTP server that sleeps for 5 seconds and return an OK response.

The objective is findingo out the bottlenecks for such slow endpoints, and investigate strategies
to boost the upper limits of requests served per second.


## Running the Servers

To start each server, run:

    make golang
    make nodejs
    make python

All of them listen on :8000.


## Benchmarking

Use the `wrk` tool to load the servers:

    make load target=http://target.server:8000

It runs a heavy load of requests against the server for 30 seconds and display the results.
