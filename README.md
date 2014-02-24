= issh =
--------

Insecure SSH

Copyright (c) 2014 James Andariese

A secure shell client which uses known private keys to easily allow
for forced commands to be run.

Usage
-----

`issh <username> <hostname> <key-seed>`

  Connect to host <hostname> using the private key based on the
  seed token, <key-seed>.  This depends on there being a forced
  command setup on the remote end for the key.  Use -K to print out
  a suitable public key.

`issh -K <key-seed>`

  Print out the public key generated from the private key based on
  the seed token, <key-seed>

Description
-----------

The goal of this project is to make forced commands easier to manage by making a few assumptions that cause it to be extremely insecure.

First off, there is no host key checking.  The reasoning for this is simple: if you have a load balanced cluster, and want to do something
[horrible] to monitor it such as use ssh then forced commands are the way to go -- if you use forced commands, you will end up with ssh -o StrictHostKeyChecking=no -q -i select-a-command.pub -l user host eventually.  You'll likely reuse this key on multiple hosts for multiple purposes to avoid having to generate a new key every time.  You'll likely also have forgotten -q so you'll get a warning every time until you add the host key.  You'll add it first as your user (or root, probably) then try to figure out why your service data is still filled with warnings about SSH rather than the expected output.  Then you'll finally add the host key with su nagios.  Then, one fateful day, that box will die and you have to do it all over again with its replacement.  This will happen a few times until you just add the -q.

No host key checking.  It's good to have but you'll spend a lot of time bypassing it anyway.

Second, the key is mostly known to everyone (p, q, g are always known with issh) and the other is easily guessed because it's almost certainly both English and related to the command being run.  The point of issh is to allow it to be put into scripts and to allow things that are relatively harmless to be run.  If it's considerably more costly to do this than to just serve up a DDoS on ssh itself, it shouldn't be in a forced command run through issh.  One example of how to use this is to have the "ipvs" key mapped to "ipvsadm -L -n".

The third (and also only "assumption", actually) is in the second cause of insecurity: this assumes you're not doing stupid and dangerous things in forced commands.  You shouldn't be anyway but you really shouldn't be with issh.

You have been warned.

Tutorial
--------

In this walkthrough, we will export "ps auxww" via issh.  It will run as root.

1. Generate a key

  `issh -K processes`

  `command="exit 1",no-port-forwarding,no-X11-forwarding,no-pty ssh-dss AAAAB3NzaC1kc3MAAACBALoKZiM+G538VNup8FaIhU2Vvc8v4DMxYx5yZ5KaXKWfIzzEINXwbE+YxiOgRsoGqWw4sTvwPUoQH5M9Ve6Qw3c3LSQ6plhRi9rBuqmpnZuvjzVZkC0RSXnI4SbRWZ2g8XY3XKnMgXfOCLoRAGL6Cav7IuDk2kUnKdrL4X7zt+FdAAAAFQCXsf7T2nGz5JqoBcG9BIgPinFXcQAAAIAVxDWPdHOrhBwZRDPXS57N/qgt8fIt0e7R1L1KV+9dsPLvF4+u+yaEf21mAnwlLBfhZtLqR1cBS74TWLyjhCpyHcfLV99h9ZaJ5bd8G/FNCixg040gia5lI1mlYaF3oaf99wGFWfY/FqkiTPHxGJmnIMNlcenUEYeFKOHhlmWDPwAAAIEAmNzJtTZdi05AqVxZ2xpVLEt9iKIegGlnbfccnI0NHxSQewW+cyEdQYZcq66e9goSXoXX12kbCGzhsNCWrRr7gvnU65Y+GtdCtL7tFiNT/V5TgkdQ4ZhPwpckuGh2a9MsG+zxQuYyAV/w99ThJtf4CBzODA/IA3pkNvVqdSrUgpM= issh generated from seed: 'ps'`

2. Copy and paste the output into .ssh/authorized_keys on the remote host.  Note to change the command to `command="ps auxww"`.

3. Use issh to connect

  `issh root guadalajara processes`

  `USER    PID  %CPU %MEM    VSZ   RSS TT  STAT STARTED         TIME COMMAND
root     11 400.0  0.0      0    64  -  RL   31Jan14 115003:30.60 [idle]
james 80332   3.7  0.3  38796 15784  0  R+    7:17PM      0:17.35 emacs README.md (emacs-24.3)`

Note that the command is not part of the invocation.  By default, issh will run exit 1 with its key (same as the default that it prints with -K) to help protect against negligence.  You cannot specify a command except through command= in the authorized_keys file.

