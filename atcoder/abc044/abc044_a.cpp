#include <iostream>

using namespace std;

int cost(int N, int K, int X, int Y){
    int cost = 0;
    for (int i = 1; i <= N; i++){
        if(i <= K){
            cost += X;
        }else{
            cost += Y;
        }
    }
    return cost;
}


int main()
{
    int N, K, X, Y;
    cin >> N;
    cin >> K;
    cin >> X;
    cin >> Y;
    
    cout <<  cost(N, K, X, Y) << endl;

    return 0;
}

