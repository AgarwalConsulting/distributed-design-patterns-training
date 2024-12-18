# Example

Formula for Minimum Quorum Size: (n/2 + 1)

A, B, C

## Prepare Phase (Bell)

- A -> B [Prepare: 1]
- A -> C [Prepare: 1]

- B -> A [Prom: 1]
- C -> A [Prom: 1]

---

## Propose Phase (Bell)

- A -> B [Prop: 1; Bell]
- B -> A [Accept: 1]

---

## Prepare Phase (Cat)

- B -> A [Prepare: 2]
- B -> C [Prepare: 2]

- B -> A [Prom: 2]
- C -> A [Prom: 2, Clock: 0]

---

## Propose Phase

- B -> C [Prop: 2; Bell/Cat]

<!--
-  -> A [Accept: 1] -->

---

## Propose Phase (Bell)

- A -> C [Prop: 1; Bell] (Delay)
- C -> A [Accept: 1] (Delay)
