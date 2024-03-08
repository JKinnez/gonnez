package tokenizer_test

import (
	"github.com/JKinnez/gonnez/tokenizer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tokenizer.Generator", func() {
	Describe("GenerateSymetricKey", func() {
		It("should return a string", func() {
			key, err := tokenizer.GenerateSymmetricKey()

			Expect(err).To(BeNil())
			Expect(key).ToNot(BeEmpty())
		})
	})

	Describe("GenerateKeyPair", func() {
		It("should return a string", func() {
			private, public, err := tokenizer.GenerateKeyPair()

			Expect(err).To(BeNil())
			Expect(private).ToNot(BeEmpty())
			Expect(public).ToNot(BeEmpty())
		})
	})

})
