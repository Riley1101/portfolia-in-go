---
{
  "id": "3799b12d-a396-42c9-8c65-2a801ecb0b93",
  "title": "Creating and traversing tries",
  "description": "Tries, also known as prefix trees, are tree-like data structures used for quick string searching.",
  "slug": "creating-and-traversing-tries",
  "publishedAt": "2023-07-07T17:43:00.000Z"
}
---

Today we gonna talk about Tries and how you can implement tries in this article.

Tries, or prefix trees, are tree-like data structures that are often used for quick string searching and storage. Each node in a trie stands for a character in a given string, and the edges linking the nodes represent the potential characters that can follow a given character. Tries are especially useful for autocomplete and dictionary functions because they allow for rapid searching and retrieval of words based on a given prefix.

![](https://cdn.sanity.io/images/m9whymrq/production/ecb6a784eae4b92dbcead26f7153ea0b96b11be5-4753x2856.png)

While binary search trees are often used for storing strings, they can also be used to store any sortable data. This is because binary search trees use a comparison function to determine the order of the data, allowing for efficient searching and retrieval based on the sorted order.

However, as mentioned before, string comparison can be relatively expensive since each character in the string must be compared in sequence until a difference is found.

For example, let's compare the strings `apple` and `banana`. We start by comparing the first character of each string, `a` and `b`. Since `a` comes before `b` in the ASCII character set, `apple` is considered less than "banana".

Sequential comparison can be relatively expensive since each character in the string must be compared in sequence until a difference is found. For example, to compare the strings `cart` and `care`, we start by comparing the first character of each string, `c`, and `c`. Since they are the same, we move on to the second characters, `a` and `a`, which are also the same. We then compare the third characters, `r`, and `e`, which are different. This tells us that the `cart` is less than `care` in alphabetical order.

## Tries

Tries are a fascinating data structure that can be used to store and search for strings efficiently. In a trie, each node represents a character, and the edges represent the potential characters that can follow. As a result, searching for a string in a trie only requires traversing the trie once, making it faster than binary search trees for string-related operations.

### Creating a trie

Creating a trie involves using an array or map to represent each character in a string. A map is commonly used for simplicity, as it associates each character with its corresponding node in the trie. To insert a new string, you start at the root node and move down the tree, creating new nodes as needed for each character.

For example, let's insert the string `cat` into a trie. We begin at the root node and check if its children contain a node for the character `c`. Since the trie is empty, the map does not contain a node for `c`. We create a new node for "c" and add it as a child of the root node. We then move down to the `c` node and repeat the process for the characters `a` and `t`, creating new nodes as necessary.

Using a map for each node makes it easy to retrieve nodes and determine if a string is present in the trie. You can traverse down the tree, checking if each character in the string has a corresponding node. There are also more advanced variations of the trie, such as compressed tries, which reduce the memory required to store a large number of strings.

Here is what a trie node would look like.

```typescript
export class Node {
  value: string;
  isComplete: boolean;
  children: Map<string, Node>;
  parent: Node | null;
  constructor(value: string, isComplete = false) {
    this.parent = null;
    this.value = value;
    this.children = new Map();
    this.isComplete = isComplete;
  }
}
export class Tries {
  root: Node;
  constructor() {
    this.root = new Node(" ");
  }
  insert(word: string) {
    let current = this.root;
    for (const letter of word) {
      if (!current.children.has(letter)) {
        let newNode = new Node(letter);
        newNode.parent = current;
        current.children.set(letter, newNode);
      }
      current = current.children.get(letter);
    }
    current.isComplete = true;
    return current;
  }
  getValues() {
    console.log(this.root);
    return this.root;
  }
}
```

## Traversing a trie

Traversing a trie involves starting at the root node and moving down the tree, following the edges that correspond to the characters in the string you're searching for. This means that the search process is more efficient compared to other data structures such as hash tables and binary trees. In fact, tries are specifically designed for speedy and effective string searches.

If the string is present in the trie, you will eventually reach a node that has a `isComplete` flag set to `true`, indicating that the string is complete and present in the trie. However, if the string is not present in the trie, the search process will continue until the end of the string is reached, and the same `isComplete` flag will be set to `false`.

Moreover, tries can be used not only for string search operations but also for other applications such as autocomplete and spell-checking. In addition, tries can be implemented in various ways depending on the specific requirements of the application, such as using bitmaps for space efficiency or implementing compressed tries for faster search times.

If you are not familiar with how traversing works check out my other article on how traversing in tree data structures [Traversing on tree data structures](https://www.arkar.space/articles/exploring-tree-traversals-depth-first-search-and-breadth-first-search-algorithms).

```typescript
export function traverseNodes(node: Node) {
  let arr = [];
  if (!node) return [];
  let queue = [node];
  while (queue.length) {
    let current = queue.pop();
    arr.push(current.value);
    current.children.forEach((item) => {
      queue.push(item);
    });
  }
  return arr;
}
```
