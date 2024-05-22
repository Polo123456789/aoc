CXX=/usr/bin/clang++-18
CXXFLAGS=-std=c++17 \
			-stdlib=libc++ \
			-fsanitize=address \
			-Weverything \
			-Wno-c++98-compat-pedantic \
			-Wno-c++98-compat \
			-Wno-global-constructors \
			-Wno-exit-time-destructors \
			-Wno-documentation-unknown-command

SRC=$(wildcard *.cpp)
OBJ=$(SRC:.cpp=.o)
OUT=main

all: $(OUT)

$(OUT): $(OBJ)
	@echo "Linking $@"
	@$(CXX) $(CXXFLAGS) -o $@ $^

%.o: %.cpp
	@echo "Compiling $<"
	@$(CXX) $(CXXFLAGS) -c -o $@ $<

clean:
	rm -f $(OBJ) $(OUT)
