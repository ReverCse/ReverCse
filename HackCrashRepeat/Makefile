# My Makefile for CPSC4050 using some advanced Make Concepts

## Compiler
CC=g++
## Linker
LD=$(CC)
## Flags
CPPFLAGS = -Wall -g -std=c++11 -DLINUX
LFLAGS = -lglut -L/usr/lib -L/usr/X11R6/lib -lXi -lXmu -lGL -lGLU -lm -fopenmp
MACFLAGS = -framework glut -framework opengl
INCLUDE = -I include
SRC = $(wildcard src/*.cpp)
OBJECTS = $(SRC:.cpp=.o)
MAIN = assn3

all: $(MAIN)

$(MAIN): $(OBJECTS)
	$(CC) $(CPPFLAGS) $(INCLUDE) -o $(MAIN) $(OBJECTS) $(LFLAGS)

.cpp.o:
	$(CC) $(CPPFLAGS) $(INCLUDE) $(LFLAGS) -c $< -o $@

mac:
	$(CC) $(CPPFLAGS) $(MACFLAGS) $(INCLUDE) $(SRC) -o mac
	rm -rf *.dSYM

clean:
	rm -f $(wildcard src/*.o)
	rm -f $(MAIN)

run:
	./assn3 models/model1.obj

depend :
	makedepend ${PROGFILES}
