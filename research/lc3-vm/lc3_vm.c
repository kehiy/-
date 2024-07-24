#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>

#define MEMORY_MAX (1 << 16)
uint16_t memory[MEMORY_MAX]; /* 65536 locations */

enum {
  R_R0 = 0,
  R_R1,
  R_R2,
  R_R3,
  R_R4,
  R_R5,
  R_R6,
  R_R7,
  R_PC, /* program counter */
  R_COND,
  R_COUNT
};

uint16_t reg[R_COUNT];

enum {
  FL_POS = 1 << 0, /* P */
  FL_ZRO = 1 << 1, /* Z */
  FL_NEG = 1 << 2, /* N */
};

enum {
  OP_BR = 0, /* branch */
  OP_ADD,    /* add  */
  OP_LD,     /* load */
  OP_ST,     /* store */
  OP_JSR,    /* jump register */
  OP_AND,    /* bitwise and */
  OP_LDR,    /* load register */
  OP_STR,    /* store register */
  OP_RTI,    /* unused */
  OP_NOT,    /* bitwise not */
  OP_LDI,    /* load indirect */
  OP_STI,    /* store indirect */
  OP_JMP,    /* jump */
  OP_RES,    /* reserved (unused) */
  OP_LEA,    /* load effective address */
  OP_TRAP    /* execute trap */
};

enum
{
    TRAP_GETC = 0x20,  /* get character from keyboard, not echoed onto the terminal */
    TRAP_OUT = 0x21,   /* output a character */
    TRAP_PUTS = 0x22,  /* output a word string */
    TRAP_IN = 0x23,    /* get character from keyboard, echoed onto the terminal */
    TRAP_PUTSP = 0x24, /* output a byte string */
    TRAP_HALT = 0x25   /* halt the program */
};

uint16_t sign_extend(uint16_t x, int bit_count) {
  /*
      The immediate mode value has only 5 bits, but it needs to be added to a
     16-bit number. To do the addition, those 5 bits need to be extended to 16
     to match the other number. For positive numbers, we can simply fill in 0’s
     for the additional bits. For negative numbers, this causes a problem. For
     example, -1 in 5 bits is 1 1111. If we just extended it with 0’s, this is
     0000 0000 0001 1111 which is equal to 31. Sign extension corrects this
     problem by filling in 0’s for positive numbers and 1’s for negative
     numbers, so that original values are preserved.
  */
  if ((x >> (bit_count - 1)) & 1) {
    x |= (0xFFFF << bit_count);
  }
  return x;
}

void update_flags(uint16_t r) {
  if (reg[r] == 0) {
    reg[R_COND] = FL_ZRO;
  } else if (reg[r] >> 15) /* a 1 in the left-most bit indicates negative */
  {
    reg[R_COND] = FL_NEG;
  } else {
    reg[R_COND] = FL_POS;
  }
}

int main(int argc, const char *argv[]) {
  if (argc < 2) {
    printf("Usage: lc3 [image-file1] ...\n");
    return 2;
  }

  for (int i = 1; i < argc; ++i) {
    if (!read_image(argv[i])) {
      printf("failed to load image %s\n", argv[i]);
      return 1;
    }
  }

  // At the same time we can only set one condition flag, for start we set it to
  // Z (0).
  reg[R_COND] = FL_ZRO;

  enum { PC_START = 0x3000 }; // program counter start position.
  reg[R_PC] = PC_START;

  int running = 1;
  while (running) {
    // (fetch) decode execute
    uint16_t instr = mem_read(reg[R_PC]++);
    uint16_t op = instr >> 12; // TODO: why?

    // fetch (decode) execute
    switch (op) {
    // fetch decode (execute)
    case OP_ADD:

        /* destination register (DR) */
        uint16_t r0 = (instr >> 9) & 0x7; //TODO: why & with 0x7?

        /* first operand (SR1) */
        uint16_t r1 = (instr >> 6) & 0x7; //TODO: why & with 0x7?

        /* whether we are in immediate mode */
        uint16_t imm_flag = (instr >> 5) & 0x1;

        if (imm_flag)
        {
            uint16_t imm5 = sign_extend(instr & 0x1F, 5); //TODO: why & instr with 0x1F?
            reg[r0] = reg[r1] + imm5;
        }
        else
        {
            uint16_t r2 = instr & 0x7; //TODO: why & instr with 0x7?
            reg[r0] = reg[r1] + reg[r2];
        }

        update_flags(r0);

      break;

    case OP_AND:
        {
            uint16_t r0 = (instr >> 9) & 0x7;
            uint16_t r1 = (instr >> 6) & 0x7;
            uint16_t imm_flag = (instr >> 5) & 0x1;

            if (imm_flag)
            {
                uint16_t imm5 = sign_extend(instr & 0x1F, 5);
                reg[r0] = reg[r1] & imm5;
            }
            else
            {
                uint16_t r2 = instr & 0x7;
                reg[0] = reg[r1] & reg[r2];
            }

            update_flags(r0);
        }

      break;
    case OP_NOT:
        {
            uint16_t r0 = (instr >> 9) & 0x7;
            uint16_t r1 = (instr >> 6) & 0x7;
            
            reg[r0] = ~reg[r1];

            update_flags(r0);
        }

      break;
    case OP_BR:

        {
            uint16_t pc_offset = sign_extend(instr & 0x1FF, 9);
            uint16_t cond_flag = (instr >> 9) & 0x7;
            if (cond_flag & reg[R_COND]) //TODO: find out how it works bit by bit.
            {
                reg[R_PC] += pc_offset;
            }
        }

      break;
    case OP_JMP:

        {
            uint16_t r1 = (instr >> 6) & 0x7;
            reg[R_PC] = reg[r1];
        }

      break;
    case OP_JSR:

        {
            uint16_t long_flag = (instr >> 11) & 1;
            reg[R_R7] = reg[R_PC];
            if (long_flag)
            {
                uint16_t long_pc_offset = sign_extend(instr & 0x7FF, 11);
                reg[R_PC] += long_pc_offset;  /* JSR */
            }
            else
            {
                uint16_t r1 = (instr >> 6) & 0x7;
                reg[R_PC] = reg[r1]; /* JSRR */
            }
        }

      break;
    case OP_LD:

        {
            uint16_t r0 = (instr >> 9) & 0x7;
            uint16_t pc_offset = sign_extend(instr & 0x1FF, 9);
            reg[r0] = mem_read(reg[R_PC + pc_offset]);
            update_flags(r0);
        }

      break;

    case OP_LDI:

        {
            uint16_t r0 = (instr >> 9) & 0x7;
            uint16_t pc_offset = sign_extend(instr & 0x1FF, 9);
            reg[r0] = mem_read(mem_read(reg[R_PC] + pc_offset));
            update_flags(r0);
        }

      break;
    case OP_LDR:

        {
            uint16_t r0 = (instr >> 9) & 0x7;
            uint16_t r1 = (instr >> 6) & 0x7;
            uint16_t offset = sign_extend(instr & 0x3F, 6);
            reg[r0] = mem_read(reg[r1] + offset);
            update_flags(r0);
        }

      break;
    case OP_LEA:

        {
            uint16_t r0 = (instr >> 9) & 0x7;
            uint16_t offset = sign_extend(instr & 0x3F, 9);

            reg[r0] = reg[R_PC] + offset;

            update_flags(r0);
        }

      break;
    case OP_ST:

        {
            uint16_t r0 = (instr >> 9) & 0x7;
            uint16_t pc_offset = sign_extend(instr & 0x1FF, 9);
            mem_write(reg[R_PC] + pc_offset, reg[r0]);
        }

      break;
    case OP_STI:

        {
            uint16_t r0 = (instr >> 9) & 0x7;
            uint16_t pc_offset = sign_extend(instr & 0x1FF, 9);
            mem_write(mem_read(reg[R_PC] + pc_offset), reg[r0]);
        }

      break;
    case OP_STR:
        
        {
            uint16_t r0 = (instr >> 9) & 0x7;
            uint16_t r1 = (instr >> 6) & 0x7;
            uint16_t offset = sign_extend(instr & 0x3F, 6);
            mem_write(reg[r1] + offset, reg[r0]);
        }

      break;
    case OP_TRAP:
      break;
    case OP_RES:
        abort();

    case OP_RTI:
        abort();

    default:
      break;
    }
  }
}
