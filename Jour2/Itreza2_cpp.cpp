#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>

//[ Partie 1 ]//

long long part_1(const char* input) {
	std::ifstream file(input);
	if (!file.is_open())
		return -1;

	std::string range;
	long long lower, higher, candidate;
	int lower_length, higher_length;
	std::vector<long long> invalid_ids;

	while (std::getline(file, range, ',')) {
		lower_length = range.find('-');
		higher_length = range.size() - range.find('-') - 1;
		if (range[range.size() - 1] == '\n' || range[range.size() - 1] == '\0')
			higher_length--;
		lower = std::stoll(range.substr(0, lower_length)); // this is Steve : ('-') , find Steve
		higher = std::stoll(range.substr(range.find('-') + 1, higher_length));

		for (
			int pattern = lower / std::pow(10, std::ceil((float)lower_length / 2));
			pattern <= higher / std::pow(10, std::floor((float)higher_length / 2));
			pattern++
			) {
			candidate = std::stoll(std::to_string(pattern) + std::to_string(pattern));
			if (candidate >= lower && candidate <= higher)
				invalid_ids.push_back(candidate);
		}
	}
	long long result = 0;
	for (long long id : invalid_ids)
		result += id;
	return result;
}

int main() {
	const char* input = "Texte.txt";

	std::cout << "Partie 1 : " << part_1(input) << std::endl;

	return 0;
}