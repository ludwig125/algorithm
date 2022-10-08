#include <iostream>
#include <algorithm>
#include <vector>

using namespace std;
int main()
{
    int N;
    cin >> N;
    
    vector<int> A;
    int num;
    for(int i=0; i<N; i++){
        cin >> num;
        A.push_back(num);
    }

    std::sort(A.begin(), A.end() );

    int min_cost;

    int cost, target_num;
    for(int i = A.front(); i <= A.back(); i++){
//        cout << i << endl;
        cost = 0;
        target_num = i;
    
        for(vector<int>::iterator it = A.begin(); it != A.end(); it++){
//            cout << "A: "  << *it << endl;
            cost += (*it - target_num) * (*it - target_num);
        }
//        cout << "cost: " << cost << endl;
        if ( i == A.front() || min_cost > cost){
            min_cost = cost;
        }
    }
    cout << min_cost << endl;

    return 0;
}
