#include <iostream>
#include "doublelinkedlist.h"

using namespace std;

int main()
{
  List L1,L2,L3,Lb,La;
   createList(L1);
   createList(L2);
   createList(L3);
   address ad;
   createNewElm(ad, 1);
   inserFirst(L1, ad);
   createNewElm(ad, 3);
   insertLast(L1, ad);
   cout<<"List 1 : ";
   printInfo(L1);
   cout<<endl;
   cout<<"insert First 0 : ";
   createNewElm(ad, 0);
   inserFirst(L1, ad);
   printInfo(L1);
   createNewElm(ad, 2);
   address q = findElement(L1, 1);
   insertAfter(L1, q, ad);
   cout<<"Insert After 2 : ";
   printInfo(L1);
   createNewElm(ad, 4);
   insertLast(L1, ad);
   cout<<"Insert Last 4 : ";
   printInfo(L1);

   cout<<"List 1 : ";
   printInfo(L1);
   cout<<endl;


   int i = 5;
   while (i<=10) {
    createNewElm(ad, i);
    i = i+1;
    insertLast(L2, ad);
   }
   cout<<"List 2 : ";
   printInfo(L2);
   address p;
   deleteFirst(L2, p);
   cout<<"Delete First : ";
   printInfo(L2);
   deleteLast(L2, p);
   cout<<"Delete Last : ";
   printInfo(L2);
   address w = findElement(L2, 7);
   deleteAfter(L2, w, p);
   cout<<"Delete After 7 : ";
   printInfo(L2);
   cout<<"List 2 : ";
   printInfo(L2);
   cout<<endl;
   concat(L1,L2,L3);
   cout<<"concatenate List 1 dan List 2 : ";
   printInfo(L3);
   cout<<endl;
   i = 1;
   createList(La);
   while(i<=10) {
    createNewElm(ad, i);
    i = i+1;
    insertLast(La, ad);
   }
   cout<<"List La : ";
   printInfo(La);

   double a = median(La);
   cout<<"Median List La : "<<a;
   cout<<endl;

     i = 1;
   createList(Lb);
   while(i<=7) {
    createNewElm(ad, i);
    i = i+1;
    insertLast(Lb, ad);
   }
   cout<<"List La : ";
   printInfo(Lb);

   double b = median(Lb);
   cout<<"Median List Lb : "<<b;



    return 0;
}

//Nama : Gunawan Nur Ahmad
//NIM : 1301194292
