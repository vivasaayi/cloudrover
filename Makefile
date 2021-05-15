.PHONY:all

dockerbuild:
	docker build -t cloudrover .

dockerrun:
	docker-compose up cloudrover