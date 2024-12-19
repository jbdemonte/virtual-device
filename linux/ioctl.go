package linux

import "unsafe"

// https://github.com/torvalds/linux/blob/master/include/uapi/asm-generic/ioctl.h

// based on https://gist.github.com/artizirk/d08a889164701b1c0e0428e280674429

const (
	_IOC_NRBITS   = 8
	_IOC_TYPEBITS = 8
	_IOC_SIZEBITS = 14
	_IOC_DIRBITS  = 2

	_IOC_NRSHIFT   = 0
	_IOC_TYPESHIFT = _IOC_NRSHIFT + _IOC_NRBITS
	_IOC_SIZESHIFT = _IOC_TYPESHIFT + _IOC_TYPEBITS
	_IOC_DIRSHIFT  = _IOC_SIZESHIFT + _IOC_SIZEBITS

	_IOC_NONE  = 0
	_IOC_WRITE = 1
	_IOC_READ  = 2
)

func _IOC(dir, typ, nr, size uintptr) uintptr {
	return ((dir << _IOC_DIRSHIFT) |
		(typ << _IOC_TYPESHIFT) |
		(nr << _IOC_NRSHIFT) |
		(size << _IOC_SIZESHIFT))
}

func _IOC_TYPECHECK(t interface{}) uintptr {
	return uintptr(unsafe.Sizeof(t))
}

func _IO(typ byte, nr uintptr) uintptr {
	return _IOC(_IOC_NONE, uintptr(typ), nr, 0)
}

func _IOR(typ byte, nr uintptr, t interface{}) uintptr {
	return _IOC(_IOC_READ, uintptr(typ), nr, _IOC_TYPECHECK(t))
}

func _IOW(typ byte, nr uintptr, t interface{}) uintptr {
	return _IOC(_IOC_WRITE, uintptr(typ), nr, _IOC_TYPECHECK(t))
}

func _IOWR(typ byte, nr uintptr, t interface{}) uintptr {
	return _IOC(_IOC_READ|_IOC_WRITE, uintptr(typ), nr, _IOC_TYPECHECK(t))
}
