// +build ignore

#include <linux/types.h>
#ifdef asm_inline
#undef asm_inline
#define asm_inline asm
#endif

#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>

char LICENSE[] SEC("license") = "GPL";

//TODO:
SEC("kprobe/sys_execve")
int self(void *ctx) {
    return 0;
}
