package errors_test

import (
	"fmt"
	"io"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/neighborly/go-errors"
	pkgErrors "github.com/pkg/errors"
)

var _ = Describe("Errors", func() {
	It("creates an error", func() {
		msg := "an error with a message"
		err := errors.New(msg)

		Expect(errors.Code(err)).To(Equal(errors.InternalError))
		Expect(err.Error()).To(Equal(msg))
		Expect(errors.DisplayMessage(err)).To(Equal("Internal Error"))
		Expect(fmt.Sprintf("%+v", errors.StackTrace(err))).To(ContainSubstring("/errors_test.go:"))
	})

	It("creates an NotFound error", func() {
		msg := "an error with a message"
		err := errors.NotFound.New(msg)

		Expect(errors.Code(err)).To(Equal(errors.NotFound))
		Expect(err.Error()).To(Equal(msg))
		Expect(errors.DisplayMessage(err)).To(Equal("Not Found"))
		Expect(fmt.Sprintf("%+v", errors.StackTrace(err))).To(ContainSubstring("/errors_test.go:"))
	})

	It("creates an InternalError error", func() {
		msg := "an error with a message"
		err := errors.InternalError.New(msg)

		Expect(errors.Code(err)).To(Equal(errors.InternalError))
		Expect(err.Error()).To(Equal(msg))
		Expect(errors.DisplayMessage(err)).To(Equal("Internal Error"))
		Expect(fmt.Sprintf("%+v", errors.StackTrace(err))).To(ContainSubstring("/errors_test.go:"))
	})

	It("creates an InvalidArgument error", func() {
		msg := "an error with a message"
		err := errors.InvalidArgument.New(msg)

		Expect(errors.Code(err)).To(Equal(errors.InvalidArgument))
		Expect(err.Error()).To(Equal(msg))
		Expect(errors.DisplayMessage(err)).To(Equal("Invalid Argument"))
		Expect(fmt.Sprintf("%+v", errors.StackTrace(err))).To(ContainSubstring("/errors_test.go:"))
	})

	It("creates an Unauthenticated error", func() {
		msg := "an error with a message"
		err := errors.Unauthenticated.New(msg)

		Expect(errors.Code(err)).To(Equal(errors.Unauthenticated))
		Expect(err.Error()).To(Equal(msg))
		Expect(errors.DisplayMessage(err)).To(Equal("Unauthenticated"))
		Expect(fmt.Sprintf("%+v", errors.StackTrace(err))).To(ContainSubstring("/errors_test.go:"))
	})

	It("creates an PermissionDenied error", func() {
		msg := "an error with a message"
		err := errors.PermissionDenied.New(msg)

		Expect(errors.Code(err)).To(Equal(errors.PermissionDenied))
		Expect(err.Error()).To(Equal(msg))
		Expect(errors.DisplayMessage(err)).To(Equal("Permission Denied"))
		Expect(fmt.Sprintf("%+v", errors.StackTrace(err))).To(ContainSubstring("/errors_test.go:"))
	})

	It("creates an Unknown error", func() {
		msg := "an error with a message"
		err := errors.Unknown.New(msg)

		Expect(errors.Code(err)).To(Equal(errors.Unknown))
		Expect(err.Error()).To(Equal(msg))
		Expect(errors.DisplayMessage(err)).To(Equal("Unknown"))
		Expect(fmt.Sprintf("%+v", errors.StackTrace(err))).To(ContainSubstring("/errors_test.go:"))
	})

	It("wraps an error", func() {
		msg := "an error"
		err := errors.Wrap(errors.New(msg), "some other message")

		Expect(errors.Code(err)).To(Equal(errors.InternalError))
		Expect(err.Error()).To(Equal("some other message: " + msg))
		Expect(fmt.Sprintf("%+v", errors.StackTrace(err))).To(ContainSubstring("/errors_test.go:"))
	})

	It("can get error cause", func() {
		err := io.EOF
		wrapped := errors.Wrap(err, "some other message")

		Expect(wrapped.Error()).To(Equal("some other message: " + err.Error()))
		Expect(errors.Cause(wrapped)).To(Equal(err))
		Expect(errors.DisplayMessage(err)).To(Equal("Internal Error"))
	})

	It("can get error cause with multiple wraps", func() {
		err := io.EOF
		wrapped := errors.Wrap(err, "some other message")
		wrappedTwo := errors.Wrap(wrapped, "last error")

		Expect(errors.Cause(wrappedTwo)).To(Equal(err))
		Expect(errors.DisplayMessage(err)).To(Equal("Internal Error"))
	})

	It("can get error cause of a pkg/errors", func() {
		msg := "an error"
		err := pkgErrors.New(msg)
		wrapped := pkgErrors.Wrap(err, "some other message")

		Expect(errors.Cause(wrapped)).To(Equal(err))
		Expect(errors.DisplayMessage(err)).To(Equal("Internal Error"))
	})

	It("allows to change display message", func() {
		msg := "an error with a message"
		err := errors.WithDisplayMessage(errors.New(msg), "We had a problem")

		Expect(errors.Code(err)).To(Equal(errors.InternalError))
		Expect(errors.DisplayMessage(err)).To(Equal("We had a problem"))
	})

	It("allows to change display message of custom error", func() {
		msg := "an error with a message"
		err := errors.WithDisplayMessage(errors.NotFound.New(msg), "The record was not found")

		Expect(errors.Code(err)).To(Equal(errors.NotFound))
		Expect(errors.DisplayMessage(err)).To(Equal("The record was not found"))
	})

	It("can get stack trace of non go-error error", func() {
		err := io.EOF

		Expect(fmt.Sprintf("%+v", errors.StackTrace(err))).To(ContainSubstring("/errors_test.go:"))
	})

	Describe("#Code", func() {
		It("should return a wrapped error code", func() {
			err := errors.Join(fmt.Errorf("test"), errors.InvalidArgument.New("invalid argument"))
			Expect(errors.Code(err)).To(Equal(errors.InvalidArgument))
		})
	})
})
