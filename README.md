# TruthTableSimulator

A truth table simulator and logical equivalence calculator written in Go for CPS213.

## Features
- Generates complete truth tables for Boolean functions in Sum of Products (SOP) or Product of Sums (POS) form.
- Determines whether two Boolean functions are logically equivalent across the union of their variable domains.
- Supports single-letter variables (e.g., `a`, `b`, `c`), complement notation (e.g., `a'`, `b'`), AND (`.` or adjacent letters), and OR (`+`).

## Usage

Run the interactive CLI:
```bash
make run
```

## Testing

Run unit tests:
```bash
make test
```
