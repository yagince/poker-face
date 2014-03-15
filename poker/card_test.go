package poker

import (
	. "github.com/r7kamura/gospel"
	"testing"
	"."
)

func TestDescribe(t *testing.T) {
	Describe(t, "Choice#ShouldSave", func() {

		Context("Reset is true", func() {

			It("should be false", func() {
				choice := poker.Choice{Reset: true}
				Expect(choice.ShouldSave()).To(Equal, false)
			})

		})

		Context("Open is true", func() {

			It("should be false", func() {
				choice := poker.Choice{Open: true}
				Expect(choice.ShouldSave()).To(Equal, false)
			})

		})

		Context("Reset & Open is false both", func() {

			It("should be true", func() {
				choice := poker.Choice{}
				Expect(choice.ShouldSave()).To(Equal, true)
			})

		})

	})
}
