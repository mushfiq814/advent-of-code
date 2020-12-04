/**
 * Day 1 Part 1
 * find two entries that sum to 2020 and then
 * multiply those two numbers together.
 */ 

#include <iostream>
#include <fstream>
#include <string>
#include <vector>

using namespace std;

int productOf2020Sum2(vector<int> &arr)
{
	int i, j;
	for (i = 0; i < arr.size(); i++)
		for (j = 1; j < arr.size(); j++)
			if (arr[i] + arr[j] == 2020) 
				return arr[i] * arr[j];
}

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

		cout << productOf2020Sum2(arr) << endl;

		file.close();
	}
}