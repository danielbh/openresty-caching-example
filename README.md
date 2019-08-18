# openresty-caching-example
An example for using caching with nginx/openresty


cd nginx

docker build . -t=reverseproxy

cd ..

cd node

docker build . -t=nodeapp

cd ..

docker-compose up

curl localhost:8080

docker build --target builder
