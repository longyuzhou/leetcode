// Given a binary tree containing digits from 0-9 only,
// each root-to-leaf path could represent a number.
// An example is the root-to-leaf path 1->2->3 which represents the number 123.
// Find the total sum of all root-to-leaf numbers.
// Note: A leaf is a node with no children.
// Example:
//   Input: [1,2,3]
//       1
//      / \
//     2   3
//   Output: 25
//   Explanation:
//   The root-to-leaf path 1->2 represents the number 12.
//   The root-to-leaf path 1->3 represents the number 13.
//   Therefore, sum = 12 + 13 = 25.
//
// Solution: dfs

var sumNumbers = function (root) {
  let sum = 0;
  const dfs = (node, num) => {
    if (node == null) {
      return;
    }
    num = num * 10 + node.val;
    if (node.left == null && node.right == null) {
      sum += num;
    } else {
      dfs(node.left, num);
      dfs(node.right, num);
    }
  };
  dfs(root, 0);
  return sum;
};

const test = (root, expect) => {
  const actual = sumNumbers(root);
  if (actual != expect) {
    console.log(`fail! expect ${expect}, got ${actual}`);
  } else {
    console.log('pass');
  }
};

test(
  {
    val: 1,
    left: { val: 2 },
    right: { val: 3 },
  },
  25
);

test(
  {
    val: 4,
    left: { val: 9, left: { val: 5 }, right: { val: 1 } },
    right: { val: 0 },
  },
  1026
);
