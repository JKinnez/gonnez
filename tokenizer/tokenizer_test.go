package tokenizer_test

import (
	"os"

	"github.com/JKinnez/gonnez/tokenizer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tokenizer.Flows", func() {
	Describe("SymetricKey Bearer token flow", func() {
		Context("with minimum options", func() {
			It("should create a token and read it", func() {
				key, err := tokenizer.GenerateSymmetricKey()
				Expect(err).To(BeNil())

				err = os.Setenv("GONNEZ_SYMMETRIC_KEY", key)
				Expect(err).To(BeNil())

				subject := "3x0"
				t := tokenizer.New(tokenizer.Config{})
				symetricBearerToken, err := t.GenerateSymetricBearerToken(subject, tokenizer.Durations.OneDay)
				Expect(err).To(BeNil())
				Expect(symetricBearerToken).ToNot(BeEmpty())
				Expect(symetricBearerToken).To(ContainSubstring("Bearer "))

				jsonToken, err := t.ReadSymetricBearerToken(symetricBearerToken)
				Expect(err).To(BeNil())
				Expect(jsonToken.Subject).To(Equal(subject))
			})
		})

		Context("with many options", func() {
			It("should create a token and read it", func() {
				key, err := tokenizer.GenerateSymmetricKey()
				Expect(err).To(BeNil())

				err = os.Setenv("SYMMETRIC_KEY", key)
				Expect(err).To(BeNil())

				subject := "3x0"
				t := tokenizer.New(tokenizer.Config{
					Audience:           "audience",
					Issuer:             "issuer",
					Location:           "Europe/Madrid",
					SymetricKeyEnvName: "SYMMETRIC_KEY",
				})
				symetricBearerToken, err := t.GenerateSymetricBearerToken(subject, tokenizer.Durations.OneDay)
				Expect(err).To(BeNil())
				Expect(symetricBearerToken).ToNot(BeEmpty())
				Expect(symetricBearerToken).To(ContainSubstring("Bearer "))

				jsonToken, err := t.ReadSymetricBearerToken(symetricBearerToken)
				Expect(err).To(BeNil())
				Expect(jsonToken.Subject).To(Equal(subject))
			})
		})
	})

	Describe("SymetricKey token flow", func() {
		Context("with minimum options", func() {
			It("should create a token and read it", func() {
				key, err := tokenizer.GenerateSymmetricKey()
				Expect(err).To(BeNil())

				err = os.Setenv("GONNEZ_SYMMETRIC_KEY", key)
				Expect(err).To(BeNil())

				subject := "3x0"
				t := tokenizer.New(tokenizer.Config{})
				symetricToken, err := t.GenerateSymetricToken(subject, tokenizer.Durations.OneDay)
				Expect(err).To(BeNil())
				Expect(symetricToken).ToNot(BeEmpty())

				jsonToken, err := t.ReadSymetricToken(symetricToken)
				Expect(err).To(BeNil())
				Expect(jsonToken.Subject).To(Equal(subject))
			})
		})

		Context("with many options", func() {
			It("should create a token and read it", func() {
				key, err := tokenizer.GenerateSymmetricKey()
				Expect(err).To(BeNil())

				err = os.Setenv("SYMMETRIC_KEY", key)
				Expect(err).To(BeNil())

				subject := "3x0"
				t := tokenizer.New(tokenizer.Config{
					Audience:           "audience",
					Issuer:             "issuer",
					Location:           "Europe/Madrid",
					SymetricKeyEnvName: "SYMMETRIC_KEY",
				})
				symetricToken, err := t.GenerateSymetricToken(subject, tokenizer.Durations.OneDay)
				Expect(err).To(BeNil())
				Expect(symetricToken).ToNot(BeEmpty())

				jsonToken, err := t.ReadSymetricToken(symetricToken)
				Expect(err).To(BeNil())
				Expect(jsonToken.Subject).To(Equal(subject))
			})
		})
	})

	Describe("Private/Public Bearer token flow", func() {
		Context("with minimum options", func() {
			It("should create a token and read it", func() {
				public, private, err := tokenizer.GenerateKeyPair()
				Expect(err).To(BeNil())

				err = os.Setenv("GONNEZ_PRIVATE_KEY", private)
				Expect(err).To(BeNil())

				err = os.Setenv("GONNEZ_PUBLIC_KEY", public)
				Expect(err).To(BeNil())

				subject := "3x0"
				t := tokenizer.New(tokenizer.Config{})
				privateBearerToken, err := t.GeneratePrivateBearerToken(subject, tokenizer.Durations.OneDay)
				Expect(err).To(BeNil())
				Expect(privateBearerToken).ToNot(BeEmpty())
				Expect(privateBearerToken).To(ContainSubstring("Bearer "))

				jsonToken, err := t.ReadPrivateBearerToken(privateBearerToken)
				Expect(err).To(BeNil())
				Expect(jsonToken.Subject).To(Equal(subject))
			})
		})

		Context("with many options", func() {
			It("should create a token and read it", func() {
				public, private, err := tokenizer.GenerateKeyPair()
				Expect(err).To(BeNil())

				err = os.Setenv("PRIVATE_KEY", private)
				Expect(err).To(BeNil())

				err = os.Setenv("PUBLIC_KEY", public)
				Expect(err).To(BeNil())

				subject := "3x0"
				t := tokenizer.New(tokenizer.Config{
					Audience:          "audience",
					Issuer:            "issuer",
					Location:          "Europe/Madrid",
					PrivateKeyEnvName: "PRIVATE_KEY",
					PublicKeyEnvName:  "PUBLIC_KEY",
				})
				privateBearerToken, err := t.GeneratePrivateBearerToken(subject, tokenizer.Durations.OneDay)
				Expect(err).To(BeNil())
				Expect(privateBearerToken).ToNot(BeEmpty())
				Expect(privateBearerToken).To(ContainSubstring("Bearer "))

				jsonToken, err := t.ReadPrivateBearerToken(privateBearerToken)
				Expect(err).To(BeNil())
				Expect(jsonToken.Subject).To(Equal(subject))
			})
		})
	})

	Describe("Private/Public token flow", func() {
		Context("with minimum options", func() {
			It("should create a token and read it", func() {
				public, private, err := tokenizer.GenerateKeyPair()
				Expect(err).To(BeNil())

				err = os.Setenv("GONNEZ_PRIVATE_KEY", private)
				Expect(err).To(BeNil())

				err = os.Setenv("GONNEZ_PUBLIC_KEY", public)
				Expect(err).To(BeNil())

				subject := "3x0"
				t := tokenizer.New(tokenizer.Config{})
				privateToken, err := t.GeneratePrivateToken(subject, tokenizer.Durations.OneDay)
				Expect(err).To(BeNil())
				Expect(privateToken).ToNot(BeEmpty())

				jsonToken, err := t.ReadPrivateToken(privateToken)
				Expect(err).To(BeNil())
				Expect(jsonToken.Subject).To(Equal(subject))
			})
		})

		Context("with many options", func() {
			It("should create a token and read it", func() {
				public, private, err := tokenizer.GenerateKeyPair()
				Expect(err).To(BeNil())

				err = os.Setenv("PRIVATE_KEY", private)
				Expect(err).To(BeNil())

				err = os.Setenv("PUBLIC_KEY", public)
				Expect(err).To(BeNil())

				subject := "3x0"
				t := tokenizer.New(tokenizer.Config{
					Audience:          "audience",
					Issuer:            "issuer",
					Location:          "Europe/Madrid",
					PrivateKeyEnvName: "PRIVATE_KEY",
					PublicKeyEnvName:  "PUBLIC_KEY",
				})
				privateToken, err := t.GeneratePrivateToken(subject, tokenizer.Durations.OneDay)
				Expect(err).To(BeNil())
				Expect(privateToken).ToNot(BeEmpty())

				jsonToken, err := t.ReadPrivateToken(privateToken)
				Expect(err).To(BeNil())
				Expect(jsonToken.Subject).To(Equal(subject))
			})
		})
	})
})
