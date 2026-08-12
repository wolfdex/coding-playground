# 1047. Remove All Adjacent Duplicates In String

- **Problem Link:** [LeetCode 1047 - Remove All Adjacent Duplicates In String](https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/)
- **Difficulty:** Easy
- **Topics:** String, Stack, Two Pointers

---

## 📌 Problem Overview

A given string `s` consisting of lowercase English letters. A duplicate removal consists of choosing two adjacent and equal letters and removing them.

repeatedly make duplicate removals on `s` until no longer can. Return the final string after all such duplicate removals have been made.

**Example:**
- **Input:** `s = "abbaca"`
- **Execution:** `"abbaca"` $\rightarrow$ remove `"bb"` $\rightarrow$ `"aaca"` $\rightarrow$ remove `"aa"` $\rightarrow$ `"ca"`
- **Output:** `"ca"`

---

## 💡 Solutions & Performance Evolution

### Go Variant 1: Pre-Allocated Stack (`1047_variant_1_prealloc_stack.go`)
- **Concept:** Uses a slice as a standard LIFO stack. To avoid costly dynamic heap re-allocations during runtime (slice growing from `0` $\rightarrow$ `8` $\rightarrow$ `16` ...), the stack capacity is pre-allocated up front with `make([]byte, 0, len(s))`.
- **ASCII Optimization:** Uses `byte` (1 byte) instead of `rune` (4 bytes), saving 75% memory footprint since inputs are restricted to lowercase English letters.
- **Time Complexity:** $O(n)$
- **Space Complexity:** $O(n)$ — requires auxiliary memory for the separate stack slice.

---

### Go Variant 2: In-Place Two-Pointer (`1047_variant_2_in_place.go`) 🚀 *(Beats 100% Runtime)*
- **Concept:** Eliminates auxiliary slice allocation by using `writeIndex` as a virtual stack pointer directly on the converted `[]byte(s)` slice:
  - `i` acts as the **read pointer**, scanning through the input.
  - `writeIndex` acts as the **stack pointer** (top of stack is `writeIndex - 1`).
  - **Pop:** Decrement `writeIndex` when a duplicate is found (overwritten in subsequent steps).
  - **Push:** Assign `stack[writeIndex] = stack[i]` and increment `writeIndex`.
- **Pros:** Zero dynamic allocations inside the loop, minimal memory overhead, zero garbage collector pressure.
- **Time Complexity:** $O(n)$
- **Space Complexity:** $O(n)$ total / $O(1)$ auxiliary space.

---

## 📊 Performance Comparison

| Metric | Variant 1 (Pre-Allocated) | Variant 2 (In-Place) |
| :--- | :--- | :--- |
| **Approach** | Separate Stack Slice (`make`) | Virtual Stack Pointer (`writeIndex`) |
| **Auxiliary Slices** | 1 Extra Slice | **0 Extra Slices** |
| **Heap Allocations** | `len(s)` Bytes | **In-Place Reuse** |
| **Runtime** | ~1 ms | **0 ms (Beats 100.00%)** |
| **Memory** | ~8.1 MB | **~8.2 MB** |

---

## 📝 Key Takeaway for Go

> **In-Place Mutation Pattern:** When modifying strings or slices where the output length is bounded by input length ($\text{length}_{\text{out}} \le \text{length}_{\text{in}}$), reuse the backing array in-place with a `writeIndex` pointer. This bypasses secondary stack allocations entirely and yields optimal execution speed.
