describe("Place page", () => {
    it("displays the correct title", () => {
        cy.visit("http://localhost:1313/place/idaho_boise/")
        cy.get('[data-test="place-title"]').should('contain', 'Boise, Idaho')
    })

    it("has a bread crumb", () => {
        cy.visit("http://localhost:1313/place/idaho_boise/")
        cy.get('[data-test="breadcrumb"]').eq(0).within( () => {
            cy.get('[data-test="breadcrumb-home"]').should("contain", "Home")
            cy.get('[data-test="breadcrumb-country"]').should("contain", "United States")
            cy.get('[data-test="breadcrumb-state"]').should("contain", "Idaho")
            cy.get('[data-test="breadcrumb-place"]').should("contain", "Boise")

        })
    })
})