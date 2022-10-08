#include <iostream>
#include <algorithm>
#include <vector>

using namespace std;
int main()
{
    int N, L;
    cin >> N >> L;
    vector<string> v;
    string ret;
    
    string s;
    for(int i=0; i<N; i++){
        cin >> s;
        v.push_back(s);
    }

    sort(v.begin(), v.end(), less<string>() );                        
    for(vector<string>::iterator it = v.begin(); it != v.end(); it++){
        ret += *it;
    }
    cout << ret << endl;

    return 0;
}
