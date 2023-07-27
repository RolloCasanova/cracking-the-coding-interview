# One Away

There are three types of edits that can be performed on strings: insert a character, remove a character, or replace a character. Given two strings, write a function to check if they are one edit (or zero edits) away.

## Example

``` nocode
pale, ple -> true
pales, pale -> true
pale, bale -> true
pale, bake -> false
```

## Hints

<details>
    <summary>Hint #23</summary>
    Start with the easy thing. Can you check each of the conditions separately?
</details>

<details>
    <summary>Hint #97</summary>
    What is the relationship between the "insert character" option and the "remove character" option? Do these need to be two separate checks?
</details>

<details>
    <summary>Hint #130</summary>
    Can you do all three checks in a single pass?
</details>
