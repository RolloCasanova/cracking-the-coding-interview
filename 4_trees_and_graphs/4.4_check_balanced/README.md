# Check Balanced

Implement a function to check if a binary tree is balanced. For the purposes of this question, a balanced tree is defined to be a tree such that the heights of the two subtrees of any node never differ by more than one.

## Hints

<details>
    <summary> #21 </summary>
    Note that for each node, we could recursively compute the heights of each subtree. This could be somewhat inefficient, especially if you have to repeatedly compute heights for the same nodes. What if you could modify the `isBalanced` function to not only check if the tree is balanced, but also to return the height of the tree if the tree is balanced?
</details>

<details>
    <summary> #33 </summary>
    Each subtree could be balanced, but if the root node has two subtrees that are imbalanced, the tree is not balanced. Think about how to check for this condition efficiently as possible.
</details>

<details>
    <summary> #49 </summary>
    If you've developed a brute force solution, be careful about its runtime. If you are computing the height of the subtrees for each node, you could have a pretty inefficient algorithm.
</details>

<details>
    <summary> #105 </summary>
    What if you could modify the `isBalanced` function to check the height of each subtree while also checking if a subtree is balanced? Try having the function check the height of its subtree as it checks the subtrees balance. What would be the runtime in this case?
</details>

<details>
    <summary> #124 </summary>
    The height of a subtree is either balanced, or it isn't. If all possible subtrees must be balanced, what does this imply about the entire tree?
</details>
