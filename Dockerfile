from  docker.io/library/alpine

copy resourceManager /home/

run chmod +x /home/resourceManager

workdir /home/

cmd ./resourceManager