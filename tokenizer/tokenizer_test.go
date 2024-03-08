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

				timeInHours := 1
				subject := "3x0"
				t := tokenizer.New(tokenizer.EncriptionOpts{Subject: subject})
				symetricBearerToken, err := t.GenerateSymetricBearerToken(timeInHours)
				Expect(err).To(BeNil())
				Expect(symetricBearerToken).ToNot(BeEmpty())
				Expect(symetricBearerToken).To(ContainSubstring("Bearer "))

				r := tokenizer.NewReader(tokenizer.ReaderOptions{Token: symetricBearerToken})
				jsonToken, err := r.ReadSymetricBearerToken()
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

				timeInHours := 1
				subject := "3x0"
				t := tokenizer.New(tokenizer.EncriptionOpts{
					Audience:           "audience",
					Issuer:             "issuer",
					Subject:            subject,
					Location:           "Europe/Madrid",
					SymetricKeyEnvName: "SYMMETRIC_KEY",
				})
				symetricBearerToken, err := t.GenerateSymetricBearerToken(timeInHours)
				Expect(err).To(BeNil())
				Expect(symetricBearerToken).ToNot(BeEmpty())
				Expect(symetricBearerToken).To(ContainSubstring("Bearer "))

				r := tokenizer.NewReader(tokenizer.ReaderOptions{
					Token:              symetricBearerToken,
					Issuer:             "issuer",
					Location:           "Europe/Madrid",
					SymetricKeyEnvName: "SYMMETRIC_KEY",
				})
				jsonToken, err := r.ReadSymetricBearerToken()
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

				timeInHours := 1
				subject := "3x0"
				t := tokenizer.New(tokenizer.EncriptionOpts{Subject: subject})
				symetricBearerToken, err := t.GenerateSymetricToken(timeInHours)
				Expect(err).To(BeNil())
				Expect(symetricBearerToken).ToNot(BeEmpty())

				r := tokenizer.NewReader(tokenizer.ReaderOptions{Token: symetricBearerToken})
				jsonToken, err := r.ReadSymetricToken()
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

				timeInHours := 1
				subject := "3x0"
				t := tokenizer.New(tokenizer.EncriptionOpts{
					Audience:           "audience",
					Issuer:             "issuer",
					Subject:            subject,
					Location:           "Europe/Madrid",
					SymetricKeyEnvName: "SYMMETRIC_KEY",
				})
				symetricBearerToken, err := t.GenerateSymetricToken(timeInHours)
				Expect(err).To(BeNil())
				Expect(symetricBearerToken).ToNot(BeEmpty())

				r := tokenizer.NewReader(tokenizer.ReaderOptions{
					Token:              symetricBearerToken,
					Issuer:             "issuer",
					Location:           "Europe/Madrid",
					SymetricKeyEnvName: "SYMMETRIC_KEY",
				})
				jsonToken, err := r.ReadSymetricToken()
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

				timeInHours := 1
				subject := "3x0"
				t := tokenizer.New(tokenizer.EncriptionOpts{Subject: subject})
				privateBearerToken, err := t.GeneratePrivateBearerToken(timeInHours)
				Expect(err).To(BeNil())
				Expect(privateBearerToken).ToNot(BeEmpty())
				Expect(privateBearerToken).To(ContainSubstring("Bearer "))

				r := tokenizer.NewReader(tokenizer.ReaderOptions{Token: privateBearerToken})
				jsonToken, err := r.ReadPrivateBearerToken()
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

				timeInHours := 1
				subject := "3x0"
				t := tokenizer.New(tokenizer.EncriptionOpts{
					Audience:          "audience",
					Issuer:            "issuer",
					Subject:           subject,
					Location:          "Europe/Madrid",
					PrivateKeyEnvName: "PRIVATE_KEY",
				})
				privateBearerToken, err := t.GeneratePrivateBearerToken(timeInHours)
				Expect(err).To(BeNil())
				Expect(privateBearerToken).ToNot(BeEmpty())
				Expect(privateBearerToken).To(ContainSubstring("Bearer "))

				r := tokenizer.NewReader(tokenizer.ReaderOptions{
					Token:            privateBearerToken,
					Issuer:           "issuer",
					Location:         "Europe/Madrid",
					PublicKeyEnvName: "PUBLIC_KEY",
				})
				jsonToken, err := r.ReadPrivateBearerToken()
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

				timeInHours := 1
				subject := "3x0"
				t := tokenizer.New(tokenizer.EncriptionOpts{Subject: subject})
				privateBearerToken, err := t.GeneratePrivateToken(timeInHours)
				Expect(err).To(BeNil())
				Expect(privateBearerToken).ToNot(BeEmpty())

				r := tokenizer.NewReader(tokenizer.ReaderOptions{Token: privateBearerToken})
				jsonToken, err := r.ReadPrivateToken()
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

				timeInHours := 1
				subject := "3x0"
				t := tokenizer.New(tokenizer.EncriptionOpts{
					Audience:          "audience",
					Issuer:            "issuer",
					Subject:           subject,
					Location:          "Europe/Madrid",
					PrivateKeyEnvName: "PRIVATE_KEY",
				})
				privateBearerToken, err := t.GeneratePrivateToken(timeInHours)
				Expect(err).To(BeNil())
				Expect(privateBearerToken).ToNot(BeEmpty())

				r := tokenizer.NewReader(tokenizer.ReaderOptions{
					Token:            privateBearerToken,
					Issuer:           "issuer",
					Location:         "Europe/Madrid",
					PublicKeyEnvName: "PUBLIC_KEY",
				})
				jsonToken, err := r.ReadPrivateToken()
				Expect(err).To(BeNil())
				Expect(jsonToken.Subject).To(Equal(subject))
			})
		})
	})
})
