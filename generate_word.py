import sys
import itertools

try:
    from itertools import imap
except ImportError:
    imap=map

string = sys.argv[1]
length = sys.argv[2]

file = open("word-list.txt","w+")
a = []
for i in imap(''.join, itertools.product(string, repeat=int(length))):
    a.append(i+"\n")
file.writelines(a)
