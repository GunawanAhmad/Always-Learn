#include "doublelinkedlist.h"
#include<iostream>

void createList(List &L) {
    first(L) = NULL;
    last(L) = NULL;
}
void createNewElm(address &p, infotype x) {
    p = new elmlist;
    info(p) = x;
    next(p) = NULL;
    prev(p) = NULL;
}

bool emptyList(List L) {
    if (first(L)==NULL) {
        return true;
    } else {
        return false;
    }
}

void inserFirst (List &L, address p) {
    bool cek = emptyList(L);
     if (cek == true) {
        first(L) = p;
        last(L) = p;
    } else {
         next(p) = first(L);
         prev(first(L)) = p;
         first(L) = p;
    }

}

void insertAfter(List &L, address prec, address &p) {
    next(p) = next(prec);
    prev(next(p)) = p;
    next(prec) = p;
    prev(p) = prec;
}

void insertLast(List &L, address p) {
    bool cek = emptyList(L);
    if (cek == true) {
        first(L) = p;
        last(L) = p;
    } else {
        prev(p) = last(L);
        next(last(L)) = p;
        last(L) = p;
    }

}

void deleteFirst(List &L, address p) {
    p = first(L);
    first(L) = next(first(L));
    next(p) = NULL;
}

void deleteLast(List &L, address p) {
    p = last(L);
    next(prev(last(L))) = NULL;
    prev(last(L)) = last(L);



}

address findElement(List L, infotype x) {
    address p = first(L);
    if (info(p) == x) {
        return p;
    } else {
        while (p != NULL && (info(p)) != x) {
            p = next(p);
            if (info(p)==x) {
                return p;
            }

        }
    }
}

void printInfo(List L) {
    address p = first(L);
    while (p != NULL) {
        cout<<info(p)<<" ";
        p = next(p);
    }
    cout<<endl;
}

void deleteAfter(List &L, address prec, address &p) {
    p = next(prec);
    prev(next(p)) = prec;
    next(prec) = next(p);
    next(p) = NULL;
    prev(p) = NULL;

}

void concat(List L1, List L2, List &L3) {
    L3 = L1;
    address p = last(L3);
    address q = first(L2);
    next(p) = q;
    address a = next(p);
    while (q != NULL) {
        next(a) = next(q);
        a = next(a);
        q = next(q);
    }

}

void ascending(List &L, address p) {
    if (first(L) ==NULL) {
        first(L) = p;
        last(L) = p;
    } else if (info(p)< info(first(L))) {
        inserFirst(L, p);
    } else if (info(p)>info(last(L))) {
        insertLast(L, p);
    } else {
        address q;
        q = first(L);
        while(info(p)<info(q)) {
            q = next(q);
        }
        insertAfter(L,q,p);
    }
}

double median(List L) {
    int i = 0;
    address p = first(L);
    while (p!=NULL) {
        i = i+1;
        p = next(p);
    }
    if (i%2 == 0) {
        int hasil = i/2;
        int hasil2 = hasil +1;
        address q = first(L);
        int j=1;
        for (int i=0; i<hasil-1; i++) {
            q = next(q);

        }
        double h1 = double(info(q));
        p = first(L);
        for (int k=0; k<hasil2-1; k++) {
            p = next(p);
        }
        double h2 = double(info(p));
        double rent = (h1+h2)/2;
        return rent;

    } else {
        address a = first(L);
        double hasil = (double)(i/2);
        double i = 1.0;
        while(a != NULL && i<=hasil) {
            i = i+1;
            a = next(a);
        }
        double h3 = (double)info(a);
        return (double)h3;
    }
}


//Nama : Gunawan Nur Ahmad
//NIM : 1301194292
