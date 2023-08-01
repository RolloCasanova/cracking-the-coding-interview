# Partition

Write code to partition a linked list around a value x, such that all nodes less than x come before all nodes greater than or equal to x. If x is contained within the list, the values of x only need to be after the elements less than x (see below). The partition element x can appear anywhere in the "right partition"; it does not need to appear between the left and right partitions.

## Example

``` nocode
Input:  3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1 [partition = 5]
Output: 3 -> 1 -> 2 -> 10 -> 5 -> 5 -> 8
```

## Hints

<details>
    <summary>Hint #3</summary>
    Consider that the elements don't have to stay in the same relative order. We only need to ensure that elements less than the pivot must be before elements greater than the pivot. Does that help you come up with more solutions?
</details>

<details>
    <summary>Hint #24</summary>
    Observe that `3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1` is equivalent to `3 -> 1 -> 2 -> 10 -> 5 -> 5 -> 8`. What's the relationship between these two linked lists? Can you see a way of transforming the first into the second?
</details>
