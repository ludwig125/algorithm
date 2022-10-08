#include <iostream>
#include <algorithm>
#include <vector>

using namespace std;

int kaijo(int n){
    if (n > 0) {
        return n * kaijo(n - 1);
    }
    return 1;
}

int kaijo_n_k(int n, int k){
    if (n > k) {
        return n * kaijo_n_k(n - 1, k);
    }
    return 1;
}



int main()
{
    int n;
    cin >> n;
//    cout << kaijo(n) << endl;

    int k;
    cin >> k;
    cout << kaijo_n_k(n, k) << endl;

    return 0;
}
