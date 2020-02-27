#ifndef CDLL_1301194292_H_INCLUDED
#define CDLL_1301194292_H_INCLUDED
#define next(p) p->next
#define info(p) p->info
#define prev(p) p->prev
#define first(L) L.first
#include<iostream>

typedef char infotype;
typedef struct elmlist *address;
struct List {
    address first;
};
struct elmlist {
    infotype info;
    address next;
    address prev;

};

bool isEmpty(List L);
void createList(List &L);
void createNewElmt(infotype x, address &p);
void insertFirst(List &L, address p);
void insertAfter(address prec, address p);
void deleteFirst(List &L, address &p);
void deleteAfter(address prec, address &p);
int countWord(char arr[], List L);
void printInfo(List L);
address findElement(List L, infotype x);
#endif // CDLL_1301194292_H_INCLUDED
