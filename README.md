# gomachine

A very basic CPU simulator.

Custom instruction set.

Assembler included.

Built with Go.

## Getting Started

1. Clone this repo.

2. Write assembly.

   **loop.gm**

   ```
   MOV r1 3
   MOV r2 2
   ADD r3 r1 r2
   loop:
   INC r4
   ADD r0 r0 r3
   CMP r4 10
   JLT loop
   ```

3. Compile `gomachine`.

   ```
   $ make
   ```

4. Run it.

   ```
   $ ./gomachine loop.gm
   ```

## Reference

TODO

### TODO:

- debug flag (log machine steps)
- user input
- cli help message
- docs
