// Code generated by 'ccgo stdio/gen.c -crt-import-path "" -export-defines "" -export-enums "" -export-externs X -export-fields F -export-structs "" -export-typedefs "" -hide _OSSwapInt16,_OSSwapInt32,_OSSwapInt64 -o stdio/stdio_linux_386.go -pkgname stdio', DO NOT EDIT.

package stdio

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
	BUFSIZ                = 8192
	EOF                   = -1
	FILENAME_MAX          = 4096
	FOPEN_MAX             = 16
	L_ctermid             = 9
	L_tmpnam              = 20
	P_tmpdir              = "/tmp"
	SEEK_CUR              = 1
	SEEK_END              = 2
	SEEK_SET              = 0
	TMP_MAX               = 238328
	X_ATFILE_SOURCE       = 1
	X_BITS_STDIO_LIM_H    = 1
	X_BITS_TYPESIZES_H    = 1
	X_BITS_TYPES_H        = 1
	X_BSD_SIZE_T_         = 0
	X_BSD_SIZE_T_DEFINED_ = 0
	X_DEFAULT_SOURCE      = 1
	X_FEATURES_H          = 1
	X_FILE_OFFSET_BITS    = 64
	X_GCC_SIZE_T          = 0
	X_IOFBF               = 0
	X_IOLBF               = 1
	X_IONBF               = 2
	X_IO_EOF_SEEN         = 0x0010
	X_IO_ERR_SEEN         = 0x0020
	X_IO_USER_LOCK        = 0x8000
	X_POSIX_C_SOURCE      = 200809
	X_POSIX_SOURCE        = 1
	X_SIZET_              = 0
	X_SIZE_T              = 0
	X_SIZE_T_             = 0
	X_SIZE_T_DECLARED     = 0
	X_SIZE_T_DEFINED      = 0
	X_SIZE_T_DEFINED_     = 0
	X_STDC_PREDEF_H       = 1
	X_STDIO_H             = 1
	X_SYS_CDEFS_H         = 1
	X_SYS_SIZE_T_H        = 0
	X_T_SIZE              = 0
	X_T_SIZE_             = 0
	X_VA_LIST_DEFINED     = 0
	I386                  = 1
	Linux                 = 1
	Unix                  = 1
)

type Ptrdiff_t = int32 /* <builtin>:3:26 */

type Size_t = uint32 /* <builtin>:9:23 */

type Wchar_t = int32 /* <builtin>:15:24 */

type X__builtin_va_list = uintptr /* <builtin>:46:14 */
type X__float128 = float64        /* <builtin>:47:21 */

// Wide character type.
//    Locale-writers should change this as necessary to
//    be big enough to hold unique values not between 0 and 127,
//    and not (wchar_t) -1, for each defined multibyte character.

// Define this type if we are doing the whole job,
//    or if we want this type in particular.

//  In 4.3bsd-net2, leave these undefined to indicate that size_t, etc.
//     are already defined.
//  BSD/OS 3.1 and FreeBSD [23].x require the MACHINE_ANSI_H check here.
//  NetBSD 5 requires the I386_ANSI_H and X86_64_ANSI_H checks here.

// A null pointer constant.

// Copyright (C) 1989-2018 Free Software Foundation, Inc.
//
// This file is part of GCC.
//
// GCC is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 3, or (at your option)
// any later version.
//
// GCC is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// Under Section 7 of GPL version 3, you are granted additional
// permissions described in the GCC Runtime Library Exception, version
// 3.1, as published by the Free Software Foundation.
//
// You should have received a copy of the GNU General Public License and
// a copy of the GCC Runtime Library Exception along with this program;
// see the files COPYING3 and COPYING.RUNTIME respectively.  If not, see
// <http://www.gnu.org/licenses/>.

// ISO C Standard:  7.15  Variable arguments  <stdarg.h>

// Define __gnuc_va_list.

type X__gnuc_va_list = X__builtin_va_list /* stdarg.h:40:27 */

// Define the standard macros for the user,
//    if this invocation was from the user program.

// bits/types.h -- definitions of __*_t types underlying *_t types.
//    Copyright (C) 2002-2018 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <http://www.gnu.org/licenses/>.

// Never include this file directly; use <sys/types.h> instead.

// Copyright (C) 1991-2018 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <http://www.gnu.org/licenses/>.

// Determine the wordsize from the preprocessor defines.

// Convenience types.
type X__u_char = uint8   /* types.h:30:23 */
type X__u_short = uint16 /* types.h:31:28 */
type X__u_int = uint32   /* types.h:32:22 */
type X__u_long = uint32  /* types.h:33:27 */

// Fixed-size types, underlying types depend on word size and compiler.
type X__int8_t = int8     /* types.h:36:21 */
type X__uint8_t = uint8   /* types.h:37:23 */
type X__int16_t = int16   /* types.h:38:26 */
type X__uint16_t = uint16 /* types.h:39:28 */
type X__int32_t = int32   /* types.h:40:20 */
type X__uint32_t = uint32 /* types.h:41:22 */
type X__int64_t = int64   /* types.h:46:44 */
type X__uint64_t = uint64 /* types.h:47:46 */

// Smallest types with at least a given width.
type X__int_least8_t = X__int8_t     /* types.h:51:18 */
type X__uint_least8_t = X__uint8_t   /* types.h:52:19 */
type X__int_least16_t = X__int16_t   /* types.h:53:19 */
type X__uint_least16_t = X__uint16_t /* types.h:54:20 */
type X__int_least32_t = X__int32_t   /* types.h:55:19 */
type X__uint_least32_t = X__uint32_t /* types.h:56:20 */
type X__int_least64_t = X__int64_t   /* types.h:57:19 */
type X__uint_least64_t = X__uint64_t /* types.h:58:20 */

// quad_t is also 64 bits.
type X__quad_t = int64    /* types.h:65:37 */
type X__u_quad_t = uint64 /* types.h:66:46 */

// Largest integral types.
type X__intmax_t = int64   /* types.h:74:37 */
type X__uintmax_t = uint64 /* types.h:75:46 */

// The machine-dependent file <bits/typesizes.h> defines __*_T_TYPE
//    macros for each of the OS types we define below.  The definitions
//    of those macros must use the following macros for underlying types.
//    We define __S<SIZE>_TYPE and __U<SIZE>_TYPE for the signed and unsigned
//    variants of each of the following integer types on this machine.
//
// 	16		-- "natural" 16-bit type (always short)
// 	32		-- "natural" 32-bit type (always int)
// 	64		-- "natural" 64-bit type (long or long long)
// 	LONG32		-- 32-bit type, traditionally long
// 	QUAD		-- 64-bit type, always long long
// 	WORD		-- natural type of __WORDSIZE bits (int or long)
// 	LONGWORD	-- type of __WORDSIZE bits, traditionally long
//
//    We distinguish WORD/LONGWORD, 32/LONG32, and 64/QUAD so that the
//    conventional uses of `long' or `long long' type modifiers match the
//    types we define, even when a less-adorned type would be the same size.
//    This matters for (somewhat) portably writing printf/scanf formats for
//    these types, where using the appropriate l or ll format modifiers can
//    make the typedefs and the formats match up across all GNU platforms.  If
//    we used `long' when it's 64 bits where `long long' is expected, then the
//    compiler would warn about the formats not matching the argument types,
//    and the programmer changing them to shut up the compiler would break the
//    program's portability.
//
//    Here we assume what is presently the case in all the GCC configurations
//    we support: long long is always 64 bits, long is always word/address size,
//    and int is always 32 bits.

// We want __extension__ before typedef's that use nonstandard base types
//    such as `long long' in C89 mode.
// bits/typesizes.h -- underlying types for *_t.  Linux/x86-64 version.
//    Copyright (C) 2012-2018 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <http://www.gnu.org/licenses/>.

// See <bits/types.h> for the meaning of these macros.  This file exists so
//    that <bits/types.h> need not vary across different GNU platforms.

// X32 kernel interface is 64-bit.

// Number of descriptors that can fit in an `fd_set'.

type X__dev_t = X__u_quad_t                /* types.h:143:25 */ // Type of device numbers.
type X__uid_t = uint32                     /* types.h:144:25 */ // Type of user identifications.
type X__gid_t = uint32                     /* types.h:145:25 */ // Type of group identifications.
type X__ino_t = uint32                     /* types.h:146:25 */ // Type of file serial numbers.
type X__ino64_t = X__u_quad_t              /* types.h:147:27 */ // Type of file serial numbers (LFS).
type X__mode_t = uint32                    /* types.h:148:26 */ // Type of file attribute bitmasks.
type X__nlink_t = uint32                   /* types.h:149:27 */ // Type of file link counts.
type X__off_t = int32                      /* types.h:150:25 */ // Type of file sizes and offsets.
type X__off64_t = X__quad_t                /* types.h:151:27 */ // Type of file sizes and offsets (LFS).
type X__pid_t = int32                      /* types.h:152:25 */ // Type of process identifications.
type X__fsid_t = struct{ F__val [2]int32 } /* types.h:153:26 */ // Type of file system IDs.
type X__clock_t = int32                    /* types.h:154:27 */ // Type of CPU usage counts.
type X__rlim_t = uint32                    /* types.h:155:26 */ // Type for resource measurement.
type X__rlim64_t = X__u_quad_t             /* types.h:156:28 */ // Type for resource measurement (LFS).
type X__id_t = uint32                      /* types.h:157:24 */ // General type for IDs.
type X__time_t = int32                     /* types.h:158:26 */ // Seconds since the Epoch.
type X__useconds_t = uint32                /* types.h:159:30 */ // Count of microseconds.
type X__suseconds_t = int32                /* types.h:160:31 */ // Signed count of microseconds.

type X__daddr_t = int32 /* types.h:162:27 */ // The type of a disk address.
type X__key_t = int32   /* types.h:163:25 */ // Type of an IPC key.

// Clock ID used in clock and timer functions.
type X__clockid_t = int32 /* types.h:166:29 */

// Timer ID returned by `timer_create'.
type X__timer_t = uintptr /* types.h:169:12 */

// Type to represent block size.
type X__blksize_t = int32 /* types.h:172:29 */

// Types from the Large File Support interface.

// Type to count number of disk blocks.
type X__blkcnt_t = int32       /* types.h:177:28 */
type X__blkcnt64_t = X__quad_t /* types.h:178:30 */

// Type to count file system blocks.
type X__fsblkcnt_t = uint32        /* types.h:181:30 */
type X__fsblkcnt64_t = X__u_quad_t /* types.h:182:32 */

// Type to count file system nodes.
type X__fsfilcnt_t = uint32        /* types.h:185:30 */
type X__fsfilcnt64_t = X__u_quad_t /* types.h:186:32 */

// Type of miscellaneous file system fields.
type X__fsword_t = int32 /* types.h:189:28 */

type X__ssize_t = int32 /* types.h:191:27 */ // Type of a byte count, or error.

// Signed long type used in system calls.
type X__syscall_slong_t = int32 /* types.h:194:33 */
// Unsigned long type used in system calls.
type X__syscall_ulong_t = uint32 /* types.h:196:33 */

// These few don't really vary by system, they always correspond
//    to one of the other defined types.
type X__loff_t = X__off64_t /* types.h:200:19 */ // Type of file sizes and offsets (LFS).
type X__caddr_t = uintptr   /* types.h:201:14 */

// Duplicates info from stdint.h but this is used in unistd.h.
type X__intptr_t = int32 /* types.h:204:25 */

// Duplicate info from sys/socket.h.
type X__socklen_t = uint32 /* types.h:207:23 */

// C99: An integer type that can be accessed as an atomic entity,
//    even in the presence of asynchronous interrupts.
//    It is not currently necessary for this to be machine-specific.
type X__sig_atomic_t = int32 /* types.h:212:13 */

// bits/types.h -- definitions of __*_t types underlying *_t types.
//    Copyright (C) 2002-2018 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <http://www.gnu.org/licenses/>.

// Never include this file directly; use <sys/types.h> instead.

// Integral type unchanged by default argument promotions that can
//    hold any value corresponding to members of the extended character
//    set, as well as at least one value that does not correspond to any
//    member of the extended character set.

// Conversion state information.
type X__mbstate_t = struct {
	F__count int32
	F__value struct{ F__wch uint32 }
} /* __mbstate_t.h:21:3 */

// The tag name of this struct is _G_fpos_t to preserve historic
//    C++ mangled names for functions taking fpos_t arguments.
//    That name should not be used in new code.
type X_G_fpos_t = struct {
	F__pos   X__off_t
	F__state X__mbstate_t
} /* __fpos_t.h:10:9 */

// The tag name of this struct is _G_fpos_t to preserve historic
//    C++ mangled names for functions taking fpos_t arguments.
//    That name should not be used in new code.
type X__fpos_t = X_G_fpos_t /* __fpos_t.h:14:3 */

// bits/types.h -- definitions of __*_t types underlying *_t types.
//    Copyright (C) 2002-2018 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <http://www.gnu.org/licenses/>.

// Never include this file directly; use <sys/types.h> instead.

// The tag name of this struct is _G_fpos64_t to preserve historic
//    C++ mangled names for functions taking fpos_t and/or fpos64_t
//    arguments.  That name should not be used in new code.
type X_G_fpos64_t = struct {
	F__pos   X__off64_t
	F__state X__mbstate_t
} /* __fpos64_t.h:10:9 */

// bits/types.h -- definitions of __*_t types underlying *_t types.
//    Copyright (C) 2002-2018 Free Software Foundation, Inc.
//    This file is part of the GNU C Library.
//
//    The GNU C Library is free software; you can redistribute it and/or
//    modify it under the terms of the GNU Lesser General Public
//    License as published by the Free Software Foundation; either
//    version 2.1 of the License, or (at your option) any later version.
//
//    The GNU C Library is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
//    Lesser General Public License for more details.
//
//    You should have received a copy of the GNU Lesser General Public
//    License along with the GNU C Library; if not, see
//    <http://www.gnu.org/licenses/>.

// Never include this file directly; use <sys/types.h> instead.

// The tag name of this struct is _G_fpos64_t to preserve historic
//    C++ mangled names for functions taking fpos_t and/or fpos64_t
//    arguments.  That name should not be used in new code.
type X__fpos64_t = X_G_fpos64_t /* __fpos64_t.h:14:3 */

type X_IO_FILE = struct {
	F_flags          int32
	F_IO_read_ptr    uintptr
	F_IO_read_end    uintptr
	F_IO_read_base   uintptr
	F_IO_write_base  uintptr
	F_IO_write_ptr   uintptr
	F_IO_write_end   uintptr
	F_IO_buf_base    uintptr
	F_IO_buf_end     uintptr
	F_IO_save_base   uintptr
	F_IO_backup_base uintptr
	F_IO_save_end    uintptr
	F_markers        uintptr
	F_chain          uintptr
	F_fileno         int32
	F_flags2         int32
	F_old_offset     X__off_t
	F_cur_column     uint16
	F_vtable_offset  int8
	F_shortbuf       [1]int8
	F_lock           uintptr
	F_offset         X__off64_t
	F_codecvt        uintptr
	F_wide_data      uintptr
	F_freeres_list   uintptr
	F_freeres_buf    uintptr
	F__pad5          Size_t
	F_mode           int32
	F_unused2        [40]int8
} /* __FILE.h:4:1 */

type X__FILE = X_IO_FILE /* __FILE.h:5:25 */

// The opaque type of streams.  This is the definition used elsewhere.
type FILE = X_IO_FILE /* FILE.h:7:25 */

// These macros are used by bits/stdio.h and internal headers.

// Many more flag bits are defined internally.

type Va_list = X__gnuc_va_list /* stdio.h:52:24 */

type Off_t = X__off64_t /* stdio.h:65:19 */

type Ssize_t = X__ssize_t /* stdio.h:77:19 */

// The type of the second argument to `fgetpos' and `fsetpos'.
type Fpos_t = X__fpos64_t /* stdio.h:86:20 */

// If we are compiling with optimizing read this file.  It contains
//    several optimizing inline functions and macros.

var _ int8 /* gen.c:2:13: */
