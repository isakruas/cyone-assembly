# Cyone Assembly Language Documentation

## Overview

Cyone is an assembly language specifically designed for interacting with the Cyone Kernel, a system developed for 8-bit microcontrollers. The kernel provides access to preconfigured resources and enables the creation of complex programs that exceed the typical memory constraints of these microcontrollers.

## Table of Contents

1. [Kernel Overview](#kernel-overview)
2. [Program Structure](#program-structure)
3. [Variable Declaration](#variable-declaration)
4. [Value Assignment](#value-assignment)
5. [Reading Values](#reading-values)
6. [Conditional Structures](#conditional-structures)
7. [Goto Command](#goto-command)
8. [Kernel Resource Calls](#kernel-resource-calls)
9. [Direct Memory Manipulation](#direct-memory-manipulation)
10. [Finalization Block](#finalization-block)
11. [Complete Example](#complete-example)
12. [Notes and Considerations](#notes-and-considerations)

## Kernel Overview

The Cyone Kernel facilitates advanced programming by managing resources and extending the capabilities of 8-bit microcontrollers. It provides preconfigured peripherals and system functions to simplify complex tasks.

## Program Structure

A Cyone program is organized into code blocks located at specific memory addresses. Each block contains instructions and declarations.

### Block Declaration

```cyone
block <address> { 
    // Instructions and declarations 
}
```

- `<address>`: Memory address where the block begins.
- Example: `block 0x0100 { ... }`

## Variable Declaration

Variables in Cyone are associated with specific memory addresses.

### Syntax

```cyone
loc <name> at <address>;
```

- `<name>`: The variable's name.
- `<address>`: The memory address for the variable.
- Example: `loc x at 0x0000;`

## Value Assignment

Assign values to variables or directly to memory addresses.

### Syntax

```cyone
<variable> = <value>;           // Assign value to a variable
mem[<address>] = <value>;      // Directly assign value to a memory 
```

- `<variable>`: The name of the variable.
- `<value>`: The value to assign (in hexadecimal).
- `<address>`: The memory address.
- Example: `x = 0x0A;`
- Direct memory assignment: `mem[0x0002] = 0xFF;`

## Reading Values

Retrieve values from variables or memory addresses.

### Syntax

```cyone
<variable> = mem[<address>];   // Read from memory into a variable
```

- `<variable>`: The name of the variable to store the read value.
- `<address>`: The memory address to read from.
- Example: `result = mem[0x0002];`

## Conditional Structures

Execute code based on conditions.

### Syntax

```cyone
if (<condition>) { 
    // Instructions if the condition is true 
} else {
    // Instructions if the condition is false
}
```

- `<condition>`: An expression that evaluates to true or false.
- Example: 
  ```cyone
  if (result > 0x0F) {
      // Code if result is greater than 15
  } else {
      // Code if result is 15 or less
  }
  ```

## Goto Command

Unconditionally transfers execution to a specified block.

### Syntax

```cyone
goto <address>;
```

- `<address>`: Memory address of the target block.
- Example: `goto 0x0200;`

## Kernel Resource Calls

Invoke functions provided by the Cyone Kernel for various system operations.

### Syntax

```cyone
call <function_name> (<arguments>);
```

- `<function_name>`: Name of the kernel-provided function.
- `<arguments>`: Parameters required by the function.
- Example: `call DRAW_RECTANGLE (0x10, 0x10, 0x20, 0x05);`

## Direct Memory Manipulation

Access and modify memory directly.

### Syntax

```cyone
mem[<address>] = <value>;      // Write value to address
<variable> = mem[<address>];   // Read value from address
```

- `<address>`: Memory address to manipulate.
- `<value>`: Value to write.
- Example: `mem[0x0002] = 0xFF;`

## Complete Example

```cyone
// Variable declaration and memory association
loc x at 0x0000;         // Variable x at address 0x0000
loc y at 0x0001;         // Variable y at address 0x0001
loc z at 0x0002;         // Variable z at address 0x0002
loc result at 0x0003;    // Variable result at address 0x0003
loc flag at 0x0004;      // Variable flag at address 0x0004
loc temp at 0x0005;      // Temporary variable for intermediate calculations

// Main code block - start of loop
start at 0x0100;

block 0x0100 {
    // Initialize variables
    x = 0x07;           // Assign 7 (0x07) to variable x
    y = 0x03;           // Assign 3 (0x03) to variable y
    z = 0x00;           // Zero variable z
    temp = 0x00;        // Zero temporary variable temp
    
    // Calculate the sum of x and y
    result = x + y;     // Sum x and y, store in result

    // Check if result is a multiple of 3
    temp = result % 0x03; // Calculate remainder of result divided by 3
    if (temp == 0x00) {
        flag = 0x01;   // If remainder is 0, result is a multiple of 3
    } else {
        flag = 0x00;   // Otherwise, not a multiple of 3
    }
    
    // Test flag and perform conditional jump
    if (flag == 0x01) {
        goto 0x0200;  // Jump to block 0x0200 if flag is 1
    }
    
    goto 0x0300;      // If condition is not met, jump to 0x0300
}

// Block for when result is a multiple of 3
block 0x0200 {
    // Store result in an extra address
    mem[0x0006] = result;  // Store result in 0x0006
    // Perform additional operation
    temp = result * 0x02;  // Double result and store in temp
    
    // Call a drawing function with specific arguments
    call DRAW_RECTANGLE (0x10, 0x10, temp, 0x05);
    
    // Mark end of processing for this iteration
    mem[0x0007] = 0x01;  // Indicate that result has been processed
    goto 0x0500;        // Jump to the end of loop block
}

// Alternative block if result is not a multiple of 3
block 0x0300 {
    // Store result in an extra address
    mem[0x0006] = result;  // Store result in 0x0006
    // Perform different operation
    temp = result - 0x01;  // Subtract 1 from result and store in temp
    
    // Call a drawing function with different arguments
    call DRAW_CIRCLE (0x20, 0x20, temp);
    
    // Mark end of processing for this iteration
    mem[0x0007] = 0x00;  // Indicate that result has been processed
    goto 0x0500;        // Jump to the end of loop block
}

// End of loop block
block 0x0500 {
    // Check if the program should continue looping
    if (mem[0x0007] == 0x01) {
        // If flag was set to 1, program should stop or perform another action
        goto 0x0600;  // Jump to end block
    } else {
        // Otherwise, restart the loop
        goto 0x0100;
    }
}

// End block
block 0x0600 {
    // Set final flag value and end
    flag = mem[0x0007];
    
    // Clear variables
    x = 0x00;
    y = 0x00;
    z = 0x00;
    temp = 0x00;
    
    // Optionally, restart or end the program
    goto 0x0100; // May restart or end the program
}
```

## Notes and Considerations

- Ensure proper memory address allocation to avoid conflicts.
- Pay attention to the alignment of code blocks to ensure correct execution flow.
- Kernel function names and available resources may vary depending on the specific implementation.

This documentation provides a comprehensive overview of the Cyone Assembly Language, including syntax and usage examples. For further details, consult the Cyone Kernel technical reference or related resources.
