#include <iostream>
#include "cdll_1301194292.h"
using namespace std;

int main()
{
    List L;
    createList(L);
    address ad, a;
    createNewElmt('a', ad);
    insertFirst(L, ad);
    createNewElmt('t', ad);
    insertFirst(L, ad);
    a = findElement(L, 'a');
    createNewElmt('t', ad);
    insertAfter(a, ad);
    a = findElement(L, 't');
    createNewElmt('c', ad);
    insertAfter(a, ad);
    createNewElmt('a', ad);
    insertAfter(prev(first(L)), ad);
    createNewElmt('s', ad);
    insertAfter(prev(first(L)), ad);
    createNewElmt('c', ad);
    insertAfter(prev(first(L)), ad);
    createNewElmt('a', ad);
    insertAfter(prev(first(L)), ad);

    printInfo(L);
    char arr[3] = {'c', 'a', 't'};
    int i = countWord(arr, L);
    cout<<i<<endl;


}
