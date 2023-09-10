# Animal Shelter

An animal shelter, which holds only dogs and cats, operates on a strictly "first in, first out" basis. People must adopt either the "oldest" (based on arrival time) of all animals at the shelter, or they can select whether they would prefer a dog or a cat (and will receive the oldest animal of that type). They cannot select which specific animal they would like. Create the data structures to maintain this system and implement operations such as `enqueue`, `dequeueAny`, `dequeueDog`, and `dequeueCat`. You may use the built-in `LinkedList` data structure.

## Hints

<details>
    <summary>#22</summary>
    Observe that the oldest animal can always be found at the front of the queue.
</details>

<details>
    <summary>#56</summary>
    We could maintain a `LinkedList` for dogs and a `LinkedList` for cats, and pop from the front of the appropriate list.
</details>

<details>
    <summary>#63</summary>
    How would you maintain a `LinkedList` so that `enqueue` and `dequeueAny` work in `O(1)` time?
</details>
