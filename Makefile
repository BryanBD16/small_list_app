run:
	go run main.go

test:
	go test ./... -v

build:
	go build .

curl-metrics:
	curl -s http://localhost:3000/metrics | grep -E 'get_requests_total|add_requests_total|clear_requests_total'

curl-get:
	curl http://localhost:3000/element -i

curl-add:
	curl -X POST -H "Content-Type: application/json" -d '{"Name":"test item","Description":"This is a test item"}' http://localhost:3000/element/add

curl-clear:
	curl http://localhost:3000/element/clear

mysql-docker:
	docker run --name mysql-db -p 3306:3306 -e MYSQL_ROOT_PASSWORD=a -e MYSQL_DATABASE=list_app -d mysql

mysql-start:
	docker start mysql-db

mysql-connect:
	mysql --host 127.0.0.1 --user root -p

start-prometheus:
	sudo docker run -d --name=prometheus -p 9090:9090 -v /home/bryan-blais-dupuis/Downloads/prometheus-3.1.0.linux-amd64/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus

loki-start:
	docker run -d --name=loki -p 3100:3100 grafana/loki:latest

loki-stop:
	docker stop loki && docker rm loki
	
grafana-start:
	docker stop grafana || true && docker rm grafana || true
	docker run -d --name=grafana -p 3001:3000 -e "GF_SECURITY_ADMIN_PASSWORD=admin" grafana/grafana:latest

grafana-stop:
	docker stop grafana && docker rm grafana

# Create the Docker network
create-network:
	sudo docker network create monitoring

# Connect Prometheus, Loki, and Grafana to the monitoring network
connect-containers:
	sudo docker network connect monitoring prometheus
	sudo docker network connect monitoring loki
	sudo docker network connect monitoring grafana
