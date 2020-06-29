FROM waziup/wazigate-static-server:latest

COPY . /root

#This command puts only the required files in the index.zip file which is good for production
#RUN zip /index.zip /root/docker-compose.yml /root/package.json

#This command puts everything in the index.zip file which is good for development
RUN zip -q -r /index.zip /root/