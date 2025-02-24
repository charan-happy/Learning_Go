/*
This problem was asked by Google.

Given the root to a binary tree, implement serialize(root), which serializes the tree into a string, and deserialize(s), which deserializes the string back into the tree.

For example, given the following Node class

class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

The following test should pass:

node = Node('root', Node('left', Node('left.left')), Node('right'))
assert deserialize(serialize(node)).left.left.val == 'left.left'
*/

// Bruteforce

package main

import (
    "fmt"
    "strings"
)

// Brute-force serialize: turn tree into a string
func serializeBrute(root *Node) string {
    if root == nil {
        return "null" // Mark empty nodes
    }
    // Build string: value,left,right
    result := root.Val
    result += "," + serializeBrute(root.Left)
    result += "," + serializeBrute(root.Right)
    return result
}

// Brute-force deserialize: string back to tree
func deserializeBrute(data string) *Node {
    // Split string into array
    vals := strings.Split(data, ",")
    index := 0 // Track our position

    // Helper to build tree recursively
    var buildTree func() *Node
    buildTree = func() *Node {
        if index >= len(vals) || vals[index] == "null" {
            index++ // Move past "null"
            return nil
        }
        // Make a new node with current value
        node := &Node{Val: vals[index]}
        index++ // Move to next value
        node.Left = buildTree()  // Build left subtree
        node.Right = buildTree() // Build right subtree
        return node
    }
    return buildTree()
}

// Time Complexity : O(n * L) ≈ O(n²) << Serialization O(n) << Deserialization 
// Space Complexity : O(n + S) << Serializatioon 0(n) << Deserialization

// optimized solution
package main

import (
    "fmt"
    "strings"
)

// Optimized serialize
func serialize(root *Node) string {
    var parts []string // Collect parts in a slice
    var traverse func(*Node)
    traverse = func(node *Node) {
        if node == nil {
            parts = append(parts, "null")
            return
        }
        parts = append(parts, node.Val) // Add value
        traverse(node.Left)             // Left subtree
        traverse(node.Right)            // Right subtree
    }
    traverse(root)
    return strings.Join(parts, ",") // Join once at the end
}

// Optimized deserialize
func deserialize(data string) *Node {
    vals := strings.Split(data, ",")
    index := 0 // Position tracker

    var buildTree func() *Node
    buildTree = func() *Node {
        if index >= len(vals) || vals[index] == "null" {
            index++
            return nil
        }
        node := &Node{Val: vals[index]}
        index++
        node.Left = buildTree()
        node.Right = buildTree()
        return node
    }
    return buildTree()
}

// Time complexity : O(n) << serialization 0(n) << Deserialization
// Space complexity : O(n + S) << serialization 0(n) << Deserialization


// Ref : https://x.com/i/grok?conversation=1876532656982372670