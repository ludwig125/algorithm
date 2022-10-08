#include <iostream>
#include <cmath>

using namespace std;


long long int kaijo(int x){
    if (x > 0) {
        return x * kaijo(x - 1);
    }
    return 1;
}



// 高さh、幅wのマス目を移動する場合の数を返す
long long int kumi(int h, int w) {
    // 高さh、幅wのマス目を移動する場合の数は
    // (h-1 + w-1) C (h-1) ※  Cは組み合わせ
    return kaijo(h+w-2)/kaijo(h-1)/kaijo(w-1);

//    long long int tmp1 = kaijo(h+w-2);
//    cout << "tmp1"<< tmp1 << endl;
//    long long int tmp2 = kaijo(h-1);
//    cout << "tmp2"<< tmp2 << endl;
//    long long int tmp3 = kaijo(w-1);
//    cout << "tmp3"<< tmp3 << endl;
//    return tmp1/tmp2/tmp3;
    

}

long long int kakeawase(int H, int W, int B, int i){
    long first  = kumi(i, B);
    long second = kumi(H-(i-1), W-B);
    cout << first << " " << second << endl;
    return first * second;
}

int main()
{
    long MOD_NUM = pow(10, 9.0) + 7;

    int H, W, A, B;
    cin >> H >> W >> A >> B;
    long long int baai=0;
    for (int i=0; i < H-A; i++){
        baai += kakeawase(H, W, B, i+1);
        if (baai > MOD_NUM) {
            baai -= MOD_NUM;
        }
    }
    cout << baai << endl;    

    return 0;
}
