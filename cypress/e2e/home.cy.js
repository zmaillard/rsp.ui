describe("Home page", () => {
  it("displays all 4 countries on the home page", () => {
    cy.visit("http://localhost:1313")
    cy.get('[data-test="country-tab-item"]')
        .eq(0)
        .within(() => {
          cy.get('[data-test="country-tab-item-name"]').should("contain", "United States")

        })
      cy.get('[data-test="country-tab-item"]')
          .eq(1)
          .within(() => {
              cy.get('[data-test="country-tab-item-name"]').should("contain", "Canada")

          })

      cy.get('[data-test="country-tab-item"]')
          .eq(2)
          .within(() => {
              cy.get('[data-test="country-tab-item-name"]').should("contain", "Costa Rica")

          })
      cy.get('[data-test="country-tab-item"]')
          .eq(3)
          .within(() => {
              cy.get('[data-test="country-tab-item-name"]').should("contain", "México")

          })
  })

    it("displays browse by state on the home page", () => {
        cy.visit("http://localhost:1313")
        cy.get('[data-test="country-browse-by-subdivision"]').should("contain", "Browse By State")
    })

    it("displays browse by state on the home page", () => {
        cy.visit("http://localhost:1313")
        cy.get('[data-test="country-browse-by-highway"]').should("contain", "Browse By Highway")
    })
})