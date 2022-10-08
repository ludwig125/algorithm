#include <iostream>

using namespace std;

int sig(int n){
    if (n > 0){
        return n + sig(n - 1);
    }
    return 0;
}


int main()
{
    int N;
    cin >> N;
    
    cout << sig(N) << endl;

    return 0;
}
