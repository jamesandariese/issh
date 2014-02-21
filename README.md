issh
====

Insecure SSH

Copyright (c) 2014 James Andariese

A secure shell client which uses known private keys to easily allow
for forced commands to be run.

Usage:
issh <username> <hostname> <key-seed>

  Connect to host <hostname> using the private key based on the
  seed token, <key-seed>.  This depends on there being a forced
  command setup on the remote end for the key.  Use -K to print out
  a suitable public key.

issh -K <key-seed>

  Print out the public key generated from the private key based on
  the seed token, <key-seed>
