package bytecode

// This package has all of the bytecode we require for compiling the language.
// We will use Stack Compiling, Not using physical register as we don't have access to them

// Instructions are slice of Byte.
type Instructions []byte

// Opcode is OneByte instruction
type Opcode byte
