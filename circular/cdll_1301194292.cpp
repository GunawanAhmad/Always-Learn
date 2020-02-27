#include <iostream>
#include "cdll_1301194292.h"

using namespace std;

bool isEmpty(List L) {
    if (first(L)==NULL) {
        return true;
    } else {
        return false;
    }
}
void createList(List &L) {
    first(L) = NULL;
}
void createNewElmt(infotype x, address &p) {
    p = new elmlist;
    info(p) = x;
    next(p) = NULL;
    prev(p) = NULL;

}
void insertFirst(List &L, address p) {
    bool a = isEmpty(L);
    if (a==true) {
        first(L) = p;
        next(first(L)) = p;
        prev(first(L)) = p;
    } else {
        next(prev(first(L))) = p;
        prev(p) = prev(first(L));
        prev(first(L)) = p;
        next(p) = first(L);
        first(L) = p;
    }

}

void insertAfter(address prec, address p) {
        prev(next(prec)) = p;
        next(p) = next(prec);
        next(prec) = p;
        prev(p) = prec;
}

void printInfo(List L) {
    address q = first(L);
    cout<<info(q)<<" ";
    q = next(q);
    while (q != first(L)) {
        cout<<info(q)<<" ";
        q = next(q);
    }
    cout<<endl;

}

address findElement(List L, infotype x) {
    address p = first(L);
    if (info(p)==x) {
        return p;
    }
    while (next(p)!=first(L) && info(p)!=x) {
        p = next(p);
        if (info(p)==x) {
            return p;
        }
    }
    return NULL;
}

void deleteFirst(List &L, address &p) {
    p = first(L);
    prev(next(p)) = prev(p);
    next(prev(p)) = next(p);
    first(L) = next(p);
    next(p) = NULL;
    prev(p) = NULL;
}

void deleteAfter(address prec, address &p) {
    p = next(prec);
    next(prec) = next(p);
    prev(next(p)) = prec;
    next(p) = NULL;
    prev(p) = NULL;
}

int countWord(char arr[], List L) {
    address a,b,c;
    a = first(L);
    b = next(a);
    c = next(b);
    int hasil = 0;

    while (b!= first(L)) {
        if (arr[0]==info(a) && arr[1]==info(b) && arr[2]==info(c)) {
            hasil  = hasil + 1;
        }
        a = next(a);
        b = next(b);
        c = next(c);
    }
    return hasil;
}
