describe("Country page", () => {
  it("displays all 4 countries on the united states page", () => {
    cy.visit("http://localhost:1313/country/united-states")
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

    it("displays browse by state on the country page", () => {
        cy.visit("http://localhost:1313/country/united-states")
        cy.get('[data-test="country-browse-by-subdivision"]').should("contain", "Browse By State")
    })

    it("displays browse by highway on the country page", () => {
        cy.visit("http://localhost:1313/country/united-states")
        cy.get('[data-test="country-browse-by-highway"]').should("contain", "Browse By Highway")
    })

    it("has a bread crumb", () => {
        cy.visit("http://localhost:1313/country/canada")
        cy.get('[data-test="breadcrumb"]').eq(0).within( () => {
            cy.get('[data-test="breadcrumb-home"]').should("contain", "Home")
            cy.get('[data-test="breadcrumb-country"]').should("contain", "Canada")

        })
    })
})