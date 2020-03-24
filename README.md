./build.sh
docker build . -t webdebug
docker run --rm -p 8888:8888 webdebug
