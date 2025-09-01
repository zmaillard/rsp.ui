describe('States Page Tests', () => {
    beforeEach(() => {
        cy.visit('/state/')
    })

    it('should load the page title and breadcrumbs', () => {
        cy.get('[data-cy="state-index-title"]').contains('States').should('be.visible')
        cy.get('[data-cy="breadcrumb"]').should('exist')
        cy.get('[data-cy="breadcrumb"] a').contains('Home').should('exist')
        cy.get('[data-cy="breadcrumb"] span').contains('States').should('exist')
    })

    it('should display countries with state listings', () => {
        cy.get('[data-cy="country-united-states"]').should('be.visible')
        cy.get('[data-cy="country-count-united-states"]').invoke('text').should(v => {
            expect(Number.isInteger(+v), 'input should be a number').to.eq(true)
        })
        cy.get('[data-cy="country-canada"]').should('be.visible')
        cy.get('[data-cy="country-count-canada"]').invoke('text').should(v => {
            expect(Number.isInteger(+v), 'input should be a number').to.eq(true)
        })
    })

    it('should display US states with image counts', () => {
        const expectedStates = ['idaho', 'california', 'oregon', 'washington', 'nevada']

        expectedStates.forEach(state => {
            cy.get(`[data-cy="state-${state}"]`).should('exist')
            cy.get(`[data-cy="state-count-${state}"]`).invoke('text').should(v => {
                expect(Number.isInteger(+v), 'input should be a number').to.eq(true)
            })
        })
    })


    it('should navigate to a specific state page when clicked', () => {
        cy.get('[data-cy="state-link-california"]').click()
        cy.url().should('include', '/state/california/')
    })

    it('should display the correct count of states for United States', () => {
        // Get all US state elements and verify count is reasonable
        cy.get('a:contains("United States")').parent().find('ul > li').should('have.length.at.least', 40)
    })
})