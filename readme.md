docker build -t go-lang .
docker run -it -p 8000:8000 go-lang:latest

localhost:8000/api/encrypt , localhost:8000/api/decrypt
