all:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o urantiabook .

dev:
	GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o urantiabook .
	- docker image rm localhost:32000/urantiabook/urantiabook 
	docker build -t localhost:32000/urantiabook/urantiabook .
	docker push localhost:32000/urantiabook/urantiabook

beta:
	GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o urantiabook .
	- docker image rm reg.urantiatech.com/urantiabook/urantiabook 
	docker build -t reg.urantiatech.com/urantiabook/urantiabook .
	docker push reg.urantiatech.com/urantiabook/urantiabook

prod:
	GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o urantiabook .
	- docker image rm reg.urantiatech.com/urantiabook/urantiabook:prod 
	docker build -t reg.urantiatech.com/urantiabook/urantiabook:prod .
	docker push reg.urantiatech.com/urantiabook/urantiabook:prod
