describe("State page", () => {
  it("displays the correct title", () => {
      cy.visit("http://localhost:1313/state/alberta")
      cy.get('[data-test="state-title"]').should('contain', 'Alberta')
  })

    it("has a bread crumb", () => {
        cy.visit("http://localhost:1313/state/alberta")
        cy.get('[data-test="breadcrumb"]').eq(0).within( () => {
            cy.get('[data-test="breadcrumb-home"]').should("contain", "Home")
            cy.get('[data-test="breadcrumb-country"]').should("contain", "Canada")
            cy.get('[data-test="breadcrumb-state"]').should("contain", "Alberta")

        })
    })
})