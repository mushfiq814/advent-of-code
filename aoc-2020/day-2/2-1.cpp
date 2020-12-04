/**
 * Day 2 Part 1
 * Each line of the input file gives the password policy and then the password.
 * The password policy indicates the lowest and highest number of 
 * times a given letter must appear for the password to be valid.
 * For example, 1-3 a means that the password must contain a at 
 * least 1 time and at most 3 times.
 * 
 * How many passwords are valid according to their policies?
 */ 

#include <iostream>
#include <fstream>
#include <string>
#include <vector>

using namespace std;

int main(int argc, char ** argv) {
	// create file object and read input file
	fstream file;
	file.open("input.txt", ios::in);
	
	// check whether the file is open
	if (file.is_open()) {
		vector <int> arr;
		string tp;

		// build list of nums
		while (getline(file, tp))
			arr.push_back(stoi(tp));

		

		file.close();
	}
}