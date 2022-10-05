// Code generated by 'ccgo uuid/uuid/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -o uuid/uuid/uuid_darwin_amd64.go -pkgname uuid', DO NOT EDIT.

package uuid

import (
	"math"
	"reflect"
	"sync/atomic"
	"unsafe"
)

var _ = math.Pi
var _ reflect.Kind
var _ atomic.Value
var _ unsafe.Pointer

const (
	X_BSD_I386__TYPES_H_                   = 0
	X_BSD_MACHINE__TYPES_H_                = 0
	X_CDEFS_H_                             = 0
	X_DARWIN_FEATURE_64_BIT_INODE          = 1
	X_DARWIN_FEATURE_ONLY_UNIX_CONFORMANCE = 1
	X_DARWIN_FEATURE_UNIX_CONFORMANCE      = 3
	X_FILE_OFFSET_BITS                     = 64
	X_LP64                                 = 1
	X_Nonnull                              = 0
	X_Null_unspecified                     = 0
	X_Nullable                             = 0
	X_SYS__PTHREAD_TYPES_H_                = 0
	X_SYS__TYPES_H_                        = 0
	X_UUID_STRING_T                        = 0
	X_UUID_T                               = 0
	X_UUID_UUID_H                          = 0
)

type Ptrdiff_t = int64 /* <builtin>:3:26 */

type Size_t = uint64 /* <builtin>:9:23 */

type Wchar_t = int32 /* <builtin>:15:24 */

type X__int128_t = struct {
	Flo int64
	Fhi int64
} /* <builtin>:21:43 */ // must match modernc.org/mathutil.Int128
type X__uint128_t = struct {
	Flo uint64
	Fhi uint64
} /* <builtin>:22:44 */ // must match modernc.org/mathutil.Int128

type X__builtin_va_list = uintptr /* <builtin>:46:14 */
type X__float128 = float64        /* <builtin>:47:21 */

// Public include file for the UUID library
//
// Copyright (C) 1996, 1997, 1998 Theodore Ts'o.
//
// %Begin-Header%
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, and the entire permission notice in its entirety,
//    including the disclaimer of warranties.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
// 3. The name of the author may not be used to endorse or promote
//    products derived from this software without specific prior
//    written permission.
//
// THIS SOFTWARE IS PROVIDED ``AS IS'' AND ANY EXPRESS OR IMPLIED
// WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
// OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE, ALL OF
// WHICH ARE HEREBY DISCLAIMED.  IN NO EVENT SHALL THE AUTHOR BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT
// OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR
// BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
// LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE
// USE OF THIS SOFTWARE, EVEN IF NOT ADVISED OF THE POSSIBILITY OF SUCH
// DAMAGE.
// %End-Header%

// Copyright (c) 2003-2007 Apple Inc. All rights reserved.
//
// @APPLE_OSREFERENCE_LICENSE_HEADER_START@
//
// This file contains Original Code and/or Modifications of Original Code
// as defined in and that are subject to the Apple Public Source License
// Version 2.0 (the 'License'). You may not use this file except in
// compliance with the License. The rights granted to you under the License
// may not be used to create, or enable the creation or redistribution of,
// unlawful or unlicensed copies of an Apple operating system, or to
// circumvent, violate, or enable the circumvention or violation of, any
// terms of an Apple operating system software license agreement.
//
// Please obtain a copy of the License at
// http://www.opensource.apple.com/apsl/ and read it before using this file.
//
// The Original Code and all software distributed under the License are
// distributed on an 'AS IS' basis, WITHOUT WARRANTY OF ANY KIND, EITHER
// EXPRESS OR IMPLIED, AND APPLE HEREBY DISCLAIMS ALL SUCH WARRANTIES,
// INCLUDING WITHOUT LIMITATION, ANY WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE, QUIET ENJOYMENT OR NON-INFRINGEMENT.
// Please see the License for the specific language governing rights and
// limitations under the License.
//
// @APPLE_OSREFERENCE_LICENSE_HEADER_END@

// Copyright (c) 2000-2018 Apple Inc. All rights reserved.
//
// @APPLE_OSREFERENCE_LICENSE_HEADER_START@
//
// This file contains Original Code and/or Modifications of Original Code
// as defined in and that are subject to the Apple Public Source License
// Version 2.0 (the 'License'). You may not use this file except in
// compliance with the License. The rights granted to you under the License
// may not be used to create, or enable the creation or redistribution of,
// unlawful or unlicensed copies of an Apple operating system, or to
// circumvent, violate, or enable the circumvention or violation of, any
// terms of an Apple operating system software license agreement.
//
// Please obtain a copy of the License at
// http://www.opensource.apple.com/apsl/ and read it before using this file.
//
// The Original Code and all software distributed under the License are
// distributed on an 'AS IS' basis, WITHOUT WARRANTY OF ANY KIND, EITHER
// EXPRESS OR IMPLIED, AND APPLE HEREBY DISCLAIMS ALL SUCH WARRANTIES,
// INCLUDING WITHOUT LIMITATION, ANY WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE, QUIET ENJOYMENT OR NON-INFRINGEMENT.
// Please see the License for the specific language governing rights and
// limitations under the License.
//
// @APPLE_OSREFERENCE_LICENSE_HEADER_END@
// Copyright 1995 NeXT Computer, Inc. All rights reserved.
// Copyright (c) 1991, 1993
//	The Regents of the University of California.  All rights reserved.
//
// This code is derived from software contributed to Berkeley by
// Berkeley Software Design, Inc.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
// 3. All advertising materials mentioning features or use of this software
//    must display the following acknowledgement:
//	This product includes software developed by the University of
//	California, Berkeley and its contributors.
// 4. Neither the name of the University nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
// OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
// HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
// LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
// OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
// SUCH DAMAGE.
//
//	@(#)cdefs.h	8.8 (Berkeley) 1/9/95

// This SDK is designed to work with clang and specific versions of
// gcc >= 4.0 with Apple's patch sets

// Compatibility with compilers and environments that don't support compiler
// feature checking function-like macros.

// The __CONCAT macro is used to concatenate parts of symbol names, e.g.
// with "#define OLD(foo) __CONCAT(old,foo)", OLD(foo) produces oldfoo.
// The __CONCAT macro is a bit tricky -- make sure you don't put spaces
// in between its arguments.  __CONCAT can also concatenate double-quoted
// strings produced by the __STRING macro, but this only works with ANSI C.

// __unused denotes variables and functions that may not be used, preventing
// the compiler from warning about it if not used.

// __used forces variables and functions to be included even if it appears
// to the compiler that they are not used (and would thust be discarded).

// __cold marks code used for debugging or that is rarely taken
// and tells the compiler to optimize for size and outline code.

// __deprecated causes the compiler to produce a warning when encountering
// code using the deprecated functionality.
// __deprecated_msg() does the same, and compilers that support it will print
// a message along with the deprecation warning.
// This may require turning on such warning with the -Wdeprecated flag.
// __deprecated_enum_msg() should be used on enums, and compilers that support
// it will print the deprecation warning.
// __kpi_deprecated() specifically indicates deprecation of kernel programming
// interfaces in Kernel.framework used by KEXTs.

// __unavailable causes the compiler to error out when encountering
// code using the tagged function of variable.

// Delete pseudo-keywords wherever they are not available or needed.

// We use `__restrict' as a way to define the `restrict' type qualifier
// without disturbing older software that is unaware of C99 keywords.

// Compatibility with compilers and environments that don't support the
// nullability feature.

// __disable_tail_calls causes the compiler to not perform tail call
// optimization inside the marked function.

// __not_tail_called causes the compiler to prevent tail call optimization
// on statically bound calls to the function.  It has no effect on indirect
// calls.  Virtual functions, objective-c methods, and functions marked as
// "always_inline" cannot be marked as __not_tail_called.

// __result_use_check warns callers of a function that not using the function
// return value is a bug, i.e. dismissing malloc() return value results in a
// memory leak.

// __swift_unavailable causes the compiler to mark a symbol as specifically
// unavailable in Swift, regardless of any other availability in C.

// __abortlike is the attribute to put on functions like abort() that are
// typically used to mark assertions. These optimize the codegen
// for outlining while still maintaining debugability.

// Declaring inline functions within headers is error-prone due to differences
// across various versions of the C language and extensions.  __header_inline
// can be used to declare inline functions within system headers.  In cases
// where you want to force inlining instead of letting the compiler make
// the decision, you can use __header_always_inline.
//
// Be aware that using inline for functions which compilers may also provide
// builtins can behave differently under various compilers.  If you intend to
// provide an inline version of such a function, you may want to use a macro
// instead.
//
// The check for !__GNUC__ || __clang__ is because gcc doesn't correctly
// support c99 inline in some cases:
// http://gcc.gnu.org/bugzilla/show_bug.cgi?id=55965

// Compiler-dependent macros that bracket portions of code where the
// "-Wunreachable-code" warning should be ignored. Please use sparingly.

// Compiler-dependent macros to declare that functions take printf-like
// or scanf-like arguments.  They are null except for versions of gcc
// that are known to support the features properly.  Functions declared
// with these attributes will cause compilation warnings if there is a
// mismatch between the format string and subsequent function parameter
// types.

// Source compatibility only, ID string not emitted in object file

// __alloc_size can be used to label function arguments that represent the
// size of memory that the function allocates and returns. The one-argument
// form labels a single argument that gives the allocation size (where the
// arguments are numbered from 1):
//
// void	*malloc(size_t __size) __alloc_size(1);
//
// The two-argument form handles the case where the size is calculated as the
// product of two arguments:
//
// void	*calloc(size_t __count, size_t __size) __alloc_size(1,2);

// COMPILATION ENVIRONMENTS -- see compat(5) for additional detail
//
// DEFAULT	By default newly complied code will get POSIX APIs plus
//		Apple API extensions in scope.
//
//		Most users will use this compilation environment to avoid
//		behavioral differences between 32 and 64 bit code.
//
// LEGACY	Defining _NONSTD_SOURCE will get pre-POSIX APIs plus Apple
//		API extensions in scope.
//
//		This is generally equivalent to the Tiger release compilation
//		environment, except that it cannot be applied to 64 bit code;
//		its use is discouraged.
//
//		We expect this environment to be deprecated in the future.
//
// STRICT	Defining _POSIX_C_SOURCE or _XOPEN_SOURCE restricts the
//		available APIs to exactly the set of APIs defined by the
//		corresponding standard, based on the value defined.
//
//		A correct, portable definition for _POSIX_C_SOURCE is 200112L.
//		A correct, portable definition for _XOPEN_SOURCE is 600L.
//
//		Apple API extensions are not visible in this environment,
//		which can cause Apple specific code to fail to compile,
//		or behave incorrectly if prototypes are not in scope or
//		warnings about missing prototypes are not enabled or ignored.
//
// In any compilation environment, for correct symbol resolution to occur,
// function prototypes must be in scope.  It is recommended that all Apple
// tools users add either the "-Wall" or "-Wimplicit-function-declaration"
// compiler flags to their projects to be warned when a function is being
// used without a prototype in scope.

// These settings are particular to each product.
// Platform: MacOSX
// #undef __DARWIN_ONLY_UNIX_CONFORMANCE (automatically set for 64-bit)

// The __DARWIN_ALIAS macros are used to do symbol renaming; they allow
// legacy code to use the old symbol, thus maintaining binary compatibility
// while new code can use a standards compliant version of the same function.
//
// __DARWIN_ALIAS is used by itself if the function signature has not
// changed, it is used along with a #ifdef check for __DARWIN_UNIX03
// if the signature has changed.  Because the __LP64__ environment
// only supports UNIX03 semantics it causes __DARWIN_UNIX03 to be
// defined, but causes __DARWIN_ALIAS to do no symbol mangling.
//
// As a special case, when XCode is used to target a specific version of the
// OS, the manifest constant __ENVIRONMENT_MAC_OS_X_VERSION_MIN_REQUIRED__
// will be defined by the compiler, with the digits representing major version
// time 100 + minor version times 10 (e.g. 10.5 := 1050).  If we are targeting
// pre-10.5, and it is the default compilation environment, revert the
// compilation environment to pre-__DARWIN_UNIX03.

// symbol suffixes used for symbol versioning

// symbol versioning macros

// symbol release macros
// Copyright (c) 2010 Apple Inc. All rights reserved.
//
// @APPLE_OSREFERENCE_LICENSE_HEADER_START@
//
// This file contains Original Code and/or Modifications of Original Code
// as defined in and that are subject to the Apple Public Source License
// Version 2.0 (the 'License'). You may not use this file except in
// compliance with the License. The rights granted to you under the License
// may not be used to create, or enable the creation or redistribution of,
// unlawful or unlicensed copies of an Apple operating system, or to
// circumvent, violate, or enable the circumvention or violation of, any
// terms of an Apple operating system software license agreement.
//
// Please obtain a copy of the License at
// http://www.opensource.apple.com/apsl/ and read it before using this file.
//
// The Original Code and all software distributed under the License are
// distributed on an 'AS IS' basis, WITHOUT WARRANTY OF ANY KIND, EITHER
// EXPRESS OR IMPLIED, AND APPLE HEREBY DISCLAIMS ALL SUCH WARRANTIES,
// INCLUDING WITHOUT LIMITATION, ANY WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE, QUIET ENJOYMENT OR NON-INFRINGEMENT.
// Please see the License for the specific language governing rights and
// limitations under the License.
//
// @APPLE_OSREFERENCE_LICENSE_HEADER_END@

// POSIX.1 requires that the macros we test be defined before any standard
// header file is included.  This permits us to convert values for feature
// testing, as necessary, using only _POSIX_C_SOURCE.
//
// Here's a quick run-down of the versions:
//  defined(_POSIX_SOURCE)		1003.1-1988
//  _POSIX_C_SOURCE == 1L		1003.1-1990
//  _POSIX_C_SOURCE == 2L		1003.2-1992 C Language Binding Option
//  _POSIX_C_SOURCE == 199309L		1003.1b-1993
//  _POSIX_C_SOURCE == 199506L		1003.1c-1995, 1003.1i-1995,
//					and the omnibus ISO/IEC 9945-1: 1996
//  _POSIX_C_SOURCE == 200112L		1003.1-2001
//  _POSIX_C_SOURCE == 200809L		1003.1-2008
//
// In addition, the X/Open Portability Guide, which is now the Single UNIX
// Specification, defines a feature-test macro which indicates the version of
// that specification, and which subsumes _POSIX_C_SOURCE.

// Deal with IEEE Std. 1003.1-1990, in which _POSIX_C_SOURCE == 1L.

// Deal with IEEE Std. 1003.2-1992, in which _POSIX_C_SOURCE == 2L.

// Deal with various X/Open Portability Guides and Single UNIX Spec.

// Deal with all versions of POSIX.  The ordering relative to the tests above is
// important.

// POSIX C deprecation macros
// Copyright (c) 2010 Apple Inc. All rights reserved.
//
// @APPLE_OSREFERENCE_LICENSE_HEADER_START@
//
// This file contains Original Code and/or Modifications of Original Code
// as defined in and that are subject to the Apple Public Source License
// Version 2.0 (the 'License'). You may not use this file except in
// compliance with the License. The rights granted to you under the License
// may not be used to create, or enable the creation or redistribution of,
// unlawful or unlicensed copies of an Apple operating system, or to
// circumvent, violate, or enable the circumvention or violation of, any
// terms of an Apple operating system software license agreement.
//
// Please obtain a copy of the License at
// http://www.opensource.apple.com/apsl/ and read it before using this file.
//
// The Original Code and all software distributed under the License are
// distributed on an 'AS IS' basis, WITHOUT WARRANTY OF ANY KIND, EITHER
// EXPRESS OR IMPLIED, AND APPLE HEREBY DISCLAIMS ALL SUCH WARRANTIES,
// INCLUDING WITHOUT LIMITATION, ANY WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE, QUIET ENJOYMENT OR NON-INFRINGEMENT.
// Please see the License for the specific language governing rights and
// limitations under the License.
//
// @APPLE_OSREFERENCE_LICENSE_HEADER_END@

// Set a single macro which will always be defined and can be used to determine
// the appropriate namespace.  For POSIX, these values will correspond to
// _POSIX_C_SOURCE value.  Currently there are two additional levels corresponding
// to ANSI (_ANSI_SOURCE) and Darwin extensions (_DARWIN_C_SOURCE)

// If the developer has neither requested a strict language mode nor a version
// of POSIX, turn on functionality provided by __STDC_WANT_LIB_EXT1__ as part
// of __DARWIN_C_FULL.

// long long is not supported in c89 (__STRICT_ANSI__), but g++ -ansi and
// c99 still want long longs.  While not perfect, we allow long longs for
// g++.

// ****************************************
//
//  Public darwin-specific feature macros
//

// _DARWIN_FEATURE_64_BIT_INODE indicates that the ino_t type is 64-bit, and
// structures modified for 64-bit inodes (like struct stat) will be used.

// _DARWIN_FEATURE_64_ONLY_BIT_INODE indicates that the ino_t type may only
// be 64-bit; there is no support for 32-bit ino_t when this macro is defined
// (and non-zero).  There is no struct stat64 either, as the regular
// struct stat will already be the 64-bit version.

// _DARWIN_FEATURE_ONLY_VERS_1050 indicates that only those APIs updated
// in 10.5 exists; no pre-10.5 variants are available.

// _DARWIN_FEATURE_ONLY_UNIX_CONFORMANCE indicates only UNIX conforming API
// are available (the legacy BSD APIs are not available)

// _DARWIN_FEATURE_UNIX_CONFORMANCE indicates whether UNIX conformance is on,
// and specifies the conformance level (3 is SUSv3)

// This macro casts away the qualifier from the variable
//
// Note: use at your own risk, removing qualifiers can result in
// catastrophic run-time failures.

// __XNU_PRIVATE_EXTERN is a linkage decoration indicating that a symbol can be
// used from other compilation units, but not other libraries or executables.

// Architecture validation for current SDK

// Similar to OS_ENUM/OS_CLOSED_ENUM/OS_OPTIONS/OS_CLOSED_OPTIONS
//
// This provides more advanced type checking on compilers supporting
// the proper extensions, even in C.

// Copyright (c) 2003-2007 Apple Inc. All rights reserved.
//
// @APPLE_OSREFERENCE_LICENSE_HEADER_START@
//
// This file contains Original Code and/or Modifications of Original Code
// as defined in and that are subject to the Apple Public Source License
// Version 2.0 (the 'License'). You may not use this file except in
// compliance with the License. The rights granted to you under the License
// may not be used to create, or enable the creation or redistribution of,
// unlawful or unlicensed copies of an Apple operating system, or to
// circumvent, violate, or enable the circumvention or violation of, any
// terms of an Apple operating system software license agreement.
//
// Please obtain a copy of the License at
// http://www.opensource.apple.com/apsl/ and read it before using this file.
//
// The Original Code and all software distributed under the License are
// distributed on an 'AS IS' basis, WITHOUT WARRANTY OF ANY KIND, EITHER
// EXPRESS OR IMPLIED, AND APPLE HEREBY DISCLAIMS ALL SUCH WARRANTIES,
// INCLUDING WITHOUT LIMITATION, ANY WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE, QUIET ENJOYMENT OR NON-INFRINGEMENT.
// Please see the License for the specific language governing rights and
// limitations under the License.
//
// @APPLE_OSREFERENCE_LICENSE_HEADER_END@

// Copyright (c) 2000-2003 Apple Computer, Inc. All rights reserved.
//
// @APPLE_OSREFERENCE_LICENSE_HEADER_START@
//
// This file contains Original Code and/or Modifications of Original Code
// as defined in and that are subject to the Apple Public Source License
// Version 2.0 (the 'License'). You may not use this file except in
// compliance with the License. The rights granted to you under the License
// may not be used to create, or enable the creation or redistribution of,
// unlawful or unlicensed copies of an Apple operating system, or to
// circumvent, violate, or enable the circumvention or violation of, any
// terms of an Apple operating system software license agreement.
//
// Please obtain a copy of the License at
// http://www.opensource.apple.com/apsl/ and read it before using this file.
//
// The Original Code and all software distributed under the License are
// distributed on an 'AS IS' basis, WITHOUT WARRANTY OF ANY KIND, EITHER
// EXPRESS OR IMPLIED, AND APPLE HEREBY DISCLAIMS ALL SUCH WARRANTIES,
// INCLUDING WITHOUT LIMITATION, ANY WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE, QUIET ENJOYMENT OR NON-INFRINGEMENT.
// Please see the License for the specific language governing rights and
// limitations under the License.
//
// @APPLE_OSREFERENCE_LICENSE_HEADER_END@

// This header file contains integer types.  It's intended to also contain
// flotaing point and other arithmetic types, as needed, later.

type X__int8_t = int8     /* _types.h:37:33 */
type X__uint8_t = uint8   /* _types.h:41:33 */
type X__int16_t = int16   /* _types.h:42:33 */
type X__uint16_t = uint16 /* _types.h:43:33 */
type X__int32_t = int32   /* _types.h:44:33 */
type X__uint32_t = uint32 /* _types.h:45:33 */
type X__int64_t = int64   /* _types.h:46:33 */
type X__uint64_t = uint64 /* _types.h:47:33 */

type X__darwin_intptr_t = int64   /* _types.h:49:33 */
type X__darwin_natural_t = uint32 /* _types.h:50:33 */

// The rune type below is declared to be an ``int'' instead of the more natural
// ``unsigned long'' or ``long''.  Two things are happening here.  It is not
// unsigned so that EOF (-1) can be naturally assigned to it and used.  Also,
// it looks like 10646 will be a 31 bit standard.  This means that if your
// ints cannot hold 32 bits, you will be in trouble.  The reason an int was
// chosen over a long is that the is*() and to*() routines take ints (says
// ANSI C), but they use __darwin_ct_rune_t instead of int.  By changing it
// here, you lose a bit of ANSI conformance, but your programs will still
// work.
//
// NOTE: rune_t is not covered by ANSI nor other standards, and should not
// be instantiated outside of lib/libc/locale.  Use wchar_t.  wchar_t and
// rune_t must be the same type.  Also wint_t must be no narrower than
// wchar_t, and should also be able to hold all members of the largest
// character set plus one extra value (WEOF). wint_t must be at least 16 bits.

type X__darwin_ct_rune_t = int32 /* _types.h:70:33 */ // ct_rune_t

// mbstate_t is an opaque object to keep conversion state, during multibyte
// stream conversions.  The content must not be referenced by user programs.
type X__mbstate_t = struct {
	_           [0]uint64
	F__mbstate8 [128]int8
} /* _types.h:79:3 */

type X__darwin_mbstate_t = X__mbstate_t /* _types.h:81:33 */ // mbstate_t

type X__darwin_ptrdiff_t = int64 /* _types.h:84:33 */ // ptr1 - ptr2

type X__darwin_size_t = uint64 /* _types.h:92:33 */ // sizeof()

type X__darwin_va_list = X__builtin_va_list /* _types.h:98:33 */ // va_list

type X__darwin_wchar_t = int32 /* _types.h:104:33 */ // wchar_t

type X__darwin_rune_t = X__darwin_wchar_t /* _types.h:109:33 */ // rune_t

type X__darwin_wint_t = int32 /* _types.h:112:33 */ // wint_t

type X__darwin_clock_t = uint64        /* _types.h:117:33 */ // clock()
type X__darwin_socklen_t = X__uint32_t /* _types.h:118:33 */ // socklen_t (duh)
type X__darwin_ssize_t = int64         /* _types.h:119:33 */ // byte count or error
type X__darwin_time_t = int64          /* _types.h:120:33 */ // time()

// Type definitions; takes common type definitions that must be used
// in multiple header files due to [XSI], removes them from the system
// space, and puts them in the implementation space.

type X__darwin_blkcnt_t = X__int64_t                    /* _types.h:55:25 */ // total blocks
type X__darwin_blksize_t = X__int32_t                   /* _types.h:56:25 */ // preferred block size
type X__darwin_dev_t = X__int32_t                       /* _types.h:57:25 */ // dev_t
type X__darwin_fsblkcnt_t = uint32                      /* _types.h:58:25 */ // Used by statvfs and fstatvfs
type X__darwin_fsfilcnt_t = uint32                      /* _types.h:59:25 */ // Used by statvfs and fstatvfs
type X__darwin_gid_t = X__uint32_t                      /* _types.h:60:25 */ // [???] process and group IDs
type X__darwin_id_t = X__uint32_t                       /* _types.h:61:25 */ // [XSI] pid_t, uid_t, or gid_t
type X__darwin_ino64_t = X__uint64_t                    /* _types.h:62:25 */ // [???] Used for 64 bit inodes
type X__darwin_ino_t = X__darwin_ino64_t                /* _types.h:64:26 */ // [???] Used for inodes
type X__darwin_mach_port_name_t = X__darwin_natural_t   /* _types.h:68:28 */ // Used by mach
type X__darwin_mach_port_t = X__darwin_mach_port_name_t /* _types.h:69:35 */ // Used by mach
type X__darwin_mode_t = X__uint16_t                     /* _types.h:70:25 */ // [???] Some file attributes
type X__darwin_off_t = X__int64_t                       /* _types.h:71:25 */ // [???] Used for file sizes
type X__darwin_pid_t = X__int32_t                       /* _types.h:72:25 */ // [???] process and group IDs
type X__darwin_sigset_t = X__uint32_t                   /* _types.h:73:25 */ // [???] signal set
type X__darwin_suseconds_t = X__int32_t                 /* _types.h:74:25 */ // [???] microseconds
type X__darwin_uid_t = X__uint32_t                      /* _types.h:75:25 */ // [???] user IDs
type X__darwin_useconds_t = X__uint32_t                 /* _types.h:76:25 */ // [???] microseconds
type X__darwin_uuid_t = [16]uint8                       /* _types.h:77:25 */
type X__darwin_uuid_string_t = [37]int8                 /* _types.h:78:17 */

// Copyright (c) 2003-2013 Apple Inc. All rights reserved.
//
// @APPLE_OSREFERENCE_LICENSE_HEADER_START@
//
// This file contains Original Code and/or Modifications of Original Code
// as defined in and that are subject to the Apple Public Source License
// Version 2.0 (the 'License'). You may not use this file except in
// compliance with the License. The rights granted to you under the License
// may not be used to create, or enable the creation or redistribution of,
// unlawful or unlicensed copies of an Apple operating system, or to
// circumvent, violate, or enable the circumvention or violation of, any
// terms of an Apple operating system software license agreement.
//
// Please obtain a copy of the License at
// http://www.opensource.apple.com/apsl/ and read it before using this file.
//
// The Original Code and all software distributed under the License are
// distributed on an 'AS IS' basis, WITHOUT WARRANTY OF ANY KIND, EITHER
// EXPRESS OR IMPLIED, AND APPLE HEREBY DISCLAIMS ALL SUCH WARRANTIES,
// INCLUDING WITHOUT LIMITATION, ANY WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE, QUIET ENJOYMENT OR NON-INFRINGEMENT.
// Please see the License for the specific language governing rights and
// limitations under the License.
//
// @APPLE_OSREFERENCE_LICENSE_HEADER_END@

// Copyright (c) 2000-2018 Apple Inc. All rights reserved.
//
// @APPLE_OSREFERENCE_LICENSE_HEADER_START@
//
// This file contains Original Code and/or Modifications of Original Code
// as defined in and that are subject to the Apple Public Source License
// Version 2.0 (the 'License'). You may not use this file except in
// compliance with the License. The rights granted to you under the License
// may not be used to create, or enable the creation or redistribution of,
// unlawful or unlicensed copies of an Apple operating system, or to
// circumvent, violate, or enable the circumvention or violation of, any
// terms of an Apple operating system software license agreement.
//
// Please obtain a copy of the License at
// http://www.opensource.apple.com/apsl/ and read it before using this file.
//
// The Original Code and all software distributed under the License are
// distributed on an 'AS IS' basis, WITHOUT WARRANTY OF ANY KIND, EITHER
// EXPRESS OR IMPLIED, AND APPLE HEREBY DISCLAIMS ALL SUCH WARRANTIES,
// INCLUDING WITHOUT LIMITATION, ANY WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE, QUIET ENJOYMENT OR NON-INFRINGEMENT.
// Please see the License for the specific language governing rights and
// limitations under the License.
//
// @APPLE_OSREFERENCE_LICENSE_HEADER_END@
// Copyright 1995 NeXT Computer, Inc. All rights reserved.
// Copyright (c) 1991, 1993
//	The Regents of the University of California.  All rights reserved.
//
// This code is derived from software contributed to Berkeley by
// Berkeley Software Design, Inc.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
// 3. All advertising materials mentioning features or use of this software
//    must display the following acknowledgement:
//	This product includes software developed by the University of
//	California, Berkeley and its contributors.
// 4. Neither the name of the University nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE REGENTS AND CONTRIBUTORS ``AS IS'' AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED.  IN NO EVENT SHALL THE REGENTS OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
// OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
// HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
// LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
// OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
// SUCH DAMAGE.
//
//	@(#)cdefs.h	8.8 (Berkeley) 1/9/95

// pthread opaque structures

type X__darwin_pthread_handler_rec = struct {
	F__routine uintptr
	F__arg     uintptr
	F__next    uintptr
} /* _pthread_types.h:57:1 */

type X_opaque_pthread_attr_t = struct {
	F__sig    int64
	F__opaque [56]int8
} /* _pthread_types.h:63:1 */

type X_opaque_pthread_cond_t = struct {
	F__sig    int64
	F__opaque [40]int8
} /* _pthread_types.h:68:1 */

type X_opaque_pthread_condattr_t = struct {
	F__sig    int64
	F__opaque [8]int8
} /* _pthread_types.h:73:1 */

type X_opaque_pthread_mutex_t = struct {
	F__sig    int64
	F__opaque [56]int8
} /* _pthread_types.h:78:1 */

type X_opaque_pthread_mutexattr_t = struct {
	F__sig    int64
	F__opaque [8]int8
} /* _pthread_types.h:83:1 */

type X_opaque_pthread_once_t = struct {
	F__sig    int64
	F__opaque [8]int8
} /* _pthread_types.h:88:1 */

type X_opaque_pthread_rwlock_t = struct {
	F__sig    int64
	F__opaque [192]int8
} /* _pthread_types.h:93:1 */

type X_opaque_pthread_rwlockattr_t = struct {
	F__sig    int64
	F__opaque [16]int8
} /* _pthread_types.h:98:1 */

type X_opaque_pthread_t = struct {
	F__sig           int64
	F__cleanup_stack uintptr
	F__opaque        [8176]int8
} /* _pthread_types.h:103:1 */

type X__darwin_pthread_attr_t = X_opaque_pthread_attr_t             /* _pthread_types.h:109:39 */
type X__darwin_pthread_cond_t = X_opaque_pthread_cond_t             /* _pthread_types.h:110:39 */
type X__darwin_pthread_condattr_t = X_opaque_pthread_condattr_t     /* _pthread_types.h:111:43 */
type X__darwin_pthread_key_t = uint64                               /* _pthread_types.h:112:23 */
type X__darwin_pthread_mutex_t = X_opaque_pthread_mutex_t           /* _pthread_types.h:113:40 */
type X__darwin_pthread_mutexattr_t = X_opaque_pthread_mutexattr_t   /* _pthread_types.h:114:44 */
type X__darwin_pthread_once_t = X_opaque_pthread_once_t             /* _pthread_types.h:115:39 */
type X__darwin_pthread_rwlock_t = X_opaque_pthread_rwlock_t         /* _pthread_types.h:116:41 */
type X__darwin_pthread_rwlockattr_t = X_opaque_pthread_rwlockattr_t /* _pthread_types.h:117:45 */
type X__darwin_pthread_t = uintptr                                  /* _pthread_types.h:118:34 */

// Copyright (c) 2003-2012 Apple Inc. All rights reserved.
//
// @APPLE_OSREFERENCE_LICENSE_HEADER_START@
//
// This file contains Original Code and/or Modifications of Original Code
// as defined in and that are subject to the Apple Public Source License
// Version 2.0 (the 'License'). You may not use this file except in
// compliance with the License. The rights granted to you under the License
// may not be used to create, or enable the creation or redistribution of,
// unlawful or unlicensed copies of an Apple operating system, or to
// circumvent, violate, or enable the circumvention or violation of, any
// terms of an Apple operating system software license agreement.
//
// Please obtain a copy of the License at
// http://www.opensource.apple.com/apsl/ and read it before using this file.
//
// The Original Code and all software distributed under the License are
// distributed on an 'AS IS' basis, WITHOUT WARRANTY OF ANY KIND, EITHER
// EXPRESS OR IMPLIED, AND APPLE HEREBY DISCLAIMS ALL SUCH WARRANTIES,
// INCLUDING WITHOUT LIMITATION, ANY WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE, QUIET ENJOYMENT OR NON-INFRINGEMENT.
// Please see the License for the specific language governing rights and
// limitations under the License.
//
// @APPLE_OSREFERENCE_LICENSE_HEADER_END@
// Copyright (c) 2003-2007 Apple Inc. All rights reserved.
//
// @APPLE_OSREFERENCE_LICENSE_HEADER_START@
//
// This file contains Original Code and/or Modifications of Original Code
// as defined in and that are subject to the Apple Public Source License
// Version 2.0 (the 'License'). You may not use this file except in
// compliance with the License. The rights granted to you under the License
// may not be used to create, or enable the creation or redistribution of,
// unlawful or unlicensed copies of an Apple operating system, or to
// circumvent, violate, or enable the circumvention or violation of, any
// terms of an Apple operating system software license agreement.
//
// Please obtain a copy of the License at
// http://www.opensource.apple.com/apsl/ and read it before using this file.
//
// The Original Code and all software distributed under the License are
// distributed on an 'AS IS' basis, WITHOUT WARRANTY OF ANY KIND, EITHER
// EXPRESS OR IMPLIED, AND APPLE HEREBY DISCLAIMS ALL SUCH WARRANTIES,
// INCLUDING WITHOUT LIMITATION, ANY WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE, QUIET ENJOYMENT OR NON-INFRINGEMENT.
// Please see the License for the specific language governing rights and
// limitations under the License.
//
// @APPLE_OSREFERENCE_LICENSE_HEADER_END@

type Uuid_t = X__darwin_uuid_t /* _uuid_t.h:31:25 */

type Uuid_string_t = X__darwin_uuid_string_t /* uuid.h:43:33 */

var sUUID_NULL = Uuid_t{uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0), uint8(0)} /* uuid.h:49:1 */

var _ int8 /* gen.c:2:13: */
