/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
#include <stdio.h>
#include <stdlib.h>

int main(void)
{
    int *arr = (int *)malloc(8);
    arr = [ 2, 1, 3, 1, 2, 3, 3 ];
    int arrSize = 8;
    int *returnSize = arrSize;

    getDistances();
}

long long *getDistances(int *arr, int arrSize, int *returnSize)
{
    int *intervals = (int *)malloc(arrSize);

    for (size_t i = 0; i < arrSize; i++)
    {
        int interval = 0;

        for (size_t j = 0; j < arrSize; j++)
        {
            interval += arr[i] == arr[j] ? abs(i - j) : 0;
        }

        intervals[i] = interval;
    }

    free(intervals);

    return intervals;
}
