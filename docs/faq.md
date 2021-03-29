# FAQ - Frequently Asked Questions

## Secure tracing

When Tracee-eBPF reads information from user programs it is subject to a race condition where the user program might be able to change the arguments after Tracee has read them. For example, a program invoked `execve("/bin/ls", NULL, 0)`, Tracee picked that up and will report that, then the program changed the first argument from `/bin/ls` to `/bin/bash`, and this is what the kernel will execute. To mitigate this, Tracee also provide "LSM" (Linux Security Module) based events, for example, the `bprm_check` event which can be reported by Tracee and cross-referenced with the reported regular syscall event.