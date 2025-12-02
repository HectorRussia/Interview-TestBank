# Sort Simple

A small guide and exercise for sorting lists in Python — words, integers, and decimals.

## Overview

Python provides two common ways to sort lists:

- `list.sort()` — sorts the list in-place and returns `None`.
- `sorted(iterable)` — returns a new sorted list and leaves the original unchanged.

Both use Python's Timsort algorithm (stable, adaptive), with average/worst-case time complexity O(n log n).

## Quick Examples

In-place sorting with `list.sort()`:

```python
elements = [5, 3, 6, 2, 1]
elements.sort()
print(elements)  # [1, 2, 3, 5, 6]

words = ["grape", "apple", "banana", "orange"]
words.sort()
print(words)  # ['apple', 'banana', 'grape', 'orange']
```

Using `sorted()` to keep the original list:

```python
orig = [3, 1, 2]
new = sorted(orig)
print(orig)  # [3, 1, 2]
print(new)   # [1, 2, 3]
```

## Challenge (Functions to Implement)

Implement three small helper functions (typical exercise for beginners):

- `sort_words(words: List[str]) -> List[str]` — return a sorted list of strings (ascending, lexicographic).
- `sort_numbers(numbers: List[int]) -> List[int]` — return a sorted list of integers (ascending).
- `sort_decimals(numbers: List[float]) -> List[float]` — return a sorted list of floats (ascending).

Example implementations:

```python
from typing import List

def sort_words(words: List[str]) -> List[str]:
	return sorted(words)

def sort_numbers(numbers: List[int]) -> List[int]:
	return sorted(numbers)

def sort_decimals(numbers: List[float]) -> List[float]:
	return sorted(numbers)
```

These use `sorted()` so the original lists are not modified.

## Time and Space Complexity

- Time: O(n log n) — Timsort, where `n` is the number of elements.
- Space: O(n) in the worst case when using `sorted()` (returns a new list). `list.sort()` is typically O(1) extra space (in-place) apart from temporary buffers.

## How to run / test

1. Create a small script (e.g. `example.py`) that imports and calls the functions above.
2. Run it with Python:

```bash
py example.py
# or
python example.py
```

Example `example.py`:

```python
from sort_simple import sort_words, sort_numbers, sort_decimals

print(sort_words(['grape', 'apple', 'banana']))
print(sort_numbers([5, 2, 9, 1]))
print(sort_decimals([3.14, 2.71, 0.5]))
```

## Notes & Tips

- For case-insensitive string sorting, use: `sorted(words, key=str.lower)`.
- To sort descending, pass `reverse=True`.
- Use the `key` parameter for custom comparisons (e.g., by length: `key=len`).

## Contributing

If you'd like to add tests or alternate implementations (e.g., demonstrate `list.sort()` in-place vs `sorted()`), open a PR or add a small `tests/` script.

---

Happy sorting! 🚀
