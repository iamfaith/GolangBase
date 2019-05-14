SHELL = /bin/bash
start:
	chmod +x ./docker/run-init.sh
	./docker/run-init.sh
	docker-compose start

up:
	export $(cut -d= -f1 conf/app.env)
	source ./conf/app.env
	chmod +x ./docker/run-init.sh
	./docker/run-init.sh
	docker-compose stop
	docker system prune -f
	docker-compose up -d

up-force:
	export $(cut -d= -f1 conf/app.env)
	source ./conf/app.env
	chmod +x ./docker/run-init.sh
	./docker/run-init.sh
	docker-compose stop
	docker system prune -f
	docker-compose up -d
	
stopall:
	docker stop $(docker ps -q)
