package main

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	pb "github.com/tetsuya/microservice-learning/src/userservice/protogen"
)

var _ = Describe("Main", func() {
	Context("Request with total price", func() {
		var (
			resp *pb.User
			err  error
		)

		BeforeEach(func() {
			s := server{}
			id := int32(123)
			req := &pb.GetUserRequest{Id: id}
			resp, err = s.GetUser(context.Background(), req)
		})

		It("will return no error", func() {
			Expect(err).ShouldNot(HaveOccurred())
		})
		It("will return requested ID", func() {
			Expect(resp.Id).To(Equal(int32(123)))
		})
		It("will return mocked name", func() {
			Expect(resp.Name).To(Equal("John Smith"))
		})
		It("will return mocked email", func() {
			Expect(resp.Email).To(Equal("johnsmith@example.com"))
		})
	})
})
