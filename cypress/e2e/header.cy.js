describe("Header", () => {
    it("displays the correct links on the navbar", () => {
        cy.visit("http://localhost:1313")
        cy.get('[data-test="navbar-home-link"]')
            .eq(0)
            .within(() => {
            })
    })
})