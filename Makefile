build:
	docker build -t my-go-fiber-app-v1:v1 .

run:
	docker run -d -p 3000:3000 --name fiber-app my-go-fiber-app-v1:v1
