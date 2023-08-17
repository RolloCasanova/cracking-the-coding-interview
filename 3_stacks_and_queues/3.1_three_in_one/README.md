# Three in One

Describe how could you use a single array to implement three stacks?

## Hints

<details>
    <summary>Hint #2</summary>
    A stack is simply a data structure in which the most recently added elements are removed first. Can you simulate a single stack using an array? Remember that there are many possible solutions, and there are tradeoffs of each.
</details>

<details>
    <summary>Hint #12</summary>
    We could simulate three stacks in an array by just allocating the first third of the array to the first stack, the second third to the second stack, and the final third to the third stack. One might actually be much bigger than the others though. Can we be more flexible with the divisions?
</details>

<details>
    <summary>Hint #38</summary>
    If you want to allow for flexible divisions, you can shift stacks around. Can you ensure that all available capacity is used?
</details>

<details>
    <summary>Hint #58</summary>
    Try thinking about the array as circular, such that the end of the array "wraps around" to the start of the array.
</details>
