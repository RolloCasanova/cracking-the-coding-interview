# Stack of Plates

## Problem

Imagine a (literal) stack of plates. If the stack gets too high, it might topple. Therefore, in real life, we would likely start a new stack when the previous stack exceeds some threshold.

Implement a data structure `SetOfStacks` that mimics this. `SetOfStacks` should be composed of several stacks and should create a new stack once the previous one exceeds capacity. `SetOfStacks.push()` and `SetOfStacks.pop()` should behave identically to a single stack (that is, `pop()` should return the same values as it would if there were just a single stack).

## Follow Up

Implement a function `popAt(int index)` which performs a pop operation on a specific sub-stack.

## Hints

<details>
    <summary>#64</summary>
    You will need to keep track of the size of each sub-stack. When one stack is full, you may need to create a new stack.
</details>

<details>
    <summary>#81</summary>
    Popping an element at a specific sub-stack will mean that some stacks are not at full capacity. Is this an issue? There's no right answer, but you should think about how to handle this.
</details>
