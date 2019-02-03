target = 'http://localhost:8000'

requirements:
	sudo pip install gunicorn
	sudo pip install gevent
	sudo apt install wrk -y

load:
	wrk -t8 -c2400 -d30s --timeout 10 ${target}

golang:
	go run server.go

nodejs:
	node server.js

python:
	gunicorn -k gevent server
