./build.sh

docker build . -t sudhan/webdebug

docker run --rm -p 8888:8888 sudhan/webdebug
