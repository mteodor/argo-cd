package grpc

import (
	"github.com/mennanov/fmutils"
	googleproto "google.golang.org/protobuf/proto"
	googleprotoreflect "google.golang.org/protobuf/reflect/protoreflect"
)

/*
	Takes an existing message and a new message and overwrites the existing fm.Paths with the
	values found in new
*/
func Mask(existing googleproto.Message, new googleproto.Message, msk fmutils.NestedMask) {
	exrft := existing.ProtoReflect()
	newrft := new.ProtoReflect()

	if exrft.Type() != newrft.Type() {
		return
	}

	fields := newrft.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		fd := fields.Get(i)
		m, ok := msk[string(fd.Name())]

		if ok {
			if len(m) == 0 {
				exrft.Set(fd, newrft.Get(fd))
				continue
			}

			if fd.Kind() == googleprotoreflect.MessageKind {
				Mask(exrft.Get(fd).Message().Interface(), newrft.Get(fd).Message().Interface(), m)
			}
		}
	}
}