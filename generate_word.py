import sys
import itertools

try:
    from itertools import imap
except ImportError:
    imap=map

string = "AB"
length = "5"

file = open("word-list.txt","w+")
a = []
for i in imap(''.join, itertools.product(string, repeat=int(length))):
    a.append(i+"\n")
file.writelines(a)
