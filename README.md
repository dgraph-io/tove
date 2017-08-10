# Tove: Crash vulnerability tests for Badger using the ALICE framework

Tove is named after the Badger character from Lewis Carroll's Alice in Wonderland.

ALICE is a framework developed by the University of Wisconsin - Madison. It's
accompanied by a paper named [All File Systems Are Not Created Equal: On the
Complexity of Crafting Crash-Consistent
Applications](http://research.cs.wisc.edu/adsl/Publications/alice-osdi14.pdf).

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
```
