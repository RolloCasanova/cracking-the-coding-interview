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

#### Chapter 2: Linked Lists (8/8)

|  #  | Problem                    | Solution         | Tests         | Solved |
|-----|----------------------------|:----------------:|:-------------:|:------:|
| 2.1 | [Remove Dups][p2.1]        | [solution][s2.1] | [tests][t2.1] |   ✅   |
| 2.2 | [Return Kth to Last][p2.2] | [solution][s2.2] | [tests][t2.2] |   ✅   |
| 2.3 | [Delete Middle Node][p2.3] | [solution][s2.3] | [tests][t2.3] |   ✅   |
| 2.4 | [Partition][p2.4]          | [solution][s2.4] | [tests][t2.4] |   ✅   |
| 2.5 | [Sum Lists][p2.5]          | [solution][s2.5] | [tests][t2.5] |   ✅   |
| 2.6 | [Palindrome][p2.6]         | [solution][s2.6] | [tests][t2.6] |   ✅   |
| 2.7 | [Intersection][p2.7]       | [solution][s2.7] | [tests][t2.7] |   ✅   |
| 2.8 | [Loop Detection][p2.8]     | [solution][s2.8] | [tests][t2.8] |   ✅   |

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
[p2.7]: 2_linked_lists/2.7_intersection/README.md
[s2.7]: 2_linked_lists/2.7_intersection/main.go
[t2.7]: 2_linked_lists/2.7_intersection/main_test.go
[p2.8]: 2_linked_lists/2.8_loop_detection/README.md
[s2.8]: 2_linked_lists/2.8_loop_detection/main.go
[t2.8]: 2_linked_lists/2.8_loop_detection/main_test.go

#### Chapter 3: Stacks and Queues (6/6)

|  #  | Problem                  | So3lution        | Tests         | Solved |
|-----|--------------------------|:----------------:|:-------------:|:------:|
| 3.1 | [Three in One][p3.1]     | [solution][s3.1] | [tests][t3.1] |   ✅   |
| 3.2 | [Stack Min][p3.2]        | [solution][s3.2] | [tests][t3.2] |   ✅   |
| 3.3 | [Stack of Plates][p3.3]  | [solution][s3.3] | [tests][t3.3] |   ✅   |
| 3.4 | [Queue via Stacks][p3.4] | [solution][s3.4] | [tests][t3.4] |   ✅   |
| 3.5 | [Sort Stack][p3.5]       | [solution][s3.5] | [tests][t3.5] |   ✅   |
| 3.6 | [Animal Shelter][p3.6]   | [solution][s3.6] | [tests][t3.6] |   ✅   |

<!-- References: Chapter 3 -->
[p3.1]: 3_stacks_and_queues/3.1_three_in_one/README.md
[s3.1]: 3_stacks_and_queues/3.1_three_in_one/main.go
[t3.1]: 3_stacks_and_queues/3.1_three_in_one/main_test.go
[p3.2]: 3_stacks_and_queues/3.2_stack_min/README.md
[s3.2]: 3_stacks_and_queues/3.2_stack_min/main.go
[t3.2]: 3_stacks_and_queues/3.2_stack_min/main_test.go
[p3.3]: 3_stacks_and_queues/3.3_stack_of_plates/README.md
[s3.3]: 3_stacks_and_queues/3.3_stack_of_plates/main.go
[t3.3]: 3_stacks_and_queues/3.3_stack_of_plates/main_test.go
[p3.4]: 3_stacks_and_queues/3.4_queue_via_stacks/README.md
[s3.4]: 3_stacks_and_queues/3.4_queue_via_stacks/main.go
[t3.4]: 3_stacks_and_queues/3.4_queue_via_stacks/main_test.go
[p3.5]: 3_stacks_and_queues/3.5_sort_stack/README.md
[s3.5]: 3_stacks_and_queues/3.5_sort_stack/main.go
[t3.5]: 3_stacks_and_queues/3.5_sort_stack/main_test.go
[p3.6]: 3_stacks_and_queues/3.6_animal_shelter/README.md
[s3.6]: 3_stacks_and_queues/3.6_animal_shelter/main.go
[t3.6]: 3_stacks_and_queues/3.6_animal_shelter/main_test.go

<!--
#### Chapter 4: Trees and Graphs (0/12)

|  #  | Problem                       | Solution          | Tests          | Solved |
|-----|-------------------------------|:-----------------:|:--------------:|:------:|
| 4.1 | [Route Between Nodes][p4.1]   | [solution][s4.1]  | [tests][t4.1]  |   ✅   |
| 4.2 | [Minimal Tree][p4.2]          | [solution][s4.2]  | [tests][t4.2]  |   ✅   |
| 4.3 | [List of Depths][p4.3]        | [solution][s4.3]  | [tests][t4.3]  |   ✅   |
| 4.4 | [Check Balanced][p4.4]        | [solution][s4.4]  | [tests][t4.4]  |   ✅   |
| 4.5 | [Validate BST][p4.5]          | [solution][s4.5]  | [tests][t4.5]  |   ✅   |
| 4.6 | [Successor][p4.6]             | [solution][s4.6]  | [tests][t4.6]  |   ✅   |
| 4.7 | [Build Order][p4.7]           | [solution][s4.7]  | [tests][t4.7]  |   ✅   |
| 4.8 | [First Common Ancestor][p4.8] | [solution][s4.8]  | [tests][t4.8]  |   ✅   |
| 4.9 | [BST Sequences][p4.9]         | [solution][s4.9]  | [tests][t4.9]  |   ✅   |
| 4.10 | [Check Subtree][p4.10]       | [solution][s4.10] | [tests][t4.10] |   ✅   |
| 4.11 | [Random Node][p4.11]         | [solution][s4.11] | [tests][t4.11] |   ✅   |
| 4.12 | [Paths with Sum][p4.12]      | [solution][s4.12] | [tests][t4.12] |   ✅   |
-->

<!-- References: Chapter 4 -->
<!--
[p4.1]: 4_trees_and_graphs/4.1_route_between_nodes/README.md
[s4.1]: 4_trees_and_graphs/4.1_route_between_nodes/main.go
[t4.1]: 4_trees_and_graphs/4.1_route_between_nodes/main_test.go
[p4.2]: 4_trees_and_graphs/4.2_minimal_tree/README.md
[s4.2]: 4_trees_and_graphs/4.2_minimal_tree/main.go
[t4.2]: 4_trees_and_graphs/4.2_minimal_tree/main_test.go
[p4.3]: 4_trees_and_graphs/4.3_list_of_depths/README.md
[s4.3]: 4_trees_and_graphs/4.3_list_of_depths/main.go
[t4.3]: 4_trees_and_graphs/4.3_list_of_depths/main_test.go
[p4.4]: 4_trees_and_graphs/4.4_check_balanced/README.md
[s4.4]: 4_trees_and_graphs/4.4_check_balanced/main.go
[t4.4]: 4_trees_and_graphs/4.4_check_balanced/main_test.go
[p4.5]: 4_trees_and_graphs/4.5_validate_bst/README.md
[s4.5]: 4_trees_and_graphs/4.5_validate_bst/main.go
[t4.5]: 4_trees_and_graphs/4.5_validate_bst/main_test.go
[p4.6]: 4_trees_and_graphs/4.6_successor/README.md
[s4.6]: 4_trees_and_graphs/4.6_successor/main.go
[t4.6]: 4_trees_and_graphs/4.6_successor/main_test.go
[p4.7]: 4_trees_and_graphs/4.7_build_order/README.md
[s4.7]: 4_trees_and_graphs/4.7_build_order/main.go
[t4.7]: 4_trees_and_graphs/4.7_build_order/main_test.go
[p4.8]: 4_trees_and_graphs/4.8_first_common_ancestor/README.md
[s4.8]: 4_trees_and_graphs/4.8_first_common_ancestor/main.go
[t4.8]: 4_trees_and_graphs/4.8_first_common_ancestor/main_test.go
[p4.9]: 4_trees_and_graphs/4.9_bst_sequences/README.md
[s4.9]: 4_trees_and_graphs/4.9_bst_sequences/main.go
[t4.9]: 4_trees_and_graphs/4.9_bst_sequences/main_test.go
[p4.10]: 4_trees_and_graphs/4.10_check_subtree/README.md
[s4.10]: 4_trees_and_graphs/4.10_check_subtree/main.go
[t4.10]: 4_trees_and_graphs/4.10_check_subtree/main_test.go
[p4.11]: 4_trees_and_graphs/4.11_random_node/README.md
[s4.11]: 4_trees_and_graphs/4.11_random_node/main.go
[t4.11]: 4_trees_and_graphs/4.11_random_node/main_test.go
[p4.12]: 4_trees_and_graphs/4.12_paths_with_sum/README.md
[s4.12]: 4_trees_and_graphs/4.12_paths_with_sum/main.go
[t4.12]: 4_trees_and_graphs/4.12_paths_with_sum/main_test.go
-->
