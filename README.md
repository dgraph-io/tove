# Tove: Crash vulnerability tests for Badger using the ALICE framework

Tove is named after the Badger character from Lewis Carroll's Alice in Wonderland.

ALICE is a framework developed by the University of Wisconsin - Madison. It's
accompanied by a paper named [All File Systems Are Not Created Equal: On the
Complexity of Crafting Crash-Consistent
Applications](http://research.cs.wisc.edu/adsl/Publications/alice-osdi14.pdf).

**It's important to read this paper to understand how to interpret ALICE output
and write/modify ALICE tests.**

## ALICE Installation

Clone ALICE from `git@github.com:dgraph-io/alice.git`. This is a Dgraph fork of
ALICE that implements some additional features and tweaks to get ALICE running
with Badger.

```
$ git clone git@github.com:dgraph-io/alice.git
```

Some dependencies also have to be installed (`libunwind` and `bitvector`). On Arch Linux:

```
# pacman -S libunwind
# pip2 install bitvector
```

ALICE also relies on environment vars to run. `ALICE_HOME` must be set to the
directory that you cloned ALICE into. `ALICE_HOME/bin` must be in your path.
E.g. on Arch Linux:

```
$ echo 'export ALICE_HOME=/home/$USER/alice' >> ~/.bashrc
$ echo 'export PATH=$ALICE_HOME/bin:$PATH' >> ~/.bashrc
```

The `alice-strace` program must then be installed:

```
$ cd $ALICE_HOME/alice-strace
$ ./configure
$ make
# make install
```

By default, the auto load safe path for GDB is too strict for debugging Go
applications. One way to fix this is to effectively disable it by allowing
loading of any scripts (*this could potentially be a security risk*):

```
$ echo "set auto-load safe-path /" >> ~/.gdbinit
```

## Running the Tests

`cd` into the directory containing a particular test, then run the `run.sh`
script. E.g.:

```
$ cd badger
$ ./run.sh
start:big
stop:big
GDB warnings while processing
/home/petsta/go/src/github.com/dgraph-io/tove/badger/bin/workload: Loading Go
Runtime support.
-------------------------------------------------------------------------------
ALICE tool version 0.0.1. Please go through the documentation, particularly the
listed caveats and limitations, before deriving any inferences from this tool.
-------------------------------------------------------------------------------
Parsing traces to determine logical operations ...
Warning: trunc_disk_ops called for the same initial and final size, 0
Logical operations:
0       stdout("'start:big\n'")
1       creat("LOCK", parent=1309541, mode='0666', inode=1310326)
2       append("LOCK", offset=0, count=5, inode=1310326)
3       creat("MANIFEST", parent=1309541, mode='0666', inode=1310327)
4       trunc("MANIFEST", initial_size=0, inode=1310327, final_size=0)
5       fsync(".", size=100, inode=1309541)
6       creat("000000.vlog", parent=1309541, mode='0666', inode=1310332)
7       fsync(".", size=120, inode=1309541)
8       append("000000.vlog", offset=0, count=1048670, inode=1310332)
9       append("000000.vlog", offset=1048670, count=1048670, inode=1310332)
10      append("000000.vlog", offset=2097340, count=1048670, inode=1310332)
11      append("000000.vlog", offset=3146010, count=1048670, inode=1310332)
12      append("000000.vlog", offset=4194680, count=1048670, inode=1310332)
13      fsync("000000.vlog", size=5243350, inode=1310332)
14      creat("000001.vlog", parent=1309541, mode='0666', inode=1310337)
15      fsync(".", size=140, inode=1309541)
16      append("000001.vlog", offset=0, count=1048670, inode=1310337)
17      append("000001.vlog", offset=1048670, count=2097340, inode=1310337)
18      unlink("000000.vlog", parent=1309541, inode=1310332, size=5243350, hardlinks=1)
19      creat("000001.sst", parent=1309541, mode='0666', inode=1310342)
20      fsync(".", size=160, inode=1309541)
21      append("000001.sst", offset=0, count=470, inode=1310342)
22      file_sync_range("000001.sst", count=470, offset=0, inode=1310342)
23      append("MANIFEST", offset=0, count=8, inode=1310327)
24      fsync("MANIFEST", size=8, inode=1310327)
25      unlink("LOCK", parent=1309541, inode=1310326, size=5, hardlinks=1)
26      fsync(".", size=160, inode=1309541)
27      fsync(".", size=160, inode=1309541)
28      stdout("'stop:big\n'")
-------------------------------------
Finding vulnerabilities...
Warning: trunc_disk_ops called for the same initial and final size, 0
Warning: trunc_disk_ops called for the same initial and final size, 0
(Dynamic vulnerability) Atomicity: Operation 9(???)
(Dynamic vulnerability) Atomicity: Operation 10(???)
(Dynamic vulnerability) Atomicity: Operation 11(???)
(Dynamic vulnerability) Atomicity: Operation 12(???)
(Dynamic vulnerability) Atomicity: Operation 16(???)
(Dynamic vulnerability) Atomicity: Operation 17(???)
(Dynamic vulnerability) Atomicity: Operation 18(,  semi-truncated (3 count splits))
(Dynamic vulnerability) Atomicity: Operation 23(garbage written, zero written)
(Static vulnerability) Atomicity: Operation /usr/lib/go/src/syscall/asm_linux_amd64.s:27[syscall.Syscall] (,garbage written,???, semi-truncated (3 count splits) ,zero written)
Done finding vulnerabilities.
```
