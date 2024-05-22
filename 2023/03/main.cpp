#include <iostream>

#include "engine_schematic.hpp"

int main(void) {
	engine_schematic engine(std::cin);
	// std::cout << engine << "\n---\n";
	std::cout << "Part 1: " << engine.get_sum_of_parts() << "\n";
	return 0;
}
