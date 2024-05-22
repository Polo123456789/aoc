#ifndef ENGINE_SCHEMATIC_HPP
#define ENGINE_SCHEMATIC_HPP

#include <vector>
#include <string>
#include <iostream>

class engine_schematic {
   public:
	using map_type = std::vector<std::string>;
	using size_type = map_type::size_type;

	// I need negative indices
	using index_type = map_type::difference_type;

	struct position {
		index_type x;
		index_type y;
	};

	constexpr static char empty = '.';

	explicit engine_schematic(std::istream &in);

	// Returns the character at the given coordinates or empty if out of range
	[[nodiscard]] char get_or_empty(position p) const;

	// Returns the character without checking if it's in range
	[[nodiscard]] char get(position p) const;

	[[nodiscard]] bool		 in_range(position p) const;
	[[nodiscard]] index_type width(void) const;
	[[nodiscard]] index_type height(void) const;

	[[nodiscard]] int get_number_at(position p, index_type count) const;

	[[nodiscard]] int get_sum_of_parts(void) const;
	[[nodiscard]] int get_sum_of_gear_ratios(void) const;
   private:
	[[nodiscard]] index_type count_digits_at(position p) const;
	[[nodiscard]] bool has_symbol_around(position p) const;

	// Indexed by [y][x]
	map_type map;

	friend std::ostream &operator<<(std::ostream &o, engine_schematic &e);
};

#endif // ENGINE_SCHEMATIC_HPP
