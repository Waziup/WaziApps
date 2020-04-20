FROM python:slim-buster

COPY . /root/src/

RUN  apt-get update \
     && apt-get install -y \
     curl \
     zip \
     && cd /root/src/ \
     && zip /index.zip docker-compose.yml package.json

#----------------------------#

# Uncomment For development
# ENTRYPOINT ["tail", "-f", "/dev/null"]
# sudo docker exec -it waziup.hello-world-python sh

# Uncomment For production
ENTRYPOINT ["python", "/root/src/main.py"]

