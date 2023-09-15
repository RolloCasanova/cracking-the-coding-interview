# Minimal Tree

Given a sorted (increasing order) array with unique integer elements, write an algorithm to create a binary search tree with minimal height.

## Hints

<details>
    <summary> #19 </summary>
        Start with the middle element of the array. That would be the root of the tree. The left half of the array will become the left subtree, and the right half of the array will become the right subtree. Can you apply this to a recursive algorithm?
</details>

<details>
    <summary> #73 </summary>
        Think about creating the tree recursively. Try to find the root first. Then, what would the left child be? The right child? Recurse through each subarray.
</details>

<details>
    <summary> #116 </summary>
        Each parent node requires a left and right child. The root should be the middle of the array, since the array is sorted. So, the root's left child will be the middle of the left half of the array, and the root's right child will be the middle of the right half of the array. Can you apply this recursively?
</details>
