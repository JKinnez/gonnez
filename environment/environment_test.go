package environment_test

import (
	"os"

	"github.com/JKinnez/gonnez/environment"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Environment", func() {
	Describe("IsProduction", func() {
		Context("when the environment is production", func() {
			It("should return true", func() {
				_ = os.Setenv("ENV", "production")
				environment.Init(os.Getenv("ENV"))
				Expect(environment.IsProduction()).To(BeTrue())
			})
		})

		Context("when the environment is not production", func() {
			It("should return false", func() {
				_ = os.Setenv("ENV", "development")
				environment.Init(os.Getenv("ENV"))
				Expect(environment.IsProduction()).To(BeFalse())
			})

			It("should return false", func() {
				_ = os.Setenv("ENV", "test")
				environment.Init(os.Getenv("ENV"))
				Expect(environment.IsProduction()).To(BeFalse())
			})

			It("should return false", func() {
				_ = os.Setenv("ENV", "ci")
				environment.Init(os.Getenv("ENV"))
				Expect(environment.IsProduction()).To(BeFalse())
			})
		})
	})

	Describe("IsDevelopment", func() {
		Context("when the environment is development", func() {
			It("should return true", func() {
				_ = os.Setenv("ENV", "development")
				environment.Init(os.Getenv("ENV"))
				Expect(environment.IsDevelopment()).To(BeTrue())
			})
		})

		Context("when the environment is not development", func() {
			It("should return false", func() {
				_ = os.Setenv("ENV", "production")
				environment.Init(os.Getenv("ENV"))
				Expect(environment.IsDevelopment()).To(BeFalse())
			})

			It("should return false", func() {
				_ = os.Setenv("ENV", "test")
				environment.Init(os.Getenv("ENV"))
				Expect(environment.IsDevelopment()).To(BeFalse())
			})

			It("should return false", func() {
				_ = os.Setenv("ENV", "ci")
				environment.Init(os.Getenv("ENV"))
				Expect(environment.IsDevelopment()).To(BeFalse())
			})
		})
	})

	Describe("IsTest", func() {
		Context("when the environment is test", func() {
			It("should return true", func() {
				_ = os.Setenv("ENV", "test")
				environment.Init(os.Getenv("ENV"))
				Expect(environment.IsTest()).To(BeTrue())
			})
		})

		Context("when the environment is not test", func() {
			It("should return false", func() {
				_ = os.Setenv("ENV", "production")
				environment.Init(os.Getenv("ENV"))
				Expect(environment.IsTest()).To(BeFalse())
			})

			It("should return false", func() {
				_ = os.Setenv("ENV", "development")
				environment.Init(os.Getenv("ENV"))
				Expect(environment.IsTest()).To(BeFalse())
			})

			It("should return false", func() {
				_ = os.Setenv("ENV", "ci")
				environment.Init(os.Getenv("ENV"))
				Expect(environment.IsTest()).To(BeFalse())
			})
		})
	})

	Describe("IsCI", func() {
		Context("when the environment is ci", func() {
			It("should return true", func() {
				_ = os.Setenv("ENV", "ci")
				environment.Init(os.Getenv("ENV"))
				Expect(environment.IsCI()).To(BeTrue())
			})
		})

		Context("when the environment is not ci", func() {
			It("should return false", func() {
				_ = os.Setenv("ENV", "production")
				environment.Init(os.Getenv("ENV"))
				Expect(environment.IsCI()).To(BeFalse())
			})

			It("should return false", func() {
				_ = os.Setenv("ENV", "development")
				environment.Init(os.Getenv("ENV"))
				Expect(environment.IsCI()).To(BeFalse())
			})

			It("should return false", func() {
				_ = os.Setenv("ENV", "test")
				environment.Init(os.Getenv("ENV"))
				Expect(environment.IsCI()).To(BeFalse())
			})
		})
	})
})
