package tokenizer_test

import (
	"github.com/JKinnez/gonnez/tokenizer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tokenizer.Flows", func() {
	Describe("SymetricKey Bearer token flow", func() {
		Context("whit no options", func() {
			It("should return an error", func() {
				subject := "3x0"
				t := tokenizer.New(tokenizer.Config{})
				_, err := t.GenerateSymetricBearerToken(subject, tokenizer.Durations.OneDay)
				Expect(err).ToNot(BeNil())
				Expect(err).To(MatchError(tokenizer.ErrNotSymetricKey))
			})
		})

		Context("with minimum options", func() {
			It("should create a token and read it", func() {
				key, err := tokenizer.GenerateSymmetricKey()
				Expect(err).To(BeNil())

				subject := "3x0"
				t := tokenizer.New(tokenizer.Config{
					SymetricKey: key,
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

		Context("with many options", func() {
			It("should create a token and read it", func() {
				key, err := tokenizer.GenerateSymmetricKey()
				Expect(err).To(BeNil())

				subject := "3x0"
				t := tokenizer.New(tokenizer.Config{
					Audience:    "audience",
					Issuer:      "issuer",
					Location:    "Europe/Madrid",
					SymetricKey: key,
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
		Context("whit no options", func() {
			It("should return an error", func() {
				subject := "3x0"
				t := tokenizer.New(tokenizer.Config{})
				_, err := t.GenerateSymetricToken(subject, tokenizer.Durations.OneDay)
				Expect(err).ToNot(BeNil())
				Expect(err).To(MatchError(tokenizer.ErrNotSymetricKey))
			})
		})

		Context("with minimum options", func() {
			It("should create a token and read it", func() {
				key, err := tokenizer.GenerateSymmetricKey()
				Expect(err).To(BeNil())

				subject := "3x0"
				t := tokenizer.New(tokenizer.Config{
					SymetricKey: key,
				})
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

				subject := "3x0"
				t := tokenizer.New(tokenizer.Config{
					Audience:    "audience",
					Issuer:      "issuer",
					Location:    "Europe/Madrid",
					SymetricKey: key,
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
		Context("with no private key option set", func() {
			It("should return an error on encryption", func() {
				subject := "3x0"
				t := tokenizer.New(tokenizer.Config{})
				_, err := t.GeneratePrivateBearerToken(subject, tokenizer.Durations.OneDay)
				Expect(err).ToNot(BeNil())
				Expect(err).To(MatchError(tokenizer.ErrNotPrivateKey))
			})
		})

		Context("with no public key option set", func() {
			It("should return an error on decryption", func() {
				_, private, err := tokenizer.GenerateKeyPair()
				Expect(err).To(BeNil())

				subject := "3x0"
				t := tokenizer.New(tokenizer.Config{
					PrivateKey: private,
				})
				privateBearerToken, err := t.GeneratePrivateBearerToken(subject, tokenizer.Durations.OneDay)
				Expect(err).To(BeNil())
				Expect(privateBearerToken).ToNot(BeEmpty())
				Expect(privateBearerToken).To(ContainSubstring("Bearer "))

				_, err = t.ReadPrivateBearerToken(privateBearerToken)
				Expect(err).ToNot(BeNil())
				Expect(err).To(MatchError(tokenizer.ErrNotPublicKey))
			})
		})

		Context("with minimum options", func() {
			It("should create a token and read it", func() {
				public, private, err := tokenizer.GenerateKeyPair()
				Expect(err).To(BeNil())

				subject := "3x0"
				t := tokenizer.New(tokenizer.Config{
					PublicKey:  public,
					PrivateKey: private,
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

		Context("with many options", func() {
			It("should create a token and read it", func() {
				public, private, err := tokenizer.GenerateKeyPair()
				Expect(err).To(BeNil())

				subject := "3x0"
				t := tokenizer.New(tokenizer.Config{
					Audience:   "audience",
					Issuer:     "issuer",
					Location:   "Europe/Madrid",
					PrivateKey: private,
					PublicKey:  public,
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
		Context("with no private key option set", func() {
			It("should return an error on encryption", func() {
				subject := "3x0"
				t := tokenizer.New(tokenizer.Config{})
				_, err := t.GeneratePrivateToken(subject, tokenizer.Durations.OneDay)
				Expect(err).ToNot(BeNil())
				Expect(err).To(MatchError(tokenizer.ErrNotPrivateKey))
			})
		})

		Context("with no public key option set", func() {
			It("should return an error on decryption", func() {
				_, private, err := tokenizer.GenerateKeyPair()
				Expect(err).To(BeNil())

				subject := "3x0"
				t := tokenizer.New(tokenizer.Config{
					PrivateKey: private,
				})
				privateToken, err := t.GeneratePrivateToken(subject, tokenizer.Durations.OneDay)
				Expect(err).To(BeNil())
				Expect(privateToken).ToNot(BeEmpty())
				Expect(privateToken).ToNot(BeEmpty())

				_, err = t.ReadPrivateBearerToken(privateToken)
				Expect(err).ToNot(BeNil())
				Expect(err).To(MatchError(tokenizer.ErrNotPublicKey))
			})
		})

		Context("with minimum options", func() {
			It("should create a token and read it", func() {
				public, private, err := tokenizer.GenerateKeyPair()
				Expect(err).To(BeNil())

				subject := "3x0"
				t := tokenizer.New(tokenizer.Config{
					PublicKey:  public,
					PrivateKey: private,
				})
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

				subject := "3x0"
				t := tokenizer.New(tokenizer.Config{
					Audience:   "audience",
					Issuer:     "issuer",
					Location:   "Europe/Madrid",
					PrivateKey: private,
					PublicKey:  public,
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
