/*
 * The purpose of this file is to define macros
 * that tracee.bpf.c relies on which are defined
 * in linux kernel headers but not in vmlinux.h 
 * 
 * vmlinux.h is generated from BTF information
 * in vmlinux but this does not include macros.
 * 
 */ 

#ifndef __TRACEE_MISSING_MACROS_H__
#define __TRACEE_MISSING_MACROS_H__

#define ULLONG_MAX	(~0ULL)

#endif