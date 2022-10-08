#include <iostream>
#include <string>
using namespace std;

void key0(string &str){
    str.push_back('0'); 
}

void key1(string &str){
    str.push_back('1'); 
}

void keyb(string &str){
//    str.pop_back();  // c++11 g++  -std=c++11  abc043_b.cpp
    if (str.empty()){
        return;
    }
    str.erase(str.end()-1);
}

int main()
{
    string str = "";

    string in_str;
    cin >> in_str;
    for (int i = 0; i < in_str.size(); i++){
//        cout << in_str[i] << endl;
        if (in_str[i] == '0'){
            key0(str);
        }else if (in_str[i] == '1'){
            key1(str);
        }else if (in_str[i] == 'B'){
            keyb(str);
        }
    }
    cout << str << endl;
    return 0;
}
