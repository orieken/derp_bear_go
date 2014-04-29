package derp_bear_go_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sourcegraph/go-selenium"
	"fmt"
)

var _ = Describe("SimpleForm", func() {

    var webDriver selenium.WebDriver
    var err error

    BeforeEach(func() {
        caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
        if webDriver, err = selenium.NewRemote(caps, "http://localhost:9000/wd/hub"); err != nil {
            fmt.Printf("Failed to open session: %s\n", err)
            return
        }
        webDriver.Get("http://derp-bear.herokuapp.com/")

    })

    AfterEach(func(){
        webDriver.Quit()
    })

    Describe("Derp-Bear", func() {
        Context("Landing page", func() {
            It("has a title", func() {
            Expect(webDriver.Title()).To(Equal("Welcome to Derp-Bear"))
            })
        })

        Context("Simple Form Example", func() {
            It("has successful submission message", func() {

                var simpleFormExampleLink selenium.WebElement
                simpleFormExampleLink, err = webDriver.FindElement(selenium.ById, "simple_form_example")
                simpleFormExampleLink.Click()

                var firstName selenium.WebElement
                firstName, err = webDriver.FindElement(selenium.ById, "first_name")
                firstName.SendKeys("Oscar")

                var lastName selenium.WebElement
                lastName, err = webDriver.FindElement(selenium.ById, "last_name")
                lastName.SendKeys("Rieken")

                var emailAddress selenium.WebElement
                emailAddress, err = webDriver.FindElement(selenium.ById, "email")
                emailAddress.SendKeys("orieken@derp-bear.com")

                var submitButton selenium.WebElement
                submitButton, err = webDriver.FindElement(selenium.ByCSSSelector, "[type=submit]")
                submitButton.Click()

                var successMessage selenium.WebElement
                successMessage, err = webDriver.FindElement(selenium.ByTagName, "h1")

                Expect(successMessage.Text()).To(Equal("Thank you Oscar for submitting the form"))
            })
        })
     })
})
