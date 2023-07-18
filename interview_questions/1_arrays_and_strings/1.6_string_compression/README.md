# String Compression

Implement a method to perform basic string compression using the counts of repeated characters. For example, the string `aabcccccaaa` would become `a2b1c5a3`. If the "compressed" string would not become smaller than the original string, your method should return the original string. You can assume the string has only uppercase and lowercase letters (a-z).

## Hints

<details>
    <summary>Hint #92</summary>
    Do the easy thing first. Compress the string, then compare the lengths.
</details>

<details>
    <summary>Hint #110</summary>
    Be careful that you aren't repeatedly concatenating strings together. This can be very inefficient.
</details>

<details>
    <summary>Hint #120</summary>
    Think about how you can optimize the runtime. Is there some way you can combine the loops?
</details>
