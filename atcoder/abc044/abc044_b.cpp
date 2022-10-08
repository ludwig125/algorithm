#include <iostream>
#include <string>

using namespace std;

int main()
{
    string w;
    cin >> w;
    
//    cout << w << endl;
//    cout << w.size() << endl;
    int char_num;
    for (int i = 0; i < w.size(); i++){
//        cout << w[i] << endl;
        char_num = 0;
        for (int j = 0; j < w.size(); j++){
            //cout << w[i] << w[j] << endl;        
            if(w[i] == w[j]){
                char_num += 1;
            }
        }
//        cout << "char_num " << char_num << endl;
        if(char_num % 2 != 0){
            cout << "No" << endl;
            return 0;
        }
    }
    cout << "Yes" << endl;

    return 0;
}
