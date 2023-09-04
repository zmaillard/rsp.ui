describe("County page", () => {
  it("displays the correct title", () => {
      cy.visit("http://localhost:1313/county/idaho_ada-county/")
      cy.get('[data-test="county-title"]').should('contain', 'Ada County, Idaho')
  })

    it("has a bread crumb", () => {
        cy.visit("http://localhost:1313/county/idaho_ada-county/")
        cy.get('[data-test="breadcrumb"]').eq(0).within( () => {
            cy.get('[data-test="breadcrumb-home"]').should("contain", "Home")
            cy.get('[data-test="breadcrumb-country"]').should("contain", "United States")
            cy.get('[data-test="breadcrumb-state"]').should("contain", "Idaho")
            cy.get('[data-test="breadcrumb-county"]').should("contain", "Ada County")
        })
    })
})