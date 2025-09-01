describe('Homepage Tests', () => {
    beforeEach(() => {
        cy.visit('/')
    })

    it('should load the title', () => {
        cy.get('h1').contains('roadsign.pictures').should('be.visible')
        cy.get('h2').contains('A gallery of road sign photographs').should('be.visible')
    })

    it('should display Browse By State section', () => {
        cy.get('[data-cy="browse-by-state-title"]').contains("Browse By State").should('be.visible')
        cy.get('[data-cy="states-container"]').should('be.visible')
        // Check if state links exist
        cy.get('[data-cy="state-link-california"]').should('exist')
        cy.get('[data-cy="state-link-oregon"]').should('exist')
        cy.get('[data-cy="state-link-washington"]').should('exist')
    })

    it('should display Browse By Highway section', () => {
        cy.get('[data-cy="browse-by-highway-title"]').contains("Browse By Highway").should('be.visible')
        cy.get('[data-cy="highways-container"]').should('be.visible')
        // Check if highway type links exist
        cy.get('[data-cy="highway-link-interstate-highway"]').should('exist')
        cy.get('[data-cy="highway-link-us-highway"]').should('exist')
        cy.get('[data-cy="highway-link-california-state-highway"]').should('exist')
    })

    it('should navigate to a state page when clicked', () => {
        cy.get('[data-cy="state-link-california"]').click()
        cy.url().should('include', '/state/california/')
    })

    it('should navigate to a highway type page when clicked', () => {
        cy.get('[data-cy="highway-link-interstate-highway"]').click()
        cy.url().should('include', '/highwaytype/interstate-highway/')
    })
})
