docker build ./go -t=goapp -f=./go/Dockerfile
docker build ./nginx -t=reverseproxy -f=./nginx/Dockerfile
