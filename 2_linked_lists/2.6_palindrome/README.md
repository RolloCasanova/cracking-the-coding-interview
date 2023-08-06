# Palindrome

Implement a function to check if a linked list is a palindrome.

## Hints

<details>
    <summary>Hint #5</summary>
    A palindrome is something which is the same when written forwards and backwards. What if you reversed the linked list?
</details>

<details>
    <summary>Hint #13</summary>
    Try using a stack.
</details>

<details>
    <summary>Hint #29</summary>
    Assume you have the length of the linked list. Can you implement this recursively?
</details>

<details>
    <summary>Hint #61</summary>
    In the recursive approach (we have the length of the list), the middle is the base case: isPermutation(middle) is true. The node x to the immediate left of the middle: What can that node do to check if x -> middle -> y forms a palindrome? Now suppose that checks out. What about the previous node a? If x -> middle -> y is a palindrome, how can it check that a -> x -> middle -> y -> b is a palindrome?
</details>

<details>
    <summary>Hint #101</summary>
    Go back to the previous hint. Remember: There are ways to return multiple values. You can do this with a new class.
</details>
