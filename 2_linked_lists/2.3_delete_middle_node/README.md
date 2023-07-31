# Delete Middle Node

Implement an algorithm to delete a node in the middle (i.e., any node but the first and last node, not necessarily the exact middle) of a singly linked list, given only access to that node.

## Example

Input: the node c from the linked list `a->b->c->d->e->f`

Result: nothing is returned, but the new linked list looks like `a->b->d->e->f`

## Hints

<details>
    <summary>Hint #72</summary>
    Picture the list `1->5->9->12`. Removing `9` would make it look like `1->5->12`. You only have access to the `9` node. Can you make it look like the correct answer?
</details>
