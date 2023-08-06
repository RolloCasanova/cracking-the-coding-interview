# Cracking the coding interview

A collection of solutions for the book "Cracking the coding interview" by Gayle Laakmann McDowell (6th edition) made by me in Go just for fun.

## How to run

```bash
go run <chapter>/<exercise>/main.go [optional-params]
```

e.g.

```bash
go run 1_arrays_and_strings/1.1_is_unique/main.go abcde
```

## How to test

### A single exercise

```bash
go test -v ./<chapter>/<exercise>/...
```

e.g.

```bash
go test -v ./1_arrays_and_strings/1.1_is_unique/...
```

### All exercises of a chapter

``` bash
go test -v ./<chapter>/...
```

e.g.

``` bash
go test -v ./1_arrays_and_strings/...
```

### All exercises

```bash
go test -v ./...
```

### Progress

#### Chapter 1: Arrays and Strings (9/9)

|  #  | Problem                        | Solution         | Tests         | Solved |
|-----|--------------------------------|:----------------:|:-------------:|:------:|
| 1.1 | [Is Unique][p1.1]              | [solution][s1.1] | [tests][t1.1] |   ✅   |
| 1.2 | [Check Permutation][p1.2]      | [solution][s1.2] | [tests][t1.2] |   ✅   |
| 1.3 | [URLify][p1.3]                 | [solution][s1.3] | [tests][t1.3] |   ✅   |
| 1.4 | [Palindrome Permutation][p1.4] | [solution][s1.4] | [tests][t1.4] |   ✅   |
| 1.5 | [One Away][p1.5]               | [solution][s1.5] | [tests][t1.5] |   ✅   |
| 1.6 | [String Compression][p1.6]     | [solution][s1.6] | [tests][t1.6] |   ✅   |
| 1.7 | [Rotate Matrix][p1.7]          | [solution][s1.7] | [tests][t1.7] |   ✅   |
| 1.8 | [Zero Matrix][p1.8]            | [solution][s1.8] | [tests][t1.8] |   ✅   |
| 1.9 | [String Rotation][p1.9]        | [solution][s1.9] | [tests][t1.9] |   ✅   |

<!-- References: Chapter 1-->
[p1.1]: 1_arrays_and_strings/1.1_is_unique/README.md
[s1.1]: 1_arrays_and_strings/1.1_is_unique/main.go
[t1.1]: 1_arrays_and_strings/1.1_is_unique/main_test.go
[p1.2]: 1_arrays_and_strings/1.2_check_permutation/README.md
[s1.2]: 1_arrays_and_strings/1.2_check_permutation/main.go
[t1.2]: 1_arrays_and_strings/1.2_check_permutation/main_test.go
[p1.3]: 1_arrays_and_strings/1.3_urlify/README.md
[s1.3]: 1_arrays_and_strings/1.3_urlify/main.go
[t1.3]: 1_arrays_and_strings/1.3_urlify/main_test.go
[p1.4]: 1_arrays_and_strings/1.4_palindrome_permutation/README.md
[s1.4]: 1_arrays_and_strings/1.4_palindrome_permutation/main.go
[t1.4]: 1_arrays_and_strings/1.4_palindrome_permutation/main_test.go
[p1.5]: 1_arrays_and_strings/1.5_one_away/README.md
[s1.5]: 1_arrays_and_strings/1.5_one_away/main.go
[t1.5]: 1_arrays_and_strings/1.5_one_away/main_test.go
[p1.6]: 1_arrays_and_strings/1.6_string_compression/README.md
[s1.6]: 1_arrays_and_strings/1.6_string_compression/main.go
[t1.6]: 1_arrays_and_strings/1.6_string_compression/main_test.go
[p1.7]: 1_arrays_and_strings/1.7_rotate_matrix/README.md
[s1.7]: 1_arrays_and_strings/1.7_rotate_matrix/main.go
[t1.7]: 1_arrays_and_strings/1.7_rotate_matrix/main_test.go
[p1.8]: 1_arrays_and_strings/1.8_zero_matrix/README.md
[s1.8]: 1_arrays_and_strings/1.8_zero_matrix/main.go
[t1.8]: 1_arrays_and_strings/1.8_zero_matrix/main_test.go
[p1.9]: 1_arrays_and_strings/1.9_string_rotation/README.md
[s1.9]: 1_arrays_and_strings/1.9_string_rotation/main.go
[t1.9]: 1_arrays_and_strings/1.9_string_rotation/main_test.go

#### Chapter 2: Linked Lists (6/8)

|  #  | Problem                    | Solution         | Tests         | Solved |
|-----|----------------------------|:----------------:|:-------------:|:------:|
| 2.1 | [Remove Dups][p2.1]        | [solution][s2.1] | [tests][t2.1] |   ✅   |
| 2.2 | [Return Kth to Last][p2.2] | [solution][s2.2] | [tests][t2.2] |   ✅   |
| 2.3 | [Delete Middle Node][p2.3] | [solution][s2.3] | [tests][t2.3] |   ✅   |
| 2.4 | [Partition][p2.4]          | [solution][s2.4] | [tests][t2.4] |   ✅   |
| 2.5 | [Sum Lists][p2.5]          | [solution][s2.5] | [tests][t2.5] |   ✅   |
| 2.6 | [Palindrome][p2.6]         | [solution][s2.6] | [tests][t2.6] |   ✅   |
<!--
| 2.7 | [Intersection][p2.7]       | [solution][s2.7] | [tests][t2.7] |   ✅   |
| 2.8 | [Loop Detection][p2.8]     | [solution][s2.8] | [tests][t2.8] |   ✅   | -->

<!-- References: Chapter 2 -->
[p2.1]: 2_linked_lists/2.1_remove_dups/README.md
[s2.1]: 2_linked_lists/2.1_remove_dups/main.go
[t2.1]: 2_linked_lists/2.1_remove_dups/main_test.go
[p2.2]: 2_linked_lists/2.2_return_kth_to_last/README.md
[s2.2]: 2_linked_lists/2.2_return_kth_to_last/main.go
[t2.2]: 2_linked_lists/2.2_return_kth_to_last/main_test.go
[p2.3]: 2_linked_lists/2.3_delete_middle_node/README.md
[s2.3]: 2_linked_lists/2.3_delete_middle_node/main.go
[t2.3]: 2_linked_lists/2.3_delete_middle_node/main_test.go
[p2.4]: 2_linked_lists/2.4_partition/README.md
[s2.4]: 2_linked_lists/2.4_partition/main.go
[t2.4]: 2_linked_lists/2.4_partition/main_test.go
[p2.5]: 2_linked_lists/2.5_sum_lists/README.md
[s2.5]: 2_linked_lists/2.5_sum_lists/main.go
[t2.5]: 2_linked_lists/2.5_sum_lists/main_test.go
[p2.6]: 2_linked_lists/2.6_palindrome/README.md
[s2.6]: 2_linked_lists/2.6_palindrome/main.go
[t2.6]: 2_linked_lists/2.6_palindrome/main_test.go
<!--
[p2.7]: 2_linked_lists/2.7_intersection/README.md
[s2.7]: 2_linked_lists/2.7_intersection/main.go
[t2.7]: 2_linked_lists/2.7_intersection/main_test.go
[p2.8]: 2_linked_lists/2.8_loop_detection/README.md
[s2.8]: 2_linked_lists/2.8_loop_detection/main.go
[t2.8]: 2_linked_lists/2.8_loop_detection/main_test.go
-->
