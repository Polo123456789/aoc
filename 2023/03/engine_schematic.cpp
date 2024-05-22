#include "engine_schematic.hpp"

#include <algorithm>
#include <cctype>
#include <string>
#include <array>

constexpr std::array<engine_schematic::position, 8> positions_around = {
	{ { -1, -1 }, { 0, -1 }, { 1, -1 },
	  { -1, 0 },             { 1, 0 },
	  { -1, 1 },  { 0, 1 },   { 1, 1 } }
};

engine_schematic::engine_schematic(std::istream &in) {
	std::string line;
	while (std::getline(in, line)) {
		map.emplace_back(std::move(line));
		line.clear();
	}
}

std::ostream &operator<<(std::ostream &o, engine_schematic &e) {
	for (const auto &l : e.map) {
		o << l << '\n';
	}
	return o;
}

[[nodiscard]] char engine_schematic::get(position p) const {
	auto x = static_cast<size_type>(p.x);
	auto y = static_cast<size_type>(p.y);
	return map[y][x];
}

[[nodiscard]] bool engine_schematic::in_range(position p) const {
	if (p.x >= this->width() || p.x < 0) {
		return false;
	}

	if (p.y >= this->height() || p.y < 0) {
		return false;
	}

	return true;
}

[[nodiscard]] char engine_schematic::get_or_empty(position p) const {
	if (!in_range(p)) {
		return empty;
	}
	return this->get(p);
}

[[nodiscard]] auto engine_schematic::width(void) const -> index_type {
	if (map.size() == 0) {
		return 0;
	}
	// Assuming all lines are the same width
	return static_cast<index_type>(map[0].size());
}

[[nodiscard]] auto engine_schematic::height(void) const -> index_type {
	return static_cast<index_type>(map.size());
}

[[nodiscard]] auto
	engine_schematic::count_digits_at(position p) const -> index_type {
	index_type count = 0;
	while (p.x < this->width() && std::isdigit(this->get(p))) {
		count++;
		p.x++;
	}
	return count;
}

static bool is_simbol(char c) {
	return !std::isdigit(c) && c != engine_schematic::empty;
}

[[nodiscard]] bool engine_schematic::has_symbol_around(position p) const {
	return std::any_of(positions_around.begin(), positions_around.end(),
					   [this, p](const auto &pos) {
						   auto new_pos = position{
							   p.x + pos.x,
							   p.y + pos.y
						   };
						   return is_simbol(this->get_or_empty(new_pos));
					   });
}


[[nodiscard]] int engine_schematic::get_number_at(position p, index_type count) const {
	auto y = static_cast<size_type>(p.y);
	auto x = static_cast<size_type>(p.x);
	auto ccount = static_cast<size_type>(count);

	return std::stoi(map[y].substr(x, ccount));
}

// Here for debugging
using index_type = engine_schematic::index_type;
[[maybe_unused]] static void print_all_around(const engine_schematic &e,
							 index_type				 x_start,
							 index_type				 x_end,
							 index_type				 y_pos) {
	for (auto y = y_pos - 1; y <= y_pos + 1; y++) {
		for (auto x = x_start - 1; x <= x_end; x++) {
			std::cout << e.get_or_empty({x, y});
		}
		std::cout << '\n';
	}
}

[[nodiscard]] int engine_schematic::get_sum_of_parts(void) const {
	int sum = 0;
	for (index_type y = 0; y < this->height(); y++) {
		for (index_type x = 0; x < this->width(); x++) {
			auto count = this->count_digits_at({x, y});
			if (count != 0) {
				for (auto i = x; i < x + count; i++) {
					if (this->has_symbol_around({i, y})) {
						auto num = this->get_number_at({x, y}, count);
						sum += num;
						break;
					}
				}
			}
			x += count;
		}
	}
	return sum;
}

[[nodiscard]] int engine_schematic::get_sum_of_gear_ratios(void) const {
	return 0;
}
