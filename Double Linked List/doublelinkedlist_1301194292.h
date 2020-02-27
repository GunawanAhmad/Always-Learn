#ifndef DOUBLELINKEDLIST_H_INCLUDED
#define DOUBLELINKEDLIST_H_INCLUDED
#include <iostream>
#define first(L) L.first
#define last(L) L.last
#define next(P) P->next
#define prev(p) p->prev
#define info(P) P->info

using namespace std;
typedef int infotype;
typedef struct elmlist *address;

struct elmlist {
infotype info;
address next;
address prev;
};

struct List{
address first;
address last;
};

void createList(List &L);
void createNewElm(address&p, infotype x);
bool emptyList(List L);
void inserFirst (List &L, address p);
void insertAfter(List &L, address prec, address &p);
void insertLast(List &L, address p);
void deleteFirst(List &L, address p);
void deleteLast(List &L, address p);
address findElement(List L, infotype x);
void printInfo(List L);
void deleteAfter(List &L, address prec, address &p);
void concat(List L1, List L2, List &L3);
void ascending(List &L, address p);
double median(List L);



#endif // DOUBLELINKEDLIST_H_INCLUDED

//Nama : Gunawan Nur Ahmad
//NIM : 1301194292
