/**
 * Day 1 Part 2
 * find three entries that sum to 2020 and then
 * multiply those three numbers together.
 */ 

#include <iostream>
#include <fstream>
#include <string>
#include <vector>

using namespace std;

int productOf2020Sum3(vector<int> &arr)
{
	int i, j, k;
	for (i = 0; i < arr.size(); i++)
		for (j = 1; j < arr.size(); j++)
			for (k = 1; k < arr.size(); k++)
				if (arr[i] + arr[j] + arr[k] == 2020) 
					return arr[i] * arr[j] * arr[k];
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

		cout << productOf2020Sum3(arr) << endl;

		file.close();
	}
}