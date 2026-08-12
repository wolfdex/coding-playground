# leetcode 20. Valid Parentheses

- **Problem Link:** [LeetCode 20 - Valid Parentheses](https://leetcode.com/problems/valid-parentheses/)
- **Difficulty:** Easy
- **Topics:** String, Stack

---

## 📌 Problem Overview

Given a string `s` containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is valid.

An input string is valid if:
1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.

---

## 💡 Evolution & Solution Variants

In this directory, I document the step-by-step optimization of solving this problem using a **Stack** data structure in Go.

### Go Variant 1: Idiomatic & Readable (`0020_valid_parentheses_variant_1_rune_map.go`)
- **Concept:** Iterates over the string using `range s` (working with `rune`), using a `map[rune]rune` to lookup matching brackets and a slice as a standard LIFO stack.
- **Pros:** High readability, idiomatic Go code, native support for Unicode character ranges.
- **Trade-off:**
  - `rune` (int32) uses 4 bytes per character instead of 1 byte.
  - Dynamically growing slice causes multiple re-allocations as elements are appended.

---

### Go Variant 2: Memory & Performance Optimized (`0020_valid_parentheses_variant_2_byte_prealloc.go`)
- **Concept:** Optimizes space and memory allocations based on hardware-level mechanics:
  1. **Pre-Allocation:** Uses `make([]byte, 0, len(s))` to allocate the slice's capacity upfront. The underlying array is created exactly once on the stack/heap, eliminating dynamic re-allocations during runtime.
  2. **Type Downscaling (`rune` → `byte`):** Since parentheses belong to the ASCII subset, switching from `rune` (4 bytes) to `byte` (1 byte) reduces the stack's memory footprint by 75%.
- **Pros:** Significant decrease in GC pressure, lower memory consumption, and optimal time complexity ($O(n)$).

---

## 📊 Summary Comparison by leetcode

| Metric | Variant 1 (`rune`) | Variant 2 (`byte` + Pre-alloc) |
| :--- | :--- | :--- |
| **Data Type** | `rune` (4 Bytes) | `byte` (1 Byte) |
| **Slice Allocations** | Dynamic (`append` triggers growth) | Zero dynamic re-allocations (`cap = len(s)`) |
| **Memory Footprint** | Higher | **Minimal (Optimized)** |
| **Time Complexity** | $O(n)$ | $O(n)$ |
| **Space Complexity**| $O(n)$ | $O(n)$ |

---

## 📝 Key Takeaway for Go

> **Performance Rule:** When working with ASCII strings, prefer `byte` over `rune`. Always pre-allocate slices with `make([]T, 0, capacity)` when the maximum upper bound of elements is known beforehand to avoid costly slice growth re-allocations.
