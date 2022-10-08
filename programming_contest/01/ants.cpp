#include <iostream>

using namespace std;

double abs(int L, int x){
    double from_mid = (double)L/2 - x;
    if(from_mid < 0){
        return from_mid * (-1);
    }
    return from_mid;
}

int main()
{

    int L, n;
    cin >> L;
    cin >> n;

    int x[n];
    for(int i = 0; i < n; i++){
        cin >> x[i];
    }

    double min_from_mid = (double)L;
    double max_from_mid = (double)0;
    double from_mid;
    for(int i = 0; i < n; i++){
        from_mid = abs(L, x[i]);
        if(from_mid < min_from_mid){
            min_from_mid = from_mid;
        }
        if(from_mid > max_from_mid){
            max_from_mid = from_mid;
        }
    }
    cout << "min " << (int)(L/2 - min_from_mid) << endl;
    cout << "max " << (int)(L/2 + max_from_mid) << endl;

    return 0;
}
