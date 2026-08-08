describe('Highway Names Page Tests', () => {
    context('Named highway page', () => {
        beforeEach(() => {
            cy.visit('/highwaynames/albany-corvallis-highway-no-31/')
        })

        it('should display highway name title', () => {
            cy.get('h1').should('be.visible')
            cy.get('h1').should('contain', 'Albany')
        })

        it('should have working breadcrumb navigation', () => {
            cy.get('nav[aria-label="Breadcrumb"]').should('be.visible')
            
            // Test home breadcrumb link
            cy.get('nav[aria-label="Breadcrumb"] a[href="/"]').click()
            cy.url().should('eq', Cypress.config().baseUrl)
        })

        it('should have working global navigation link', () => {
            cy.visit('/highwaynames/albany-corvallis-highway-no-31/')
            
            // Test Map link from global nav
            cy.contains('nav a', 'Map').click()
            cy.url().should('include', '/map')
        })

        it('should display sign content', () => {
            // Named highway pages show signs along that route
            cy.get('a[href*="/sign/"]').should('exist')
        })

        it('should navigate to sign page', () => {
            // Click on a sign link
            cy.get('a[href*="/sign/"]').first().click()
            cy.url().should('include', '/sign/')
        })

        it('should display highway associations if present', () => {
            // Check for highway links or references
            cy.get('body').then($body => {
                if ($body.find('a[href*="/highway/"]').length > 0) {
                    cy.get('a[href*="/highway/"]').first().should('be.visible')
                }
            })
        })
    })

    context('Named highway with highway link', () => {
        it('should navigate to associated highway if linked', () => {
            cy.visit('/highwaynames/albany-corvallis-highway-no-31/')
            
            cy.get('body').then($body => {
                if ($body.find('a[href*="/highway/"]').length > 0) {
                    cy.get('a[href*="/highway/"]').first().click()
                    cy.url().should('include', '/highway/')
                }
            })
        })
    })

    context('Named highway pagination', () => {
        it('should handle pagination if highway has many signs', () => {
            cy.visit('/highwaynames/albany-corvallis-highway-no-31/')
            
            cy.get('body').then($body => {
                if ($body.find('[data-cy="pagination"]').length > 0) {
                    cy.get('[data-cy="pagination"]').should('be.visible')
                }
            })
        })
    })
})
