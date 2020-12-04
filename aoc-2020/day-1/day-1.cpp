#include <iostream>
#include <fstream>
#include <string>
using namespace std;

int main(int argc, char ** argv) {
	// create file object and read input file
	fstream file;
	file.open("input.txt", ios::in);
	
	// check whether the file is open
	if (file.is_open()) {
		string tp;

		// read data from file object and put it into string.
		while (getline(file, tp))
		{
			// print the data of the string
			cout << tp << "\n";
		}

		// close the file object.
		file.close();
	}
}