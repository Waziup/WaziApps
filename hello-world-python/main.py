#!/usr/bin/python
import usock
import os
from urllib.parse import urlparse

#---------------------#

usock.sockAddr = "/root/app/proxy.sock"


#Path to the root of the code
PATH = os.path.dirname(os.path.abspath(__file__))

#---------------------#


def index(url, body=""):
    return 200, b"Salam Goloooo"


usock.routerGET("/", index)

#------------------#


def ui(url, body=''):
    filename = urlparse(url).path.replace("/ui/", "")
    if (len(filename) == 0):
        filename = 'index.html'
    try:
        with open(PATH + '/ui/' + filename, mode='rb') as file:
            return 200, file.read()
    except Exception as e:
        print("Error: ", e)
        return 404, b"File not found"


usock.routerGET("/ui/(.*)", ui)
usock.routerPOST("/ui/(.*)", ui)

#------------------#


def time(url, body=""):
    import datetime
    dateAndTime = datetime.datetime.now().strftime("%B %d %Y %H:%M:%S")

    out = str.encode(dateAndTime)
    return 200, out


usock.routerGET("/time", time)

#------------------#

if __name__ == "__main__":
    usock.start()