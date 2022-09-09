#include <iostream>
#include <math.h>
#include <string>

using namespace std;


int main() {
    FILE* fptr;
    
    fptr = fopen("random.txt", "w");

    int i;
    if(fptr == NULL) {
        cout << "Error";
        exit(1);
    }

    unsigned long long int value = pow(10, 18);

    for(i = 0; i < 1000 * 20 * 1000; i++) {
        unsigned long long int val = rand() % value;
        fprintf(fptr, "%d ", val);
    }

    fclose(fptr);
    cout << "numbers generated successfully !!";
    return 0;
}