This is Readme
--------------
docker build -t go-lang .

docker run -it -p 8000:8000 go-lang:latest
ENCRYPT
-------
localhost:8000/api/encrypt 

Example input json  body -> {"name":"zingyy99"}


DECRYPT
--------
 localhost:8000/api/decrypt

