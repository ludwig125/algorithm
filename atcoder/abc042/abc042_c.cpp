#include <iostream>
#include <algorithm>
#include <vector>

using namespace std;
int main()
{
    int N, K;
    cin >> N >> K;
    
    vector<int> v;
    int num;
    // $B7y$$$J?t;z$rEPO?(B
    for(int i=0; i<K; i++){
        cin >> num;
        v.push_back(num);
    }

    int n, ans = 0;
    bool flg = false;
    ans = N;
    while(!flg){
        N = ans;
        while(N!=0){
            // $B$R$H7eL\$+$i=g$K(Bn$B$KF~$l$F7y$$$J?t;z$H0lCW$7$F$$$J$$$+Hf3S(B
            n = N % 10;
            for(vector<int>::iterator it = v.begin(); it != v.end(); it++){
                if(n != *it){
                    flg = true;
                }else {
                    flg = false;
                    break;
                }
            }
            if(!flg){
                break;
            }
            N /= 10;
        }
        if(flg){
            cout << ans << endl;
            exit;
        }else{
            ans++;
        }
    }

    return 0;
}
