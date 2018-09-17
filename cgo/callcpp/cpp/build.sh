#!/bin/bash

# C call cpp
#g++ -Wall -c person.cpp wrapper.cpp
#gcc -Wall -c hello.c
#g++ -o test *.o
#rm *.o

# Cpp to lib
#g++ -Wall -c person.cpp wrapper.cpp
#ar -rv libhello.a *.o
#rm *.o
#gcc -Wall -c hello.c
#g++ -o test *.o -L. -lhello
#rm *.o

# All to lib
g++ -Wall -c person.cpp wrapper.cpp
gcc -Wall -c hello.c
#ar -crv libhello.a *.o
ar -crsv libhello.a *.o
rm *.o

