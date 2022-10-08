#include <iostream>
#include <vector>
//#include <algorithm>

using namespace std;

int main()
{

    int n;
    cin >> n;

    int a[n];
    for(int i = 0; i < n; i++){
        cin >> a[i];
    }


    int tmp;    
    for(int i = 0; i < n-1; i++){
        for(int j = i+1; j < n; j++){
            if(a[i] < a[j]){
                tmp = a[i];
                a[i] = a[j];
                a[j] = tmp;
            }
        }
    }
//    sort(a, a + n-1, greater<int>());
//    for(int k = 0; k < n; k++){
//        cout << a[k] << endl;
//    }

    int side[3];
    for(int i = 0; i < n-2; i++){
        side[0] = a[i];
        for(int j = i+1; j < n-1; j++){
            side[1] = a[j];
            for(int k = j+1; k < n; k++){
                side[2] = a[k];
                if (side[0] < (side[1]+side[2])){
                    cout << "side0: " << side[0] << ", side1: " << side[1] << ", side2: " << side[2]
                    << " sum " << side[0] + side[1] + side[2] << endl;
                    return 0;
                }
            }
        }
    }
    cout << 0 << endl;
    return 0;
}
