package iyaml

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/yaml.v2"
)

var _ = Context("IYAML", func() {
	Context("New", func() {
		It("should return IYAML", func() {
			/* arrange/act/assert */
			Expect(New()).
				Should(Not(BeNil()))
		})
	})
	Context("Marshal", func() {
		Context("doesn't err", func() {
			It("should return expected result", func() {
				/* arrange */
				providedStruct := struct {
					Field1 string
					Field2 int
				}{
					Field1: "dummyString1",
					Field2: 1000,
				}

				expectedBytes, _ := yaml.Marshal(providedStruct)

				objectUnderTest := New()

				/* act */
				actualBytes, actualErr := objectUnderTest.Marshal(providedStruct)

				/* assert */
				Expect(actualBytes).To(Equal(expectedBytes))
				Expect(actualErr).To(BeNil())
			})
		})
	})
	Context("Unmarshal", func() {
		Context("errs", func() {
			It("should return expected err", func() {
				/* arrange */
				erroneousInput := []byte("$$")
				objectUnderTest := New()

				expectedErr := yaml.Unmarshal(erroneousInput, &struct{}{})

				/* act */
				actualErr := objectUnderTest.Unmarshal(erroneousInput, &struct{}{})

				/* assert */
				Expect(actualErr).To(Equal(expectedErr))
			})
		})
		Context("doesn't err", func() {
			It("should unmarshal expected object", func() {
				/* arrange */
				expectedStruct := struct {
					Field1 string
					Field2 int
				}{
					Field1: "dummyString1",
					Field2: 1000,
				}

				providedBytes, err := yaml.Marshal(expectedStruct)
				if nil != err {
					Fail(err.Error())
				}

				actualStruct := struct {
					Field1 string
					Field2 int
				}{}

				objectUnderTest := New()

				/* act */
				actualErr := objectUnderTest.Unmarshal(providedBytes, &actualStruct)

				/* assert */
				Expect(actualStruct).To(Equal(expectedStruct))
				Expect(actualErr).To(BeNil())
			})
		})
	})
})
