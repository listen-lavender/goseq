#!/usr/bin/env python
# coding=utf8

import os, sys, time
import requests, json, httplib

def patch_send():
    old_send= httplib.HTTPConnection.send
    def new_send(self, data):
        print(data)
        return old_send(self, data) #return is not necessary, but never hurts, in case the library is changed
    httplib.HTTPConnection.send = new_send

patch_send()

def get_seq():

    url = "http://39.97.171.91:5690/seq/independent/dev-yunxun/next"
    # res = requests.get(url, headers=headers)
    res = requests.get(url)
    res = json.loads(res.content)

    print(res)

if __name__ == "__main__":
    get_seq()

